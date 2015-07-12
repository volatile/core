package core

import "net/http"

// Context contains the data used on the wire of each request.
type Context struct {
	index    int
	Request  *http.Request
	Response Response
}

// Next calls the next handler in the stack.
func (c *Context) Next() {
	c.index++
	handlers[c.index](c)
}

// Response contains the data that will be written in the final response.
// It is used by handlers to avoid writing the main http.ResponseWriter multiple times.
type Response struct {
	Status int
	Header http.Header
	Body   []byte
}
