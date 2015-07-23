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
		index:          -1, // Begin with -1 because the NextWriter will increment index before calling the first handler.
	}

	// Enter the handlers stack.
	// We use a binder to set the c.written flag on first write and break handlers chain.
	c.ResponseWriter = ResponseWriterBinder{
		Writer:         c.ResponseWriter,
		ResponseWriter: c.ResponseWriter,
		BeforeWrite:    func([]byte) { c.written = true },
	}

	// Sau to the client to "keep-alive" by default.
	c.ResponseWriter.Header().Set("Connection", "keep-alive")

	c.Next()
}
