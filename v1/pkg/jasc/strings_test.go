package jasc_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/foxcapades/gojasc/v1/pkg/jasc"
	"github.com/foxcapades/tally-go/v1/tally"
)

func TestSerializeString(t *testing.T) {
	Convey("SerializeString", t, func() {
		Convey("Empty string", func() {
			input := ""
			out, err := jasc.SerializeString(input)

			So(err, ShouldBeNil)
			So(out, ShouldResemble, []byte{'#'})
		})

		Convey("Single character string", func() {
			input := "!"
			out, err := jasc.SerializeString(input)

			So(err, ShouldBeNil)
			So(out, ShouldResemble, []byte{'$', '&', '!'})
		})

		Convey("Multi-character string", func() {
			input := "Hello world"
			out, err := jasc.SerializeString(input)

			So(err, ShouldBeNil)
			So(out, ShouldResemble, []byte("$0Hello world"))
		})
	})
}

func TestDeserializeString(t *testing.T) {
	Convey("DeserializeString", t, func() {
		Convey("Empty string", func() {
			input := "#"
			offset := tally.UTally(0)
			out, err := jasc.DeserializeString([]byte(input), &offset)

			So(err, ShouldBeNil)
			So(out, ShouldEqual, "")
		})

		Convey("Single character string", func() {
			input := "$&!"
			offset := tally.UTally(0)
			out, err := jasc.DeserializeString([]byte(input), &offset)

			So(err, ShouldBeNil)
			So(out, ShouldResemble, "!")
		})

		Convey("Multi-character string", func() {
			input := "$0Hello world"
			offset := tally.UTally(0)
			out, err := jasc.DeserializeString([]byte(input), &offset)

			So(err, ShouldBeNil)
			So(out, ShouldResemble, "Hello world")
		})

		Convey("Invalid empty string", func() {
			input := ""
			offset := tally.UTally(0)
			out, err := jasc.DeserializeString([]byte(input), &offset)

			So(err, ShouldPointTo, jasc.ErrNoHeader)
			So(out, ShouldEqual, "")
		})
	})
}
