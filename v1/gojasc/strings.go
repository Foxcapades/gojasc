package gojasc

// MicroString represents a string with a max size of 57 serialized bytes.
type MicroString string

//func (m MicroString) SerializeJASC() []byte {
//
//}

// TinyString represents a string with a max size of 3,249 serialized bytes.
type TinyString string

// ShortString represents a string with a max size of 185,193 serialized bytes.
type ShortString string

// MidString represents a string with a max size of 10,556,001 serialized bytes.
type MidString string

// LargeString represents a string with a max size of 601,692,057 serialized
// bytes.
type LargeString string

// HugeString represents a string with a max size of 34,296,447,249 serialized
// bytes.
type HugeString string