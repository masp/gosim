// Package atomic replaces the standard library's internal/runtime/atomic
// package in translated code.
//
// Outside of the runtime, the standard library only uses a few functions from
// internal/runtime/atomic (for example in sync.Pool). This package implements
// those functions on top of gosimruntime.
package atomic

import "github.com/jellevandenhooff/gosim/gosimruntime"

//go:norace
func LoadAcquintptr(ptr *uintptr) uintptr {
	if gosimruntime.AtomicYield {
		gosimruntime.Yield()
	}
	return *ptr
}

//go:norace
func StoreReluintptr(ptr *uintptr, val uintptr) {
	if gosimruntime.AtomicYield {
		gosimruntime.Yield()
	}
	*ptr = val
}
