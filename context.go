package core

import "net/http"

// Context contains the data used on the wire of each request.
type Context struct {
	index    int
	Request  *http.Request
	Response Response // Used by handlers to avoid writing the http.ResponseWriter multiple times.
}

// Next calls the next handler in the stack.
func (c *Context) Next() {
	c.index++
	handlers[c.index](c)
}

// Response represents the response from an HTTP request.
type Response struct {
	Status int
	Header http.Header
	Body   []byte
}
