package j57

import (
	"github.com/foxcapades/tally-go/v1/tally"
	"unsafe"
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Float32 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝

// SerializeFloat32 converts the given float32 value to a base57 encoded byte
// slice.
//
// The byte slice returned may be 1-7 bytes in length.
func SerializeFloat32(v float32) []byte {
	return SerializeUint32(unsafeF32U(v))
}

// AppendFloat32 converts the given float32 value to base57 and writes it
// to the given buffer starting at off.Cur().
//
// The given offset value will be incremented as the buffer is written and after
// this function call will be at the next writable offset position.
//
// WARNING: This method makes no attempt to verify that the given byte buffer is
// actually long enough to hold the serialized value.  The buffer size should be
// at least SizeFloat32(v) in length.
func AppendFloat32(v float32, buf []byte, off *tally.UTally) int {
	return AppendUint32(unsafeF32U(v), buf, off)
}

func DeserializeFloat32(v []byte, off *tally.UTally) (out float32, err error) {
	t, e := DeserializeUint32(v, off)
	return unsafeU32F(t), e
}

// SizeFloat32 returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given float32 value.
//
// This size includes the byte needed for the number size header.
func SizeFloat32(v float32) uint {
	return SizeUint32(unsafeF32U(v))
}

func unsafeU32F(v uint32) float32 {
	return *(*float32)(unsafe.Pointer(&v))
}
func unsafeF32U(v float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&v))
}
