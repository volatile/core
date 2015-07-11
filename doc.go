// Package core is the perfect foundation for any web application.
// It allows you to connect all and only the components you need in a flexible and efficient way.
//
// A handlers stack (middlewares) is used to pass data in line, from the first to the last handler.
// So you can perform actions downstream, then filter and manipulate the response upstream.
//
// No handlers are bundled in the core package.
//
// Example
//
// In "app.go":
//
//  package main
//
//  import (
//  	"net/http"
//
//  	"github.com/volatile/core"
//  )
//
//  func main() {
//  	core.Use(func(c *core.Context) {
//  		c.Response = []byte("Hello, World!")
//  	})
//
//  	core.Run()
//  }
//
// Run the server
//
//  $ go run app.go [-p port]
//
// The application is reachable at "http://localhost:8080/".
package core
