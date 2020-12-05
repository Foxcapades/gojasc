package j57_test

import (
	"fmt"
	"github.com/foxcapades/gojasc/v1/pkg/j57"
	"github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestSerializeComplex128(t *testing.T) {
	tests := [...]struct {
		i complex128
		o1 uint64
		o2 uint64
	}{
		{
			i: complex(math.MaxFloat64, math.MaxFloat64),
			o1: 0x7FEF_FFFF_FFFF_FFFF,
			o2: 0x7FEF_FFFF_FFFF_FFFF,
		},
		{
			i: complex(math.MaxFloat32, math.MaxFloat32),
			o1: 0x47EF_FFFF_E000_0000,
			o2: 0x47EF_FFFF_E000_0000,
		},
		{
			i: complex(1, math.MaxFloat32),
			o1: 0x3FF0_0000_0000_0000,
			o2: 0x47EF_FFFF_E000_0000,
		},
		{
			i: complex(math.MaxFloat32, 1),
			o1: 0x47EF_FFFF_E000_0000,
			o2: 0x3FF0_0000_0000_0000,
		},
		{
			i: 0,
			o1: 0,
			o2: 0,
		},
		{
			i: complex(math.SmallestNonzeroFloat32, 1),
			o1: 0x36A0_0000_0000_0000,
			o2: 0x3FF0_0000_0000_0000,
		},
		{
			i: complex(1, math.SmallestNonzeroFloat32),
			o1: 0x3FF0_0000_0000_0000,
			o2: 0x36A0_0000_0000_0000,
		},
		{
			i: complex(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32),
			o1: 0x36A0_0000_0000_0000,
			o2: 0x36A0_0000_0000_0000,
		},
		{
			i: complex(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64),
			o1: 1,
			o2: 1,
		},
	}

	Convey("SerializeComplex64", t, func() {
		for _, test := range tests {

			check := append(j57.SerializeUint64(test.o1), j57.SerializeUint64(test.o2)...)
			Convey(fmt.Sprintf("(%f) -> %s", test.i, string(check)), func() {
				So(j57.SerializeComplex128(test.i), ShouldResemble, check)
			})
		}
	})
}

func TestAppendComplex128(t *testing.T) {
	Convey("AppendComplex64", t, func() {
		tests := [...]struct {
			i complex128
			o1 uint64
			o2 uint64
			s int
		}{
			{
				i: complex(math.MaxFloat64, math.MaxFloat64),
				o1: 0x7FEF_FFFF_FFFF_FFFF,
				o2: 0x7FEF_FFFF_FFFF_FFFF,
				s: 24,
			},
			{
				i: complex(math.MaxFloat32, math.MaxFloat32),
				o1: 0x47EF_FFFF_E000_0000,
				o2: 0x47EF_FFFF_E000_0000,
				s: 24,
			},
			{
				i: complex(1, math.MaxFloat32),
				o1: 0x3FF0_0000_0000_0000,
				o2: 0x47EF_FFFF_E000_0000,
				s: 24,
			},
			{
				i: complex(math.MaxFloat32, 1),
				o1: 0x47EF_FFFF_E000_0000,
				o2: 0x3FF0_0000_0000_0000,
				s: 24,
			},
			{
				o1: 0,
				o2: 0,
				s: 2,
			},
			{
				i: complex(math.SmallestNonzeroFloat32, 1),
				o1: 0x36A0_0000_0000_0000,
				o2: 0x3FF0_0000_0000_0000,
				s: 24,
			},
			{
				i: complex(1, math.SmallestNonzeroFloat32),
				o1: 0x3FF0_0000_0000_0000,
				o2: 0x36A0_0000_0000_0000,
				s: 24,
			},
			{
				i: complex(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32),
				o1: 0x36A0_0000_0000_0000,
				o2: 0x36A0_0000_0000_0000,
				s: 24,
			},
			{
				i: complex(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64),
				o1: 1,
				o2: 1,
				s: 4,
			},
		}

		for _, test := range tests {
			check := append(j57.SerializeUint64(test.o1), j57.SerializeUint64(test.o2)...)
			buff := make([]byte, test.s)
			off := tally.UTally(0)

			Convey(fmt.Sprintf("(%f) -> %s", test.i, string(check)), func() {
				So(j57.AppendComplex128(test.i, buff, &off), ShouldEqual, test.s)
				So(buff, ShouldResemble, check)
				So(off, ShouldEqual, test.s)
			})
		}
	})
}
