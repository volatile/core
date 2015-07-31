package core

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	headerKey := "foo"
	headerValue := "bar"
	status := http.StatusForbidden
	body := "foobar"

	Use(func(c *Context) {
		c.ResponseWriter.Header().Set(headerKey, headerValue)
		c.ResponseWriter.WriteHeader(status)
		c.ResponseWriter.Write([]byte(body))
	})

	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	handlers.ServeHTTP(w, r)

	if w.Code != status {
		t.Errorf("status code: want %q, got %q", status, w.Code)
	}

	headerFooValue := w.Header().Get(headerKey)
	if headerFooValue != headerValue {
		t.Errorf("header: want %q, got %q", headerValue, headerFooValue)
	}

	wbody := w.Body.String()
	if wbody != body {
		t.Errorf("body: want %q, got %q", body, wbody)
	}
}
