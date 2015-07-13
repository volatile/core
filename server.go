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
)

// Run starts the server for listening and serving.
func Run() {
	if handlers == nil {
		panic("core: the handlers stack cannot be empty")
	}

	flag.BoolVar(&Production, "production", false, "run the server in production environment")
	flag.IntVar(&port, "port", 8080, "the port to listen on")
	flag.Parse()

	panic(http.ListenAndServe(":"+strconv.Itoa(port), handlers))
}
