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
	}

	// Throw the fresh context in the handlers stack.
	handlers[0](c)

	// Send the final response.
	w.Write(c.Response)
}
