package j57

import (
	"github.com/foxcapades/tally-go/v1/tally"
	"unsafe"
)

func SerializeBytes(v []byte) []byte {
	sz := len(v)

	if sz == 0 {
		return []byte{min}
	}

	buffer := make([]byte, SizeBytes(v))
	offset := tally.UTally(0)
	subSerializeBytes(v, sz, buffer, &offset)

	return buffer
}

func AppendBytes(v []byte, buf []byte, off *tally.UTally) (wrote int) {
	sz := len(v)

	if sz == 0 {
		buf[off.Inc()] = min
		return 1
	}

	return subSerializeBytes(v, sz, buf, off)
}

func subSerializeBytes(
	in []byte,
	sz int,
	buf []byte,
	off *tally.UTally,
) (wrote int) {
	blockCnt := sz / chunkSize
	overflow := sz % chunkSize
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

	if overflow > 0 {
		i := 0
		for ; i < overflow; i++ {
			blockBuf[i] = in[pos]
			pos++
		}
		for ; i < 8; i++ {
			blockBuf[i] = 0
		}

		wrote += AppendUint64(unsafeS2U64(&blockBuf), buf, off)
	}

	return
}

func DeserializeBytes(buf []byte, off *tally.UTally) ([]byte, error) {
	if len(buf) < 1 {
		return []byte{}, ErrNoHeader
	}

	chunks, err := DeserializeUint32(buf, off)
	if err != nil {
		return []byte{}, err
	}

	bytes, err := DeserializeUint64(buf, off)
	if err != nil {
		return []byte{}, err
	}

	stage := make([]byte, bytes)

	pos := tally.ITally(0)
	for i := uint32(0); i < chunks; i++ {
		ch, err := DeserializeUint64(buf, off)
		if err != nil {
			return []byte{}, err
		}

		tmp := unsafeU642S(ch)
		copy(stage[pos.Add(chunkSize):], tmp)
	}

	return stage, nil
}

func SizeBytes(v []byte) uint {
	sz := len(v)
	if sz == 0 {
		return 1
	}

	hold := *(*[]byte)(unsafe.Pointer(&v))
	blocks := sz / chunkSize
	extras := sz % chunkSize

	blockSize := uint(0)

	for i := 0; i < blocks; i++ {
		blockSize += SizeUint64(*(*uint64)(unsafe.Pointer(&hold[i*8])))
	}
	if extras > 0 {
		tmp := [chunkSize]byte{}
		for i, b := range hold[sz-extras:] {
			tmp[i] = b
		}
		blockSize += SizeUint64(*(*uint64)(unsafe.Pointer(&tmp)))
	}

	sizePrefix := SizeUint64(uint64(sz))
	blockPrefix := SizeUint64(uint64(blockSize))

	return blockSize + sizePrefix + blockPrefix
}
