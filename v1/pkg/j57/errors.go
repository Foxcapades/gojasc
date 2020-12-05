package j57

var (
	ErrNoHeader          = newJASCError("no size header in JASC value", 0)
	ErrInvalidJASCFormat = newJASCError("invalid format for a JASC value", 0)
)

var (
	TxtErrInvalidJASCByte = "unrecognized JASC byte value"
)

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

type JASCError interface {
	error

	Pos() int
	Kind() ValueKind
}

func IsJASCError(e error) bool {
	_, ok := e.(JASCError)
	return ok
}

type jascError struct {
	err  string
	pos  int
	kind ValueKind
}

func (j *jascError) Pos() int {
	return j.pos
}

func (j *jascError) Error() string {
	return j.err
}

func NewJASCFormatError(pos int, kind ValueKind) error {
	return &jascError{
		err:  "invalid serialized byte value for type " + kind.String(),
		pos:  pos,
		kind: kind,
	}
}

func NewJASCByteError(pos int) error {
	return newJASCError(TxtErrInvalidJASCByte, pos)
}

func newJASCError(err string, pos int) error {
	return &jascError{err: err, pos: pos}
}
