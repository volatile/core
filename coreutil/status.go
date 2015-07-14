package coreutil

import (
	"net/http"
	"reflect"
)

// ResponseStatus returns the HTTP response status.
// Remember that the status is only set by the server after ResponseWriter.WriteHeader() has been called.
func ResponseStatus(w http.ResponseWriter) int {
	return int(httpResponseStruct(reflect.ValueOf(w)).FieldByName("status").Int())
}

// httpResponseStruct returns the response structure after going trough all theintermediary ResponseWriterBinder.
func httpResponseStruct(v reflect.Value) reflect.Value {
	switch v.Type().String() {
	case "core.ResponseWriterBinder":
		return httpResponseStruct(v.FieldByName("ResponseWriter").Elem())
	case "*http.response":
		return v.Elem()
	default:
		panic("core: call of httpResponseStruct on unknown interface type")
	}
}
