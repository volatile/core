package core

import (
	"flag"
	"net/http"
	"strconv"
)

var (
	// Production defines if the server must be using production settings.
	// It can be used by handlers to provide different logic for this environment.
	Production bool
	port       int
	beforeRun  []func()
)

func init() {
	flag.BoolVar(&Production, "production", false, "run the server in production environment")
	flag.IntVar(&port, "port", 8080, "the port to listen on")
	flag.Parse()
}

// BeforeRun adds a function that will be triggered before runnning the server.
func BeforeRun(f func()) {
	beforeRun = append(beforeRun, f)
}

// Run starts the server for listening and serving.
func Run() {
	if handlers == nil {
		panic("core: the handlers stack cannot be empty")
	}

	// Add a last handler to prevent "index out of range" errors if the previous last handler in stack calls Next().
	Use(func(c *Context) {
		http.NotFound(c.ResponseWriter, c.Request)
	})

	for _, f := range beforeRun {
		f()
	}

	panic(http.ListenAndServe(":"+strconv.Itoa(port), handlers))
}
