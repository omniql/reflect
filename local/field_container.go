package local

import (
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
)

type fieldContainer struct {
	id          string
	app         *applicationContainer
	parent      *oType
	value       *oType
	position    hybrids.FieldNumber
	name        string
	hybridsType hybrids.Types
	items       *itemsContainer
}

func (f *fieldContainer) ID() string {
	return f.id
}

func (f *fieldContainer) Application() reflect.ApplicationContainer {
	return f.app
}

func (f *fieldContainer) Parent() reflect.OType {
	return f.parent
}

func (f *fieldContainer) ValueType() reflect.OType {
	return f.value
}

func (f *fieldContainer) Position() hybrids.FieldNumber {
	return f.position
}

//field name
func (f *fieldContainer) Name() string {
	return f.name
}

//the underlying data type
func (f *fieldContainer) HybridType() hybrids.Types {
	return f.hybridsType
}

//if is a vector the item type
func (f *fieldContainer) Items() reflect.ItemsContainer {
	return f.items
}
