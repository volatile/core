package coreutil

import (
	"net/http/httptest"
	"testing"

	"github.com/volatile/core"
)

func TestResponseWriterBinder(t *testing.T) {
	headerKey := "Content-Type"
	headerValue := "text/plain"
	body := "foobar"

	w := httptest.NewRecorder()
	c := &core.Context{
		ResponseWriter: w,
	}

	BindResponseWriter(c.ResponseWriter, c, func(p []byte) {
		c.ResponseWriter.Header().Set(headerKey, headerValue)
	})

	c.ResponseWriter.Write([]byte(body))

	if w.Body.String() != body {
		t.Errorf("response writer binder: body: want %q, got %q", body, w.Body.String())
	}

	if w.Header().Get(headerKey) != headerValue {
		t.Errorf("response writer binder: %q header: want %q, got %q", headerKey, headerValue, w.Header().Get(headerKey))
	}
}

func TestSetDetectedContentType(t *testing.T) {
	headerKey := "Content-Type"
	headerValue := "text/html; charset=utf-8"

	w := httptest.NewRecorder()
	w.WriteHeader(403)

	SetDetectedContentType(w, []byte("<!DOCTYPE html>"))

	if w.Header().Get(headerKey) != headerValue {
		t.Errorf("set detected content type: want %q, got %q", headerValue, w.Header().Get(headerKey))
	}
}
