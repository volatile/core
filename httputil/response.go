package httputil

import (
	"io"
	"net/http"
	"reflect"

	"github.com/volatile/core"
)

// responseWriterBinder represents a binder that redirects the downstream response writing into a writer that will take care itself to write back the original response writer.
type responseWriterBinder struct {
	io.Writer
	http.ResponseWriter
	before []func([]byte) // Stores a set of functions that will be triggered just before writing the repsonse.
}

// Write calls the upstream writer after executing the functions in the "before" fields.
func (w responseWriterBinder) Write(p []byte) (int, error) {
	for _, f := range w.before {
		f(p)
	}
	return w.Writer.Write(p)
}

// BindResponseWriter redirects the downstream response writing into the writer "w", that will take care itself to write back the original response writer.
// "before" can be a set of functions that will be triggered juste before writing the repsonse.
func BindResponseWriter(w io.Writer, c *core.Context, before ...func([]byte)) {
	c.ResponseWriter = responseWriterBinder{w, c.ResponseWriter, before}
}

// ResponseStatus returns the HTTP response status.
// Remember that the status is only set by the server after "WriteHeader()" has been called.
func ResponseStatus(w http.ResponseWriter) int {
	return int(httpResponseStruct(reflect.ValueOf(w)).FieldByName("status").Int())
}

// httpResponseStruct returns the response structure after going trough all the intermediary response writers.
func httpResponseStruct(v reflect.Value) reflect.Value {
	switch v.Type().String() {
	case "*http.response":
		return v.Elem()
	default:
		return httpResponseStruct(v.FieldByName("ResponseWriter").Elem())
	}
}

// SetDetectedContentType detects, sets and returns the response "Conten-Type" header value.
func SetDetectedContentType(w http.ResponseWriter, p []byte) string {
	ct := w.Header().Get("Content-Type")
	if ct == "" {
		ct = http.DetectContentType(p)
		w.Header().Set("Content-Type", ct)
	}
	return ct
}
