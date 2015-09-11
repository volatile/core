package core

import "net/http"

// Context contains all the data needed during the serving flow.
// It contains the standard http.ResponseWriter and *http.Request.
// The Data field can be used to pass all kind of data through the handler stack.
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Data           map[string]interface{}
	index          int           // Keeps the actual handler index.
	handlersStack  HandlersStack // Keeps the reference to the actual handlers stack.
	written        bool          // A flag to know if the response has been written.
}

// Next calls the next handler in the stack, but only if the response isn't already written.
func (c *Context) Next() {
	// Call the next handler only if there is one and the response hasn't been written.
	if !c.written && len(c.handlersStack)-1 > c.index {
		c.index++
		c.handlersStack[c.index](c)
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
