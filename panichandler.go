package panichandler

import (
	"unsafe"

	"sync"

	"github.com/agiledragon/gomonkey/v2"
)

//go:linkname fatalpanic runtime.fatalpanic
func fatalpanic(msgs *_panic)

//nolint:structcheck
//go:linkname _panic runtime._panic
type _panic struct {
	argp      unsafe.Pointer // pointer to arguments of deferred call run during panic; cannot move - known to liblink
	arg       interface{}    // argument to panic
	link      *_panic        // link to earlier panic
	recovered bool           // whether this panic is over
	aborted   bool           // the panic was aborted
}

// Handler will be put on the panic handling stack
//
// the first parameter will contain the variable passed into panic()
//
// what ever you will return will be handed over to the next panic() handler up until the final panic
//
// if you return an instance of IgnorePanic{} the panic will be ignored and execution continues
// note that this will mostly result in more panics
type Handler func(interface{}) interface{}

var handlers []Handler
var mu sync.Mutex

// IgnorePanic can be returned in the Handler to ignore the panic and continue execution
type IgnorePanic struct{}

func init() {
	patch()
}

func patch() {
	var patches *gomonkey.Patches
	patches = gomonkey.ApplyFunc(fatalpanic, func(v *_panic) {
		patches.Reset()
		defer patch()
		mu.Lock()
		defer mu.Unlock()
		for _, handler := range handlers {
			if v.arg = handler(v.arg); v.arg != nil {
				if _, ok := v.arg.(IgnorePanic); ok {
					return
				}
			}
		}
		fatalpanic(v)
	})
}

// OnPanic adds the specified handler function to the panic handler stack
func OnPanic(fn Handler) {
	mu.Lock()
	handlers = append(handlers, fn)
	mu.Unlock()
}
