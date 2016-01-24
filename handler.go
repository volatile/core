package core

import "net/http"

// Handler represents a request handler.
type Handler func(*Context)

// HandlersStack contains a set of handlers.
type HandlersStack struct {
	Handlers     []Handler // The handlers stack.
	PanicHandler Handler   // The handler called in case of panic. Useful to send custom server error information. Context.Data["panic"] contains the panic error.
}

// defaultHandlersStack contains the default handlers stack used for serving.
var defaultHandlersStack = NewHandlersStack()

// NewHandlersStack returns a new NewHandlersStack.
func NewHandlersStack() *HandlersStack {
	return new(HandlersStack)
}

// Use adds a handler to the handlers stack.
func (hs *HandlersStack) Use(h Handler) {
	hs.Handlers = append(hs.Handlers, h)
}

// Use adds a handler to the default handlers stack.
func Use(h Handler) {
	defaultHandlersStack.Use(h)
}

// HandlePanic sets the panic handler of the handlers stack.
// Context.Data["panic"] contains the panic error.
func (hs *HandlersStack) HandlePanic(h Handler) {
	hs.PanicHandler = h
}

// HandlePanic sets the panic handler of the default handlers stack.
// Context.Data["panic"] contains the panic error.
func HandlePanic(h Handler) {
	defaultHandlersStack.HandlePanic(h)
}

// ServeHTTP makes a context for the request, sets some "good practice" default headers and enters the handlers stack.
func (hs *HandlersStack) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	defer c.Recover() // Always recover form panics.
	c.Next()          // Enter the handlers stack.
}
