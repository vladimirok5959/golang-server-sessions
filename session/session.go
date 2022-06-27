package session

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type vars struct {
	Bool   map[string]bool
	Int    map[string]int
	Int64  map[string]int64
	String map[string]string
}

type Session struct {
	w       http.ResponseWriter
	r       *http.Request
	tmpdir  string
	varlist *vars
	changed bool
	hash    string
}

func New(w http.ResponseWriter, r *http.Request, tmpdir string) *Session {
	s := Session{
		w:       w,
		r:       r,
		tmpdir:  tmpdir,
		varlist: &vars{},
		changed: false,
		hash:    "",
	}

	cookie, err := r.Cookie("session")
	if err == nil && len(cookie.Value) == 40 {
		// Load from file
		s.hash = cookie.Value
		fname := strings.Join([]string{s.tmpdir, s.hash}, string(os.PathSeparator))
		f, err := os.Open(fname)
		if err == nil {
			defer f.Close()
			dec := json.NewDecoder(f)
			err = dec.Decode(&s.varlist)
			if err == nil {
				return &s
			}

			// Update file last modify time if needs
			if info, err := os.Stat(fname); err == nil {
				if time.Since(info.ModTime()) > 30*time.Minute {
					_ = os.Chtimes(fname, time.Now(), time.Now())
				}
			}
		}
	} else {
		// Create new
		rand.Seed(time.Now().Unix())

		// Real remote IP for proxy servers
		rRemoteAddr := r.RemoteAddr
		if r.Header.Get("X-Real-IP") != "" && len(r.Header.Get("X-Real-IP")) <= 25 {
			rRemoteAddr = rRemoteAddr + ", " + strings.TrimSpace(r.Header.Get("X-Real-IP"))
		} else if r.Header.Get("X-Forwarded-For") != "" && len(r.Header.Get("X-Forwarded-For")) <= 25 {
			rRemoteAddr = rRemoteAddr + ", " + strings.TrimSpace(r.Header.Get("X-Forwarded-For"))
		}

		sign := rRemoteAddr + r.Header.Get("User-Agent") + fmt.Sprintf("%d", int64(time.Now().Unix())) + fmt.Sprintf("%d", int64(rand.Intn(9999999-99)+99))
		s.hash = fmt.Sprintf("%x", sha1.Sum([]byte(sign)))

		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    s.hash,
			Path:     "/",
			Expires:  time.Now().Add(7 * 24 * time.Hour),
			HttpOnly: true,
		})
	}

	// Init empty
	s.varlist = &vars{
		Bool:   map[string]bool{},
		Int:    map[string]int{},
		Int64:  map[string]int64{},
		String: map[string]string{},
	}

	return &s
}

func (s *Session) Close() bool {
	if !s.changed {
		return false
	}

	r, err := json.Marshal(s.varlist)
	if err == nil {
		f, err := os.Create(strings.Join([]string{s.tmpdir, s.hash}, string(os.PathSeparator)))
		if err == nil {
			defer f.Close()
			_, err = f.Write(r)
			if err == nil {
				s.changed = false
				return true
			}
		}
	}

	return false
}

func (s *Session) Destroy() error {
	if s.tmpdir == "" || s.hash == "" {
		return nil
	}
	err := os.Remove(strings.Join([]string{s.tmpdir, s.hash}, string(os.PathSeparator)))
	if err == nil {
		s.changed = false
	}
	return err
}
