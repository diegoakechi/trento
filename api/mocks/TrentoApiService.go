// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	web "github.com/trento-project/trento/web"
)

// TrentoApiService is an autogenerated mock type for the TrentoApiService type
type TrentoApiService struct {
	mock.Mock
}

// GetChecksSettingsById provides a mock function with given fields: id
func (_m *TrentoApiService) GetChecksSettingsById(id string) (*web.JSONChecksSettings, error) {
	ret := _m.Called(id)

	var r0 *web.JSONChecksSettings
	if rf, ok := ret.Get(0).(func(string) *web.JSONChecksSettings); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*web.JSONChecksSettings)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsWebServerUp provides a mock function with given fields:
func (_m *TrentoApiService) IsWebServerUp() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
