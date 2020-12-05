package j57

import "github.com/foxcapades/tally-go/v1/tally"

func SerializeBool(v bool) []byte {
	if v {
		return SerializeUint8(1)
	}

	return SerializeUint8(0)
}

func AppendBool(v bool, buf []byte, off *tally.UTally) int {
	if v {
		return AppendUint8(1, buf, off)
	} else {
		return AppendUint8(0, buf, off)
	}
}

func DeserializeBool(buf []byte, off *tally.UTally) (bool, error) {
	if tmp, e := DeserializeUint8(buf, off); e != nil {
		return false, e
	} else if tmp == 0 {
		return false, nil
	} else if tmp == 1 {
		return true, nil
	} else {
		return false, ErrInvalidJASCFormat
	}
}

func SizeBool(v bool) uint {
	if v {
		return 2
	} else {
		return 1
	}
}
