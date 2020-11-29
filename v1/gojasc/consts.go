package gojasc

const (
	Base         = 57
	min     byte = '#'
	max     byte = '['
	conMask byte = 128

	bTrue  = min + 1
	bFalse = min
)

type ValueKind uint8

const (
	KindNone ValueKind = iota
	KindBool
	KindString
	KindUint8
	KindUint16
	KindUint32
	KindUint64
)

func (v ValueKind) String() string {
	switch v {
	case KindString:
		return "string"
	case KindBool:
		return "bool"
	case KindUint8:
		return "uint8"
	case KindUint16:
		return "uint16"
	case KindUint32:
		return "uint32"
	case KindUint64:
		return "uint64"
	}
	return "none"
}
