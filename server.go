package core

import (
	"flag"
	"net/http"
)

var (
	// Production allows handlers know whether the server is running in a production environment.
	Production bool

	// Address is the TCP network address on which the server is listening and serving.
	Address string

	// beforeRun stores a set of functions that are triggered just before running the server.
	beforeRun []func()
)

// BeforeRun adds a function that will be triggered just before running the server.
func BeforeRun(f func()) {
	beforeRun = append(beforeRun, f)
}

// Run starts the server for listening and serving.
func Run() {
	// Add a last handler to prevent "index out of range" errors if the previous last handler calls Next.
	Use(func(c *Context) {
		http.NotFound(c.ResponseWriter, c.Request)
	})

	flag.BoolVar(&Production, "production", false, "run the server in production environment")
	flag.StringVar(&Address, "address", ":8080", "the address to listen and serving on")
	flag.Parse()

	for _, f := range beforeRun {
		f()
	}

	panic(http.ListenAndServe(Address, handlers))
}
