// +build ignore

package main

import (
	"os"

	"encoding/json"

	"reflect"

	"github.com/talon-one/go-panichandler"
)

func main() {
	panichandler.OnPanic(func(v interface{}) interface{} {
		json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
			"expected": "reflect: call of reflect.Value.Interface on zero Value",
			"actual":   v,
		})
		os.Exit(0)
		return nil
	})

	var a interface{}
	v := reflect.ValueOf(a)
	v.Interface()
}
