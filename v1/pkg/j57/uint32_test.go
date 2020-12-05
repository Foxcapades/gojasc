package j57_test

import (
	"fmt"
	"github.com/foxcapades/gojasc/v1/pkg/j57"
	"github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║     Uint32 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

func TestSerializeUint32(t *testing.T) {
	Convey("SerializeUint32", t, func() {
		Convey("Input: 0", func() {
			So(string(j57.SerializeUint32(0)), ShouldEqual, "#")
		})
		Convey("Input: 1", func() {
			So(string(j57.SerializeUint32(1)), ShouldEqual, "$$")
		})
		Convey("Input: 255", func() {
			So(string(j57.SerializeUint32(255)), ShouldEqual, "%'>")
		})
		Convey("Input: 65,535", func() {
			So(string(j57.SerializeUint32(65_535)), ShouldEqual, "&7,M")
		})
		Convey("Input: 185,194", func() {
			So(string(j57.SerializeUint32(185_194)), ShouldEqual, "'$##$")
		})
		Convey("Input: 10,556,100", func() {
			So(string(j57.SerializeUint32(10_556_100)), ShouldEqual, "($##$M")
		})
		Convey("Input: 4,294,967,295", func() {
			So(string(j57.SerializeUint32(4_294_967_295)), ShouldEqual, ")**TS+;")
		})
	})
}

func TestSerializeUint32Into(t *testing.T) {
	Convey("AppendUint32", t, func() {
		Convey("Input: 0", func() {
			buf := make([]byte, 8)
			off := tally.UTally(0)
			j57.AppendUint32(0, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "#")
		})
		Convey("Input: 1", func() {
			buf := make([]byte, 10)
			off := tally.UTally(0)
			j57.AppendUint32(1, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "$$")
		})
		Convey("Input: 255", func() {
			buf := make([]byte, 10)
			off := tally.UTally(0)
			j57.AppendUint32(255, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "%'>")
		})
		Convey("Input: 65,535", func() {
			buf := make([]byte, 10)
			off := tally.UTally(0)
			j57.AppendUint32(65_535, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "&7,M")
		})
		Convey("Input: 185,194", func() {
			buf := make([]byte, 10)
			off := tally.UTally(0)
			j57.AppendUint32(185_194, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "'$##$")
		})
		Convey("Input: 10,556,100", func() {
			buf := make([]byte, 10)
			off := tally.UTally(0)
			j57.AppendUint32(10_556_100, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "($##$M")
		})
		Convey("Input: 4,294,967,295", func() {
			buf := make([]byte, 10)
			off := tally.UTally(0)
			j57.AppendUint32(4_294_967_295, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, ")**TS+;")
		})

		Convey("Repeated use", func() {
			buf := make([]byte, 6)
			off := tally.UTally(0)
			j57.AppendUint32(1, buf, &off)
			j57.AppendUint32(2, buf, &off)
			j57.AppendUint32(3, buf, &off)

			So(buf, ShouldResemble, []byte("$$$%$&"))
		})
	})
}

func TestDeserializeUint32(t *testing.T) {
	Convey("DeserializeUint32", t, func() {
		Convey("Input: 0", func() {
			off := tally.UTally(0)
			out, err := j57.DeserializeUint32([]byte("#"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 0)
		})
		Convey("Input: 1", func() {
			off := tally.UTally(0)
			out, err := j57.DeserializeUint32([]byte("$$"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 1)
		})
		Convey("Input: 255", func() {
			off := tally.UTally(0)
			out, err := j57.DeserializeUint32([]byte("%'>"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 255)
		})
		Convey("Input: 65,535", func() {
			off := tally.UTally(0)
			out, err := j57.DeserializeUint32([]byte("&7,M"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 65535)
		})
		Convey("Input: 185,194", func() {
			off := tally.UTally(0)
			out, err := j57.DeserializeUint32([]byte("'$##$"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 185_194)
		})
		Convey("Input: 10,556,100", func() {
			off := tally.UTally(0)
			out, err := j57.DeserializeUint32([]byte("($##$M"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, 10_556_100)
		})
		Convey("Input: 4,294,967,295", func() {
			off := tally.UTally(0)
			out, err := j57.DeserializeUint32([]byte(")**TS+;"), &off)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, uint32(4294967295))
		})
	})
}

func ExampleDeserializeUint32() {
	input := []byte("'$##$)**TS+;($##$M")
	offset := tally.UTally(0)

	for int(offset.Cur()) < len(input) {
		fmt.Println(j57.DeserializeUint32(input, &offset))
	}

	// Output:
	// 185194 <nil>
	// 4294967295 <nil>
	// 10556100 <nil>
}
