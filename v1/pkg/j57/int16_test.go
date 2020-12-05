package j57_test

import (
	"fmt"
	. "github.com/foxcapades/gojasc/v1/pkg/j57"
	. "github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSerializeInt16(t *testing.T) {
	Convey("SerializeInt16", t, func() {
		tests := [...]struct {
			i int16
			o []byte
		}{
			{i: 32767, o: []byte{38, 45, 39, 84}},
			{i: 127, o: []byte{37, 37, 48}},
			{o: []byte{35}},
			{i: -128, o: []byte{38, 55, 42, 64}},
			{i: -32768, o: []byte{38, 45, 39, 85}},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.i), func() {
				So(SerializeInt16(test.i), ShouldResemble, test.o)
			})
		}
	})
}

func TestSerializeInt16Into(t *testing.T) {
	Convey("AppendInt16", t, func() {
		tests := [...]struct {
			input  int16
			output []byte
		}{
			{input: 32767, output: []byte{39, 45, 39, 84}},
			{input: 127, output: []byte{37, 37, 48}},
			{output: []byte{35}},
			{input: -128, output: []byte{38, 55, 42, 64}},
			{input: -32768, output: []byte{39, 45, 39, 85}},
		}

		buf := [32]byte{}
		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.input), func() {
				var tally UTally

				So(AppendInt16(test.input, buf[:], &tally), ShouldResemble, len(test.output))
			})
		}
	})
}

func TestDeserializeInt16(t *testing.T) {
	Convey("DeserializeInt16", t, func() {
		tests := [...]struct {
			o int16
			i []byte
		}{
			{i: []byte{37, 37, 48}, o: 127},
			{i: []byte{35}},
			{i: []byte{38, 55, 42, 64}, o: -128},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.o), func() {
				var tally UTally

				t, e := DeserializeInt16(test.i, &tally)

				So(e, ShouldBeNil)
				So(t, ShouldEqual, test.o)
			})
		}
	})
}

func TestSizeInt16(t *testing.T) {
	Convey("SizeInt16", t, func() {
		tests := [...]struct {
			i int16
			o int
		}{
			{0, 1},
			{1, 2},
			{57, 3},
			{3249, 4},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%d) -> %d", test.i, test.o), func() {
				So(SizeInt16(test.i), ShouldEqual, test.o)
			})
		}
	})
}
