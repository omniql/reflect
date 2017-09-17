package local

import (
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
)

type applicationContainer struct {
	path             string
	version          string
	rType            hybrids.ResourceIDType
	resourceMap      map[string]*resourceContainer
	resourceIndex    []*resourceContainer
	externalAppMap   map[string]*externalApplicationContainer
	externalAppIndex []*externalApplicationContainer
}

func (a *applicationContainer) Path() string {
	return a.path
}

func (a *applicationContainer) ResourceIDType() hybrids.ResourceIDType {
	return a.rType
}

func (a *applicationContainer) Version() string {
	return a.version
}

func (a *applicationContainer) ResourceCount() int {
	return len(a.resourceIndex)
}

func (a *applicationContainer) ImportsCount() int {
	return len(a.externalAppIndex)
}

func (a *applicationContainer) LookupResources() reflect.LookupResources {
	return a
}

func (a *applicationContainer) LookupImports() reflect.LookupImports {
	return a
}

func (a *applicationContainer) ImportByAlias(alias string) (r reflect.ExternalApplicationContainer, ok bool) {
	r, ok = a.externalAppMap[reflect.ToLower(alias)]
	return
}

func (a *applicationContainer) ImportByPosition(position uint16) (r reflect.ExternalApplicationContainer, ok bool) {
	if int(position) > len(a.externalAppIndex) {
		return
	}
	r = a.externalAppIndex[int(position)]
	ok = true
	return
}

func (a *applicationContainer) ResourceByName(name string) (r reflect.ResourceContainer, ok bool) {
	r, ok = a.resourceMap[reflect.ToLower(name)]
	return
}

func (a *applicationContainer) ResourceByPosition(position uint16) (r reflect.ResourceContainer, ok bool) {
	if int(position) > len(a.resourceIndex) {
		return
	}
	r = a.resourceIndex[int(position)]
	ok = true
	return
}
