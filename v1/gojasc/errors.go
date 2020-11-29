package gojasc

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
}

type jascError struct {
	err string
	pos int
}

func (j *jascError) Pos() int {
	return j.pos
}

func (j *jascError) Error() string {
	return j.err
}

func NewJASCByteError(pos int) error {
	return newJASCError(TxtErrInvalidJASCByte, pos)
}

func newJASCError(err string, pos int) error {
	return &jascError{err: err, pos: pos}
}
