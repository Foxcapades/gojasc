package j57_test

import (
	. "github.com/foxcapades/gojasc/v1/pkg/j57"
	tally2 "github.com/foxcapades/tally-go/v1/tally"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestSerializeFloat64(t *testing.T) {
	Convey("SerializeFloat64", t, func() {
		So(SerializeFloat64(math.MaxFloat64), ShouldResemble, SerializeUint64(9_218_868_437_227_405_311))
		So(SerializeFloat64(math.MaxFloat32), ShouldResemble, SerializeUint64(5_183_643_170_566_569_984))
		So(SerializeFloat64(math.SmallestNonzeroFloat64), ShouldResemble, SerializeUint64(1))
		So(SerializeFloat64(0), ShouldResemble, SerializeUint64(0))
	})
}

func TestDeserializeFloat64(t *testing.T) {
	Convey("DeserializeFloat64", t, func() {
		var tally tally2.UTally

		t1, e := DeserializeFloat64(SerializeUint64(9_218_868_437_227_405_311), &tally)
		So(e, ShouldBeNil)
		So(t1, ShouldEqual, math.MaxFloat64)

		tally = 0
		t2, e := DeserializeFloat64(SerializeUint64(5_183_643_170_566_569_984), &tally)
		So(e, ShouldBeNil)
		So(t2, ShouldEqual, math.MaxFloat32)

		tally = 0
		t3, e := DeserializeFloat64(SerializeUint64(1), &tally)
		So(e, ShouldBeNil)
		So(t3, ShouldEqual, math.SmallestNonzeroFloat64)

		tally = 0
		t4, e := DeserializeFloat64(SerializeUint64(0), &tally)
		So(e, ShouldBeNil)
		So(t4, ShouldEqual, 0)
	})
}
