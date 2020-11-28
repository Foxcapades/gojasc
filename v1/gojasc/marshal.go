package gojasc

import "github.com/foxcapades/tally-go/v1/tally"

type Marshaler interface {
	MarshalJASC() []byte
	MarshalJASCInto(buf []byte, off *tally.UTally)
	JASCSize() uint
}

type Unmarshaler interface {
	UnmarshalJASC(buf []byte, off *tally.UTally)
}
