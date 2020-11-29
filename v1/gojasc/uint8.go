package gojasc

import "github.com/foxcapades/tally-go/v1/tally"

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Uint8 Serialization                                                ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// SerializeUint8 converts the given uint8 value to a base57 encoded byte slice.
//
// The byte slice returned may be 1-3 bytes in length.
func SerializeUint8(v uint8) []byte {
	if v == 0 {
		return []byte{min}
	}

	sz := SizeUint8(v)
	pos := sz
	out := make([]byte, pos)

	for v > 0 {
		pos--
		out[pos] = v%Base + min
		v /= Base
	}

	// Subtract 1 since the size header should not included itself.
	out[0] = byte(sz-1) + min

	return out
}

// SerializeUint8Into converts the given uint8 value to base57 and writes it to
// the given buffer starting at off.Cur().
//
// The given offset value will be incremented as the buffer is written and after
// this function call will be at the next writable offset position.
//
// WARNING: This method makes no attempt to verify that the given byte buffer is
// actually long enough to hold the serialized value.  The buffer size should be
// at least SizeUint8(v) in length.
func SerializeUint8Into(v uint8, buf []byte, off *tally.UTally) {
	if v == 0 {
		buf[off.Inc()] = min
		return
	}

	sz := SizeUint8(v)
	pos := sz
	cur := off.Add(pos)

	for v > 0 {
		pos--
		buf[pos] = v%Base + min
		v /= Base
	}

	// Subtract 1 since the size header should not included itself.
	buf[cur] = byte(sz-1) + min
}

func DeserializeUint8(v []byte, off *tally.UTally) (uint8, error) {
	a := DeserializeUDigit(v[off.Inc()])

	if a < 0 {
		return 0, NewJASCByteError(0)
	}

	switch a {
	case 0:
		return 0, nil
	case 1:
		return DeserializeUDigit(v[off.Inc()]), nil
	default:
		return DeserializeUDigit(v[off.Inc()])*Base + DeserializeUDigit(v[off.Inc()]), nil
	}
}

// SizeUint8 returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given uint8 value.
//
// This size includes the byte needed for the number size header.
func SizeUint8(v uint8) uint {
	switch true {
	case v == 0:
		return 1
	case v < Base:
		return 2
	default:
		return 3
	}
}
