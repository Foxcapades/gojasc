package j57_test

import (
	"fmt"
	. "github.com/foxcapades/gojasc/v1/pkg/j57"
	. "github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"math/bits"
	"testing"
)

func TestSerializeInt(t *testing.T) {
	type tt struct {
		i    int64
		o    []byte
		size int
	}

	tests := []tt{
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
	}

	if bits.UintSize == 64 {
		tests = append(tests,
			tt{
				i:    -128,
				o:    []byte{46, 85, 89, 53, 82, 62, 75, 53, 47, 55, 44, 76},
				size: 12,
			},
			tt{
				i:    -32768,
				o:    []byte{46, 85, 89, 53, 82, 62, 75, 53, 47, 45, 42, 40},
				size: 12,
			},
			tt{
				i:    -2147483648,
				o:    []byte{46, 85, 89, 53, 82, 62, 71, 77, 79, 59, 71, 49},
				size: 12,
			},
			tt{
				i:    9223372036854775807,
				o:    []byte{46, 60, 62, 44, 58, 77, 55, 44, 41, 45, 40, 90},
				size: 12,
			},
			tt{
				i:    -9223372036854775808,
				o:    []byte{46, 60, 62, 44, 58, 77, 55, 44, 41, 45, 40, 91},
				size: 12,
			},
		)
	} else {
		tests = append(tests,
			tt{
				i:    -128,
				o:    []byte{41, 42, 42, 84, 83, 41, 46},
				size: 7,
			},
			tt{
				i:    -32768,
				o:    []byte{41, 42, 42, 84, 73, 38, 67},
				size: 7,
			},
			tt{
				i:    -2147483648,
				o:    []byte{41, 38, 67, 59, 87, 67, 76},
				size: 7,
			},
		)
	}

	Convey("SerializeInt", t, func() {
		buf := [32]byte{}
		for _, test := range tests {
			Convey(fmt.Sprintf("(%d) -> %s", test.i, string(test.o)), func() {
				var tally UTally
				So(AppendInt(int(test.i), buf[:], &tally), ShouldResemble, test.size)
			})
		}
	})
}

func TestSerializeIntInto(t *testing.T) {
	type tt struct {
		i int64
		o []byte
		s int
	}

	var tests []tt

	Convey("AppendInt", t, func() {
		if bits.UintSize == 32 {
			tests = []tt{
				{i: 2147483647, o: []byte{41, 38, 67, 59, 87, 67, 75}, s: 7},
				{i: 32767, o: []byte{39, 45, 39, 84}, s: 4},
				{i: 127, o: []byte{37, 37, 48}, s: 3},
				{o: []byte{35}, s: 1},
				{i: -128, o: []byte{41, 42, 42, 84, 83, 41, 46}, s: 7},
				{i: -32768, o: []byte{41, 42, 42, 84, 73, 38, 67}, s: 7},
				{i: -2147483648, o: []byte{41, 38, 67, 59, 87, 67, 76}, s: 7},
			}
		} else {
			tests = []tt{
				{
					i: 9223372036854775807,
					o: []byte{46, 60, 62, 44, 58, 77, 55, 44, 41, 45, 40, 90},
					s: 12,
				},
				{
					i: 2147483647,
					o: []byte{41, 38, 67, 59, 87, 67, 75},
					s: 7,
				},
				{
					i: 32767,
					o: []byte{38, 45, 39, 84},
					s: 4,
				},
				{
					i: 127,
					o: []byte{37, 37, 48},
					s: 3,
				},
				{
					o: []byte{35},
					s: 1,
				},
				{
					i: -128,
					o: []byte{46, 85, 89, 53, 82, 62, 75, 53, 47, 55, 44, 76},
					s: 12,
				},
				{
					i: -32768,
					o: []byte{46, 85, 89, 53, 82, 62, 75, 53, 47, 45, 42, 40},
					s: 12,
				},
				{
					i: -2147483648,
					o: []byte{46, 85, 89, 53, 82, 62, 71, 77, 79, 59, 71, 49},
					s: 12,
				},
				{
					i: -9223372036854775808,
					o: []byte{46, 60, 62, 44, 58, 77, 55, 44, 41, 45, 40, 91},
					s: 12,
				},
			}
		}

		buf := [32]byte{}
		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.i), func() {
				var tally UTally

				So(AppendInt(int(test.i), buf[:], &tally), ShouldResemble, test.s)
			})
		}
	})
}

func TestDeserializeInt(t *testing.T) {
	type tt struct {
		i []byte
		o int64
	}

	var tests []tt

	Convey("DeserializeInt", t, func() {
		if bits.UintSize == 32 {
			tests = []tt{
				{i: []byte{41, 38, 67, 59, 87, 67, 75}, o: 2147483647},
				{i: []byte{38, 45, 39, 84}, o: 32767},
				{i: []byte{37, 37, 48}, o: 127},
				{i: []byte{35}},
				{i: []byte{41, 42, 42, 84, 83, 41, 46}, o: -128},
				{i: []byte{41, 42, 42, 84, 73, 38, 67}, o: -32768},
				{i: []byte{41, 38, 67, 59, 87, 67, 76}, o: -2147483648},
			}
		} else {
			tests = []tt{
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
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.o), func() {
				var tally UTally

				t, e := DeserializeInt(test.i, &tally)

				So(e, ShouldBeNil)
				So(t, ShouldEqual, test.o)
			})
		}
	})
}

func TestSizeInt(t *testing.T) {
	Convey("SizeInt", t, func() {
		tests := [...]struct {
			i int
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
				So(SizeInt(test.i), ShouldEqual, test.o)
			})
		}
	})
}
