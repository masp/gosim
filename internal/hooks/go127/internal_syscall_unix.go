//go:build linux

package go127

import "syscall"

func InternalSyscallUnix_fcntl(fd int32, cmd int32, args int32) (int32, int32) {
	return 0, int32(syscall.ENOSYS)
}

func InternalSyscallUnix_GetRandom(p []byte, flags uintptr) (n int, err error) {
	return GolangOrgXSysUnix_Getrandom(p, int(flags))
}

// vgetrandom reports that the vDSO getrandom is unsupported so that callers
// fall back to the (simulated) getrandom system call.
func InternalSyscallUnix_vgetrandom(p []byte, flags uint32) (ret int, supported bool) {
	return -1, false
}

func InternalSyscallUnix_Unlinkat(dirfd int, path string, flags int) error {
	return Syscall_unlinkat(dirfd, path, flags)
}

func InternalSyscallUnix_Openat(dirfd int, path string, flags int, perm uint32) (int, error) {
	return Syscall_openat(dirfd, path, flags, perm)
}
