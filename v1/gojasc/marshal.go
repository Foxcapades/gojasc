package gojasc

import "github.com/foxcapades/tally-go/v1/tally"

type Serializable interface {
	SerializeJASC() ([]byte, error)
	SerializeJASCInto(buf []byte, off *tally.UTally)
	JASCSize() uint
}

type Deserializable interface {
	DeserializeJASC(buf []byte, off *tally.UTally)
}
