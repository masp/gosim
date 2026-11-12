package go127

import "unsafe"

// weakHandles maps objects to their weak handle. weak.Make must return equal
// pointers for the same object, so handles are shared.
//
// Handles hold strong references: objects referenced by a weak pointer are
// never reclaimed inside of the simulation. This keeps weak pointers
// deterministic, since their behavior does not depend on when the garbage
// collector runs.
var weakHandles map[unsafe.Pointer]*unsafe.Pointer

func Weak_runtime_registerWeakPointer(p unsafe.Pointer) unsafe.Pointer {
	if weakHandles == nil {
		weakHandles = make(map[unsafe.Pointer]*unsafe.Pointer)
	}
	if handle, ok := weakHandles[p]; ok {
		return unsafe.Pointer(handle)
	}
	handle := new(unsafe.Pointer)
	*handle = p
	weakHandles[p] = handle
	return unsafe.Pointer(handle)
}

func Weak_runtime_makeStrongFromWeak(u unsafe.Pointer) unsafe.Pointer {
	return *(*unsafe.Pointer)(u)
}
