package core

import "net/http"

// handlersStack contains a set of handlers.
type handlersStack []func(*Context)

// handlers contains the handler stack used for serving.
var handlers handlersStack

// Use adds a handler to the handler stack.
func Use(h func(*Context)) {
	handlers = append(handlers, h)
}

// ServeHTTP makes a context for the request, sets some "good practice" default headers and enters the handler stack.
func (h handlersStack) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Make a context for the request.
	c := &Context{
		Request: r,
		index:   -1, // Begin with -1 because Next will increment the index before calling the first handler.
	}
	c.ResponseWriter = contextWriter{w, c} // Use a binder to set the context's written flag on the first write.

	// Set some "good practice" default headers.
	c.ResponseWriter.Header().Set("Cache-Control", "no-cache")
	c.ResponseWriter.Header().Set("Connection", "keep-alive")
	c.ResponseWriter.Header().Set("Vary", "Accept-Encoding")

	c.Next() // Enter the handler stack.
}
