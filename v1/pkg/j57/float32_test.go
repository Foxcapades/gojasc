package j57_test

import (
	. "github.com/foxcapades/gojasc/v1/pkg/j57"
	tally2 "github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestSerializeFloat32(t *testing.T) {
	Convey("SerializeFloat32", t, func() {
		So(SerializeFloat32(math.MaxFloat32), ShouldResemble, SerializeUint32(2_139_095_039))
		So(SerializeFloat32(math.SmallestNonzeroFloat32), ShouldResemble, SerializeUint32(1))
		So(SerializeFloat32(0), ShouldResemble, SerializeUint32(0))
	})
}

func TestDeserializeFloat32(t *testing.T) {
	Convey("DeserializeFloat32", t, func() {
		var tally tally2.UTally

		t1, e := DeserializeFloat32(SerializeUint32(2_139_095_039), &tally)
		So(e, ShouldBeNil)
		So(t1, ShouldEqual, math.MaxFloat32)

		tally = 0
		t2, e := DeserializeFloat32(SerializeUint32(1), &tally)
		So(e, ShouldBeNil)
		So(t2, ShouldEqual, math.SmallestNonzeroFloat32)

		tally = 0
		t3, e := DeserializeFloat32(SerializeUint32(0), &tally)
		So(e, ShouldBeNil)
		So(t3, ShouldEqual, 0)
	})
}
