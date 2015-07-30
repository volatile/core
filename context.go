package core

import "net/http"

// Context contains the data used on the wire of each request.
// You can pass data through handlers thanks to the Data field.
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Data           map[string]interface{} // Generic field to pass arbitrary data through the stack.
	index          int
	written        bool // Flag to know when the response is written
}

// Next calls the next handler in the stack, but only if the response is not already written.
func (c *Context) Next() {
	if !c.written {
		c.index++
		handlers[c.index](c)
	}
}

// contextWriter is a binder that sets the c.written flag on first write.
type contextWriter struct {
	http.ResponseWriter
	context *Context
}

func (w contextWriter) Write(p []byte) (int, error) {
	w.context.written = true
	return w.ResponseWriter.Write(p)
}
