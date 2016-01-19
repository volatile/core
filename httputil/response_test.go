package httputil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/volatile/core"
)

func TestResponseWriterBinder(t *testing.T) {
	headerKey := "foo"
	headerValueWant := "bar"
	bodyWant := "foobar"

	w := httptest.NewRecorder()
	c := &core.Context{
		ResponseWriter: w,
	}

	BindResponseWriter(c.ResponseWriter, c, func(p []byte) {
		c.ResponseWriter.Header().Set(headerKey, headerValueWant)
	})

	c.ResponseWriter.Write([]byte(bodyWant))

	headerValueGot := w.Header().Get(headerKey)
	if headerValueWant != headerValueGot {
		t.Errorf("header: want %q, got %q", headerValueWant, headerValueGot)
	}

	bodyGot := w.Body.String()
	if bodyWant != bodyGot {
		t.Errorf("body: want %q, got %q", bodyWant, bodyGot)
	}
}

func TestResponseStatus(t *testing.T) {
	statusWant := http.StatusForbidden
	var statusGot, customStatusGot int

	type CustomResponseWriter struct {
		http.ResponseWriter
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, http.StatusText(statusWant), statusWant)
		statusGot = ResponseStatus(w)
		customStatusGot = ResponseStatus(CustomResponseWriter{w})
	}))
	defer ts.Close()

	if _, err := http.Get(ts.URL); err != nil {
		t.Fatal(err)
	}

	if statusWant != statusGot {
		t.Errorf("http.ResponseWriter: want %d, got %d", statusWant, statusGot)
	}

	if customStatusGot != statusGot {
		t.Errorf("CustomResponseWriter: want %d, got %d", customStatusGot, statusGot)
	}
}

func TestSetDetectedContentType(t *testing.T) {
	headerKey := "Content-Type"
	headerValueWant := "text/html; charset=utf-8"

	w := httptest.NewRecorder()
	SetDetectedContentType(w, []byte("<!DOCTYPE html>"))

	headerValueGot := w.Header().Get(headerKey)
	if headerValueWant != headerValueGot {
		t.Errorf("set detected content type: want %q, got %q", headerValueWant, headerValueGot)
	}
}
