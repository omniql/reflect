package reflect

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/omniql/reflect/local"
)

type OmniTypes uint16

//go:generate stringer -type=OmniTypes
const (
	Table            OmniTypes = iota
	Enumeration
	Struct
	Union
	Resource
	ExternalResource
	Application
)

//go:generate stringer -type=UnionTypes
type UnionTypes uint16

const (
	UnionOfResources         UnionTypes = iota
	UnionOfTables
	UnionOfExternalResources
)

//go:generate mockery -name=LookupFields
//LookupFields ...
type LookupFields interface {
	ByPosition(fn hybrids.FieldNumber) (f FieldContainer, ok bool)
	ByName(fieldName string) (f FieldContainer, ok bool)
}

//go:generate mockery -name=LookupEnumeration
//LookupEnumeration ...
type LookupEnumeration interface {
	//return camelcase
	ByUint8ToString(input uint8) (value string, ok bool)
	//return camelcase
	ByUint16ToString(input uint16) (value string, ok bool)
	//accept snake case or camel case
	ByStringToUint8(input string) (value uint8, ok bool)
	//accept snake case or camel case
	ByStringToUint16(input string) (value uint16, ok bool)
}

//go:generate mockery -name=TableContainer
//Table ...
type TableContainer interface {
	ID() string
	Application() ApplicationContainer
	Name() string
	FieldCount() int
	LookupFields() LookupFields
}

//go:generate mockery -name=ResourceContainer
//Resource ...
type ResourceContainer interface {
	ID() string
	Application() ApplicationContainer
	Name() string
	Position() uint16
	Table() TableContainer
}

//go:generate mockery -name=StructContainer
//Struct ...
type StructContainer interface {
	ID() string
	Name() string
	Application() ApplicationContainer
	FieldCount() int
	LookupFields() LookupFields
}

//go:generate mockery -name=ItemsContainer
//Items ...
type ItemsContainer interface {
	ValueType() OType
	HybridType() hybrids.Types
}

//go:generate mockery -name=FieldContainer
//Field ...
type FieldContainer interface {
	//full id of this field, with the version, to allow query the backend
	ID() string

	Application() ApplicationContainer

	//table or struct
	Parent() OType

	//Otype of the field value
	//this return nil when is a scalar or vector scalar  unless tha the field is an enumeration enumeration
	ValueType() OType

	//position of this field in the table
	Position() hybrids.FieldNumber

	//field name
	Name() string

	//the underlying data type
	HybridType() hybrids.Types

	//if is a vector the item type
	Items() ItemsContainer
}

//go:generate mockery -name=EnumerationContainer
//Enumeration ...
type EnumerationContainer interface {
	Application() ApplicationContainer
	ID() string
	Name() string
	Lookup() LookupEnumeration
	//the underlying data type
	HybridType() hybrids.Types
}

//go:generate mockery -name=UnionContainer
//Union ...
type UnionContainer interface {
	Application() ApplicationContainer
	ID() string
	Name() string
	Kind() UnionTypes
	//total numbers of items in the union
	FieldCount() int
	LookupFields() LookupFields
}

//go:generate mockery -name=LookupResources
//Application ...
type LookupResources interface {
	//case insensitive
	ByName(name string) (r ResourceContainer, ok bool)
	ByPosition(position uint16) (r ResourceContainer, ok bool)
}

//go:generate mockery -name=ApplicationContainer
//Application ...
type ApplicationContainer interface {
	//example omniql.almagest.io/nebtex/omniql
	Path() string
	//branch, commit ot tag
	Version() string
	//number of resources
	ResourceCount() int
	LookupResources() LookupResources
}

//go:generate mockery -name=OType
//OType represent any omniql type
type OType interface {
	//Table, Field, Enumeration, etc..
	Kind() OmniTypes
	//if this is a vector this will return the item type
	Enumeration() EnumerationContainer
	Table() TableContainer
	Struct() StructContainer
	Union() UnionContainer
	Resource() ResourceContainer
	Application() ApplicationContainer
	ExternalResource() ExternalResourceContainer
	ExternalApplication() ExternalApplicationContainer
}

//go:generate mockery -name=ExternalResourceContainer
//ExternalResourceContainer
type ExternalResourceContainer interface {
	ID() string
	Application() ExternalApplicationContainer
	Name() string
}

//go:generate mockery -name=ExternalApplicationContainer
//ExternalApplicationContainer
type ExternalApplicationContainer interface {
	Path() string
	Version() string
	Alias() string
}

func LoadApplication(path string) {
	buider := &local.Builder{}

}
