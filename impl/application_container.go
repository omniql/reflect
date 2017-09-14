package impl

import (
	"github.com/omniql/reflect"
)

type applicationContainer struct {
	path            string
	version         string
	resourceMap     map[string]*resourceContainer
	resourceIndex   [...]*resourceContainer
}

func (a *applicationContainer) Path() string {
	return a.path
}

func (a *applicationContainer) Version() string {
	return a.version
}

func (a *applicationContainer) ResourceCount() int {
	return len(a.resourceIndex)
}

func (a *applicationContainer) LookupResources() reflect.LookupResources {
	return a
}

func (a *applicationContainer) ByName(name string) (r reflect.ResourceContainer, ok bool) {
	r, ok = a.resourceMap[reflect.ToLower(name)]
	return
}

func (a *applicationContainer) ByPosition(position uint16) (r reflect.ResourceContainer, ok bool) {
	if int(position) > len(a.resourceIndex){
		return
	}
	r = a.resourceIndex[int(position)]
	ok = true
	return
}
