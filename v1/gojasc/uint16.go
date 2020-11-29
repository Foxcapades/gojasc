package gojasc

import "github.com/foxcapades/tally-go/v1/tally"

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Uint16 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// SerializeUint16 converts the given uint16 value to a base57 encoded byte
// slice.
//
// The byte slice returned may be 1-4 bytes in length.
func SerializeUint16(v uint16) []byte {
	if v == 0 {
		return []byte{min}
	}

	sz := SizeUint16(v)
	pos := sz
	out := make([]byte, pos)

	for v > 0 {
		pos--
		out[pos] = byte(v%Base) + min
		v /= Base
	}

	out[0] = uint8(sz-1) + min

	return out
}

// SerializeUint16Into converts the given uint16 value to base57 and writes it
// to the given buffer starting at off.Cur().
//
// The given offset value will be incremented as the buffer is written and after
// this function call will be at the next writable offset position.
//
// WARNING: This method makes no attempt to verify that the given byte buffer is
// actually long enough to hold the serialized value.  The buffer size should be
// at least SizeUint16(v) in length.
func SerializeUint16Into(v uint16, buf []byte, off *tally.UTally) {
	if v == 0 {
		buf[off.Inc()] = min
		return
	}

	sz := SizeUint16(v)
	pos := sz + off.Cur()
	cur := off.Add(sz)

	for v > 0 {
		pos--
		buf[pos] = byte(v%Base) + min
		v /= Base
	}

	buf[cur] = byte(sz-1) + min
}

func DeserializeUint16(v []byte, off *tally.UTally) (out uint16, err error) {
	a := DeserializeUDigit(v[off.Inc()])

	if a < 0 {
		return 0, NewJASCByteError(0)
	}

	switch a {
	case 0:
		return 0, nil
	case 1:
		return uint16(DeserializeUDigit(v[off.Inc()])), nil
	case 2:
		return uint16(DeserializeUDigit(v[off.Inc()]))*Base +
				uint16(DeserializeUDigit(v[off.Inc()])),
			nil
	default:
		return uint16(DeserializeUDigit(v[off.Inc()]))*powU16(Base, 2) +
				uint16(DeserializeUDigit(v[off.Inc()]))*Base +
				uint16(DeserializeUDigit(v[off.Inc()])),
			nil
	}
}

// SizeUint16 returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given uint8 value.
//
// This size includes the byte needed for the number size header.
func SizeUint16(v uint16) uint {
	switch true {
	case v >= 3_249:
		return 4
	case v >= 57:
		return 3
	case v > 0:
		return 2
	default:
		return 1
	}
}

func powU16(a, b uint8) (out uint16) {
	if b == 0 {
		return 1
	}

	out = 1
	for i := uint8(0); i < b; i++ {
		out *= uint16(a)
	}

	return
}
