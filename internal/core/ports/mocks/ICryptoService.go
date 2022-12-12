// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ICryptoService is an autogenerated mock type for the ICryptoService type
type ICryptoService struct {
	mock.Mock
}

// Bcrypt provides a mock function with given fields: plaintext
func (_m *ICryptoService) Bcrypt(plaintext string) (string, error) {
	ret := _m.Called(plaintext)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(plaintext)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(plaintext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Md5 provides a mock function with given fields: plaintext
func (_m *ICryptoService) Md5(plaintext string) string {
	ret := _m.Called(plaintext)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(plaintext)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ValidateBcrypt provides a mock function with given fields: plaintext, encrypt
func (_m *ICryptoService) ValidateBcrypt(plaintext string, encrypt string) bool {
	ret := _m.Called(plaintext, encrypt)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(plaintext, encrypt)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewICryptoService interface {
	mock.TestingT
	Cleanup(func())
}

// NewICryptoService creates a new instance of ICryptoService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICryptoService(t mockConstructorTestingTNewICryptoService) *ICryptoService {
	mock := &ICryptoService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}