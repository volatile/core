/*
Package core is the perfect foundation for any web application.
It allows you to connect all and only the components you need in a flexible and efficient way.

A handlers stack (middlewares) is used to pass data in line, from the first to the last handler.
So you can perform actions downstream, then filter and manipulate the response upstream.

No handlers are bundled in this package.

Example

Here is the classic "Hello, World!" example:

	package main

	import (
		"net/http"

		"github.com/volatile/core"
	)

	func main() {
		core.Use(func(c *core.Context) {
			c.Response.Body = []byte("Hello, World!")
		})

		core.Run()
	}

Run

You can run the server with an optional port:

	$ go run app.go [-p port]

The application is reachable at "http://localhost:8080/".
*/
package core
