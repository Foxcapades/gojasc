package j57

import "github.com/foxcapades/tally-go/v1/tally"

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Int8 Serialization                                                ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// SerializeInt8 converts the given int8 value to a base57 encoded byte slice.
//
// The byte slice returned may be 1-3 bytes in length.
func SerializeInt8(v int8) []byte {
	return SerializeUint8(uint8(v))
}

// AppendInt8 converts the given int8 value to base57 and writes it to
// the given buffer starting at off.Cur().
//
// The given offset value will be incremented as the buffer is written and after
// this function call will be at the next writable offset position.
//
// WARNING: This method makes no attempt to verify that the given byte buffer is
// actually long enough to hold the serialized value.  The buffer size should be
// at least SizeInt8(v) in length.
func AppendInt8(v int8, buf []byte, off *tally.UTally) int {
	return AppendUint8(uint8(v), buf, off)
}

func DeserializeInt8(v []byte, off *tally.UTally) (int8, error) {
	tmp, err := DeserializeUint8(v, off)
	return int8(tmp), err
}

// SizeInt8 returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given int8 value.
//
// This size includes the byte needed for the number size header.
func SizeInt8(v int8) uint {
	return SizeUint8(uint8(v))
}
