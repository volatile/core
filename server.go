package core

import "net/http"

type stack []func(*Context)

var handlers stack

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

// Run starts the server for listening and serving.
func Run() {
	if handlers == nil {
		panic("core: the handlers stack cannot be empty")
	}

	http.ListenAndServe(":8080", handlers)
}
