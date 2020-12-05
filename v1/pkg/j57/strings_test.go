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
			if len(name) > 64 {
				name = "<long string omitted>"
			}

			Convey(fmt.Sprintf("I/O with SerializeBytes(%s)", name), func() {
				tmp := j57.SerializeString(tests[i])
				tal := tally.UTally(0)
				out, err := j57.DeserializeString(tmp, &tal)

				So(len(tmp), ShouldBeGreaterThanOrEqualTo, j57.SizeString(tests[i]))
				So(len(tmp)+int(float64(len(tmp))*0.11), ShouldBeGreaterThan, j57.SizeString(tests[i]))

				So(err, ShouldBeNil)
				So(out, ShouldEqual, tests[i])
			})
		}
	})
}
