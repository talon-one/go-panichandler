// +build ignore

package main

import (
	"os"

	"encoding/json"

	"github.com/talon-one/go-panichandler"
)

func main() {
	panichandler.OnPanic(func(v interface{}) interface{} {
		json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
			"expected": "Hello World",
			"actual":   v,
		})
		os.Exit(0)
		return nil
	})

	panic("Hello World")
}
