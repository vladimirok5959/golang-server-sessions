package session

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var SessionId string = ""

func TestSessionBool(t *testing.T) {
	// Set value
	request, err := http.NewRequest("GET", "/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/set" {
			sess.SetBool("some_bool", true)
			w.Write([]byte(`ok`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`404`))
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "ok" {
		t.Fatalf("bad body response, not match")
	}

	// Remember session id
	if SessionId == "" && len(recorder.Result().Cookies()) > 0 {
		SessionId = recorder.Result().Cookies()[0].Value
	}

	// Get value
	request, err = http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/get" {
			w.Write([]byte(fmt.Sprintf("%v", sess.GetBool("some_bool", false))))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`404`))
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "true" {
		t.Fatalf("bad body response, not match")
	}
}

func TestSessionInt(t *testing.T) {
	// Set value
	request, err := http.NewRequest("GET", "/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/set" {
			sess.SetInt("some_int", 5)
			w.Write([]byte(`ok`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`404`))
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "ok" {
		t.Fatalf("bad body response, not match")
	}

	// Remember session id
	if SessionId == "" && len(recorder.Result().Cookies()) > 0 {
		SessionId = recorder.Result().Cookies()[0].Value
	}

	// Get value
	request, err = http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/get" {
			w.Write([]byte(fmt.Sprintf("%d", sess.GetInt("some_int", 0))))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`404`))
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "5" {
		t.Fatalf("bad body response, not match")
	}
}

func TestSessionString(t *testing.T) {
	// Set value
	request, err := http.NewRequest("GET", "/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/set" {
			sess.SetString("some_str", "test")
			w.Write([]byte(`ok`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`404`))
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "ok" {
		t.Fatalf("bad body response, not match")
	}

	// Remember session id
	if SessionId == "" && len(recorder.Result().Cookies()) > 0 {
		SessionId = recorder.Result().Cookies()[0].Value
	}

	// Get value
	request, err = http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/get" {
			w.Write([]byte(fmt.Sprintf("%s", sess.GetString("some_str", ""))))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`404`))
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "test" {
		t.Fatalf("bad body response, not match")
	}
}

func TestSessionActualFile(t *testing.T) {
	if SessionId == "" {
		t.Fatal("SessionId is empty")
	}
	fname := "./../tmp" + string(os.PathSeparator) + SessionId
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != `{"Bool":{"some_bool":true},"Int":{"some_int":5},"String":{"some_str":"test"}}` {
		t.Fatal("actual file content, not match")
	}
	err = os.Remove(fname)
	if err != nil {
		t.Fatal(err)
	}
}
