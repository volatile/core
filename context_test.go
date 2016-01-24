package core

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRecover(t *testing.T) {
	statusWant := http.StatusInternalServerError
	bodyWant := http.StatusText(http.StatusInternalServerError) + "\n"

	hs := NewHandlersStack()
	hs.Use(func(c *Context) {
		panic("")
	})

	oldOut := os.Stdout
	log.SetOutput(ioutil.Discard)

	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	hs.ServeHTTP(w, r)

	log.SetOutput(oldOut)

	statusGot := w.Code
	if statusWant != statusGot {
		t.Errorf("status code: want %d, got %d", statusWant, statusGot)
	}

	bodyGot := w.Body.String()
	if bodyWant != bodyGot {
		t.Errorf("body: want %q, got %q", bodyWant, bodyGot)
	}
}

func TestRecoverCustom(t *testing.T) {
	statusWant := http.StatusServiceUnavailable
	bodyWant := http.StatusText(http.StatusServiceUnavailable)
	var errorWant, errorGot interface{}
	errorWant = "foobar"

	hs := NewHandlersStack()
	hs.HandlePanic(func(c *Context) {
		errorGot = c.Data["panic"]
		c.ResponseWriter.WriteHeader(statusWant)
		c.ResponseWriter.Write([]byte(bodyWant))
	})
	hs.Use(func(c *Context) {
		defer c.Recover()
		panic(errorWant)
	})

	oldOut := os.Stdout
	log.SetOutput(ioutil.Discard)

	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	hs.ServeHTTP(w, r)

	log.SetOutput(oldOut)

	if errorWant != errorGot {
		t.Errorf("panic error: want %q, got %q", errorWant, errorGot)
	}

	statusGot := w.Code
	if statusWant != statusGot {
		t.Errorf("status code: want %d, got %d", statusWant, statusGot)
	}

	bodyGot := w.Body.String()
	if bodyWant != bodyGot {
		t.Errorf("body: want %q, got %q", bodyWant, bodyGot)
	}
}
