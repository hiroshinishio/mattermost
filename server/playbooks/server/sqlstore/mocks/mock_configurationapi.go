// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-server/server/v8/playbooks/server/sqlstore (interfaces: ConfigurationAPI)

// Package mock_sqlstore is a generated GoMock package.
package mock_sqlstore

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/server/public/v8/model"
)

// MockConfigurationAPI is a mock of ConfigurationAPI interface.
type MockConfigurationAPI struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationAPIMockRecorder
}

// MockConfigurationAPIMockRecorder is the mock recorder for MockConfigurationAPI.
type MockConfigurationAPIMockRecorder struct {
	mock *MockConfigurationAPI
}

// NewMockConfigurationAPI creates a new mock instance.
func NewMockConfigurationAPI(ctrl *gomock.Controller) *MockConfigurationAPI {
	mock := &MockConfigurationAPI{ctrl: ctrl}
	mock.recorder = &MockConfigurationAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationAPI) EXPECT() *MockConfigurationAPIMockRecorder {
	return m.recorder
}

// GetConfig mocks base method.
func (m *MockConfigurationAPI) GetConfig() *model.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*model.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockConfigurationAPIMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockConfigurationAPI)(nil).GetConfig))
}
