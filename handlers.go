package core

import (
	"log"
	"net/http"
	"runtime"
)

// HandlersStack contains a set of handlers.
type HandlersStack struct {
	Handlers     []func(*Context)            // The handlers stack.
	PanicHandler func(*Context, interface{}) // The handler called in case of panic. Useful to send custom server error information.
}

// defaultHandlersStack contains the default handlers stack used for serving.
var defaultHandlersStack = NewHandlersStack()

// NewHandlersStack returns a new NewHandlersStack.
func NewHandlersStack() *HandlersStack {
	return new(HandlersStack)
}

// Use adds a handler to the handlers stack.
func (hs *HandlersStack) Use(h func(*Context)) {
	hs.Handlers = append(hs.Handlers, h)
}

// Use adds a handler to the default handlers stack.
func Use(h func(*Context)) {
	defaultHandlersStack.Use(h)
}

// HandlePanic sets the panic handler of the handlers stack.
func (hs *HandlersStack) HandlePanic(h func(*Context, interface{})) {
	hs.PanicHandler = h
}

// HandlePanic sets the panic handler of the default handlers stack.
func HandlePanic(h func(*Context, interface{})) {
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

	// Always recover form panics.
	defer hs.recover(c)

	c.Next() // Enter the handlers stack.
}

func (hs *HandlersStack) recover(c *Context) {
	if err := recover(); err != nil {
		stack := make([]byte, 64<<10)
		stack = stack[:runtime.Stack(stack, false)]
		log.Printf("%v\n%s", err, stack)

		if !c.IsWritten() {
			if hs.PanicHandler != nil {
				hs.PanicHandler(c, err)
			} else {
				http.Error(c.ResponseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
	}
}
