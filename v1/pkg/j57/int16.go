package j57

import "github.com/foxcapades/tally-go/v1/tally"

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Int16 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// SerializeInt16 converts the given int16 value to a base57 encoded byte
// slice.
//
// The byte slice returned may be 1-4 bytes in length.
func SerializeInt16(v int16) []byte {
	return SerializeUint16(uint16(v))
}

// AppendInt16 converts the given int16 value to base57 and writes it
// to the given buffer starting at off.Cur().
//
// The given offset value will be incremented as the buffer is written and after
// this function call will be at the next writable offset position.
//
// WARNING: This method makes no attempt to verify that the given byte buffer is
// actually long enough to hold the serialized value.  The buffer size should be
// at least SizeInt16(v) in length.
func AppendInt16(v int16, buf []byte, off *tally.UTally) int {
	return AppendUint16(uint16(v), buf, off)
}

func DeserializeInt16(v []byte, off *tally.UTally) (out int16, err error) {
	t, e := DeserializeUint16(v, off)
	return int16(t), e
}

// SizeInt16 returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given uint8 value.
//
// This size includes the byte needed for the number size header.
func SizeInt16(v int16) uint {
	return SizeUint16(uint16(v))
}
