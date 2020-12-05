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
// ║     Uint16 Serialization                                               ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

func TestSerializeUint16(t *testing.T) {
	Convey("SerializeUint16", t, func() {
		Convey("Input: 0", func() {
			So(string(j57.SerializeUint16(0)), ShouldEqual, "#")
		})
		Convey("Input: 1", func() {
			So(string(j57.SerializeUint16(1)), ShouldEqual, "$$")
		})
		Convey("Input: 255", func() {
			So(string(j57.SerializeUint16(255)), ShouldEqual, "%'>")
		})
		Convey("Input: 65,535", func() {
			So(string(j57.SerializeUint16(65535)), ShouldEqual, "&7,M")
		})
	})
}

func TestSerializeUint16Into(t *testing.T) {
	Convey("AppendUint16", t, func() {
		Convey("Input: 0", func() {
			buf := make([]byte, 8)
			off := tally.UTally(0)
			j57.AppendUint16(0, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "#")
		})
		Convey("Input: 1", func() {
			buf := make([]byte, 8)
			off := tally.UTally(0)
			j57.AppendUint16(1, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "$$")
		})
		Convey("Input: 255", func() {
			buf := make([]byte, 8)
			off := tally.UTally(0)
			j57.AppendUint16(255, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "%'>")
		})
		Convey("Input: 65,535", func() {
			buf := make([]byte, 8)
			off := tally.UTally(0)
			j57.AppendUint16(65535, buf, &off)
			So(string(buf[:off.Cur()]), ShouldEqual, "&7,M")
		})

		Convey("Repeated use", func() {
			buf := make([]byte, 6)
			off := tally.UTally(0)
			j57.AppendUint16(1, buf, &off)
			j57.AppendUint16(2, buf, &off)
			j57.AppendUint16(3, buf, &off)

			So(buf, ShouldResemble, []byte("$$$%$&"))
		})
	})
}

func TestDeserializeUint16(t *testing.T) {
	tests := [...]struct {
		output uint16
		input  string
	}{
		{0, "#"},
		{10, "$-"},
		{100, "%$N"},
		{1_000, "%4B"},
		{10_000, "'&'<"},
		{2, "$%"},
		{4, "$'"},
		{8, "$+"},
		{16, "$3"},
		{32, "$C"},
		{64, "%$*"},
		{128, "%%1"},
		{256, "%'?"},
		{512, "%+["},
		{1_024, "%4Z"},
		{2_048, "%FX"},
		{4_096, "&$1T"},
		{8_192, "&%@L"},
		{16_384, "&(%<"},
		{32_768, "&-'U"},
		{65_535, "&7,M"},
	}
	Convey("DeserializeUint16", t, func() {
		for _, test := range tests {

			Convey(fmt.Sprintf("(%s) -> %d", test.input, test.output), func() {
				off := tally.UTally(0)
				out, err := j57.DeserializeUint16([]byte(test.input), &off)
				So(err, ShouldBeNil)
				So(out, ShouldEqual, test.output)
			})
		}
	})
}

func ExampleDeserializeUint16() {
	input := []byte("$$%'>&7,M")
	offset := tally.UTally(0)

	for int(offset.Cur()) < len(input) {
		fmt.Println(j57.DeserializeUint32(input, &offset))
	}

	// Output:
	// 1 <nil>
	// 255 <nil>
	// 65535 <nil>
}
