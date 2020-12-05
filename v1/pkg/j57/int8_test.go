package j57_test

import (
	"fmt"
	. "github.com/foxcapades/gojasc/v1/pkg/j57"
	. "github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSerializeInt8(t *testing.T) {
	Convey("SerializeInt8", t, func() {
		tests := [...]struct {
			i int8
			o []byte
		}{
			{i: 127, o: []byte{37, 37, 48}},
			{o: []byte{35}},
			{i: -128, o: []byte{37, 37, 49}},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.i), func() {
				So(SerializeInt8(test.i), ShouldResemble, test.o)
			})
		}
	})
}

func TestSerializeInt8Into(t *testing.T) {
	Convey("AppendInt8", t, func() {
		tests := [...]struct {
			input  int8
			output []byte
			size   int
		}{
			{input: 127, output: []byte{37, 37, 48}, size: 3},
			{output: []byte{35}, size: 1},
			{input: -128, output: []byte{37, 37, 49}, size: 3},
		}

		buf := [32]byte{}
		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.input), func() {
				var tally UTally

				So(AppendInt8(test.input, buf[:], &tally), ShouldResemble, len(test.output))
			})
		}
	})
}

func TestDeserializeInt8(t *testing.T) {
	Convey("DeserializeInt8", t, func() {
		tests := [...]struct {
			o int8
			i []byte
		}{
			{i: []byte{37, 37, 48}, o: 127},
			{i: []byte{35}},
			{i: []byte{37, 37, 49}, o: -128},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("input: %d", test.o), func() {
				var tally UTally

				t, e := DeserializeInt8(test.i, &tally)

				So(e, ShouldBeNil)
				So(t, ShouldEqual, test.o)
			})
		}
	})
}

func TestSizeInt8(t *testing.T) {
	Convey("SizeInt8", t, func() {
		So(SizeInt8(0), ShouldEqual, 1)
		So(SizeInt8(56), ShouldEqual, 2)
		So(SizeInt8(57), ShouldEqual, 3)
	})
}
