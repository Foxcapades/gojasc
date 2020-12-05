package j57

import (
	"github.com/foxcapades/tally-go/v1/tally"
	"unsafe"
)

const (
	chunkSize = 8
)

func SerializeString(v string) []byte {
	in := unsafeROBytes(&v)
	sz := len(in)

	if sz == 0 {
		return []byte{min}
	}

	buffer := make([]byte, SizeBytes(in))
	offset := tally.UTally(0)
	subSerializeBytes(in, sz, buffer, &offset)

	return buffer
}

func AppendString(v string, buf []byte, off *tally.UTally) (wrote int) {
	in := unsafeROBytes(&v)
	sz := len(in)

	if sz == 0 {
		buf[off.Inc()] = min
		return 1
	}

	return subSerializeBytes(in, sz, buf, off)
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
		copy(stage[pos.Add(chunkSize):], tmp)
	}

	return string(stage), nil
}

func SizeString(v string) uint {
	sz := len(v)
	if sz == 0 {
		return 1
	}
	return SizeBytes(*(*[]byte)(unsafe.Pointer(&v)))
}

func unsafeROBytes(v *string) []byte {
	return *(*[]byte)(unsafe.Pointer(v))
}

func unsafeU642S(u uint64) (out []byte) {
	out = ((*[chunkSize]byte)(unsafe.Pointer(&u)))[:]
	return
}

func unsafeS2U64(s *[8]byte) uint64 {
	return *(*uint64)(unsafe.Pointer(s))
}
