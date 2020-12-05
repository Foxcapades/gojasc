package j57_test

import (
	"errors"
	"github.com/foxcapades/gojasc/v1/pkg/j57"
	"github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSerializeBool(t *testing.T) {
	Convey("SerializeBool", t, func() {
		So(j57.SerializeBool(true), ShouldResemble, []byte{36, 36})
		So(j57.SerializeBool(false), ShouldResemble, []byte{35})
	})
}

func TestDeserializeBool(t *testing.T) {
	Convey("DeserializeBool", t, func() {
		off := tally.UTally(0)
		a, b := j57.DeserializeBool([]byte{36, 36}, &off)
		So(b, ShouldBeNil)
		So(a, ShouldBeTrue)

		off = 0
		a, b = j57.DeserializeBool([]byte{35}, &off)
		So(b, ShouldBeNil)
		So(a, ShouldBeFalse)

		off = 0
		a, b = j57.DeserializeBool([]byte{36, 38}, &off)
		So(errors.Is(b, j57.ErrInvalidJASCFormat), ShouldBeTrue)
		So(a, ShouldBeFalse)
	})
}

func TestAppendBool(t *testing.T) {
	Convey("AppendBool", t, func() {
		buf := make([]byte, 2)
		off := tally.UTally(0)

		So(j57.AppendBool(true, buf, &off), ShouldEqual, 2)
		So(buf, ShouldResemble, []byte{36, 36})
		So(off, ShouldEqual, 2)

		off = 0
		So(j57.AppendBool(false, buf, &off), ShouldEqual, 1)
		So(buf[:1], ShouldResemble, []byte{35})
		So(off, ShouldEqual, 1)
	})
}

func TestSizeBool(t *testing.T) {
	Convey("SizeBool", t, func() {
		So(j57.SizeBool(true), ShouldEqual, 2)
		So(j57.SizeBool(false), ShouldEqual, 1)
	})
}
