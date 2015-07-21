/*
Package core is the perfect foundation for any web application.
It allows you to connect all and only the components you need in a flexible and efficient way.

A handlers (or *middlewares*) stack is used to pass data in line, from the first to the last handler.
So you can perform actions downstream, then filter and manipulate the response upstream.

No handlers are bundled in this package.

Getting started

"Hello, World!" example with request logging:

	package main

	import (
		"fmt"
		"log"
		"time"

		"github.com/volatile/core"
	)

	func main() {
		// Log
		core.Use(func(c *core.Context) {
			start := time.Now()
			c.Next()
			log.Printf(" %s  %s  %s", c.Request.Method, c.Request.URL, time.Since(start))
		})

		// Response
		core.Use(func(c *core.Context) {
			fmt.Fprint(c.ResponseWriter, "Hello, World!")
		})

		core.Run()
	}

After running, the application is reachable at "http://localhost:8080/".

Official handlers

In order of usability in you app:

- Log — Requests logging — https://github.com/volatile/log
- Compress — Responses compressing — https://github.com/volatile/compress
- CORS — Cross-Origin Resource Sharing support — https://github.com/volatile/cors
- Others are coming…

Official helpers

Helpers are just syntactic sugars to ease repetitive code and improve readability of you app.

- Route — Flexible routing helper — https://github.com/volatile/route
- Response — Readable response helper — https://github.com/volatile/response
- Others are coming…

Context

All handlers are functions that receive a context: func(*core.Context).
A Volatile context encapsulates the well known ["*http.Request"](http://golang.org/pkg/net/http/#Request) and ["http.ResponseWriter"](http://golang.org/pkg/net/http/#ResponseWriter), from the standard ["net/http"](http://golang.org/pkg/net/http/) package.

Next

Simply use the context's "Next()" method to go to the next handler.

	core.Use(func(c *core.Context) {
		c.Next()
	})

Pass data

To transmit data from a handler to another, the *coreContext has a Data field, which is a map[string]interface{}.

	// Set data
	core.Use(func(c *core.Context) {
		c.Data["id"] = 123
	})

	// Read data
	core.Use(func(c *core.Context) {
		println(c.Data["id"].(int))
	})

Response writer binding

If some of your handlers need to transform the request before sending it, they can't just use the same ResponseWriter all the stack long.
To do so, the Core provides a ResponseWriterBinder structure that has the same signature as an http.ResponseWriter, but that redirects the response upstream, to an io.Writer that will write the original http.ResponseWriter.

In other words, the ResponseWriterBinder has an output (the ResponseWriter used before setting the binder) and an input (an overwritten ResponseWriter used by the next handlers).
The Compress package (https://github.com/volatile/compress/blob/master/handler.go) is a good example.

If you need to do something just before writing the response body (like setting headers, as you can't do that after), use the BeforeWrite field.

	core.Use(func(c *core.Context) {
		// 1. Set the output
		gzw := gzip.NewWriter(c.ResponseWriter)
		defer gzw.Close()

		// 2. Set the binder
		rwb = ResponseWriterBinder{
			Writer:         gzw, // The binder output.
			ResponseWriter: c.ResponseWriter, // Keep the same Header() and WriteHeader() methods. Only the Write method change internally.
			BeforeWrite:    func(b []byte) {
				// Do something with b before writing the response body.
			},
		}

		// 3. Set the input
		c.ResponseWriter = rwb
	})

	core.Use(func(c *core.Context) {
		// The overwritten context's ResponseWriter is used in a transparent way.
		c.ResponseWriter.Write([]byte("Hello, World!"))
	})

Things to know

- When a handler writes the body of a response, it brakes the handlers chain so a c.Next() call has no effect.
- Remember that response headers must be set before the body is written. After that, trying to set a header has no effect.

Custom port

To let the server listen on a custom port, use the "-port [port]" parameter on launch.

Environments

Some handlers can have different behaviors depending on the environment the server is running.
By default, the Core suppose the server is launched in a development environment.
When running your application in a production environment, use the "-production" parameter on launch.

In your code, you have access to the core.Production flag to distinguish the environment.
*/
package core
