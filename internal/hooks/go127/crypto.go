package go127

import (
	"unsafe"

	"github.com/jellevandenhooff/gosim/gosimruntime"
)

func CryptoFips140_isBypassed() bool {
	return false
}

func CryptoFips140_setBypass() {
}

func CryptoFips140_unsetBypass() {
}

func CryptoInternalFips140_fatal(str string) {
	panic(str)
}

func CryptoInternalFips140_getIndicator() uint8 {
	return 0
}

func CryptoInternalFips140_setIndicator(uint8) {
}

func CryptoInternalFips140depsTime_monoTime() int64 {
	return gosimruntime.Nanotime()
}

func CryptoInternalSysrand_fatal(str string) {
	panic(str)
}

func CryptoRand_fatal(str string) {
	panic(str)
}

func CryptoSubtle_setDITEnabled() bool {
	return false
}

func CryptoSubtle_setDITDisabled() {
}

// CryptoInternalConstanttime_boolToUint8 replaces a function that the compiler
// implements as an intrinsic for crypto/internal/constanttime only. A bool is
// stored as a single byte that is either 0 or 1, so reading that byte converts
// without branching.
func CryptoInternalConstanttime_boolToUint8(b bool) uint8 {
	return *(*uint8)(unsafe.Pointer(&b))
}
