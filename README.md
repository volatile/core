<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/core/logo.png" alt="Volatile Core" title="Volatile Core"><br><br></p>

Volatile Core is the perfect foundation for any web application.  
It allows you to connect all and only the components you need in a flexible and efficient way.

A handlers (or *middlewares*) stack is used to pass data in line, from the first to the last handler.  
So you can perform actions downstream, then filter and manipulate the response upstream.

No handlers are bundled in this package.

## Getting started

### 1. Install

```Shell
$ go get -u github.com/volatile/core
```

### 2. Write

In `app.go`:

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

### 3. Run

```Shell
$ go run app.go [-port port] [-production]
```

The application is reachable at `http://localhost:8080/`.

## Official handlers

In order of usability in you app:

- [Log](https://github.com/volatile/log) — Requests logging
- [Compress](https://github.com/volatile/compress) — Responses compressor
- *Others are coming…*

## Official helpers

Is helper is just syntactic sugar to ease repetitive code and improve readability of you app.

- [Route](https://github.com/volatile/route) — A flexible routing helper
- [Response](https://github.com/volatile/response) — A readable response helper
- *Others are coming…*
