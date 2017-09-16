// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import reflect "github.com/omniql/reflect"

// OType is an autogenerated mock type for the OType type
type OType struct {
	mock.Mock
}

// Enumeration provides a mock function with given fields:
func (_m *OType) Enumeration() reflect.EnumerationContainer {
	ret := _m.Called()

	var r0 reflect.EnumerationContainer
	if rf, ok := ret.Get(0).(func() reflect.EnumerationContainer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.EnumerationContainer)
		}
	}

	return r0
}

// ExternalResource provides a mock function with given fields:
func (_m *OType) ExternalResource() reflect.ExternalResourceContainer {
	ret := _m.Called()

	var r0 reflect.ExternalResourceContainer
	if rf, ok := ret.Get(0).(func() reflect.ExternalResourceContainer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.ExternalResourceContainer)
		}
	}

	return r0
}

// Kind provides a mock function with given fields:
func (_m *OType) Kind() reflect.OmniTypes {
	ret := _m.Called()

	var r0 reflect.OmniTypes
	if rf, ok := ret.Get(0).(func() reflect.OmniTypes); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(reflect.OmniTypes)
	}

	return r0
}

// Resource provides a mock function with given fields:
func (_m *OType) Resource() reflect.ResourceContainer {
	ret := _m.Called()

	var r0 reflect.ResourceContainer
	if rf, ok := ret.Get(0).(func() reflect.ResourceContainer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.ResourceContainer)
		}
	}

	return r0
}

// Struct provides a mock function with given fields:
func (_m *OType) Struct() reflect.StructContainer {
	ret := _m.Called()

	var r0 reflect.StructContainer
	if rf, ok := ret.Get(0).(func() reflect.StructContainer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.StructContainer)
		}
	}

	return r0
}

// Table provides a mock function with given fields:
func (_m *OType) Table() reflect.TableContainer {
	ret := _m.Called()

	var r0 reflect.TableContainer
	if rf, ok := ret.Get(0).(func() reflect.TableContainer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.TableContainer)
		}
	}

	return r0
}

// Union provides a mock function with given fields:
func (_m *OType) Union() reflect.UnionContainer {
	ret := _m.Called()

	var r0 reflect.UnionContainer
	if rf, ok := ret.Get(0).(func() reflect.UnionContainer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.UnionContainer)
		}
	}

	return r0
}
