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
		ResponseWriter: w,
		Request:        r,
	}

	// Throw the fresh context in the handlers stack.
	handlers[0](c)
}
