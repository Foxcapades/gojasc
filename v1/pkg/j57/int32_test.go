package j57_test

import (
	"fmt"
	. "github.com/foxcapades/gojasc/v1/pkg/j57"
	. "github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSerializeInt32(t *testing.T) {
	Convey("SerializeInt32", t, func() {
		tests := [...]struct {
			i int32
			o []byte
		}{
			{i: 2147483647, o: []byte{41, 38, 67, 59, 87, 67, 75}},
			{i: 32767, o: []byte{38, 45, 39, 84}},
			{i: 127, o: []byte{37, 37, 48}},
			{o: []byte{35}},
			{i: -128, o: []byte{41, 42, 42, 84, 83, 41, 46}},
			{i: -32768, o: []byte{41, 42, 42, 84, 73, 38, 67}},
			{i: -2147483648, o: []byte{41, 38, 67, 59, 87, 67, 76}},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.i), func() {
				So(SerializeInt32(test.i), ShouldResemble, test.o)
			})
		}
	})
}

func TestSerializeInt32Into(t *testing.T) {
	Convey("AppendInt32", t, func() {
		tests := [...]struct {
			i    int32
			o    []byte
			size int
		}{
			{
				i:    2147483647,
				o:    []byte{41, 38, 67, 59, 87, 67, 75},
				size: 7,
			},
			{
				i:    32767,
				o:    []byte{39, 45, 39, 84},
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
				o:    []byte{41, 42, 42, 84, 83, 41, 46},
				size: 7,
			},
			{
				i:    -32768,
				o:    []byte{41, 42, 42, 84, 73, 38, 67},
				size: 7,
			},
			{
				i:    -2147483648,
				o:    []byte{41, 38, 67, 59, 87, 67, 76},
				size: 7,
			},
		}

		buf := [32]byte{}
		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.i), func() {
				var tally UTally

				So(AppendInt32(test.i, buf[:], &tally), ShouldResemble, test.size)
			})
		}
	})
}

func TestDeserializeInt32(t *testing.T) {
	Convey("DeserializeInt32", t, func() {
		tests := [...]struct {
			o int32
			i []byte
		}{
			{i: []byte{41, 38, 67, 59, 87, 67, 75}, o: 2147483647},
			{i: []byte{38, 45, 39, 84}, o: 32767},
			{i: []byte{37, 37, 48}, o: 127},
			{i: []byte{35}},
			{i: []byte{41, 42, 42, 84, 83, 41, 46}, o: -128},
			{i: []byte{41, 42, 42, 84, 73, 38, 67}, o: -32768},
			{i: []byte{41, 38, 67, 59, 87, 67, 76}, o: -2147483648},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.o), func() {
				var tally UTally

				t, e := DeserializeInt32(test.i, &tally)

				So(e, ShouldBeNil)
				So(t, ShouldEqual, test.o)
			})
		}
	})
}

func TestSizeInt32(t *testing.T) {
	Convey("SizeInt32", t, func() {
		tests := [...]struct {
			i int32
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
				So(SizeInt32(test.i), ShouldEqual, test.o)
			})
		}
	})
}
