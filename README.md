<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/core/logo.png" alt="Volatile Core" title="Volatile Core"><br><br></p>

Volatile Core is the perfect foundation for any web app as it's designed to have the best balance between **readability**, **flexibility** and **performance**.  

It provides a pure handler (or *middleware*) stack so you can perform actions downstream, then filter and manipulate the response upstream.

No handlers or helpers are bundled in the Core: it does one thing and does it well.  
You can find [official packages](#official-handlers) below.

For a complete **documentation**, see the [Volatile website](http://volatile.whitedevops.com).  
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

---

Note that these parameters are preset:

- `-address` to set a custom listening address.  
  The value is saved in [`Address`](https://godoc.org/github.com/volatile/core#Address).

- `-production` to switch on production environment settings.  
  Some third-party handlers may have different behaviors depending on the environment.  
  The value is saved in [`Production`](https://godoc.org/github.com/volatile/core#Production).

It's up to you to call [`flag.Parse()`](https://golang.org/pkg/flag/#Parse) in your main function if you want to use them.

---

Also, this package is fully compatible with the [`net/http.Handler`](https://golang.org/pkg/net/http/#Handler) interface:

```Go
package main

import (
	"fmt"
	"net/http"

	"github.com/volatile/core"
)

func main() {
	hs := core.NewHandlersStack()

	hs.Use(func(c *core.Context) {
		fmt.Fprint(c.ResponseWriter, "Hello, World!")
	})

	http.ListenAndServe(":8080", hs)
}
```

## Official handlers

These handlers are ready to be integrated in any of your app…

- [Compress](https://github.com/volatile/compress) — Clever response compressing
- [CORS](https://github.com/volatile/cors) — Cross-Origin Resource Sharing support
- [Log](https://github.com/volatile/log) — Requests logging
- [Secure](https://github.com/volatile/secure) — Quick security wins
- [Static](https://github.com/volatile/static) — Simple assets serving
- *Others come…*

## Official helpers

Helpers provide syntactic sugar to ease repetitive code…

- [I18n](https://github.com/volatile/i18n) — Simple internationalization
- [Response](https://github.com/volatile/response) — Readable response helper
- [Route](https://github.com/volatile/route) — Flexible routing helper
- *Others come…*
