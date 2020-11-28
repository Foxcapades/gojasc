package gojasc


func DeserializeUDigit(b byte) uint8 {
	return b - min
}

// SerializeUint8 converts the given uint8 value to a base57 encoded byte slice.
//
// The byte slice returned may be 1-3 bytes in length.
func SerializeUint8(v uint8) []byte {
	if v == 0 {
		return []byte{min}
	}

	sz  := SizeUint8(v)
	pos := sz + 1
	out := make([]byte, pos)

	for v > 0 {
		pos--
		out[pos] = v % Base + min
		v /= Base
	}

	out[0] = byte(sz) + min

	return out
}

func DeserializeUint8(v []byte) uint8 {
	a := DeserializeUDigit(v[0])

	switch a {
	case 0:
		return 0
	case 1:
		return DeserializeUDigit(v[1])
	default:
		return DeserializeUDigit(v[2]) + DeserializeUDigit(v[1]) * Base
	}
}

// SerializeUint16 converts the given uint16 value to a base57 encoded byte
// slice.
//
// The byte slice returned may be 1-4 bytes in length.
func SerializeUint16(v uint16) []byte {
	if v == 0 {
		return []byte{min}
	}

	sz  := SizeUint16(v)
	pos := sz + 1
	out := make([]byte, pos)

	for v > 0 {
		pos--
		out[pos] = byte(v % Base) + min
		v /= Base
	}

	out[0] = uint8(sz) + min

	return out
}

func DeserializeUint16(v []byte) (out uint16) {
	a := DeserializeUDigit(v[0])

	switch a {
	case 0:
		return 0
	case 1:
		return uint16(DeserializeUDigit(v[1]))
	case 2:
		return uint16(DeserializeUDigit(v[2])) + uint16(DeserializeUDigit(v[1])) * Base
	default:
		return uint16(DeserializeUDigit(v[3])) +
			uint16(DeserializeUDigit(v[2])) * Base +
			uint16(DeserializeUDigit(v[1])) * powU16(Base,2)
	}
}

// SerializeUint32 converts the given uint32 value to a base57 encoded byte
// slice.
//
// The byte slice returned may be 1-7 bytes in length.
func SerializeUint32(v uint32) []byte {
	if v == 0 {
		return []byte{min}
	}

	sz  := SizeUint32(v)
	pos := sz + 1
	out := make([]byte, pos)

	for v > 0 {
		pos--
		out[pos] = byte(v % Base) + min
		v /= Base
	}

	out[0] = uint8(sz) + min

	return out
}

func DeserializeUint32(v []byte) (out uint32) {
	sz := DeserializeUDigit(v[0])

	switch sz {
	case 0:
		return 0
	case 1:
		return uint32(DeserializeUDigit(v[1]))
	case 2:
		return uint32(DeserializeUDigit(v[2])) + uint32(DeserializeUDigit(v[1])) * Base
	}

	i := int(sz)
	for j := uint8(0); j < sz; j++{
		out += uint32(DeserializeUDigit(v[i])) * powU32(Base, j)
		i--
	}

	return out
}

// SerializeUint64 converts the given uint64 value to a base57 encoded byte
// slice.
//
// The byte slice returned may be 1-12 bytes in length.
func SerializeUint64(v uint64) []byte {
	if v == 0 {
		return []byte{min}
	}

	sz  := SizeUint64(v)
	pos := sz + 1
	out := make([]byte, pos)

	for v > 0 {
		pos--
		out[pos] = byte(v % Base) + min
		v /= Base
	}

	out[0] = uint8(sz) + min

	return out
}

func DeserializeUint64(v []byte) (out uint64) {
	sz := DeserializeUDigit(v[0])

	switch sz {
	case 0:
		return 0
	case 1:
		return uint64(DeserializeUDigit(v[1]))
	case 2:
		return uint64(DeserializeUDigit(v[2])) + uint64(DeserializeUDigit(v[1])) * Base
	}

	i := int(sz)
	for j := uint8(0); j < sz; j++{
		out += uint64(DeserializeUDigit(v[i])) * powU64(Base, j)
		i--
	}

	return out
}

func SizeUint8(v uint8) int {
	if v > Base {
		return 2
	}

	return 1
}

func SizeUint16(v uint16) int {
	switch true {
	case v > 3249:
		return 3
	case v > 57:
		return 2
	default:
		return 1
	}
}

func SizeUint32(v uint32) int {
	// 4,294,967,295
	switch true {
	case v > 601_692_057:
		return 6
	case v > 10_556_001:
		return 5
	case v > 185_193:
		return 4
	case v > 3_249:
		return 3
	case v > 57:
		return 2
	default:
		return 1
	}
}

func SizeUint64(v uint64) int {
	switch true {
	case v > 362_033_331_456_891_249:
		return 11
	case v > 6_351_461_955_384_057:
		return 10
	case v > 111_429_157_112_001:
		return 9
	case v > 1_954_897_493_193:
		return 8
	case v > 34_296_447_249:
		return 7
	case v > 601_692_057:
		return 6
	default:
		return SizeUint32(uint32(v))
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

func powU32(a, b uint8) (out uint32) {
	if b == 0 {
		return 1
	}

	out = 1
	for i := uint8(0); i < b; i++ {
		out *= uint32(a)
	}

	return
}

func powU64(a, b uint8) (out uint64) {
	if b == 0 {
		return 1
	}

	out = 1
	for i := uint8(0); i < b; i++ {
		out *= uint64(a)
	}

	return
}