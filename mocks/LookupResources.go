// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import reflect "github.com/omniql/reflect"

// LookupResources is an autogenerated mock type for the LookupResources type
type LookupResources struct {
	mock.Mock
}

// ByName provides a mock function with given fields: name
func (_m *LookupResources) ByName(name string) (reflect.ResourceContainer, bool) {
	ret := _m.Called(name)

	var r0 reflect.ResourceContainer
	if rf, ok := ret.Get(0).(func(string) reflect.ResourceContainer); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.ResourceContainer)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// ByPosition provides a mock function with given fields: position
func (_m *LookupResources) ByPosition(position uint16) (reflect.ResourceContainer, bool) {
	ret := _m.Called(position)

	var r0 reflect.ResourceContainer
	if rf, ok := ret.Get(0).(func(uint16) reflect.ResourceContainer); ok {
		r0 = rf(position)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.ResourceContainer)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(uint16) bool); ok {
		r1 = rf(position)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}
