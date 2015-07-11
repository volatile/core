<p align="center"><img src="https://s3-eu-west-1.amazonaws.com/whitedevops/volatile/volatile-core.png" alt="Volatile Core" title="Volatile Core"><br><br></p>



THIS IS A WORK IN PROGRESS SUBJECT TO CHANGES.

Volatile Core is the perfect foundation for any web application.  
It allows you to connect all and only the components you need in a flexible and efficient way.

A handlers stack (middlewares) is used to pass data in line, from the first to the last handler.  
So you can perform actions downstream, then filter and manipulate the response upstream.

No handlers are bundled in the `core` package.

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
	"net/http"

	"github.com/volatile/core"
)

func main() {
	core.Use(func(c *core.Context) {
		c.Response = []byte("Hello, World!")
	})

	core.Run()
}
```

[![GoDoc](https://godoc.org/github.com/volatile/core?status.svg)](https://godoc.org/github.com/volatile/core)

### 3. Run

```Shell
$ go run app.go [-p port]
```

The application is reachable at `http://localhost:8080/`.
