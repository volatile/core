package core

import "net/http"

// Context contains all the data needed during the serving flow.
// It contains the standard "http.ResponseWriter" and "*http.Request".
// You can pass data through handlers thanks to the Data field.
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Data           map[string]interface{} // Generic field to pass arbitrary data through the stack.
	index          int                    // Keeps the actual handler index.
	written        bool                   // A flag to know if the response has been written.
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

// Write sets the context's "written" flag before writing the response.
func (w contextWriter) Write(p []byte) (int, error) {
	w.context.written = true
	return w.ResponseWriter.Write(p)
}
