package core

import "net/http"

// Context contains the data used on the wire of each request.
type Context struct {
	index          int16
	Request        *http.Request
	Response       []byte // Used by handlers to avoid writing the http.ResponseWriter multiple times.
	responseWriter http.ResponseWriter
}

// Next calls the next handler in the stack.
func (c *Context) Next() {
	c.index++
	handlers[c.index](c)
}
