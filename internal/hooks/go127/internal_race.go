package go127

import (
	"unsafe"

	"github.com/jellevandenhooff/gosim/internal/race"
)

// These hooks are only used in -race builds (see hooksRaceGo127), where they
// forward to the runtime's race detector like internal/race did before
// go1.24.

func InternalRace_Acquire(addr unsafe.Pointer) {
	race.Acquire(addr)
}

func InternalRace_Release(addr unsafe.Pointer) {
	race.Release(addr)
}

func InternalRace_ReleaseMerge(addr unsafe.Pointer) {
	race.ReleaseMerge(addr)
}

func InternalRace_Disable() {
	race.Disable()
}

func InternalRace_Enable() {
	race.Enable()
}

func InternalRace_Read(addr unsafe.Pointer) {
	race.Read(addr)
}

func InternalRace_ReadPC(addr unsafe.Pointer, callerpc, pc uintptr) {
	race.Read(addr)
}

// t is an *internal/abi.Type
func InternalRace_ReadObjectPC(t unsafe.Pointer, addr unsafe.Pointer, callerpc, pc uintptr) {
	race.Read(addr)
}

func InternalRace_Write(addr unsafe.Pointer) {
	race.Write(addr)
}

func InternalRace_WritePC(addr unsafe.Pointer, callerpc, pc uintptr) {
	race.Write(addr)
}

// t is an *internal/abi.Type
func InternalRace_WriteObjectPC(t unsafe.Pointer, addr unsafe.Pointer, callerpc, pc uintptr) {
	race.Write(addr)
}

func InternalRace_ReadRange(addr unsafe.Pointer, len int) {
	race.ReadRange(addr, len)
}

func InternalRace_WriteRange(addr unsafe.Pointer, len int) {
	race.WriteRange(addr, len)
}

func InternalRace_Errors() int {
	return race.Errors()
}
