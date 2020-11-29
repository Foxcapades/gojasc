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
// ║     Uint16 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

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

func TestSerializeUint16Into(t *testing.T) {
	Convey("SerializeUint16Into", t, func() {
		Convey("Input: 0", func() {
			buf := make([]byte, 8)
			off := tally.UTally(0)
			gojasc.SerializeUint16Into(0, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "#")
		})
		Convey("Input: 1", func() {
			buf := make([]byte, 8)
			off := tally.UTally(0)
			gojasc.SerializeUint16Into(1, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "$$")
		})
		Convey("Input: 255", func() {
			buf := make([]byte, 8)
			off := tally.UTally(0)
			gojasc.SerializeUint16Into(255, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "%'>")
		})
		Convey("Input: 65,535", func() {
			buf := make([]byte, 8)
			off := tally.UTally(0)
			gojasc.SerializeUint16Into(65535, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "&7,M")
		})

		Convey("Repeated use", func() {
			buf := make([]byte, 6)
			off := tally.UTally(0)
			gojasc.SerializeUint16Into(1, buf, &off)
			gojasc.SerializeUint16Into(2, buf, &off)
			gojasc.SerializeUint16Into(3, buf, &off)

			So(buf, ShouldResemble, []byte("$$$%$&"))
		})
	})
}

func TestDeserializeUint16(t *testing.T) {
	Convey("DeserializeUint16", t, func() {
		Convey("Input: 0", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint16([]byte("#"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 0)
		})
		Convey("Input: 1", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint16([]byte("$$"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 1)
		})
		Convey("Input: 255", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint16([]byte("%'>"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 255)
		})
		Convey("Input: 65,535", func() {
			off := tally.UTally(0)
			out, err := gojasc.DeserializeUint16([]byte("&7,M"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 65535)
		})
	})
}

func ExampleDeserializeUint16() {
	input := []byte("$$%'>&7,M")
	offset := tally.UTally(0)

	for int(offset.Cur()) < len(input) {
		fmt.Println(gojasc.DeserializeUint32(input, &offset))
	}

	// Output:
	// 1 <nil>
	// 255 <nil>
	// 65535 <nil>
}
