package core

import "net/http"

type handlersStack []func(*Context)

var handlers handlersStack

// Use adds a handler to the handlers stack.
func Use(h func(*Context)) {
	handlers = append(handlers, h)
}

func (h handlersStack) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Init a new context for the request.
	c := &Context{
		Request: r,
		index:   -1, // Begin with -1 because "c.Next()" will increment index before calling the first handler.
	}
	c.ResponseWriter = contextWriter{w, c} // Use a binder to set the c.written flag on first write.

	// Use default headers.
	c.ResponseWriter.Header().Set("Cache-Control", "no-cache")
	c.ResponseWriter.Header().Set("Connection", "keep-alive")
	c.ResponseWriter.Header().Set("Vary", "Accept-Encoding")

	c.Next() // Enter the handlers stack.
}
