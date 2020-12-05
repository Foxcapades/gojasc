package j57

import (
	"github.com/foxcapades/tally-go/v1/tally"
	"unsafe"
)

func SerializeComplex64(v complex64) []byte {
	return SerializeUint64(unsafeC64ToU64(&v))
}

func AppendComplex64(value complex64, buffer []byte, offset *tally.UTally) int {
	return AppendUint64(unsafeC64ToU64(&value), buffer, offset)
}

func DeserializeComplex64(buffer []byte, offset *tally.UTally) (complex64, error) {
	if tmp, err := DeserializeUint64(buffer, offset); err != nil {
		return 0, err
	} else {
		return unsafeU64ToC64(&tmp), nil
	}
}

func SizeComplex64(v complex64) uint {
	return SizeUint64(unsafeC64ToU64(&v))
}

// unsafeC64ToU64 raw casts the block of memory used by the given complex64
// value as a uint64 value.
func unsafeC64ToU64(v *complex64) uint64 {
	return *(*uint64)(unsafe.Pointer(v))
}

// unsafeU64ToC64 raw casts the block of memory used by the given uint64 value
// as a complex64 value.
func unsafeU64ToC64(v *uint64) complex64 {
	return *(*complex64)(unsafe.Pointer(v))
}
