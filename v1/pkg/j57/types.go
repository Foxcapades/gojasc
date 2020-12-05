package j57

import "github.com/foxcapades/tally-go/v1/tally"

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   Boolean Serialization Functions                                      ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// BoolSerializerFn is a function type declaration matching the included
// SerializeBool function.
type BoolSerializerFn = func(bool) []byte

// BoolBufAppenderFn is a function type declaration matching the included
// SerializeBoolInto function.
type BoolBufAppenderFn = func(bool, []byte, *tally.UTally) int

// BoolDeserializerFn is a function type declaration matching the included
// DeserializeBool function.
type BoolDeserializerFn = func([]byte, *tally.UTally) (bool, error)

// BoolSizeFn is a function type declaration matching the included BoolSize
// function.
type BoolSizeFn = func(bool) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   Complex64 Serialization Functions                                    ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// C64SerializerFn is a function type declaration matching the included
// SerializeComplex64 function.
type C64SerializerFn = func(complex64) []byte

// C64BufAppenderFn is a function type declaration matching the included
// AppendComplex64 function.
type C64BufAppenderFn = func(complex64, []byte, *tally.UTally) int

// C64DeserializerFn is a function type declaration matching the included
// DeserializeComplex64 function.
type C64DeserializerFn = func([]byte, *tally.UTally) (complex64, error)

// C64SizeFn is a function type declaration matching the included Complex64Size
// function.
type C64SizeFn = func(complex64) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   Complex128 Serialization Functions                                   ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// C128SerializerFn is a function type declaration matching the included
// SerializeComplex128 function.
type C128SerializerFn = func(complex128) []byte

// C128BufAppenderFn is a function type declaration matching the included
// AppendComplex128 function.
type C128BufAppenderFn = func(complex128, []byte, *tally.UTally) int

// C128DeserializerFn is a function type declaration matching the included
// DeserializeComplex128 function.
type C128DeserializerFn = func([]byte, *tally.UTally) (complex128, error)

// C128SizeFn is a function type declaration matching the included
// Complex128Size function.
type C128SizeFn = func(complex128) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   Float32 Serialization Functions                                      ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// F32SerializerFn is a function type declaration matching the included
// SerializeFloat32 function.
type F32SerializerFn = func(float32) []byte

// F32BufAppenderFn is a function type declaration matching the included
// AppendFloat32 function.
type F32BufAppenderFn = func(float32, []byte, *tally.UTally) int

// F32DeserializerFn is a function type declaration matching the included
// DeserializeFloat32 function.
type F32DeserializerFn = func([]byte, *tally.UTally) (float32, error)

// F32SizeFn is a function type declaration matching the included Float32Size
// function.
type F32SizeFn = func(float32) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   Float64 Serialization Functions                                      ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// F64SerializerFn is a function type declaration matching the included
// SerializeFloat64 function.
type F64SerializerFn = func(float64) []byte

// F64BufAppenderFn is a function type declaration matching the included
// AppendFloat64 function.
type F64BufAppenderFn = func(float64, []byte, *tally.UTally) int

// F64DeserializerFn is a function type declaration matching the included
// DeserializeFloat64 function.
type F64DeserializerFn = func([]byte, *tally.UTally) (float64, error)

// F64SizeFn is a function type declaration matching the included Float64Size
// function.
type F64SizeFn = func(float64) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   Untyped Int Serialization Functions                                  ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// ISerializerFn is a function type declaration matching the included
// SerializeInt function.
type ISerializerFn = func(int) []byte

// IBufAppenderFn is a function type declaration matching the included
// AppendInt function.
type IBufAppenderFn = func(int, []byte, *tally.UTally) int

// IDeserializerFn is a function type declaration matching the included
// DeserializeInt function.
type IDeserializerFn = func([]byte, *tally.UTally) (int, error)

// ISizeFn is a function type declaration matching the included IntSize
// function.
type ISizeFn = func(int) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   Int8 Serialization Functions                                         ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// I8SerializerFn is a function type declaration matching the included
// SerializeInt8 function.
type I8SerializerFn = func(int8) []byte

// I8BufAppenderFn is a function type declaration matching the included
// AppendInt8 function.
type I8BufAppenderFn = func(int8, []byte, *tally.UTally) int

// I8DeserializerFn is a function type declaration matching the included
// DeserializeInt8 function.
type I8DeserializerFn = func([]byte, *tally.UTally) (int8, error)

// I8SizeFn is a function type declaration matching the included Int8Size
// function.
type I8SizeFn = func(int8) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   Int16 Serialization Functions                                        ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// I16SerializerFn is a function type declaration matching the included
// SerializeInt16 function.
type I16SerializerFn = func(int16) []byte

// I16BufAppenderFn is a function type declaration matching the included
// AppendInt16 function.
type I16BufAppenderFn = func(int16, []byte, *tally.UTally) int

// I16DeserializerFn is a function type declaration matching the included
// DeserializeInt16 function.
type I16DeserializerFn = func([]byte, *tally.UTally) (int16, error)

