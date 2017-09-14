package local

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/omniql/reflect"
)

type itemsContainer struct {
	value *oType
	hType hybrids.Types
}

func (i *itemsContainer) ValueType() reflect.OType {
	return i.value

}

func (i *itemsContainer) HybridType() hybrids.Types {
	return i.hType

}
