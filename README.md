<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/core/logo.png" alt="Volatile Core" title="Volatile Core"><br><br></p>

Volatile Core is the perfect foundation for any web application.  
It allows you to connect all and only the components you need in a flexible and efficient way.

A handlers (or *middlewares*) stack is used to pass data in line, from the first to the last handler.  
So you can perform actions downstream, then filter and manipulate the response upstream.

No handlers are bundled in this package.

For a complete **documentation**, see the Volatile website : http://volatile.whitedevops.com

## Installation

```Shell
$ go get -u github.com/volatile/core
```

## Usage

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

[![GoDoc](https://godoc.org/github.com/volatile/core?status.svg)](https://godoc.org/github.com/volatile/core)

## Official handlers

These handlers are ready to be integrated in any of your app…

- [Log](https://github.com/volatile/log) — Requests logging
- [Compress](https://github.com/volatile/compress) — Responses compressing
- [CORS](https://github.com/volatile/cors) — Cross-Origin Resource Sharing support
- [Static](https://github.com/volatile/static) — Simple assets serving
- *Others are coming…*

## Official helpers

Helpers provide syntactic sugar to ease repetitive code…

- [Route](https://github.com/volatile/route) — Flexible routing helper
- [Response](https://github.com/volatile/response) — Readable response helper
- *Others are coming…*
