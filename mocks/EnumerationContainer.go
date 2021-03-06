// Code generated by mockery v1.0.0
package mocks

import hybrids "github.com/nebtex/hybrids/golang/hybrids"
import mock "github.com/stretchr/testify/mock"
import reflect "github.com/omniql/reflect"

// EnumerationContainer is an autogenerated mock type for the EnumerationContainer type
type EnumerationContainer struct {
	mock.Mock
}

// Application provides a mock function with given fields:
func (_m *EnumerationContainer) Application() reflect.ApplicationContainer {
	ret := _m.Called()

	var r0 reflect.ApplicationContainer
	if rf, ok := ret.Get(0).(func() reflect.ApplicationContainer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.ApplicationContainer)
		}
	}

	return r0
}

// HybridType provides a mock function with given fields:
func (_m *EnumerationContainer) HybridType() hybrids.Types {
	ret := _m.Called()

	var r0 hybrids.Types
	if rf, ok := ret.Get(0).(func() hybrids.Types); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(hybrids.Types)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *EnumerationContainer) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Lookup provides a mock function with given fields:
func (_m *EnumerationContainer) Lookup() reflect.LookupEnumeration {
	ret := _m.Called()

	var r0 reflect.LookupEnumeration
	if rf, ok := ret.Get(0).(func() reflect.LookupEnumeration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.LookupEnumeration)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *EnumerationContainer) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
