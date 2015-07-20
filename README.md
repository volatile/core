<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/core/logo.png" alt="Volatile Core" title="Volatile Core"><br><br></p>

Volatile Core is the perfect foundation for any web application.  
It allows you to connect all and only the components you need in a flexible and efficient way.

A handlers (or *middlewares*) stack is used to pass data in line, from the first to the last handler.  
So you can perform actions downstream, then filter and manipulate the response upstream.

No handlers are bundled in this package.

* [Getting Started](#getting-started)
	* [1. Install](#1-install)
	* [2. Write](#2-write)
	* [3. Run](#3-run)
* [Official handlers](#official-handlers)
* [Official helpers](#official-helpers)
* [Usage](#usage)
	* [Context](#context)
		* [Next](#next)
		* [Pass data](#pass-data)
		* [Response writer binding](#response-writer-binding)
		* [Things to know](#things-to-know)
	* [Custom port](#custom-port)
	* [Environment](#environment)


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
$ go run app.go
```

The application is reachable at `http://localhost:8080/`.

## Official handlers

In order of usability in you app:

- [Log](https://github.com/volatile/log) — Requests logging
- [Compress](https://github.com/volatile/compress) — Responses compressing
- [CORS](https://github.com/volatile/cors) — Cross-Origin Resource Sharing support
- *Others are coming…*

## Official helpers

Helpers are just syntactic sugars to ease repetitive code and improve readability of you app.

- [Route](https://github.com/volatile/route) — Flexible routing helper
- [Response](https://github.com/volatile/response) — Readable response helper
- *Others are coming…*

## Usage

### Context

All handlers are functions that receive a context: `func(*core.Context)`.  
A Volatile context encapsulates the classic `*http.Request` and `http.ResponseWriter`. See the standard [`net/http`](http://golang.org/pkg/net/http/) package for these.

#### Next

To not break the handlers chain and go to the next handler, just use the context's `Next()` method.

```Go
core.Use(func(c *core.Context) {
	c.Next()
})
```

#### Pass data

To transmit data from a handler to another, the `*coreContext` has a `Data` field, which is a `map[string]interface{}`.

```Go
core.Use(func(c *core.Context) {
	c.Data["id"] = 123
})
```

#### Response writer binding

If some of your handlers need to transform the request before sending it, you can't just use the same `ResponseWriter` for each handler.  
To do so, the Core provides a `ResponseWriterBinder` structure that has the same signature as an `http.ResponseWriter`, but that redirects the response upstream, to an `io.Writer` that will write the original `http.ResponseWriter`.

In other words, the `ResponseWriterBinder` has an output (the `ResponseWriter` used before setting the binder) and an input (an overwritten context `ResponseWriter` used by the next handlers).  
You can see a real example in the [Compress](https://github.com/volatile/compress/blob/master/handler.go) package.

If you need to do something (like setting headers, as you cannot do that after) just before writing the response body, set a function on the `BeforeWrite` field.

```Go
core.Use(func(c *core.Context) {
	// 1. We set the output.
	gzw := gzip.NewWriter(c.ResponseWriter)
	defer gzw.Close()

	// 2. We set the binder.
	rwb = ResponseWriterBinder{
		Writer:         gzw, // The binder output.
		ResponseWriter: c.ResponseWriter, // To keep the same signature (but the `Write` method of the `ResponseWriter` has changed internally).
		BeforeWrite:    func(b []byte) {
			// Do something with b before writing the response body.
		},
	}

	// 3. We set the input.
	c.ResponseWriter = rwb
})

core.Use(func(c *core.Context) {
	c.ResponseWriter.Write([]byte("Hello, World!"))
})
```

#### Things to know

- When a handler writes the body of a context response, it brakes the handlers chain and a `c.Next()` call is useless.
- Remember that response headers must be set before the body is written. After that, trying to set a header has no effect.

### Custom port

To let your server listen on a custom port, use the `-port [port]` parameter on start.

### Environments

Some handlers can have different behaviors depending on the environment the server is running.  
By default, the Core suppose the server in launched in a *development* environment.  
When running your application in a *production* environment, use the `-production` parameter on start.

In your code, you have access to the `core.Production` flag to distinguish the environment.
