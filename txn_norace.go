//go:build !memdbrace
// +build !memdbrace

package memdb

// trackPointer is a noop implementation of a data race detector.
func trackPointer(*MemDB, string, interface{})  {}
func forgetPointer(*MemDB, string, interface{}) {}
