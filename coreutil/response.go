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
func SetDetectedContentType(w http.ResponseWriter, b []byte) {
	if len(w.Header().Get("Content-Type")) == 0 {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
}
