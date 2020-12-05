package j57

import (
	"github.com/foxcapades/tally-go/v1/tally"
	"unsafe"
)

const (
	stringChunkSize = 8
)

func SerializeString(v string) []byte {
	in := unsafeROBytes(&v)
	sz := len(in)

	if sz == 0 {
		return []byte{min}
	}

	buffer := make([]byte, (sz/8)*12+24)
	offset := tally.UTally(0)
	subSerializeString(in, sz, buffer, &offset)

	return buffer
}

func AppendString(v string, buf []byte, off *tally.UTally) (wrote int) {
	in := unsafeROBytes(&v)
	sz := len(in)

	if sz == 0 {
		buf[off.Inc()] = min
		return 1
	}

	return subSerializeString(in, sz, buf, off)
}

// | prefix | chunk-count | prefix | byte count |
func subSerializeString(
	in []byte,
	sz int,
	buf []byte,
	off *tally.UTally,
) (wrote int) {
	blockCnt := sz / stringChunkSize
	overflow := sz % stringChunkSize
	blockBuf := [8]byte{}

	// Add block count header
	if overflow == 0 {
		wrote += AppendUint64(uint64(blockCnt), buf, off)
	} else {
		wrote += AppendUint64(uint64(blockCnt)+1, buf, off)
	}

	// Add unpacked size header
	wrote += AppendUint64(uint64(sz), buf, off)

	pos := 0
	for i := 0; i < blockCnt; i++ {
		blockBuf[0] = in[pos]
		pos++
		blockBuf[1] = in[pos]
		pos++
		blockBuf[2] = in[pos]
		pos++
		blockBuf[3] = in[pos]
		pos++
		blockBuf[4] = in[pos]
		pos++
		blockBuf[5] = in[pos]
		pos++
		blockBuf[6] = in[pos]
		pos++
		blockBuf[7] = in[pos]
		pos++

		wrote += AppendUint64(unsafeS2U64(&blockBuf), buf, off)
	}

	i := 0
	for ; i < overflow; i++ {
		blockBuf[i] = in[pos]
		pos++
	}
	for ; i < 8; i++ {
		blockBuf[i] = 0
	}

	return wrote + AppendUint64(unsafeS2U64(&blockBuf), buf, off)
}

func DeserializeString(buf []byte, off *tally.UTally) (string, error) {
	if len(buf) < 1 {
		return "", ErrNoHeader
	}

	chunks, err := DeserializeUint32(buf, off)
	if err != nil {
		return "", err
	}

	bytes, err := DeserializeUint64(buf, off)
	if err != nil {
		return "", err
	}

	stage := make([]byte, bytes)

	pos := tally.ITally(0)
	for i := uint32(0); i < chunks; i++ {
		ch, err := DeserializeUint64(buf, off)
		if err != nil {
			return "", err
		}

		tmp := unsafeU642S(ch)
		copy(stage[pos.Add(stringChunkSize):], tmp)
	}

	return string(stage), nil
}

func SizeString(v string) uint {
	sz := len(v)
	if sz == 0 {
		return 1
	}

	return SizeUint64(uint64(sz)) + uint(sz)
}

func unsafeROBytes(v *string) []byte {
	return *(*[]byte)(unsafe.Pointer(v))
}

func unsafeU642S(u uint64) []byte {
	return ((*[stringChunkSize]byte)(unsafe.Pointer(&u)))[:]
}

func unsafeS2U64(s *[8]byte) uint64 {
	return *(*uint64)(unsafe.Pointer(s))
}
