package local

import (
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
)

type tableContainer struct {
	id         string
	app        *applicationContainer
	name       string
	fieldIndex []*fieldContainer
	fieldMap   map[string]*fieldContainer
}

func (t *tableContainer) ID() string {
	return t.id
}

func (t *tableContainer) Application() reflect.ApplicationContainer {
	return t.app
}

func (t *tableContainer) Name() string {
	return t.name
}

func (t *tableContainer) FieldCount() int {
	return len(t.fieldIndex)
}

func (t *tableContainer) LookupFields() reflect.LookupFields {
	return t
}

//ByPosition ...
func (t *tableContainer) ByPosition(fn hybrids.FieldNumber) (f reflect.FieldContainer, ok bool) {
	if int(fn) > len(t.fieldIndex) {
		return
	}
	return t.fieldIndex[int(fn)], true
}

//ByName ...
func (t *tableContainer) ByName(fieldName string) (f reflect.FieldContainer, ok bool) {
	f, ok = t.fieldMap[reflect.ToLower(fieldName)]
	return
}
