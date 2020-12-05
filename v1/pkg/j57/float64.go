package j57

import (
	"github.com/foxcapades/tally-go/v1/tally"
	"unsafe"
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Float64 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝

// SerializeFloat64 converts the given float64 value to a base57 encoded byte
// slice.
//
// The byte slice returned may be 1-7 bytes in length.
func SerializeFloat64(v float64) []byte {
	return SerializeUint64(unsafeF64U(v))
}

// AppendFloat64 converts the given float64 value to base57 and writes it
// to the given buffer starting at off.Cur().
//
// The given offset value will be incremented as the buffer is written and after
// this function call will be at the next writable offset position.
//
// WARNING: This method makes no attempt to verify that the given byte buffer is
// actually long enough to hold the serialized value.  The buffer size should be
// at least SizeFloat64(v) in length.
func AppendFloat64(v float64, buf []byte, off *tally.UTally) int {
	return AppendUint64(unsafeF64U(v), buf, off)
}

func DeserializeFloat64(v []byte, off *tally.UTally) (out float64, err error) {
	t, e := DeserializeUint64(v, off)
	return unsafeU64F(t), e
}

// SizeFloat64 returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given float64 value.
//
// This size includes the byte needed for the number size header.
func SizeFloat64(v float64) uint {
	return SizeUint64(unsafeF64U(v))
}

func unsafeU64F(v uint64) float64 {
	return *(*float64)(unsafe.Pointer(&v))
}
func unsafeF64U(v float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&v))
}
