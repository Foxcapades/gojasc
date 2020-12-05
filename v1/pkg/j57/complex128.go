package j57

import (
	"github.com/foxcapades/tally-go/v1/tally"
	"unsafe"
)

func SerializeComplex128(value complex128) []byte {
	buf := [24]byte{}
	off := tally.UTally(0)
	a, b := unsafeC128ToU128(&value)

	AppendUint64(a, buf[:], &off)
	AppendUint64(b, buf[:], &off)

	return buf[:off]
}

func AppendComplex128(value complex128, buffer []byte, offset *tally.UTally) (wrote int) {
	a, b := unsafeC128ToU128(&value)

	wrote += AppendUint64(a, buffer[:], offset)
	return wrote + AppendUint64(b, buffer[:], offset)
}

func DeserializeComplex128(buffer []byte, offset *tally.UTally) (complex128, error) {
	if tmpA, err := DeserializeUint64(buffer, offset); err != nil {
		return 0, err
	} else if tmpB, err := DeserializeUint64(buffer, offset); err != nil {
		return 0, err
	} else {
		return unsafeU128ToC128(tmpA, tmpB), nil
	}
}

func SizeComplex128(v complex128) uint {
	a, b := unsafeC128ToU128(&v)
	return SizeUint64(a) + SizeUint64(b)
}

// unsafeC128ToU128 raw casts the block of memory used by the given complex128
// value as a pair of uint64 values.
func unsafeC128ToU128(v *complex128) (a, b uint64) {
	tmp := (*[2]uint64)(unsafe.Pointer(v))
	return tmp[0], tmp[1]
}

// unsafeU128ToC128 raw casts the block of memory used by the given pair of
// uint64 values as a complex128 value.
func unsafeU128ToC128(a, b uint64) complex128 {
	tmp := [2]uint64{a, b}
	return *(*complex128)(unsafe.Pointer(&tmp))
}
