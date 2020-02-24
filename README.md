# go-panichandler

Handle panics globaly.

Golang lets you not handle panics in subroutines in an easy way.

This package uses dirty memory patching techniques to supply this feature.

Use with care!

## Usecases
* Send a panic to an error tracking service
* Gracefully close the application
* ... 

## Example
```go
package main

import (
	"os"
	"encoding/json"
	"github.com/github.com/talon-one/go-panichandler"
)

func main() {
	panichandler.OnPanic(func(v interface{}) interface{} {
		fmt.Printf("Catched a panic: %v\n", v)
		return v
	})

	ch := make(chan struct{})
	go func() {
		panic("Hello World")
		ch <- struct{}{}
	}()
	<-ch
}

```