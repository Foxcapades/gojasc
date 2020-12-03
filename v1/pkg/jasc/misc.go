package jasc

import "github.com/foxcapades/tally-go/v1/tally"

func DeserializeUDigit(b byte) uint8 {
	return b - min
}

func SerializeBool(v bool) byte {
	if v {
		return bTrue
	}

	return bFalse
}

func SerializeBoolInto(v bool, buf []byte, off *tally.UTally) {
	buf[off.Inc()] = SerializeBool(v)
}

func DeserializeBool(buf []byte, off *tally.UTally) (bool, error) {
	switch buf[off.Inc()] {
	case bFalse:
		return false, nil
	case bTrue:
		return true, nil
	default:
		return false, NewJASCFormatError(int(off.Cur()-1), KindBool)
	}
}
