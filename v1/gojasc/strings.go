package gojasc

import (
	"github.com/foxcapades/tally-go/v1/tally"
)

func SerializeString(v string) ([]byte, error) {
	sz := SizeString(v)

	if sz == 1 {
		return []byte{min}, nil
	}

	out := make([]byte, sz)
	off := tally.UTally(0)

	SerializeUint64Into(uint64(sz), out, &off)
	for i := range v {
		out[off.Inc()] = v[i]
	}

	return out, nil
}

func SerializeStringInto(v string, buf []byte, off *tally.UTally) error {
	sz := SizeString(v)

	if sz == 1 {
		buf[off.Inc()] = min
		return nil
	}

	SerializeUint64Into(uint64(sz), buf, off)
	for i := range v {
		buf[off.Inc()] = v[i]
	}

	return nil
}

func SizeString(v string) uint {
	sz := len(v)
	if sz == 0 {
		return 1
	}

	return uint(SizeUint64(uint64(sz))) + uint(sz)
}

func DeserializeString(buf []byte, off *tally.UTally) (string, error) {
	ln := len(buf)

	if ln < 1 {
		return "", ErrNoHeader
	}

	init := off.Cur()
	strSize, err := DeserializeUint64(buf, off)

	if err != nil {
		return "", err
	}

	if strSize == 0 {
		return "", nil
	}

	strSize -= uint64(off.Cur() - init)

	out := make([]byte, strSize)
	copy(out, buf[off.Add(uint(strSize)):off.Cur()])

	return string(out), nil
}
