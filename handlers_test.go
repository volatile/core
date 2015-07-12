package core

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	Use(func(c *Context) {
		c.ResponseWriter.Header().Set("foo", "bar")
		c.ResponseWriter.WriteHeader(http.StatusForbidden)
		c.ResponseWriter.Write([]byte("foobar"))
	})

	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	handlers.ServeHTTP(w, r)

	if w.Code != http.StatusForbidden {
		t.Errorf("status code: want %q, got %q", http.StatusForbidden, w.Code)
	}

	headerFooValue := w.Header().Get("foo")
	if headerFooValue != "bar" {
		t.Errorf("header: want %q, got %q", "bar", headerFooValue)
	}

	body := w.Body.String()
	if body != "foobar" {
		t.Errorf("header: want %q, got %q", "foobar", body)
	}
}
