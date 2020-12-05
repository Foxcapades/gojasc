package j57_test

import (
	"fmt"
	. "github.com/foxcapades/gojasc/v1/pkg/j57"
	. "github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"math/bits"
	"testing"
)

func TestSerializeUint(t *testing.T) {
	type tt struct {
		i uint64
		o []byte
		s int
	}

	tests := []tt{
		{
			o: []byte{35},
			s: 1,
		},
		{
			i: math.MaxUint8,
			o: []byte{37, 39, 62},
			s: 3,
		},
		{
			i: math.MaxUint16,
			o: []byte{38, 55, 44, 77},
			s: 4,
		},
		{
			i: math.MaxUint32,
			o: []byte{41, 42, 42, 84, 83, 43, 59},
			s: 7,
		},
	}

	if bits.UintSize == 64 {
		tests = append(tests,
			tt{
				i: math.MaxUint64,
				o: []byte{46, 50, 89, 53, 82, 62, 75, 53, 47, 55, 46, 89},
				s: 12,
			},
		)
	}

	Convey("SerializeUint", t, func() {
		buf := [32]byte{}
		for _, test := range tests {
			Convey(fmt.Sprintf("(%d) -> %s", test.i, string(test.o)), func() {
				var tally UTally
				So(AppendUint(uint(test.i), buf[:], &tally), ShouldResemble, test.s)
			})
		}
	})
}

func TestAppendUint(t *testing.T) {
	type tt struct {
		i uint64
		o []byte
		s int
	}

	var tests []tt

	Convey("AppendUint", t, func() {
		tests = []tt{
			{
				o: []byte{35},
				s: 1,
			},
			{
				i: math.MaxUint8,
				o: []byte{37, 39, 62},
				s: 3,
			},
			{
				i: math.MaxUint16,
				o: []byte{38, 55, 44, 77},
				s: 4,
			},
			{
				i: math.MaxUint32,
				o: []byte{41, 42, 42, 84, 83, 43, 59},
				s: 7,
			},
		}

		if bits.UintSize == 64 {
			tests = append(tests,
				tt{
					i: math.MaxUint64,
					o: []byte{46, 50, 89, 53, 82, 62, 75, 53, 47, 55, 46, 89},
					s: 12,
				},
			)
		}

		buf := [32]byte{}
		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.i), func() {
				var tally UTally

				So(AppendUint(uint(test.i), buf[:], &tally), ShouldResemble, test.s)
			})
		}
	})
}

func TestDeserializeUint(t *testing.T) {
	type tt struct {
		i []byte
		o uint64
	}

	var tests []tt

	Convey("DeserializeUint", t, func() {
		tests = []tt{
			{
				i: []byte{35},
			},
			{
				o: math.MaxUint8,
				i: []byte{37, 39, 62},
			},
			{
				o: math.MaxUint16,
				i: []byte{38, 55, 44, 77},
			},
			{
				o: math.MaxUint32,
				i: []byte{41, 42, 42, 84, 83, 43, 59},
			},
		}

		if bits.UintSize == 64 {
			tests = append(tests,
				tt{
					o: math.MaxUint64,
					i: []byte{46, 85, 89, 53, 82, 62, 75, 53, 47, 55, 46, 89},
				},
			)
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.o), func() {
				var tally UTally

				t, e := DeserializeUint(test.i, &tally)

				So(e, ShouldBeNil)
				So(t, ShouldEqual, test.o)
			})
		}
	})
}

func TestSizeUint(t *testing.T) {
	type tt struct {
		i uint64
		o int
	}

	Convey("SizeUint", t, func() {
		tests := []tt {
			// Note: these are multiples of 57
			{0, 1},
			{1, 2},
			{57, 3},
			{3_249, 4},
			{185_193, 5},
			{10_556_001, 6},
			{601_692_057, 7},
		}
		if bits.UintSize == 64 {
			tests = append(tests,
				tt{34_296_447_249, 8},
				tt{1_954_897_493_193, 9},
				tt{111_429_157_112_001, 10},
				tt{6_351_461_955_384_057, 11},
				tt{362_033_331_456_891_249, 12},
			)
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%d) -> %d", test.i, test.o), func() {
				So(SizeUint(uint(test.i)), ShouldEqual, test.o)
			})
		}
	})
}
