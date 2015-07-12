package core

import "net/http"

var handlers stack

type stack []func(*Context)

// Use adds a handler to the handlers stack.
func Use(m func(*Context)) {
	handlers = append(handlers, m)
}

func (m stack) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Init a new context for the request.
	c := &Context{
		Request: r,
		Response: Response{
			Status: http.StatusOK,
			Header: make(http.Header),
		},
	}

	// Throw the fresh context in the handlers stack.
	handlers[0](c)

	// Send the final response.
	copyHeader(w.Header(), c.Response.Header)
	w.WriteHeader(c.Response.Status)
	w.Write(c.Response.Body)
}

// copyHeader is inspired by copyHeader from net/http/httputil/reverseproxy.go
func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
