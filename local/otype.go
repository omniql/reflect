package local

import (
	"github.com/omniql/reflect"
)

type oType struct {
	kind        reflect.OmniTypes
	application *applicationContainer
	enum        *enumerationContainer
	table       *tableContainer
	str         *structContainer
	field       *fieldContainer
	res         *resourceContainer
	union       *unionContainer
	er          *externalResourceContainer
	ea          *externalApplicationContainer
}

func (o *oType) Kind() reflect.OmniTypes {
	return o.kind
}

func (o *oType) Enumeration() reflect.EnumerationContainer {
	return o.enum
}

func (o *oType) Table() reflect.TableContainer {
	return o.table
}

func (o *oType) Struct() reflect.StructContainer {
	return o.str
}

func (o *oType) Resource() reflect.ResourceContainer {
	return o.res
}

func (o *oType) Union() reflect.UnionContainer {
	return o.union
}

func (o *oType) ExternalResource() reflect.ExternalResourceContainer {
	return o.er
}
