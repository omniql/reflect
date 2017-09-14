package local

import (
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
)

type enumerationContainer struct {
	stringMap  map[string]uint16
	reverseMap [...]string
	hType      hybrids.Types
}

func (e *enumerationContainer) ByUint8ToString(input uint8) (value string, ok bool) {
	if int(input) > len(e.reverseMap) {
		return
	}

	value = e.reverseMap[int(input)]
	ok = true
	return
}

func (e *enumerationContainer) ByUint16ToString(input uint16) (value string, ok bool) {
	if int(input) > len(e.reverseMap) {
		return
	}

	value = e.reverseMap[int(input)]
	ok = true
	return
}

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

func (e *enumerationContainer) ByStringToUint16(input string) (value uint16, ok bool) {
	camel := reflect.ToLower(input)
	value, ok = e.stringMap[camel]
	return
}

func (e *enumerationContainer) Lookup() reflect.LookupEnumeration {
	return e
}

func (e *enumerationContainer) HybridType() hybrids.Types {
	return e.hType
}
