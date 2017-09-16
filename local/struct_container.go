package local

import (
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
)

type structContainer struct {
	id         string
	app        *applicationContainer
	name       string
	fieldIndex []*fieldContainer
	fieldMap   map[string]*fieldContainer
}

func (t *structContainer) ID() string {
	return t.id
}

func (t *structContainer) Application() reflect.ApplicationContainer {
	return t.app
}

func (t *structContainer) Name() string {
	return t.name
}

func (t *structContainer) FieldCount() int {
	return len(t.fieldIndex)
}

func (t *structContainer) LookupFields() reflect.LookupFields {
	return t
}

//ByPosition ...
func (t *structContainer) FieldByPosition(fn hybrids.FieldNumber) (f reflect.FieldContainer, ok bool) {
	if int(fn) > len(t.fieldIndex) {
		return
	}
	return t.fieldIndex[int(fn)], true
}

//ByName ...
func (t *structContainer) FieldByName(fieldName string) (f reflect.FieldContainer, ok bool) {
	f, ok = t.fieldMap[reflect.ToLower(fieldName)]
	return
}
