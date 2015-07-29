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

// SetDetectedContentType detects and sets and returns the response content type.
func SetDetectedContentType(c *core.Context, b []byte) string {
	ct := c.ResponseWriter.Header().Get("Content-Type")
	if len(ct) == 0 {
		ct := http.DetectContentType(b)
		c.ResponseWriter.Header().Set("Content-Type", ct)
	}
	return ct
}
