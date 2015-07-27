package coreutil

import (
	"net/http"
	"reflect"

	"github.com/volatile/core"
)

// ResponseStatus returns the HTTP response status.
// Remember that the status is only set by the server after ResponseWriter.WriteHeader() has been called.
func ResponseStatus(c *core.Context) int {
	return int(httpResponseStruct(reflect.ValueOf(c.ResponseWriter)).FieldByName("status").Int())
}

// httpResponseStruct returns the response structure after going trough all the intermediary responseWriter Binders.
func httpResponseStruct(v reflect.Value) reflect.Value {
	switch v.Type().String() {
	case "core.ResponseWriterBinder":
		return httpResponseStruct(v.FieldByName("ResponseWriter").Elem())
	case "*http.response":
		return v.Elem()
	default:
		panic("coreutil: call of httpResponseStruct on unknown interface type")
	}
}

// SetDetectedContentType detects and sets the correct content type.
func SetDetectedContentType(c *core.Context, b []byte) {
	if len(c.ResponseWriter.Header().Get("Content-Type")) == 0 {
		c.ResponseWriter.Header().Set("Content-Type", http.DetectContentType(b))
	}
}
