package jasc

var (
	ErrNoHeader          = newJASCError("no size header in JASC value", 0)
	ErrInvalidJASCFormat = newJASCError("invalid format for a JASC value", 0)
)

var (
	TxtErrInvalidJASCByte = "unrecognized JASC byte value"
)

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
