// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// ExternalApplicationContainer is an autogenerated mock type for the ExternalApplicationContainer type
type ExternalApplicationContainer struct {
	mock.Mock
}

// Alias provides a mock function with given fields:
func (_m *ExternalApplicationContainer) Alias() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m *ExternalApplicationContainer) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Version provides a mock function with given fields:
func (_m *ExternalApplicationContainer) Version() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}