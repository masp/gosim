package go127

import (
	"github.com/jellevandenhooff/gosim/gosimruntime"
)

// Since go1.24 sync.Mutex is implemented in internal/sync, which has its own
// copies of the runtime hooks that package sync used to have.

func InternalSync_runtime_SemacquireMutex(addr *uint32, lifo bool, skipframes int) {
	gosimruntime.Semacquire(addr, lifo)
}

func InternalSync_runtime_Semrelease(addr *uint32, handoff bool, skipframes int) {
	gosimruntime.Semrelease(addr)
}

func InternalSync_runtime_canSpin(i int) bool {
	return false
}

func InternalSync_runtime_doSpin() {
	panic("no")
}

func InternalSync_runtime_nanotime() int64 {
	return gosimruntime.Nanotime()
}

func InternalSync_runtime_rand() uint64 {
	return gosimruntime.Fastrand64()
}

func InternalSync_throw(str string) {
	panic(str)
}

func InternalSync_fatal(str string) {
	panic(str)
}

func Sync_runtime_SemacquireWaitGroup(addr *uint32, synctestDurable bool) {
	gosimruntime.Semacquire(addr, false)
}
