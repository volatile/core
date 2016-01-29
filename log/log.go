// Package log implements a simple logging package for handlers usage.
package log

import (
	"log"
	"runtime"
	"strings"
)

// Stack logs the error err with the stack trace.
func Stack(err interface{}) {
	stack := make([]byte, 64<<10)
	stack = stack[:runtime.Stack(stack, false)]

	log.Printf("%v\n%s", err, stack)
}

// StackWithCaller logs the error err with the caller package name and the stack trace.
func StackWithCaller(err interface{}) {
	stack := make([]byte, 64<<10)
	stack = stack[:runtime.Stack(stack, false)]

	if pack, ok := callerPackage(); ok {
		log.Printf("%s: %v\n%s", pack, err, stack)
	} else {
		log.Printf("%v\n%s", err, stack)
	}
}

func callerPackage() (pack string, ok bool) {
	var pc uintptr
	if pc, _, _, ok = runtime.Caller(2); !ok {
		return
	}
	path := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	pack = strings.Split(path[len(path)-1], ".")[0]
	return
}
