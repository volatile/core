<p align="center"><br><br><br><img src="https://s3-eu-west-1.amazonaws.com/whitedevops/volatile/volatile-core.png" alt="Volatile Core" title="Volatile Core"><br><br><br></p>

THIS IS A WORK IN PROGRESS SUBJECT TO CHANGES

Volatile Core is the perfect foundation for any web application.  
It allows you to connect all and only the components you need in a flexible and efficient way.

A handlers stack (middlewares) is used to pass data in line, from the first to the last handler.  
So you can perform actions downstream, then filter and manipulate the response upstream.

No handlers are bundled in the Core package.

# Examples

## "Hello, World!"

Because all goes on from there…

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

## Real case

Volatile is an everyday use and production ready micro framework.  
So let's finish with a real situation where you would need **logging**, **compression**, **CORS**, **sessions**.

```Go
package main

import (
	"net/http"

	"github.com/volatile/core"
	"github.com/volatile/compress"
	"github.com/volatile/log"
	"github.com/volatile/router"
	"github.com/volatile/sessions"
)

var maintRouter = router.New()

func main() {
	// TODO
}
```
