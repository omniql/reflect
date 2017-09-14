// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import reflect "github.com/omniql/reflect"

// ExternalResourceContainer is an autogenerated mock type for the ExternalResourceContainer type
type ExternalResourceContainer struct {
	mock.Mock
}

// Application provides a mock function with given fields:
func (_m *ExternalResourceContainer) Application() reflect.ExternalApplicationContainer {
	ret := _m.Called()

	var r0 reflect.ExternalApplicationContainer
	if rf, ok := ret.Get(0).(func() reflect.ExternalApplicationContainer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.ExternalApplicationContainer)
		}
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *ExternalResourceContainer) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *ExternalResourceContainer) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
