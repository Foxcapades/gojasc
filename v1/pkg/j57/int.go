package j57

import (
	"github.com/foxcapades/tally-go/v1/tally"
	"math/bits"
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Int Serialization                                                  ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// SerializeUint converts the given uint value to a base57 encoded byte slice.
//
// The byte slice returned may be 1-3 bytes in length.
func SerializeInt(v int) []byte {
	if bits.UintSize == 32 {
		// for 32 bit mode, mark the serialized value as being 32bit so the output
		// value will remain the same when parsed on both 32 and 64 bit machines.
		tmp := SerializeUint32(uint32(v))
		tmp[0] |= 127
	}

	return SerializeUint64(uint64(v))
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
func AppendInt(v int, buf []byte, off *tally.UTally) (wrote int) {
	if bits.UintSize == 32 {
		// for 32 bit mode, mark the serialized value as being 32bit so the output
		// value will remain the same when parsed on both 32 and 64 bit machines.
		init := off.Cur()
		wrote = AppendUint32(uint32(v), buf, off)
		buf[init] |= 127

		return
	}

	return AppendUint64(uint64(v), buf, off)
}

func DeserializeInt(v []byte, off *tally.UTally) (int, error) {
	// This value was serialized using SerializeInt on a 32bit system.
	if v[off.Cur()]&bit32Flag > 0 {
		v[off.Cur()] &= ^bit32Flag
		tmp, err := DeserializeUint32(v, off)
		return int(tmp), err
	}

	tmp, err := DeserializeUint64(v, off)
	return int(tmp), err
}

// SizeUint returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given uint value.
//
// This size includes the byte needed for the number size header.
func SizeInt(v int) uint {
	if bits.UintSize == 32 {
		return SizeUint32(uint32(v))
	} else {
		return SizeUint64(uint64(v))
	}
}
