// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import reflect "github.com/omniql/reflect"

// ApplicationContainer is an autogenerated mock type for the ApplicationContainer type
type ApplicationContainer struct {
	mock.Mock
}

// LookupResources provides a mock function with given fields:
func (_m *ApplicationContainer) LookupResources() reflect.LookupResources {
	ret := _m.Called()

	var r0 reflect.LookupResources
	if rf, ok := ret.Get(0).(func() reflect.LookupResources); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.LookupResources)
		}
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m *ApplicationContainer) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ResourceCount provides a mock function with given fields:
func (_m *ApplicationContainer) ResourceCount() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// Version provides a mock function with given fields:
func (_m *ApplicationContainer) Version() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
