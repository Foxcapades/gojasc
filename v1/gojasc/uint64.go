package gojasc

import "github.com/foxcapades/tally-go/v1/tally"

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Uint64 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝

// SerializeUint64 converts the given uint64 value to a base57 encoded byte
// slice.
//
// The byte slice returned may be 1-12 bytes in length.
func SerializeUint64(v uint64) []byte {
	if v == 0 {
		return []byte{min}
	}

	sz := SizeUint64(v)
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

// SerializeUint64Into converts the given uint64 value to base57 and writes it
// to the given buffer starting at off.Cur().
//
// The given offset value will be incremented as the buffer is written and after
// this function call will be at the next writable offset position.
//
// WARNING: This method makes no attempt to verify that the given byte buffer is
// actually long enough to hold the serialized value.  The buffer size should be
// at least SizeUint64(v) in length.
func SerializeUint64Into(v uint64, buf []byte, off *tally.UTally) {
	if v == 0 {
		buf[off.Inc()] = min
		return
	}

	sz := SizeUint64(v)
	pos := sz + off.Cur()
	cur := off.Add(sz)

	for v > 0 {
		pos--
		buf[pos] = byte(v%Base) + min
		v /= Base
	}

	buf[cur] = byte(sz-1) + min
}

func DeserializeUint64(v []byte, off *tally.UTally) (out uint64, err error) {
	sz := int(DeserializeUDigit(v[off.Inc()]))

	if sz < 0 {
		return 0, NewJASCByteError(0)
	} else if sz == 0 {
		return 0, nil
	}

	i := sz + int(off.Add(uint(sz)))

	for j := 0; j < sz; j++ {
		i--
		out += uint64(DeserializeUDigit(v[i])) * powU64(j)
	}

	return out, nil
}

// SizeUint64 returns the number of bytes needed in a byte buffer to hold the
// serialized form of the given uint64 value.
//
// This size includes the byte needed for the number size header.
func SizeUint64(v uint64) uint {
	switch true {
	case v >= 362_033_331_456_891_249:
		return 12
	case v >= 6_351_461_955_384_057:
		return 11
	case v >= 111_429_157_112_001:
		return 10
	case v >= 1_954_897_493_193:
		return 9
	case v >= 34_296_447_249:
		return 8
	case v >= 601_692_057:
		return 7
	default:
		return SizeUint32(uint32(v))
	}
}

func powU64(a int) (out uint64) {
	if a == 0 {
		return 1
	}

	tmp := uint64(Base)
	out = 1
	for i := 0; i < a; i++ {
		out *= tmp
	}

	return
}
