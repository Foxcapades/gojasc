package gojasc

import "github.com/foxcapades/tally-go/v1/tally"

func DeserializeUDigit(b byte) uint8 {
	return b - min
}

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
	cur := off.Add(uint(pos))

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
func SizeUint8(v uint8) int {
	switch true {
	case v == 0:
		return 1
	case v < Base:
		return 2
	default:
		return 3
	}
}

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
	pos := sz
	cur := off.Add(uint(pos))

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
func SizeUint16(v uint16) int {
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
	pos := sz
	cur := off.Add(uint(pos))

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
func SizeUint32(v uint32) int {
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
	pos := sz
	cur := off.Add(uint(pos))

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
func SizeUint64(v uint64) int {
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
