// Package maps replaces the standard library's internal/runtime/maps package
// in translated code.
//
// Outside of the runtime, the standard library only uses a constant from
// internal/runtime/maps (in hash/maphash).
package maps

import "unsafe"

// Use64BitHash mirrors internal/runtime/maps.Use64BitHash.
const Use64BitHash = unsafe.Sizeof(uintptr(0)) == 8
