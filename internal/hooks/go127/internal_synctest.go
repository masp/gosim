package go127

import "unsafe"

// testing/synctest bubbles are not supported inside of the simulation. The
// hooks below behave as if no goroutine is ever inside of a bubble, which is
// what the standard library expects when synctest is not in use.

func InternalSynctest_Run(f func()) {
	panic("gosim not implemented: testing/synctest is not supported inside of a simulation")
}

func InternalSynctest_Wait() {
	panic("gosim not implemented: testing/synctest is not supported inside of a simulation")
}

func InternalSynctest_IsInBubble() bool {
	return false
}

func InternalSynctest_associate(p unsafe.Pointer) int {
	return 0 // synctest.Unbubbled
}

func InternalSynctest_disassociate(b unsafe.Pointer) {
}

func InternalSynctest_isAssociated(p unsafe.Pointer) bool {
	return false
}

func InternalSynctest_acquire() any {
	return nil
}

func InternalSynctest_release(b any) {
}

func InternalSynctest_inBubble(b any, f func()) {
	f()
}
