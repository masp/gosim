//go:build linux

package go127

import (
	"syscall"
	"unsafe"

	"github.com/jellevandenhooff/gosim/internal/simulation"
)

//go:nocheckptr
//go:uintptrescapes
func GolangOrgXSysUnix_RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	return simulation.RawSyscall6(trap, a1, a2, a3, a4, a5, a6)
}

// vgetrandom reports that the vDSO getrandom is unsupported so that callers
// fall back to the (simulated) getrandom system call.
func GolangOrgXSysUnix_vgetrandom(p []byte, flags uint32) (ret int, supported bool) {
	return -1, false
}

func GolangOrgXSysUnix_runtime_getAuxv() []uintptr {
	return nil
}

func GolangOrgXSysUnix_SyscallNoError(trap, a1, a2, a3 uintptr) (r1, r2 uintptr) {
	panic("gosim not implemented")
}

func GolangOrgXSysUnix_RawSyscallNoError(trap, a1, a2, a3 uintptr) (r1, r2 uintptr) {
	panic("gosim not implemented")
}

func GolangOrgXSysUnix_Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	return simulation.RawSyscall6(trap, a1, a2, a3, 0, 0, 0)
}

func GolangOrgXSysUnix_Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	return simulation.RawSyscall6(trap, a1, a2, a3, a4, a5, a6)
}

func GolangOrgXSysUnix_RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	panic("gosim not implemented")
}

func GolangOrgXSysUnix_gettimeofday(timeval unsafe.Pointer) syscall.Errno {
	panic("gosim not implemented")
}
