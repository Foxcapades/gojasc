package j57

import "github.com/foxcapades/tally-go/v1/tally"

const (
	bit32Flag uint8 = 128
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Int32 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝

// SerializeInt32 converts the given int32 value to a base57 encoded byte
// slice.
//
// The byte slice returned may be 1-7 bytes in length.
func SerializeInt32(v int32) []byte {
	return SerializeUint32(uint32(v))
}

// AppendInt32 converts the given int32 value to base57 and writes it
// to the given buffer starting at off.Cur().
//
// The given offset value will be incremented as the buffer is written and after
// this function call will be at the next writable offset position.
//
// WARNING: This method makes no attempt to verify that the given byte buffer is
// actually long enough to hold the serialized value.  The buffer size should be
// at least SizeInt32(v) in length.
func AppendInt32(v int32, buf []byte, off *tally.UTally) int {
	return AppendUint32(uint32(v), buf, off)
}

func DeserializeInt32(v []byte, off *tally.UTally) (out int32, err error) {
	// This value was serialized using SerializeInt on a 32bit system.
	// We can ignore the value since we are doing a 32 bit deserialization anyway.
	if v[off.Cur()]&bit32Flag > 0 {
		v[off.Cur()] &= ^bit32Flag
	}

	t, e := DeserializeUint32(v, off)
	return int32(t), e
}

// SizeInt32 returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given int32 value.
//
// This size includes the byte needed for the number size header.
func SizeInt32(v int32) uint {
	return SizeUint32(uint32(v))
}
