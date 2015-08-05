package log

import (
	"log"
	"runtime"
)

// Stack logs the package pack with an error err and the stack trace.
func Stack(pack string, err error) {
	const size = 64 << 10
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]
	log.Printf("%s: %v\n%s", pack, err, buf)
}
