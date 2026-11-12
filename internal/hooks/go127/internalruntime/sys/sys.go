// Package sys replaces the standard library's internal/runtime/sys package in
// translated code.
//
// Outside of the runtime, the standard library only uses DITSupported from
// internal/runtime/sys (in crypto/subtle).
package sys

// DITSupported reports whether data independent timing can be enabled. The
// simulation does not support toggling it.
const DITSupported = false

func EnableDIT() bool  { return false }
func DITEnabled() bool { return false }
func DisableDIT()      {}