// I16SizeFn is a function type declaration matching the included Int16Size
// function.
type I16SizeFn = func(int16) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   Int32 Serialization Functions                                        ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// I32SerializerFn is a function type declaration matching the included
// SerializeInt32 function.
type I32SerializerFn = func(int32) []byte

// I32BufAppenderFn is a function type declaration matching the included
// AppendInt32 function.
type I32BufAppenderFn = func(int32, []byte, *tally.UTally) int

// I32DeserializerFn is a function type declaration matching the included
// DeserializeInt32 function.
type I32DeserializerFn = func([]byte, *tally.UTally) (int32, error)

// I32SizeFn is a function type declaration matching the included Int32Size
// function.
type I32SizeFn = func(int32) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   Int64 Serialization Functions                                        ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// I64SerializerFn is a function type declaration matching the included
// SerializeInt64 function.
type I64SerializerFn = func(int64) []byte

// I64BufAppenderFn is a function type declaration matching the included
// AppendInt64 function.
type I64BufAppenderFn = func(int64, []byte, *tally.UTally) int

// I64DeserializerFn is a function type declaration matching the included
// DeserializeInt64 function.
type I64DeserializerFn = func([]byte, *tally.UTally) (int64, error)

// I64SizeFn is a function type declaration matching the included Int64Size
// function.
type I64SizeFn = func(int64) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   String Serialization Functions                                       ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// StringSerializerFn is a function type declaration matching the included
// SerializeString function.
type StringSerializerFn = func(string) []byte

// StringBufAppenderFn is a function type declaration matching the included
// AppendString function.
type StringBufAppenderFn = func(string, []byte, *tally.UTally) int

// StringDeserializerFn is a function type declaration matching the included
// DeserializeString function.
type StringDeserializerFn = func([]byte, *tally.UTally) (string, error)

// StringSizeFn is a function type declaration matching the included StringSize
// function.
type StringSizeFn = func(string) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   Untyped Uint Serialization Functions                                 ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// USerializerFn is a function type declaration matching the included
// SerializeUint function.
type USerializerFn = func(uint) []byte

// UBufAppenderFn is a function type declaration matching the included
// AppendUint function.
type UBufAppenderFn = func(uint, []byte, *tally.UTally) int

// UDeserializerFn is a function type declaration matching the included
// DeserializeUint function.
type UDeserializerFn = func([]byte, *tally.UTally) (uint, error)

// USizeFn is a function type declaration matching the included UintSize
// function.
type USizeFn = func(uint) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   UInt8 Serialization Functions                                        ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// U8SerializerFn is a function type declaration matching the included
// SerializeUint8 function.
type U8SerializerFn = func(uint8) []byte

// U8BufAppenderFn is a function type declaration matching the included
// AppendUint8 function.
type U8BufAppenderFn = func(uint8, []byte, *tally.UTally) int

// U8DeserializerFn is a function type declaration matching the included
// DeserializeUint8 function.
type U8DeserializerFn = func([]byte, *tally.UTally) (uint8, error)

// U8SizeFn is a function type declaration matching the included Uint8Size
// function.
type U8SizeFn = func(uint8) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   UInt16 Serialization Functions                                       ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// U16SerializerFn is a function type declaration matching the included
// SerializeUint16 function.
type U16SerializerFn = func(uint16) []byte

// U16BufAppenderFn is a function type declaration matching the included
// AppendUint16 function.
type U16BufAppenderFn = func(uint16, []byte, *tally.UTally) int

// U16DeserializerFn is a function type declaration matching the included
// DeserializeUint16 function.
type U16DeserializerFn = func([]byte, *tally.UTally) (uint16, error)

// U16SizeFn is a function type declaration matching the included Uint16Size
// function.
type U16SizeFn = func(uint16) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   UInt32 Serialization Functions                                       ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// U32SerializerFn is a function type declaration matching the included
// SerializeUint32 function.
type U32SerializerFn = func(uint32) []byte

// U32BufAppenderFn is a function type declaration matching the included
// AppendUint32 function.
type U32BufAppenderFn = func(uint32, []byte, *tally.UTally) int

// U32DeserializerFn is a function type declaration matching the included
// DeserializeUint32 function.
type U32DeserializerFn = func([]byte, *tally.UTally) (uint32, error)

// U32SizeFn is a function type declaration matching the included Uint32Size
// function.
type U32SizeFn = func(uint32) uint

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║   UInt16 Serialization Functions                                       ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// U64SerializerFn is a function type declaration matching the included
// SerializeUint64 function.
type U64SerializerFn = func(uint64) []byte

// U64BufAppenderFn is a function type declaration matching the included
// AppendUint64 function.
type U64BufAppenderFn = func(uint64, []byte, *tally.UTally) int

// U64DeserializerFn is a function type declaration matching the included
// DeserializeUint64 function.
type U64DeserializerFn = func([]byte, *tally.UTally) (uint64, error)

// U64SizeFn is a function type declaration matching the included Uint64Size
// function.
type U64SizeFn = func(uint64) uint
