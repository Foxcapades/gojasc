package j57_test

import (
	"fmt"
	. "github.com/foxcapades/gojasc/v1/pkg/j57"
	. "github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// -128        = 18446744073709551488
// -32768      = 18446744073709518848
// -2147483648 = 18446744071562067968

func TestSerializeInt64(t *testing.T) {
	Convey("SerializeInt64", t, func() {
		tests := [...]struct {
			i int64
			o []byte
		}{
			{i: 9223372036854775807, o: []byte{46, 60, 62, 44, 58, 77, 55, 44, 41, 45, 40, 90}},
			{i: 2147483647, o: []byte{41, 38, 67, 59, 87, 67, 75}},
			{i: 32767, o: []byte{38, 45, 39, 84}},
			{i: 127, o: []byte{37, 37, 48}},
			{o: []byte{35}},
			{i: -128, o: []byte{46, 85, 89, 53, 82, 62, 75, 53, 47, 55, 44, 76}},
			{i: -32768, o: []byte{46, 85, 89, 53, 82, 62, 75, 53, 47, 45, 42, 40}},
			{i: -2147483648, o: []byte{46, 85, 89, 53, 82, 62, 71, 77, 79, 59, 71, 49}},
			{i: -9223372036854775808, o: []byte{46, 60, 62, 44, 58, 77, 55, 44, 41, 45, 40, 91}},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.i), func() {
				So(SerializeInt64(test.i), ShouldResemble, test.o)
			})
		}
	})
}

func TestAppendInt64(t *testing.T) {
	Convey("AppendInt64", t, func() {
		tests := [...]struct {
			i    int64
			o    []byte
			size int
		}{
			{
				i:    9223372036854775807,
				o:    []byte{46, 60, 62, 44, 58, 77, 55, 44, 41, 45, 40, 90},
				size: 12,
			},
			{
				i:    2147483647,
				o:    []byte{41, 38, 67, 59, 87, 67, 75},
				size: 7,
			},
			{
				i:    32767,
				o:    []byte{38, 45, 39, 84},
				size: 4,
			},
			{
				i:    127,
				o:    []byte{37, 37, 48},
				size: 3,
			},
			{
				o:    []byte{35},
				size: 1,
			},
			{
				i:    -128,
				o:    []byte{46, 85, 89, 53, 82, 62, 75, 53, 47, 55, 44, 76},
				size: 12,
			},
			{
				i:    -32768,
				o:    []byte{46, 85, 89, 53, 82, 62, 75, 53, 47, 45, 42, 40},
				size: 12,
			},
			{
				i:    -2147483648,
				o:    []byte{46, 85, 89, 53, 82, 62, 71, 77, 79, 59, 71, 49},
				size: 12,
			},
			{
				i:    -9223372036854775808,
				o:    []byte{46, 60, 62, 44, 58, 77, 55, 44, 41, 45, 40, 91},
				size: 12,
			},
		}

		buf := [32]byte{}
		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.i), func() {
				var tally UTally

				So(AppendInt64(test.i, buf[:], &tally), ShouldResemble, test.size)
			})
		}
	})
}

func TestDeserializeInt64(t *testing.T) {
	Convey("DeserializeInt64", t, func() {
		tests := [...]struct {
			o int64
			i []byte
		}{
			{i: []byte{46, 60, 62, 44, 58, 77, 55, 44, 41, 45, 40, 90}, o: 9223372036854775807},
			{i: []byte{41, 38, 67, 59, 87, 67, 75}, o: 2147483647},
			{i: []byte{38, 45, 39, 84}, o: 32767},
			{i: []byte{37, 37, 48}, o: 127},
			{i: []byte{35}},
			{i: []byte{46, 85, 89, 53, 82, 62, 75, 53, 47, 55, 44, 76}, o: -128},
			{i: []byte{46, 85, 89, 53, 82, 62, 75, 53, 47, 45, 42, 40}, o: -32768},
			{i: []byte{46, 85, 89, 53, 82, 62, 71, 77, 79, 59, 71, 49}, o: -2147483648},
			{i: []byte{46, 60, 62, 44, 58, 77, 55, 44, 41, 45, 40, 91}, o: -9223372036854775808},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.o), func() {
				var tally UTally

				t, e := DeserializeInt64(test.i, &tally)

				So(e, ShouldBeNil)
				So(t, ShouldEqual, test.o)
			})
		}
	})
}

func TestSizeInt64(t *testing.T) {
	Convey("SizeInt64", t, func() {
		tests := [...]struct {
			i int64
			o int
		}{
			// Note: these are multiples of 57
			{0, 1},
			{1, 2},
			{57, 3},
			{3_249, 4},
			{185_193, 5},
			{10_556_001, 6},
			{601_692_057, 7},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%d) -> %d", test.i, test.o), func() {
				So(SizeInt64(test.i), ShouldEqual, test.o)
			})
		}
	})
}
