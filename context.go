package core

import (
	"io"
	"net/http"
)

// Context contains the data used on the wire of each request.
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	index          int
}

// Next calls the next handler in the stack.
func (c *Context) Next() {
	c.index++
	handlers[c.index](c)
}

// NextWriter calls the next handler in the stack with a new ResponseWriter.
// It can be used by handlers (middlewares) to transfer a new writer.
// The best example is in the "compress" package.
func (c *Context) NextWriter(w http.ResponseWriter) {
	c.ResponseWriter = w
	c.index++
	handlers[c.index](c)
}

// ResponseWriterBinder can be used by handlers to pass a new ResponseWriter to the next handlers and write back to the original ResponseWriter.
type ResponseWriterBinder struct {
	io.Writer
	http.ResponseWriter
	BeforeWrite func([]byte)
}

func (w ResponseWriterBinder) Write(b []byte) (int, error) {
	w.BeforeWrite(b)
	return w.Writer.Write(b)
}
