<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/core/logo.png" alt="Volatile Core" title="Volatile Core"><br><br></p>


Volatile Core is the perfect foundation for any web app as it's designed to have the best balance between **readability**, **flexibility** and **performance**.  

It provides a pure handler (or *middleware*) stack so you can perform actions downstream, then filter and manipulate the response upstream.

No handlers or helpers are bundled in the Core: it does one thing and does it well.  
You can find [official packages](#official-handlers) below.

For a complete **documentation**, see [the Volatile website](http://volatile.whitedevops.com).  
You can also read all the code (~100 LOC) within minutes.

## Installation

```Shell
$ go get github.com/volatile/core
```

## Usage [![GoDoc](https://godoc.org/github.com/volatile/core?status.svg)](https://godoc.org/github.com/volatile/core)

```Go
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
```

By default, your app is reachable at [localhost:8080](http://localhost:8080).

- Use the `-address` parameter to set a custom listening address.  
  The value is saved in `core.Address`.

- Use the `-production` parameter when serving in a production environment.  
  Some third-party handlers may have different behaviors depending on the environment.  
  The value is saved in `core.Production`.

## Official handlers

These handlers are ready to be integrated in any of your app…

- [Compress](https://github.com/volatile/compress) — Clever response compressing
- [CORS](https://github.com/volatile/cors) — Cross-Origin Resource Sharing support
- [Log](https://github.com/volatile/log) — Requests logging
- [Static](https://github.com/volatile/static) — Simple assets serving
- *Others come…*

## Official helpers

Helpers provide syntactic sugar to ease repetitive code…

- [Response](https://github.com/volatile/response) — Readable response helper
- [Route](https://github.com/volatile/route) — Flexible routing helper
- *Others come…*
