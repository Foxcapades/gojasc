package j57_test

import (
	"fmt"
	"github.com/drhodes/golorem"
	"github.com/foxcapades/gojasc/v1/pkg/j57"
	"github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStringSerialization(t *testing.T) {
	tests := [...]string{
		"hello",
		"goodbye",
		"ohio is for lovers",
		"hey man, nice shot",
		"what a horrible night to have a curse",
		"removal of the oaken stake",
		"rabid wolves in sheep's clothing",
		"❤🧡💛💚💙💜",
		"⁶⁶⁶☠🎃",
		lorem.Paragraph(60, 64),
	}
	Convey("String (de)serialization", t, func() {
		for i := range tests {
			name := tests[i]
			if len(name) > 32 {
				name = "<long string omitted>"
			}

			Convey(fmt.Sprintf("I/O with SerializeString(%s)", tests[i]), func() {
				tmp := j57.SerializeString(tests[i])
				tal := tally.UTally(0)
				out, err := j57.DeserializeString(tmp, &tal)

				So(err, ShouldBeNil)
				So(out, ShouldEqual, tests[i])
			})
		}
	})
}
