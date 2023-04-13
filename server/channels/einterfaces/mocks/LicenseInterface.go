// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/server/public/v8/model"
	mock "github.com/stretchr/testify/mock"
)

// LicenseInterface is an autogenerated mock type for the LicenseInterface type
type LicenseInterface struct {
	mock.Mock
}

// CanStartTrial provides a mock function with given fields:
func (_m *LicenseInterface) CanStartTrial() (bool, error) {
	ret := _m.Called()

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPrevTrial provides a mock function with given fields:
func (_m *LicenseInterface) GetPrevTrial() (*model.License, error) {
	ret := _m.Called()

	var r0 *model.License
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.License, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.License); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.License)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLicenseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewLicenseInterface creates a new instance of LicenseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLicenseInterface(t mockConstructorTestingTNewLicenseInterface) *LicenseInterface {
	mock := &LicenseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
