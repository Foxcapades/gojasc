package gojasc_test

import (
	"fmt"
	"github.com/foxcapades/gojasc/v1/gojasc"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var chars = [...]byte{
// 0    1    2    3    4     5    6    7
	'#', '$', '%', '&', '\'', '(', ')', '*',
// 8    9    10   11   12   13   14   15
	'+', ',', '-', '.', '/', '0', '1', '2',
// 16   17   18   19   20   21   22   23
	'3', '4', '5', '6', '7', '8', '9', ':',
// 24   25   26   27   28   29   30   31
	';', '<', '=', '>', '?', '@', 'A', 'B',
// 32   33   34   35   36   37   38   39
	'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
// 40   41   42   43   44   45   46   47
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R',
// 48   49   50   51   52   53   54   55
	'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
// 56
	'[',
}

func TestSerializeUInt8(t *testing.T) {
	Convey("SerializeUInt8", t, func() {
		type bt struct {
			i uint8
			o []byte
		}

		for _, pair := range []bt{
			{0, []byte("#")},
			{1, []byte("$$")},
			{2, []byte("$%")},
			{3, []byte("$&")},
			{4, []byte("$'")},
			{5, []byte("$(")},
			{6, []byte("$)")},
			{7, []byte("$*")},
			{8, []byte("$+")},
			{9, []byte("$,")},
			{10, []byte("$-")},
			{11, []byte("$.")},
			{12, []byte("$/")},
			{13, []byte("$0")},
			{14, []byte("$1")},
			{15, []byte("$2")},
			{16, []byte("$3")},
			{20, []byte("$7")},
			{32, []byte("$C")},
			{64, []byte("%$*")},
			{128, []byte("%%1")},
			{255, []byte("%'>")},
		} {
			Convey(fmt.Sprintf("Input %d", pair.i), func() {
				tmp := gojasc.SerializeUint8(pair.i)
				So(tmp, ShouldResemble, pair.o)
			})
		}
	})
}

func TestDeserializeUInt8(t *testing.T) {
	Convey("DeserializeUInt8", t, func() {
		type bt struct {
			i uint8
			o []byte
		}

		for _, pair := range []bt{
			{0, []byte("#")},
			{1, []byte("$$")},
			{2, []byte("$%")},
			{3, []byte("$&")},
			{4, []byte("$'")},
			{5, []byte("$(")},
			{6, []byte("$)")},
			{7, []byte("$*")},
			{8, []byte("$+")},
			{9, []byte("$,")},
			{10, []byte("$-")},
			{11, []byte("$.")},
			{12, []byte("$/")},
			{13, []byte("$0")},
			{14, []byte("$1")},
			{15, []byte("$2")},
			{16, []byte("$3")},
			{20, []byte("$7")},
			{32, []byte("$C")},
			{64, []byte("%$*")},
			{128, []byte("%%1")},
			{255, []byte("%'>")},
		} {
			Convey(fmt.Sprintf("Input %d", pair.i), func() {
				tmp := gojasc.DeserializeUint8(pair.o)
				So(tmp, ShouldResemble, pair.i)
			})
		}
	})
}

func TestSerializeUint16(t *testing.T) {
	Convey("SerializeUint16", t, func() {
		Convey("Input: 0", func() {
			So(string(gojasc.SerializeUint16(0)), ShouldEqual, "#")
		})
		Convey("Input: 1", func() {
			So(string(gojasc.SerializeUint16(1)), ShouldEqual, "$$")
		})
		Convey("Input: 255", func() {
			So(string(gojasc.SerializeUint16(255)), ShouldEqual, "%'>")
		})
		Convey("Input: 65,535", func() {
			So(string(gojasc.SerializeUint16(65535)), ShouldEqual, "&7,M")
		})
	})
}

func TestDeserializeUint16(t *testing.T) {
	Convey("DeserializeUint16", t, func() {
		Convey("Input: 0", func() {
			So(gojasc.DeserializeUint16([]byte("#")), ShouldEqual, 0)
		})
		Convey("Input: 1", func() {
			So(gojasc.DeserializeUint16([]byte("$$")), ShouldEqual, 1)
		})
		Convey("Input: 255", func() {
			So(gojasc.DeserializeUint16([]byte("%'>")), ShouldEqual, 255)
		})
		Convey("Input: 65,535", func() {
			So(gojasc.DeserializeUint16([]byte("&7,M")), ShouldEqual, 65535)
		})
	})
}

func TestSerializeUint32(t *testing.T) {
	Convey("SerializeUint32", t, func() {
		Convey("Input: 0", func() {
			So(string(gojasc.SerializeUint32(0)), ShouldEqual, "#")
		})
		Convey("Input: 1", func() {
			So(string(gojasc.SerializeUint32(1)), ShouldEqual, "$$")
		})
		Convey("Input: 255", func() {
			So(string(gojasc.SerializeUint32(255)), ShouldEqual, "%'>")
		})
		Convey("Input: 65,535", func() {
			So(string(gojasc.SerializeUint32(65_535)), ShouldEqual, "&7,M")
		})
		Convey("Input: 185,194", func() {
			So(string(gojasc.SerializeUint32(185_194)), ShouldEqual, "'$##$")
		})
		Convey("Input: 10,556,100", func() {
			So(string(gojasc.SerializeUint32(10_556_100)), ShouldEqual, "($##$M")
		})
		Convey("Input: 4,294,967,295", func() {
			So(string(gojasc.SerializeUint32(4_294_967_295)), ShouldEqual, ")**TS+;")
		})
	})
}

func TestDeserializeUint32(t *testing.T) {
	Convey("DeserializeUint32", t, func() {
		Convey("Input: 0", func() {
			So(gojasc.DeserializeUint32([]byte("#")), ShouldEqual, 0)
		})
		Convey("Input: 1", func() {
			So(gojasc.DeserializeUint32([]byte("$$")), ShouldEqual, 1)
		})
		Convey("Input: 255", func() {
			So(gojasc.DeserializeUint32([]byte("%'>")), ShouldEqual, 255)
		})
		Convey("Input: 65,535", func() {
			So(gojasc.DeserializeUint32([]byte("&7,M")), ShouldEqual, 65535)
		})
		Convey("Input: 185,194", func() {
			So(gojasc.DeserializeUint32([]byte("'$##$")), ShouldEqual, 185_194)
		})
		Convey("Input: 10,556,100", func() {
			So(gojasc.DeserializeUint32([]byte("($##$M")), ShouldEqual, 10_556_100)
		})
		Convey("Input: 4,294,967,295", func() {
			So(gojasc.DeserializeUint32([]byte(")**TS+;")), ShouldEqual, 4294967295)
		})
	})
}

func TestSerializeUint64(t *testing.T) {
	Convey("SerializeUint64", t, func() {
		Convey("Input: 0", func() {
			So(string(gojasc.SerializeUint64(0)), ShouldEqual, "#")
		})
		Convey("Input: 1", func() {
			So(string(gojasc.SerializeUint64(1)), ShouldEqual, "$$")
		})
		Convey("Input: 255", func() {
			So(string(gojasc.SerializeUint64(255)), ShouldEqual, "%'>")
		})
		Convey("Input: 65,535", func() {
			So(string(gojasc.SerializeUint64(65_535)), ShouldEqual, "&7,M")
		})
		Convey("Input: 185,194", func() {
			So(string(gojasc.SerializeUint64(185_194)), ShouldEqual, "'$##$")
		})
		Convey("Input: 10,556,100", func() {
			So(string(gojasc.SerializeUint64(10_556_100)), ShouldEqual, "($##$M")
		})
		Convey("Input: 601,692,657", func() {
			So(string(gojasc.SerializeUint64(601_692_657)), ShouldEqual, ")$###-A")
		})
		Convey("Input: 37,256,437,249", func() {
			So(string(gojasc.SerializeUint64(37_256_437_249)), ShouldEqual, "*$'W:2>*")
		})
		Convey("Input: 2,964,996,584,284", func() {
			So(string(gojasc.SerializeUint64(2_964_996_584_284)), ShouldEqual, "+$@<NBW<Q")
		})
		Convey("Input: 123,456,789,012,345", func() {
			So(string(gojasc.SerializeUint64(123_456_789_012_345)), ShouldEqual, ",$)+JIO@IY")
		})
		Convey("Input: 12,345,678,901,234,567", func() {
			So(string(gojasc.SerializeUint64(12_345_678_901_234_567)), ShouldEqual, "-$XP1F%)'6W")
		})
		Convey("Input: 567,890,123,456,789,012", func() {
			So(string(gojasc.SerializeUint64(567_890_123_456_789_012)), ShouldEqual, ".$C:;)<MX02%")
		})
	})
}

func TestDeserializeUint64(t *testing.T) {
	Convey("DeserializeUint64", t, func() {
		Convey("Input: 0", func() {
			So(gojasc.DeserializeUint64([]byte("#")), ShouldEqual, 0)
		})
		Convey("Input: 1", func() {
			So(gojasc.DeserializeUint64([]byte("$$")), ShouldEqual, 1)
		})
		Convey("Input: 255", func() {
			So(gojasc.DeserializeUint64([]byte("%'>")), ShouldEqual, 255)
		})
		Convey("Input: 65,535", func() {
			So(gojasc.DeserializeUint64([]byte("&7,M")), ShouldEqual, 65535)
		})
		Convey("Input: 185,194", func() {
			So(gojasc.DeserializeUint64([]byte("'$##$")), ShouldEqual, 185_194)
		})
		Convey("Input: 10,556,100", func() {
			So(gojasc.DeserializeUint64([]byte("($##$M")), ShouldEqual, 10_556_100)
		})
		Convey("Input: 601,692,657", func() {
			So(gojasc.DeserializeUint64([]byte(")$###-A")), ShouldEqual, 601_692_657)
		})
		Convey("Input: 37,256,437,249", func() {
			So(gojasc.DeserializeUint64([]byte("*$'W:2>*")), ShouldEqual, 37_256_437_249)
		})
		Convey("Input: 2,964,996,584,284", func() {
			So(gojasc.DeserializeUint64([]byte("+$@<NBW<Q")), ShouldEqual, 2_964_996_584_284)
		})
		Convey("Input: 123,456,789,012,345", func() {
			So(gojasc.DeserializeUint64([]byte(",$)+JIO@IY")), ShouldEqual, 123_456_789_012_345)
		})
		Convey("Input: 12,345,678,901,234,567", func() {
			So(gojasc.DeserializeUint64([]byte("-$XP1F%)'6W")), ShouldEqual, 12_345_678_901_234_567)
		})
		Convey("Input: 567,890,123,456,789,012", func() {
			So(gojasc.DeserializeUint64([]byte(".$C:;)<MX02%")), ShouldEqual, 567_890_123_456_789_012)
		})
	})
}
