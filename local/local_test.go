package local

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
)

func Test_Load_App(t *testing.T) {
	Convey("Scalar app", t, func() {
		app, err := LoadApplication("fixtures/scalar-app")
		So(err, ShouldBeNil)
		Convey("Test that app properties are loaded ", func() {
			So(app.Version(), ShouldEqual, "v1")
			So(app.Path(), ShouldEqual, "omniql.almagest.io/omniql/scalar-app")
			So(app.LookupImports().ImportsCount(), ShouldEqual, 2)

			i, ok := app.LookupImports().ImportByPosition(0)
			So(ok, ShouldEqual, true)
			So(i.Alias(), ShouldEqual, "profile")
			So(i.Path(), ShouldEqual, "schema.omniql.almagest.io/omniql/profile")
			So(i.Version(), ShouldEqual, "v3")
			So(i.UsedResourcesCount(), ShouldEqual, 1)

			er := i.UsedResource(0)
			So(er.Application(), ShouldEqual, i)
			So(er.ID(), ShouldEqual, "schema.omniql.almagest.io/omniql/profile/Resource/profile")
			So(er.Name(), ShouldEqual, "profile")

			i, ok = app.LookupImports().ImportByAlias("social")
			So(ok, ShouldEqual, true)
			So(i.Alias(), ShouldEqual, "social")

			_, ok = app.LookupImports().ImportByPosition(50)
			So(ok, ShouldEqual, false)

			_, ok = app.LookupResources().ResourceByPosition(20)
			So(ok, ShouldEqual, false)

			Convey("Check that all the resources were loaded", func() {
				So(app.LookupResources().ResourceCount(), ShouldEqual, 3)
			})

			Convey("Check the ScalarTable resource", func() {
				r, ok := app.LookupResources().ResourceByName("scalar_table")
				So(ok, ShouldEqual, true)

				So(r.Name(), ShouldEqual, "ScalarTable")
				So(r.Application(), ShouldEqual, app)
				So(r.Position(), ShouldEqual, 0)
				So(r.ID(), ShouldEqual, "omniql.almagest.io/omniql/scalar-app/Resource/ScalarTable")

				st := r.Table()
				bty, ok := st.LookupFields().FieldByName("bty")
				So(bty.HybridType(), ShouldEqual, hybrids.Byte)
				vs, _ := st.LookupFields().FieldByName("vs")
				So(vs.HybridType(), ShouldEqual, hybrids.VectorString)
				vr, _ := st.LookupFields().FieldByName("vr")
				So(vr.HybridType(), ShouldEqual, hybrids.VectorResourceID)

				_, ok = st.LookupFields().FieldByPosition(100)
				So(ok, ShouldEqual, false)

				enumField, ok := st.LookupFields().FieldByName("country")
				So(ok, ShouldEqual, true)
				So(enumField.Name(), ShouldEqual, "Country")
				So(enumField.Application(), ShouldEqual, app)
				So(enumField.Position(), ShouldEqual, 26)
				So(enumField.ID(), ShouldEqual, "omniql.almagest.io/omniql/scalar-app/Table/ScalarTable/Field/26")
				So(enumField.ValueType().Kind(), ShouldEqual, reflect.Enumeration)
				So(enumField.ValueType().Enumeration().Name(), ShouldEqual, "Country")
				So(enumField.ValueType().Enumeration().ID(), ShouldEqual, "omniql.almagest.io/omniql/scalar-app/Enumeration/Country")
				So(enumField.ValueType().Enumeration().Application(), ShouldEqual, app)
				So(enumField.ValueType().Enumeration().HybridType(), ShouldEqual, hybrids.Uint16)
				enumValue, ok := enumField.ValueType().Enumeration().Lookup().ByStringToUint16("usa")
				So(ok, ShouldEqual, true)
				So(enumValue, ShouldEqual, 1)
				enumValue8, ok := enumField.ValueType().Enumeration().Lookup().ByStringToUint8("None")
				So(ok, ShouldEqual, true)
				So(enumValue8, ShouldEqual, 0)
				enumValue8, ok = enumField.ValueType().Enumeration().Lookup().ByStringToUint8("WWW")
				So(ok, ShouldEqual, false)

				enumValueStr, ok := enumField.ValueType().Enumeration().Lookup().ByUint16ToString(0)
				So(ok, ShouldEqual, true)
				So(enumValueStr, ShouldEqual, "None")

				_, ok = enumField.ValueType().Enumeration().Lookup().ByUint16ToString(70)
				So(ok, ShouldEqual, false)

				enumValueStr, ok = enumField.ValueType().Enumeration().Lookup().ByUint8ToString(1)
				So(ok, ShouldEqual, true)
				So(enumValueStr, ShouldEqual, "USA")

				_, ok = enumField.ValueType().Enumeration().Lookup().ByUint8ToString(70)
				So(ok, ShouldEqual, false)

				So(enumField.Parent().Table(), ShouldEqual, st)
				So(enumField.HybridType(), ShouldEqual, hybrids.Uint16)
				Convey("Check the struct field on ScalarTable", func() {
					str, ok := st.LookupFields().FieldByName("structChild")
					So(ok, ShouldEqual, true)
					So(str.HybridType(), ShouldEqual, hybrids.Struct)
					So(str.ValueType().Struct().Name(), ShouldEqual, "ScalarStruct")
					So(str.ValueType().Struct().ID(), ShouldEqual, "omniql.almagest.io/omniql/scalar-app/Struct/ScalarStruct")
					So(str.ValueType().Struct().Application(), ShouldEqual, app)
					So(str.ValueType().Struct().Application(), ShouldEqual, app)
					So(str.ValueType().Struct().LookupFields().FieldCount(), ShouldEqual, 11)
					f, ok := str.ValueType().Struct().LookupFields().FieldByPosition(5)
					So(ok, ShouldEqual, true)
					So(f.HybridType(), ShouldEqual, hybrids.Int32)
					So(f.Name(), ShouldEqual, "Int32")
					f, ok = str.ValueType().Struct().LookupFields().FieldByPosition(55)
					So(ok, ShouldEqual, false)

					f, ok = str.ValueType().Struct().LookupFields().FieldByName("uint64")
					So(ok, ShouldEqual, true)
					So(f.HybridType(), ShouldEqual, hybrids.Uint64)
					So(f.Name(), ShouldEqual, "Uint64")

				})

				Convey("Check the vector field on ScalarTable", func() {
					str, ok := st.LookupFields().FieldByPosition(25)
					So(ok, ShouldEqual, true)
					So(str.Name(), ShouldEqual, "Children")
					So(str.HybridType(), ShouldEqual, hybrids.VectorResourceID)
					So(str.Items().HybridType(), ShouldEqual, hybrids.ResourceID)
					So(str.Items().ValueType().Resource(), ShouldEqual, r)

				})

				Convey("Check the resource union field on ScalarTable", func() {
					str, ok := st.LookupFields().FieldByName("runion")
					So(ok, ShouldEqual, true)
					So(str.HybridType(), ShouldEqual, hybrids.Union)
					So(str.ValueType().Kind(), ShouldEqual, reflect.Union)
					So(str.ValueType().Union().Name(), ShouldEqual, "ResourceUnion")
					So(str.ValueType().Union().ID(), ShouldEqual, "omniql.almagest.io/omniql/scalar-app/Union/ResourceUnion")
					So(str.ValueType().Union().Application(), ShouldEqual, app)
					So(str.ValueType().Union().ItemsKind(), ShouldEqual, reflect.UnionOfResources)
					So(str.ValueType().Union().LookupFields().FieldCount(), ShouldEqual, 3)
					f, ok := str.ValueType().Union().LookupFields().FieldByPosition(0)
					So(ok, ShouldEqual, true)
					So(f.Name(), ShouldEqual, "withField")
					So(f.ValueType().Resource().Name(), ShouldEqual, "Friend")
					_, ok = str.ValueType().Union().LookupFields().FieldByPosition(90)
					So(ok, ShouldEqual, false)
					f, ok = str.ValueType().Union().LookupFields().FieldByName("File")
					So(ok, ShouldEqual, true)
					So(f.Name(), ShouldEqual, "File")

				})

			})

			Convey("Check the Friend resource ", func() {
				r, ok := app.LookupResources().ResourceByPosition(1)
				So(ok, ShouldEqual, true)

				So(r.Name(), ShouldEqual, "Friend")
				So(r.Application(), ShouldEqual, app)
				So(r.Position(), ShouldEqual, 1)
				So(r.ID(), ShouldEqual, "omniql.almagest.io/omniql/scalar-app/Resource/Friend")

				Convey("Check the Friend table ", func() {
					t := r.Table()
					So(t.Name(), ShouldEqual, "Friend")
					So(t.Application(), ShouldEqual, app)
					So(t.ID(), ShouldEqual, "omniql.almagest.io/omniql/scalar-app/Table/Friend")
					So(t.LookupFields().FieldCount(), ShouldEqual, 4)

				})
			})

			Convey("Check the File resource ", func() {
				r, ok := app.LookupResources().ResourceByPosition(2)
				So(ok, ShouldEqual, true)

				So(r.Name(), ShouldEqual, "File")
				So(r.Application(), ShouldEqual, app)
				So(r.Position(), ShouldEqual, 2)
				So(r.ID(), ShouldEqual, "omniql.almagest.io/omniql/scalar-app/Resource/File")
			})

		})

		//rt := r.Table()
		//So(rt.LookupFields().FieldCount(), ShouldEqual, 28)

	})
}
