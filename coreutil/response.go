package coreutil

import (
	"net/http"
	"reflect"
)

// ResponseStatus returns the HTTP response status.
// Remember that the status is only set by the server after "ResponseWriter.WriteHeader()"" has been called.
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

// SetDetectedContentType detects and sets and returns the response content type.
func SetDetectedContentType(w http.ResponseWriter, b []byte) string {
	ct := w.Header().Get("Content-Type")
	if ct == "" {
		ct = http.DetectContentType(b)
		w.Header().Set("Content-Type", ct)
	}
	return ct
}
