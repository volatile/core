package core

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	statusWant := http.StatusForbidden
	headerKey := "foo"
	headerValueWant := "bar"
	bodyWant := "foobar"

	hs := NewHandlersStack()
	hs.Use(func(c *Context) { c.Next() })
	hs.Use(func(c *Context) { c.Next() })
	hs.Use(func(c *Context) {
		c.ResponseWriter.Header().Set(headerKey, headerValueWant)
		c.ResponseWriter.WriteHeader(statusWant)
		c.ResponseWriter.Write([]byte(bodyWant))
	})
	hs.Use(func(c *Context) {
		c.ResponseWriter.Write([]byte("baz"))
	})

	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	hs.ServeHTTP(w, r)

	statusGot := w.Code
	if statusWant != statusGot {
		t.Errorf("status code: want %d, got %d", statusWant, statusGot)
	}

	headerValueGot := w.Header().Get(headerKey)
	if headerValueWant != headerValueGot {
		t.Errorf("header: want %q, got %q", headerValueWant, headerValueGot)
	}

	bodyGot := w.Body.String()
	if bodyWant != bodyGot {
		t.Errorf("body: want %q, got %q", bodyWant, bodyGot)
	}
}
