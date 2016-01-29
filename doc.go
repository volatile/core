/*
Package core provides a pure handlers (or middlewares) stack so you can perform actions downstream, then filter and manipulate the response upstream.

The handlers stack

A handler is a function that receives a context.
It can be registered with Use and has the possibility to break the sream or to continue with the next handler in the stack.

Example of a logger followed by the response writing:

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

	// Run server
	core.Run()

When using Run, your app is reachable at http://localhost:8080 by default.

If you need more flexibility, you can make a new handlers stack, which is fully compatible with the net/http.Handler interface:

	hs := core.NewHandlersStack()

	hs.Use(func(c *core.Context) {
		fmt.Fprint(c.ResponseWriter, "Hello, World!")
	})

	http.ListenAndServe(":8080", hs)

Flags

These flags are preset:

	-address
		The address to listen and serving on.
		Value is saved in Address.
	-production
		Run the server in production environment.
		Some third-party handlers may have different behaviors depending on the environment.
		Value is saved in Production.

It's up to you to call
	flag.Parse()
in your main function if you want to use them.

Panic recovering

When using Run, your server always recovers from panics, logs the error with stack, and sends a 500 Internal Server Error.
If you want to use a custom handler on panic, give one to HandlePanic.

Handlers and helpers

No handlers or helpers are bundled in the core: it does one thing and does it well.
That's why you need to import all and only the handlers or helpers you need:

	compress
		Clever response compressing
		https://godoc.org/github.com/volatile/compress
	cors
		Cross-Origin Resource Sharing support
		https://godoc.org/github.com/volatile/cors
	i18n
		Simple internationalization
		https://godoc.org/github.com/volatile/i18n
	log
		Requests logging
		https://godoc.org/github.com/volatile/log
	response
		Readable response helper
		https://godoc.org/github.com/volatile/response
	route
		Flexible routing helper
		https://godoc.org/github.com/volatile/route
	secure
		Quick security wins
		https://godoc.org/github.com/volatile/secure
	static
		Simple assets serving
		https://godoc.org/github.com/volatile/static
*/
package core
