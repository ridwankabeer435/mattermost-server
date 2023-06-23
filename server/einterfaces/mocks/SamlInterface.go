// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	request "github.com/mattermost/mattermost/server/v8/channels/app/request"
	mock "github.com/stretchr/testify/mock"
)

// SamlInterface is an autogenerated mock type for the SamlInterface type
type SamlInterface struct {
	mock.Mock
}

// BuildRequest provides a mock function with given fields: relayState
func (_m *SamlInterface) BuildRequest(relayState string) (*model.SamlAuthRequest, *model.AppError) {
	ret := _m.Called(relayState)

	var r0 *model.SamlAuthRequest
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string) (*model.SamlAuthRequest, *model.AppError)); ok {
		return rf(relayState)
	}
	if rf, ok := ret.Get(0).(func(string) *model.SamlAuthRequest); ok {
		r0 = rf(relayState)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SamlAuthRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(relayState)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// CheckProviderAttributes provides a mock function with given fields: SS, ouser, patch
func (_m *SamlInterface) CheckProviderAttributes(SS *model.SamlSettings, ouser *model.User, patch *model.UserPatch) string {
	ret := _m.Called(SS, ouser, patch)

	var r0 string
	if rf, ok := ret.Get(0).(func(*model.SamlSettings, *model.User, *model.UserPatch) string); ok {
		r0 = rf(SS, ouser, patch)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ConfigureSP provides a mock function with given fields:
func (_m *SamlInterface) ConfigureSP() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DoLogin provides a mock function with given fields: c, encodedXML, relayState
func (_m *SamlInterface) DoLogin(c *request.Context, encodedXML string, relayState map[string]string) (*model.User, *model.AppError) {
	ret := _m.Called(c, encodedXML, relayState)

	var r0 *model.User
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(*request.Context, string, map[string]string) (*model.User, *model.AppError)); ok {
		return rf(c, encodedXML, relayState)
	}
	if rf, ok := ret.Get(0).(func(*request.Context, string, map[string]string) *model.User); ok {
		r0 = rf(c, encodedXML, relayState)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*request.Context, string, map[string]string) *model.AppError); ok {
		r1 = rf(c, encodedXML, relayState)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetMetadata provides a mock function with given fields:
func (_m *SamlInterface) GetMetadata() (string, *model.AppError) {
	ret := _m.Called()

	var r0 string
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func() (string, *model.AppError)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() *model.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewSamlInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewSamlInterface creates a new instance of SamlInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSamlInterface(t mockConstructorTestingTNewSamlInterface) *SamlInterface {
	mock := &SamlInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
