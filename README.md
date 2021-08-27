# Please use https://github.com/mitchellh/panicwrap as it does not use some nasty memory patching.

# go-panichandler ![](https://github.com/talon-one/go-panichandler/workflows/Test/badge.svg) [![GoDoc](https://godoc.org/github.com/Eun/yaegi-template?status.svg)](https://godoc.org/github.com/Eun/yaegi-template) [![go-report](https://goreportcard.com/badge/github.com/talon-one/go-panichandler)](https://goreportcard.com/report/github.com/talon-one/go-panichandler)

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
	"github.com/talon-one/go-panichandler"
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
