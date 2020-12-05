package j57

import "github.com/foxcapades/tally-go/v1/tally"

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Int64 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝

// SerializeInt64 converts the given int64 value to a base57 encoded byte
// slice.
//
// The byte slice returned may be 1-12 bytes in length.
func SerializeInt64(v int64) []byte {
	return SerializeUint64(uint64(v))
}

// AppendInt64 converts the given int64 value to base57 and writes it
// to the given buffer starting at off.Cur().
//
// The given offset value will be incremented as the buffer is written and after
// this function call will be at the next writable offset position.
//
// WARNING: This method makes no attempt to verify that the given byte buffer is
// actually long enough to hold the serialized value.  The buffer size should be
// at least SizeInt64(v) in length.
func AppendInt64(v int64, buf []byte, off *tally.UTally) int {
	return AppendUint64(uint64(v), buf, off)
}

func DeserializeInt64(v []byte, off *tally.UTally) (out int64, err error) {
	// This value was serialized using SerializeInt on a 32bit system.
	// We can ignore the value since we are doing a 32 bit deserialization anyway.
	if v[off.Cur()]&bit32Flag > 0 {
		v[off.Cur()] &= ^bit32Flag
	}

	t, e := DeserializeUint64(v, off)
	return int64(t), e
}

// SizeInt64 returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given int64 value.
//
// This size includes the byte needed for the number size header.
func SizeInt64(v int64) uint {
	return SizeUint64(uint64(v))
}
