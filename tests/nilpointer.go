// +build ignore

package main

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/talon-one/go-panichandler"
)

func main() {
	panichandler.OnPanic(func(v interface{}) interface{} {
		json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
			"expected": "runtime error: invalid memory address or nil pointer dereference",
			"actual":   v,
		})
		os.Exit(0)
		return nil
	})

	type Foo struct {
		Bar string
	}
	var v *Foo
	fmt.Println(v.Bar)
}
