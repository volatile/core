<p align="center"><img src="https://cloud.githubusercontent.com/assets/9503891/8646803/1fd8e456-2953-11e5-9768-8383290d486d.png" alt="Volatile Core" title="Volatile Core"><br><br></p>

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

	"github.com/volatile/core"
)

func main() {
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

- [Compress](https://github.com/volatile/compress) — Responses compressor
- *Others are coming…*
