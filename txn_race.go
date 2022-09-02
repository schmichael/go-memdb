//go:build memdbrace
// +build memdbrace

package memdb

import (
	"fmt"
	"runtime"
	"sync"
)

var pointers map[string]string

var pointerMu *sync.Mutex

func init() {
	pointers = make(map[string]string)
	pointerMu = &sync.Mutex{}
}

func trackPointer(db *MemDB, table string, p interface{}) {
	k := fmt.Sprintf("%p-%s-%p", db, table, p)

	pointerMu.Lock()
	defer pointerMu.Unlock()
	if stack, ok := pointers[k]; ok {
		// Pointer already exists in table!
		panic(fmt.Sprintf("duplicate pointer: %s\n", stack))
	}

	// Add to map
	stack := make([]byte, 500)
	pointers[k] = string(stack[:runtime.Stack(stack, false)])
}

func forgetPointer(db *MemDB, table string, p interface{}) {
	k := fmt.Sprintf("%p-%s-%p", db, table, p)

	pointerMu.Lock()
	defer pointerMu.Unlock()

	if _, ok := pointers[k]; !ok {
		panic(fmt.Sprintf("pointer not in table %s: %s\n", table, k))
	}
	delete(pointers, k)
}
