package local

import (
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
)

type unionContainer struct {
	app        *applicationContainer
	id         string
	name       string
	itemsKind  reflect.UnionTypes
	fieldIndex []*fieldContainer
	fieldMap   map[string]*fieldContainer
}

//Application ...
func (u *unionContainer) Application() reflect.ApplicationContainer {
	return u.app
}

//ID ...
func (u *unionContainer) ID() string {
	return u.id
}

//Name ...
func (u *unionContainer) Name() string {
	return u.name
}

// Kind ...
func (u *unionContainer) ItemsKind() reflect.UnionTypes {
	return u.itemsKind
}

//FieldCount ...
func (u *unionContainer) FieldCount() int {
	return len(u.fieldIndex)
}

//LookupFields ...
func (u *unionContainer) LookupFields() reflect.LookupFields {
	return u
}

//ByPosition ...
func (u *unionContainer) FieldByPosition(fn hybrids.FieldNumber) (f reflect.FieldContainer, ok bool) {
	if int(fn) > len(u.fieldIndex) {
		return
	}
	return u.fieldIndex[int(fn)], true
}

//ByName ...
func (u *unionContainer) FieldByName(fieldName string) (f reflect.FieldContainer, ok bool) {
	f, ok = u.fieldMap[reflect.ToLower(fieldName)]
	return
}
