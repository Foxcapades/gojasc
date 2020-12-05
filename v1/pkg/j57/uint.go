package j57

import (
	"github.com/foxcapades/tally-go/v1/tally"
	"math/bits"
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Uint Serialization                                                ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// SerializeUint converts the given uint value to a base57 encoded byte slice.
//
// The byte slice returned may be 1-3 bytes in length.
func SerializeUint(v uint) []byte {
	if bits.UintSize == 32 {
		return SerializeUint32(uint32(v))
	} else {
		return SerializeUint64(uint64(v))
	}
}

// AppendUint converts the given uint value to base57 and writes it to
// the given buffer starting at off.Cur().
//
// The given offset value will be incremented as the buffer is written and after
// this function call will be at the next writable offset position.
//
// WARNING: This method makes no attempt to verify that the given byte buffer is
// actually long enough to hold the serialized value.  The buffer size should be
// at least SizeUint(v) in length.
func AppendUint(v uint, buf []byte, off *tally.UTally) (wrote int) {
	if bits.UintSize == 32 {
		return AppendUint32(uint32(v), buf, off)
	} else {
		return AppendUint64(uint64(v), buf, off)
	}
}

func DeserializeUint(v []byte, off *tally.UTally) (uint, error) {
	if bits.UintSize == 32 {
		tmp, err := DeserializeUint32(v, off)
		return uint(tmp), err
	} else {
		tmp, err := DeserializeUint64(v, off)
		return uint(tmp), err
	}
}

// SizeUint returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given uint value.
//
// This size includes the byte needed for the number size header.
func SizeUint(v uint) uint {
	if bits.UintSize == 32 {
		return SizeUint32(uint32(v))
	} else {
		return SizeUint64(uint64(v))
	}
}
