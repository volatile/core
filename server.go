package core

import (
	"flag"
	"net/http"
	"strconv"
)

var port = flag.Int("port", 8080, "port to listen on")

// Run starts the server for listening and serving.
func Run() {
	if handlers == nil {
		panic("core: the handlers stack cannot be empty")
	}

	flag.Parse()

	panic(http.ListenAndServe(":"+strconv.Itoa(*port), handlers))
}
