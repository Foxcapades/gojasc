package gojasc_test

import (
	"fmt"
	"github.com/foxcapades/gojasc/v1/gojasc"
	"github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Uint8 Serialization                                                ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

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

		Convey("Repeated use", func() {
			buf := make([]byte, 6)
			off := tally.UTally(0)
			gojasc.SerializeUint8Into(1, buf, &off)
			gojasc.SerializeUint8Into(2, buf, &off)
			gojasc.SerializeUint8Into(3, buf, &off)

			So(buf, ShouldResemble, []byte("$$$%$&"))
		})
	})
}

func TestSerializeUint8Into(t *testing.T) {
	Convey("SerializeUInt8Into", t, func() {
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
				buf := make([]byte, 4)
				off := tally.UTally(0)
				gojasc.SerializeUint8Into(pair.i, buf, &off)
				So(buf[0:len(pair.o)], ShouldResemble, pair.o)
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
				off := tally.UTally(0)
				tmp, err := gojasc.DeserializeUint8(pair.o, &off)
				So(err, ShouldBeNil)
				So(tmp, ShouldEqual, pair.i)
			})
		}
	})
}

func ExampleDeserializeUint8() {
	input := []byte("#$$$%$&$'$($)$*$+")
	offset := tally.UTally(0)

	for int(offset.Cur()) < len(input) {
		fmt.Println(gojasc.DeserializeUint8(input, &offset))
	}

	// Output:
	// 0 <nil>
	// 1 <nil>
	// 2 <nil>
	// 3 <nil>
	// 4 <nil>
	// 5 <nil>
	// 6 <nil>
	// 7 <nil>
	// 8 <nil>
}
