package core

import (
	"log"
	"net/http"
	"runtime"
)

// Context contains all the data needed during the serving flow.
// It contains the standard http.ResponseWriter and *http.Request.
// The Data field can be used to pass all kind of data through the handlers stack.
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Data           map[string]interface{}
	index          int            // Keeps the actual handler index.
	handlersStack  *HandlersStack // Keeps the reference to the actual handlers stack.
	written        bool           // A flag to know if the response has been written.
}

// Written tells if the response has been written.
func (c *Context) Written() bool {
	return c.written
}

// Next calls the next handler in the stack, but only if the response isn't already written.
func (c *Context) Next() {
	// Call the next handler only if there is one and the response hasn't been written.
	if !c.Written() && c.index < len(c.handlersStack.Handlers)-1 {
		c.index++
		c.handlersStack.Handlers[c.index](c)
	}
}

// Recover recovers form panics.
// It logs the stack and uses the PanicHandler (or a classic Internal Server Error) to write the response.
func (c *Context) Recover() {
	if err := recover(); err != nil {
		stack := make([]byte, 64<<10)
		stack = stack[:runtime.Stack(stack, false)]
		log.Printf("%v\n%s", err, stack)

		if !c.Written() {
			if c.handlersStack.PanicHandler != nil {
				c.handlersStack.PanicHandler(c, err)
			} else {
				http.Error(c.ResponseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
	}
}

// contextWriter represents a binder that catches a downstream response writing and sets the context's written flag on the first write.
type contextWriter struct {
	http.ResponseWriter
	context *Context
}

// Write sets the context's written flag before writing the response.
func (w contextWriter) Write(p []byte) (int, error) {
	w.context.written = true
	return w.ResponseWriter.Write(p)
}

// WriteHeader sets the context's written flag before writing the response header.
func (w contextWriter) WriteHeader(code int) {
	w.context.written = true
	w.ResponseWriter.WriteHeader(code)
}
