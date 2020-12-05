package j57_test

import (
	"fmt"
	"github.com/drhodes/golorem"
	"github.com/foxcapades/gojasc/v1/pkg/j57"
	"github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestByteSliceSerialization(t *testing.T) {
	tests := [...][]byte{
		[]byte("hello"),
		[]byte("goodbye"),
		[]byte("ohio is for lovers"),
		[]byte("hey man, nice shot"),
		[]byte("what a horrible night to have a curse"),
		[]byte("removal of the oaken stake"),
		[]byte("rabid wolves in sheep's clothing"),
		[]byte("❤🧡💛💚💙💜"),
		[]byte("⁶⁶⁶☠🎃"),
		[]byte(lorem.Paragraph(60, 64)),
	}
	Convey("String (de)serialization", t, func() {
		for i := range tests {
			Convey(fmt.Sprintf("I/O with SerializeBytes(%s)", tests[i]), func() {
				tmp := j57.SerializeBytes(tests[i])

				So(len(tmp), ShouldEqual, j57.SizeBytes(tests[i]))

				tal := tally.UTally(0)
				out, err := j57.DeserializeBytes(tmp, &tal)

				So(err, ShouldBeNil)
				So(out, ShouldResemble, tests[i])
			})

			Convey(fmt.Sprintf("I/O with AppendBytes(%s)", tests[i]), func() {
				buf := make([]byte, j57.SizeBytes(tests[i]))
				off := tally.UTally(0)

				So(j57.AppendBytes(tests[i], buf, &off), ShouldEqual, off)

				off = 0
				out, err := j57.DeserializeBytes(buf, &off)

				So(err, ShouldBeNil)
				So(out, ShouldResemble, tests[i])
			})
		}
	})
}
