// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/server/public/v8/model"
	mock "github.com/stretchr/testify/mock"
)

// PluginStore is an autogenerated mock type for the PluginStore type
type PluginStore struct {
	mock.Mock
}

// CompareAndDelete provides a mock function with given fields: keyVal, oldValue
func (_m *PluginStore) CompareAndDelete(keyVal *model.PluginKeyValue, oldValue []byte) (bool, error) {
	ret := _m.Called(keyVal, oldValue)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.PluginKeyValue, []byte) (bool, error)); ok {
		return rf(keyVal, oldValue)
	}
	if rf, ok := ret.Get(0).(func(*model.PluginKeyValue, []byte) bool); ok {
		r0 = rf(keyVal, oldValue)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*model.PluginKeyValue, []byte) error); ok {
		r1 = rf(keyVal, oldValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompareAndSet provides a mock function with given fields: keyVal, oldValue
func (_m *PluginStore) CompareAndSet(keyVal *model.PluginKeyValue, oldValue []byte) (bool, error) {
	ret := _m.Called(keyVal, oldValue)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.PluginKeyValue, []byte) (bool, error)); ok {
		return rf(keyVal, oldValue)
	}
	if rf, ok := ret.Get(0).(func(*model.PluginKeyValue, []byte) bool); ok {
		r0 = rf(keyVal, oldValue)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*model.PluginKeyValue, []byte) error); ok {
		r1 = rf(keyVal, oldValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: pluginID, key
func (_m *PluginStore) Delete(pluginID string, key string) error {
	ret := _m.Called(pluginID, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(pluginID, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllExpired provides a mock function with given fields:
func (_m *PluginStore) DeleteAllExpired() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllForPlugin provides a mock function with given fields: PluginID
func (_m *PluginStore) DeleteAllForPlugin(PluginID string) error {
	ret := _m.Called(PluginID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(PluginID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: pluginID, key
func (_m *PluginStore) Get(pluginID string, key string) (*model.PluginKeyValue, error) {
	ret := _m.Called(pluginID, key)

	var r0 *model.PluginKeyValue
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*model.PluginKeyValue, error)); ok {
		return rf(pluginID, key)
	}
	if rf, ok := ret.Get(0).(func(string, string) *model.PluginKeyValue); ok {
		r0 = rf(pluginID, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PluginKeyValue)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(pluginID, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: pluginID, page, perPage
func (_m *PluginStore) List(pluginID string, page int, perPage int) ([]string, error) {
	ret := _m.Called(pluginID, page, perPage)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]string, error)); ok {
		return rf(pluginID, page, perPage)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []string); ok {
		r0 = rf(pluginID, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(pluginID, page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveOrUpdate provides a mock function with given fields: keyVal
func (_m *PluginStore) SaveOrUpdate(keyVal *model.PluginKeyValue) (*model.PluginKeyValue, error) {
	ret := _m.Called(keyVal)

	var r0 *model.PluginKeyValue
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.PluginKeyValue) (*model.PluginKeyValue, error)); ok {
		return rf(keyVal)
	}
	if rf, ok := ret.Get(0).(func(*model.PluginKeyValue) *model.PluginKeyValue); ok {
		r0 = rf(keyVal)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PluginKeyValue)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.PluginKeyValue) error); ok {
		r1 = rf(keyVal)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetWithOptions provides a mock function with given fields: pluginID, key, value, options
func (_m *PluginStore) SetWithOptions(pluginID string, key string, value []byte, options model.PluginKVSetOptions) (bool, error) {
	ret := _m.Called(pluginID, key, value, options)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, []byte, model.PluginKVSetOptions) (bool, error)); ok {
		return rf(pluginID, key, value, options)
	}
	if rf, ok := ret.Get(0).(func(string, string, []byte, model.PluginKVSetOptions) bool); ok {
		r0 = rf(pluginID, key, value, options)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string, []byte, model.PluginKVSetOptions) error); ok {
		r1 = rf(pluginID, key, value, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPluginStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewPluginStore creates a new instance of PluginStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPluginStore(t mockConstructorTestingTNewPluginStore) *PluginStore {
	mock := &PluginStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
