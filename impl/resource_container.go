package impl

import "github.com/omniql/reflect"

type resourceContainer struct {
	id       string
	app      *applicationContainer
	name     string
	position uint16
	table    *tableContainer
}

func (r *resourceContainer) ID() string {
	return r.id
}

func (r *resourceContainer) Name() string {
	return r.name
}

func (r *resourceContainer) Position() uint16 {
	return r.position
}

func (r *resourceContainer) Application() reflect.ApplicationContainer {
	return r.app
}

func (r *resourceContainer) Table() reflect.TableContainer {
	return r.table
}
