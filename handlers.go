package core

import "net/http"

// HandlersStack contains a set of handlers.
type HandlersStack []func(*Context)

// DefaultHandlersStack contains the default handler stack used for serving.
var DefaultHandlersStack = NewHandlersStack()

// NewHandlersStack allocates and returns a new .
func NewHandlersStack() HandlersStack {
	return make(HandlersStack, 0)
}

// Use adds a handler to the handler stack.
func (hs *HandlersStack) Use(h func(*Context)) {
	*hs = append(*hs, h)
}

// Use adds a handler to the handler stack.
func Use(h func(*Context)) {
	DefaultHandlersStack.Use(h)
}

// ServeHTTP makes a context for the request, sets some "good practice" default headers and enters the handler stack.
func (hs HandlersStack) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Make a context for the request.
	c := &Context{
		Request:       r,
		Data:          make(map[string]interface{}),
		index:         -1, // Begin with -1 because Next will increment the index before calling the first handler.
		handlersStack: hs,
	}
	c.ResponseWriter = contextWriter{w, c} // Use a binder to set the context's written flag on the first write.

	// Set some "good practice" default headers.
	c.ResponseWriter.Header().Set("Cache-Control", "no-cache")
	c.ResponseWriter.Header().Set("Connection", "keep-alive")
	c.ResponseWriter.Header().Set("Vary", "Accept-Encoding")

	c.Next() // Enter the handler stack.
}
