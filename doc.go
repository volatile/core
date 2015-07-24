/*
Package core is the perfect foundation for any web app.
It allows you to connect all and only the components you need in a flexible and efficient way.

A handlers (or "middlewares") stack is used to pass data in line, from the first to the last handler.
So you can perform actions downstream, then filter and manipulate the response upstream.

No handlers are bundled in this package.

For a complete documentation, see the Volatile website : http://volatile.whitedevops.com

Usage

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

Your app is reachable at http://localhost:8080/.
If you want to use a custom port, set the "-port" parameter when running your app.

Also, use the "-production" parameter when serving in a production environment.
Some third-party handlers may have different behaviors following the environment.

Official handlers

These handlers are ready to be integrated in any of your app…

Log — Requests logging — https://github.com/volatile/log

Compress — Responses compressing — https://github.com/volatile/compress

CORS — Cross-Origin Resource Sharing support — https://github.com/volatile/cors

Static — Simple assets serving — https://github.com/volatile/static

Others are coming…

Official helpers

Helpers provide syntactic sugar to ease repetitive code…

Route — Flexible routing helper — https://github.com/volatile/route

Response — Readable response helper — https://github.com/volatile/response

Others are coming…
*/
package core
