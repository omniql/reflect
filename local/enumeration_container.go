package local

import (
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
)

type enumerationContainer struct {
	stringMap   map[string]uint16
	stringIndex []string
	hType       hybrids.Types
	app         *applicationContainer
	id          string
	name        string
}

//ByUint8ToString ...
func (e *enumerationContainer) ByUint8ToString(input uint8) (value string, ok bool) {
	if int(input) > len(e.stringIndex) {
		return
	}

	value = e.stringIndex[int(input)]
	ok = true
	return
}

//ByUint16ToString ...
func (e *enumerationContainer) ByUint16ToString(input uint16) (value string, ok bool) {
	if int(input) > len(e.stringIndex) {
		return
	}

	value = e.stringIndex[int(input)]
	ok = true
	return
}

//ByStringToUint8 ...
func (e *enumerationContainer) ByStringToUint8(input string) (value uint8, ok bool) {
	var v uint16
	camel := reflect.ToLower(input)
	v, ok = e.stringMap[camel]
	if !ok {
		return
	}
	value = uint8(v)
	return
}

//ByStringToUint16 ...
func (e *enumerationContainer) ByStringToUint16(input string) (value uint16, ok bool) {
	camel := reflect.ToLower(input)
	value, ok = e.stringMap[camel]
	return
}

//Lookup ...
func (e *enumerationContainer) Lookup() reflect.LookupEnumeration {
	return e
}

//HybridType ...
func (e *enumerationContainer) HybridType() hybrids.Types {
	return e.hType
}

//Application ...
func (e *enumerationContainer) Application() reflect.ApplicationContainer {
	return e.app
}

//Name ...
func (e *enumerationContainer) Name() string {
	return e.name
}

//ID ...
func (e *enumerationContainer) ID() string {
	return e.id
}
