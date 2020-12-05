package j57_test

import (
	"fmt"
	"github.com/foxcapades/gojasc/v1/pkg/j57"
	"github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestSerializeComplex64(t *testing.T) {
	tests := [...]struct {
		i complex64
		o uint64
	}{
		{
			i: complex(math.MaxFloat32, math.MaxFloat32),
			o: 0x7F7F_FFFF_7F7F_FFFF,
		},
		{
			i: complex(1, math.MaxFloat32),
			o: 0x7F7F_FFFF_3F80_0000,
		},
		{
			i: complex(math.MaxFloat32, 1),
			o: 0x3F80_0000_7F7F_FFFF,
		},
		{
			i: 0,
			o: 0,
		},
		{
			i: complex(math.SmallestNonzeroFloat32, 1),
			o: 0x3F80_0000_0000_0001,
		},
		{
			i: complex(1, math.SmallestNonzeroFloat32),
			o: 0x1_3F80_0000,
		},
		{
			i: complex(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32),
			o: 0x1_0000_0001,
		},
	}

	Convey("SerializeComplex64", t, func() {
		for _, test := range tests {
			check := j57.SerializeUint64(test.o)
			Convey(fmt.Sprintf("(%f) -> %s", test.i, string(check)), func() {
				So(j57.SerializeComplex64(test.i), ShouldResemble, check)
			})
		}
	})
}

func TestAppendComplex64(t *testing.T) {
	Convey("AppendComplex64", t, func() {
		tests := [...]struct {
			i complex64
			o uint64
			s int
		}{
			{
				i: complex(math.MaxFloat32, math.MaxFloat32),
				o: 0x7F7F_FFFF_7F7F_FFFF,
				s: 12,
			},
			{
				i: complex(1, math.MaxFloat32),
				o: 0x7F7F_FFFF_3F80_0000,
				s: 12,
			},
			{
				i: complex(math.MaxFloat32, 1),
				o: 0x3F80_0000_7F7F_FFFF,
				s: 12,
			},
			{
				i: 0,
				o: 0,
				s: 1,
			},
			{
				i: complex(math.SmallestNonzeroFloat32, 1),
				o: 0x3F80_0000_0000_0001,
				s: 12,
			},
			{
				i: complex(1, math.SmallestNonzeroFloat32),
				o: 0x1_3F80_0000,
				s: 7,
			},
			{
				i: complex(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32),
				o: 0x1_0000_0001,
				s: 7,
			},
		}

		for _, test := range tests {
			check := j57.SerializeUint64(test.o)
			buff := make([]byte, test.s)
			off := tally.UTally(0)

			Convey(fmt.Sprintf("(%f) -> %s", test.i, string(check)), func() {
				So(j57.AppendComplex64(test.i, buff, &off), ShouldEqual, test.s)
				So(buff, ShouldResemble, check)
				So(off, ShouldEqual, test.s)
			})
		}
	})
}
