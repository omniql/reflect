package reflect

import (
	"github.com/nebtex/hybrids/golang/hybrids"
)

type OmniTypes uint16

const (
	Table            OmniTypes = iota
	Enumeration
	Struct
	Union
	Resource
	ExternalResource
	Application
)

type UnionTypes uint16

const (
	UnionOfResources         UnionTypes = iota
	UnionOfTables
	UnionOfExternalResources
)

//LookupFields ...
type LookupFields interface {
	FieldCount() int
	FieldByPosition(fn hybrids.FieldNumber) (f FieldContainer, ok bool)
	FieldByName(fieldName string) (f FieldContainer, ok bool)
}

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

//Table ...
type TableContainer interface {
	ID() string
	Application() ApplicationContainer
	Name() string
	LookupFields() LookupFields
}

//Resource ...
type ResourceContainer interface {
	ID() string
	Application() ApplicationContainer
	Name() string
	Position() uint16
	Table() TableContainer
}

//Struct ...
type StructContainer interface {
	ID() string
	Name() string
	Application() ApplicationContainer
	LookupFields() LookupFields
}

//Items ...
type ItemsContainer interface {
	ValueType() OType
	HybridType() hybrids.Types
}

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

//Enumeration ...
type EnumerationContainer interface {
	Application() ApplicationContainer
	ID() string
	Name() string
	Lookup() LookupEnumeration
	//the underlying data type
	HybridType() hybrids.Types
}

//Union ...
type UnionContainer interface {
	Application() ApplicationContainer
	ID() string
	Name() string
	ItemsKind() UnionTypes
	//total numbers of items in the union
	LookupFields() LookupFields
}

//Application ...
type LookupResources interface {
	ResourceCount() int
	//case insensitive
	ResourceByName(name string) (r ResourceContainer, ok bool)
	ResourceByPosition(position uint16) (r ResourceContainer, ok bool)
}

//Application ...
type ApplicationContainer interface {
	//example omniql.almagest.io/nebtex/omniql
	Path() string
	//branch, commit ot tag
	Version() string
	//number of resources
	LookupResources() LookupResources
	LookupImports() LookupImports
	ResourceIDType() hybrids.ResourceIDType
}

//LookupImports ...
type LookupImports interface {
	ImportsCount() int
	ImportByPosition(position uint16) (e ExternalApplicationContainer, ok bool)
	ImportByAlias(alias string) (e ExternalApplicationContainer, ok bool)
}

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
	ExternalResource() ExternalResourceContainer
}

//ExternalResourceContainer
type ExternalResourceContainer interface {
	ID() string
	Application() ExternalApplicationContainer
	Name() string
}

//ExternalApplicationContainer
type ExternalApplicationContainer interface {
	Path() string
	Version() string
	Alias() string
	UsedResourcesCount() int
	UsedResource(pos int) ExternalResourceContainer
}

//go:generate stringer -type=OmniTypes
//go:generate stringer -type=UnionTypes
//go:generate mockery -name=LookupFields
//go:generate mockery -name=LookupEnumeration
//go:generate mockery -name=TableContainer
//go:generate mockery -name=ResourceContainer
//go:generate mockery -name=StructContainer
//go:generate mockery -name=ItemsContainer
//go:generate mockery -name=FieldContainer
//go:generate mockery -name=EnumerationContainer
//go:generate mockery -name=UnionContainer
//go:generate mockery -name=LookupResources
//go:generate mockery -name=ApplicationContainer
//go:generate mockery -name=LookupImports
//go:generate mockery -name=OType
//go:generate mockery -name=ExternalResourceContainer
//go:generate mockery -name=ExternalApplicationContainer
