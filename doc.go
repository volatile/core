/*
Package core is the perfect foundation for any web application.
It allows you to connect all and only the components you need in a flexible and efficient way.

A handlers (or middlewares) stack is used to pass data in line, from the first to the last handler.
So you can perform actions downstream, then filter and manipulate the response upstream.

No handlers are bundled in this package.

Usage

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

Run

You can run the server with an optional port:

	$ go run app.go [-port port] [-production]

The application is reachable at "http://localhost:8080/".
*/
package core
