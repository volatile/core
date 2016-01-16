/*
Package core is the perfect foundation for any web app as it's designed to have the best balance between readability, flexibility and performance.

It provides a pure handlers (or "middlewares") stack so you can perform actions downstream, then filter and manipulate the response upstream.

No handlers or helpers are bundled in the Core: it does one thing and does it well.
You can find official packages below.

For a complete documentation, see the Volatile website (http://volatile.whitedevops.com).
You can also read all the code (~100 LOC) within minutes.

Installation

In the terminal:

	$ go get github.com/volatile/core

Usage

Example of a logger followed by the response writing:

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

By default, your app is reachable at http://localhost:8080.

Flags

These flags are preset:

● -address to set a custom listening address.
The value is saved in Address.

● -production to switch on production environment settings.
Some third-party handlers may have different behaviors depending on the environment.
The value is saved in Production.

It's up to you to call flag.Parse() in your main function if you want to use them.

Panic recovering

Volatile Core recovers your server from any panic, logs the error with stack, and sends a 500 Internal Server Error.
If you want a make a custom response on panic, give it as a function to HandlePanic.

Compatibility

Volatile Core is fully compatible with the net/http.Handler interface. Just use NewHandlersStack:

	package main

	import (
		"fmt"
		"net/http"

		"github.com/volatile/core"
	)

	func main() {
		hs := core.NewHandlersStack()

		hs.Use(func(c *core.Context) {
			fmt.Fprint(c.ResponseWriter, "Hello, World!")
		})

		http.ListenAndServe(":8080", hs)
	}

Official handlers

These handlers are ready to be integrated in any of your app:

● Compress — Clever response compressing — https://github.com/volatile/compress

● CORS — Cross-Origin Resource Sharing support — https://github.com/volatile/cors

● Log — Requests logging — https://github.com/volatile/log

● Secure — Quick security wins — https://github.com/volatile/secure

● Static — Simple assets serving — https://github.com/volatile/static

Official helpers

These helpers provide syntactic sugar to ease repetitive code:

● I18n — Simple internationalization — https://github.com/volatile/i18n

● Response — Readable response helper — https://github.com/volatile/response

● Route — Flexible routing helper — https://github.com/volatile/route
*/
package core
