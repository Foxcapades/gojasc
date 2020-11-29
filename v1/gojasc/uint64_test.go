package gojasc_test

import (
	"fmt"
	"github.com/foxcapades/gojasc/v1/gojasc"
	"github.com/foxcapades/tally-go/v1/tally"
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

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Uint64 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

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

func TestSerializeUint64Into(t *testing.T) {
	Convey("SerializeUint64Into", t, func() {
		Convey("Input: 0", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(0, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "#")
		})
		Convey("Input: 1", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(1, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "$$")
		})
		Convey("Input: 255", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(255, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "%'>")
		})
		Convey("Input: 65,535", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(65_535, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "&7,M")
		})
		Convey("Input: 185,194", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(185_194, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "'$##$")
		})
		Convey("Input: 10,556,100", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(10_556_100, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "($##$M")
		})
		Convey("Input: 601,692,657", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(601_692_657, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, ")$###-A")
		})
		Convey("Input: 37,256,437,249", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(37_256_437_249, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "*$'W:2>*")
		})
		Convey("Input: 2,964,996,584,284", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(2_964_996_584_284, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "+$@<NBW<Q")
		})
		Convey("Input: 123,456,789,012,345", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(123_456_789_012_345, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, ",$)+JIO@IY")
		})
		Convey("Input: 12,345,678,901,234,567", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(12_345_678_901_234_567, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "-$XP1F%)'6W")
		})
		Convey("Input: 567,890,123,456,789,012", func() {
			buf := make([]byte, 12)
			off := tally.UTally(0)

			gojasc.SerializeUint64Into(567_890_123_456_789_012, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, ".$C:;)<MX02%")
		})

		Convey("Repeated use", func() {
			buf := make([]byte, 6)
			off := tally.UTally(0)
			gojasc.SerializeUint64Into(1, buf, &off)
			gojasc.SerializeUint64Into(2, buf, &off)
			gojasc.SerializeUint64Into(3, buf, &off)

			So(buf, ShouldResemble, []byte("$$$%$&"))
		})
	})
}

func TestDeserializeUint64(t *testing.T) {
	Convey("DeserializeUint64", t, func() {
		Convey("Input: 0", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte("#"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 0)
		})
		Convey("Input: 1", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte("$$"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 1)
		})
		Convey("Input: 255", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte("%'>"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 255)
		})
		Convey("Input: 65,535", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte("&7,M"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 65535)
		})
		Convey("Input: 185,194", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte("'$##$"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 185_194)
		})
		Convey("Input: 10,556,100", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte("($##$M"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 10_556_100)
		})
		Convey("Input: 601,692,657", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte(")$###-A"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 601_692_657)
		})
		Convey("Input: 37,256,437,249", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte("*$'W:2>*"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 37_256_437_249)
		})
		Convey("Input: 2,964,996,584,284", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte("+$@<NBW<Q"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 2_964_996_584_284)
		})
		Convey("Input: 123,456,789,012,345", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte(",$)+JIO@IY"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 123_456_789_012_345)
		})
		Convey("Input: 12,345,678,901,234,567", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte("-$XP1F%)'6W"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 12_345_678_901_234_567)
		})
		Convey("Input: 567,890,123,456,789,012", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint64([]byte(".$C:;)<MX02%"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 567_890_123_456_789_012)
		})
	})
}

func ExampleDeserializeUint64() {
	input := []byte(".$C:;)<MX02%-$XP1F%)'6W,$)+JIO@IY")
	offset := tally.UTally(0)

	for int(offset.Cur()) < len(input) {
		fmt.Println(gojasc.DeserializeUint64(input, &offset))
	}

	// Output:
	// 567890123456789012 <nil>
	// 12345678901234567 <nil>
	// 123456789012345 <nil>
}
