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
	// Is set (false)
	request, err := http.NewRequest("GET", "/isset", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/isset" {
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetBool("some_bool")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "false" {
		t.Fatalf("bad body response, not match")
	}

	// Set value
	request, err = http.NewRequest("GET", "/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/set" {
			sess.SetBool("some_bool", true)
			if _, err := w.Write([]byte(`ok`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
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
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.GetBool("some_bool", false)))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "true" {
		t.Fatalf("bad body response, not match")
	}

	// Is set (true)
	request, err = http.NewRequest("GET", "/isset", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/isset" {
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetBool("some_bool")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "true" {
		t.Fatalf("bad body response, not match")
	}

	// Del
	request, err = http.NewRequest("GET", "/del", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/del" {
			sess.DelBool("some_bool")
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetBool("some_bool")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "false" {
		t.Fatalf("bad body response, not match")
	}
}

func TestSessionInt(t *testing.T) {
	// Is set (false)
	request, err := http.NewRequest("GET", "/isset", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/isset" {
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetInt("some_int")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "false" {
		t.Fatalf("bad body response, not match")
	}

	// Set value
	request, err = http.NewRequest("GET", "/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/set" {
			sess.SetInt("some_int", 5)
			if _, err := w.Write([]byte(`ok`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
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
			if _, err := w.Write([]byte(fmt.Sprintf("%d", sess.GetInt("some_int", 0)))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "5" {
		t.Fatalf("bad body response, not match")
	}

	// Is set (true)
	request, err = http.NewRequest("GET", "/isset", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/isset" {
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetInt("some_int")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "true" {
		t.Fatalf("bad body response, not match")
	}

	// Del
	request, err = http.NewRequest("GET", "/del", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/del" {
			sess.DelInt("some_int")
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetInt("some_int")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "false" {
		t.Fatalf("bad body response, not match")
	}
}

func TestSessionInt64(t *testing.T) {
	// Is set (false)
	request, err := http.NewRequest("GET", "/isset", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/isset" {
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetInt64("some_int64")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "false" {
		t.Fatalf("bad body response, not match")
	}

	// Set value
	request, err = http.NewRequest("GET", "/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/set" {
			sess.SetInt64("some_int64", 10)
			if _, err := w.Write([]byte(`ok`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
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
			if _, err := w.Write([]byte(fmt.Sprintf("%d", sess.GetInt64("some_int64", 0)))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "10" {
		t.Fatalf("bad body response, not match")
	}

	// Is set (true)
	request, err = http.NewRequest("GET", "/isset", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/isset" {
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetInt64("some_int64")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "true" {
		t.Fatalf("bad body response, not match")
	}

	// Del
	request, err = http.NewRequest("GET", "/del", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/del" {
			sess.DelInt64("some_int64")
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetInt64("some_int64")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "false" {
		t.Fatalf("bad body response, not match")
	}
}

func TestSessionString(t *testing.T) {
	// Is set (false)
	request, err := http.NewRequest("GET", "/isset", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/isset" {
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetString("some_str")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "false" {
		t.Fatalf("bad body response, not match")
	}

	// Set value
	request, err = http.NewRequest("GET", "/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/set" {
			sess.SetString("some_str", "test")
			if _, err := w.Write([]byte(`ok`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
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
			if _, err := w.Write([]byte(sess.GetString("some_str", ""))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "test" {
		t.Fatalf("bad body response, not match")
	}

	// Is set (true)
	request, err = http.NewRequest("GET", "/isset", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/isset" {
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetString("some_str")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "true" {
		t.Fatalf("bad body response, not match")
	}

	// Del
	request, err = http.NewRequest("GET", "/del", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+SessionId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/del" {
			sess.DelString("some_str")
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.IsSetString("some_str")))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "false" {
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
	if string(bytes) != `{"Bool":{"some_bool":true},"Int":{"some_int":5},"Int64":{"some_int64":10},"String":{"some_str":"test"}}` {
		t.Fatal("actual file content, not match")
	}
	err = os.Remove(fname)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSessionDoNotCreateSessionFileForDefValues(t *testing.T) {
	// Set default values
	request, err := http.NewRequest("GET", "/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/set" {
			sess.SetBool("some_bool", false)
			sess.SetInt("some_int", 0)
			sess.SetInt64("some_int64", 0)
			sess.SetString("some_str", "")
			if _, err := w.Write([]byte(`ok`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "ok" {
		t.Fatalf("bad body response, not match")
	}

	// Remember session id
	var sessId string
	if len(recorder.Result().Cookies()) > 0 {
		sessId = recorder.Result().Cookies()[0].Value
	}
	if sessId == "" {
		t.Fatalf("session identifier is not defined")
	}

	// Get value
	request, err = http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+sessId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/get" {
			if _, err := w.Write([]byte(fmt.Sprintf(
				"(%v)(%v)(%v)(%v)",
				sess.GetBool("some_bool", false),
				sess.GetInt("some_int", 0),
				sess.GetInt64("some_int64", 0),
				sess.GetString("some_str", ""),
			))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "(false)(0)(0)()" {
		t.Fatalf("bad body response, not match")
	}

	// Check session file
	fname := "./../tmp" + string(os.PathSeparator) + sessId
	_, err = ioutil.ReadFile(fname)
	if err == nil {
		_ = os.Remove(fname)
		t.Fatalf("session file in tmp folder do not must exists")
	}
}

func TestSessionDestroy(t *testing.T) {
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
			sess.SetInt("some_var", 1)
			if _, err := w.Write([]byte(`ok`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "ok" {
		t.Fatalf("bad body response, not match")
	}

	// Remember session id
	var sessId string
	if len(recorder.Result().Cookies()) > 0 {
		sessId = recorder.Result().Cookies()[0].Value
	}
	if sessId == "" {
		t.Fatalf("session identifier is not defined")
	}

	// Get value
	request, err = http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+sessId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/get" {
			if _, err := w.Write([]byte(fmt.Sprintf("%v", sess.GetInt("some_var", 0)))); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "1" {
		t.Fatalf("bad body response, not match")
	}

	// Check destroy
	request, err = http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Cookie", "session="+sessId)
	recorder = httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := New(w, r, "./../tmp")
		defer sess.Close()
		if r.URL.Path == "/get" {
			sess.SetInt("some_var", 2)
			err := sess.Destroy()
			if err == nil {
				if _, err := w.Write([]byte(`OK`)); err != nil {
					fmt.Printf("%s\n", err.Error())
				}
			} else {
				if _, err := w.Write([]byte(`ERROR`)); err != nil {
					fmt.Printf("%s\n", err.Error())
				}
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`404`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}).ServeHTTP(recorder, request)
	if recorder.Body.String() != "OK" {
		t.Fatalf("bad body response, not match")
	}

	// Check session file
	fname := "./../tmp" + string(os.PathSeparator) + sessId
	_, err = ioutil.ReadFile(fname)
	if err == nil {
		_ = os.Remove(fname)
		t.Fatalf("session file in tmp folder do not must exists")
	}
}
