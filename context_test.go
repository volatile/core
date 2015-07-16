package core

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestResponseWriteBinder(t *testing.T) {
	in := "Foobar"
	out := new(bytes.Buffer)

	w := httptest.NewRecorder()
	binder := &ResponseWriterBinder{
		Writer:         out,
		ResponseWriter: w,
	}

	binder.Write([]byte(in))

	if out.String() != in {
		t.Errorf("header: want %q, got %q", in, out)
	}
}
