package j57_test

import (
	. "github.com/foxcapades/gojasc/v1/pkg/j57"
	"testing"
)

//goland:noinspection ALL
var __ignore_me__ interface{}

// The following tests are simply to verify the declared function types match
// the included implementations.  If the types do not match, this file will not
// compile.

func TestBoolTypes(t *testing.T) {
	__ignore_me__ = struct {
		t1 BoolBufAppenderFn
		t2 BoolDeserializerFn
		t3 BoolSerializerFn
		t4 BoolSizeFn
	}{
		AppendBool,
		DeserializeBool,
		SerializeBool,
		SizeBool,
	}
}

func TestC64Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 C64BufAppenderFn
		t2 C64DeserializerFn
		t3 C64SerializerFn
		t4 C64SizeFn
	}{
		AppendComplex64,
		DeserializeComplex64,
		SerializeComplex64,
		SizeComplex64,
	}
}

func TestC128Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 C128BufAppenderFn
		t2 C128DeserializerFn
		t3 C128SerializerFn
		t4 C128SizeFn
	}{
		AppendComplex128,
		DeserializeComplex128,
		SerializeComplex128,
		SizeComplex128,
	}
}

func TestF32Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 F32BufAppenderFn
		t2 F32DeserializerFn
		t3 F32SerializerFn
		t4 F32SizeFn
	}{
		AppendFloat32,
		DeserializeFloat32,
		SerializeFloat32,
		SizeFloat32,
	}
}

func TestF64Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 F64BufAppenderFn
		t2 F64DeserializerFn
		t3 F64SerializerFn
		t4 F64SizeFn
	}{
		AppendFloat64,
		DeserializeFloat64,
		SerializeFloat64,
		SizeFloat64,
	}
}

func TestITypes(t *testing.T) {
	__ignore_me__ = struct {
		t1 IBufAppenderFn
		t2 IDeserializerFn
		t3 ISerializerFn
		t4 ISizeFn
	}{
		AppendInt,
		DeserializeInt,
		SerializeInt,
		SizeInt,
	}
}

func TestI8Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 I8BufAppenderFn
		t2 I8DeserializerFn
		t3 I8SerializerFn
		t4 I8SizeFn
	}{
		AppendInt8,
		DeserializeInt8,
		SerializeInt8,
		SizeInt8,
	}
}

func TestI16Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 I16BufAppenderFn
		t2 I16DeserializerFn
		t3 I16SerializerFn
		t4 I16SizeFn
	}{
		AppendInt16,
		DeserializeInt16,
		SerializeInt16,
		SizeInt16,
	}
}

func TestI32Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 I32BufAppenderFn
		t2 I32DeserializerFn
		t3 I32SerializerFn
		t4 I32SizeFn
	}{
		AppendInt32,
		DeserializeInt32,
		SerializeInt32,
		SizeInt32,
	}
}

func TestI64Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 I64BufAppenderFn
		t2 I64DeserializerFn
		t3 I64SerializerFn
		t4 I64SizeFn
	}{
		AppendInt64,
		DeserializeInt64,
		SerializeInt64,
		SizeInt64,
	}
}

func TestStringTypes(t *testing.T) {
	__ignore_me__ = struct {
		t1 StringBufAppenderFn
		t2 StringDeserializerFn
		t3 StringSerializerFn
		t4 StringSizeFn
	}{
		AppendString,
		DeserializeString,
		SerializeString,
		SizeString,
	}
}

func TestUTypes(t *testing.T) {
	__ignore_me__ = struct {
		t1 UBufAppenderFn
		t2 UDeserializerFn
		t3 USerializerFn
		t4 USizeFn
	}{
		AppendUint,
		DeserializeUint,
		SerializeUint,
		SizeUint,
	}
}

func TestU8Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 U8BufAppenderFn
		t2 U8DeserializerFn
		t3 U8SerializerFn
		t4 U8SizeFn
	}{
		AppendUint8,
		DeserializeUint8,
		SerializeUint8,
		SizeUint8,
	}
}

func TestU16Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 U16BufAppenderFn
		t2 U16DeserializerFn
		t3 U16SerializerFn
		t4 U16SizeFn
	}{
		AppendUint16,
		DeserializeUint16,
		SerializeUint16,
		SizeUint16,
	}
}

func TestU32Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 U32BufAppenderFn
		t2 U32DeserializerFn
		t3 U32SerializerFn
		t4 U32SizeFn
	}{
		AppendUint32,
		DeserializeUint32,
		SerializeUint32,
		SizeUint32,
	}
}

func TestU64Types(t *testing.T) {
	__ignore_me__ = struct {
		t1 U64BufAppenderFn
		t2 U64DeserializerFn
		t3 U64SerializerFn
		t4 U64SizeFn
	}{
		AppendUint64,
		DeserializeUint64,
		SerializeUint64,
		SizeUint64,
	}
}
