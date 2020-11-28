package gojasc

import (
	"github.com/foxcapades/tally-go/v1/tally"
)

type String string

func (s String) MarshalJASC() []byte {
	sz := s.JASCSize()

	if sz == 1 {
		return []byte{min}
	}

	out := make([]byte, sz)
	off := tally.UTally(0)
	SerializeUint64Into(uint64(sz), out, &off)
	for i := range s {
		out[off.Inc()] = s[i]
	}

	return out
}

func (s String) MarshalJASCInto(buf []byte, off *tally.UTally) {
	sz := s.JASCSize()

	if sz == 1 {
		buf[off.Inc()] = min
		return
	}

	SerializeUint64Into(uint64(sz), buf, off)
	for i := range s {
		buf[off.Inc()] = s[i]
	}
}

func (s String) JASCSize() uint {
	sz := len(s)
	if sz == 0 {
		return 1
	}

	return uint(SizeUint64(uint64(sz))) + uint(sz)
}
