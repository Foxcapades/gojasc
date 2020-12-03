package jasc

import "github.com/foxcapades/tally-go/v1/tally"

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Uint32 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝

// SerializeUint32 converts the given uint32 value to a base57 encoded byte
// slice.
//
// The byte slice returned may be 1-7 bytes in length.
func SerializeUint32(v uint32) []byte {
	if v == 0 {
		return []byte{min}
	}

	sz := SizeUint32(v)
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

// SerializeUint32Into converts the given uint32 value to base57 and writes it
// to the given buffer starting at off.Cur().
//
// The given offset value will be incremented as the buffer is written and after
// this function call will be at the next writable offset position.
//
// WARNING: This method makes no attempt to verify that the given byte buffer is
// actually long enough to hold the serialized value.  The buffer size should be
// at least SizeUint32(v) in length.
func SerializeUint32Into(v uint32, buf []byte, off *tally.UTally) {
	if v == 0 {
		buf[off.Inc()] = min
		return
	}

	sz := SizeUint32(v)
	pos := sz + off.Cur()
	cur := off.Add(sz)

	for v > 0 {
		pos--
		buf[pos] = byte(v%Base) + min
		v /= Base
	}

	buf[cur] = byte(sz-1) + min
}

func DeserializeUint32(v []byte, off *tally.UTally) (out uint32, err error) {
	sz := int(DeserializeUDigit(v[off.Inc()]))

	if sz < 0 {
		return 0, NewJASCByteError(0)
	} else if sz == 0 {
		return 0, nil
	}

	i := sz + int(off.Add(uint(sz)))

	for j := 0; j < sz; j++ {
		i--
		out += uint32(DeserializeUDigit(v[i])) * powU32(j)
	}

	return out, nil
}

// SizeUint32 returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given uint32 value.
//
// This size includes the byte needed for the number size header.
func SizeUint32(v uint32) uint {
	// 4,294,967,295
	switch true {
	case v >= 601_692_057:
		return 7
	case v >= 10_556_001:
		return 6
	case v >= 185_193:
		return 5
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

func powU32(a int) (out uint32) {
	if a == 0 {
		return 1
	}

	out = 1
	for i := 0; i < a; i++ {
		out *= uint32(Base)
	}

	return
}
