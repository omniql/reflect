package local

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

)

func Test_Load_App(t *testing.T) {
	Convey("Scalar app", t, func() {

		app, err := LoadApplication("fixtures/scalar-app")
		So(err, ShouldBeNil)
		So(app.Version(), ShouldEqual, "v1")
		So(app.LookupResources().ResourceCount(), ShouldEqual, 1)
		r, _ := app.LookupResources().ResourceByName("scalar_table")
		So(r.Name(), ShouldEqual, "ScalarTable")
		rt := r.Table()
		So(rt.LookupFields().FieldCount(), ShouldEqual, 28)


	})
}

