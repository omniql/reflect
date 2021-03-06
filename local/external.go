package local

import "github.com/omniql/reflect"

type externalResourceContainer struct {
	id   string
	app  *externalApplicationContainer
	name string
}

func (e *externalResourceContainer) ID() string {
	return e.id
}

func (e *externalResourceContainer) Application() reflect.ExternalApplicationContainer {
	return e.app
}

func (e *externalResourceContainer) Name() string {
	return e.name
}

type externalApplicationContainer struct {
	path          string
	version       string
	alias         string
	usedResources []*externalResourceContainer
}

func (e *externalApplicationContainer) Path() string {
	return e.path
}

func (e *externalApplicationContainer) Version() string {
	return e.version
}

func (e *externalApplicationContainer) Alias() string {
	return e.alias
}


func (e *externalApplicationContainer) UsedResourcesCount() int {
	return len(e.usedResources)
}


func (e *externalApplicationContainer) UsedResource(pos int) reflect.ExternalResourceContainer {
	return e.usedResources[pos]
}
