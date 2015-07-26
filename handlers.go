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
		index:          -1, // Begin with -1 because "c.Next()" will increment index before calling the first handler.
	}

	// We use a binder to set the c.written flag on first write and break handlers chain.
	c.ResponseWriter = ResponseWriterBinder{
		Writer:         c.ResponseWriter,
		ResponseWriter: c.ResponseWriter,
		BeforeWrite:    func([]byte) { c.written = true },
	}

	// Use default headers.
	c.ResponseWriter.Header().Set("Cache-Control", "no-cache")
	c.ResponseWriter.Header().Set("Connection", "keep-alive")
	c.ResponseWriter.Header().Set("Vary", "Accept-Encoding")

	// Enter the handlers stack.
	c.Next()
}
