// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Oauth2UseCase is an autogenerated mock type for the Oauth2UseCase type
type Oauth2UseCase struct {
	mock.Mock
}

// AuthzFacebook provides a mock function with given fields: accessTokenClient, accessToken
func (_m *Oauth2UseCase) AuthzFacebook(accessTokenClient string, accessToken string) (bool, error) {
	ret := _m.Called(accessTokenClient, accessToken)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(accessTokenClient, accessToken)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accessTokenClient, accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOauth2UseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewOauth2UseCase creates a new instance of Oauth2UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOauth2UseCase(t mockConstructorTestingTNewOauth2UseCase) *Oauth2UseCase {
	mock := &Oauth2UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
