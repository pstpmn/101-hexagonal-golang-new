// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IServer is an autogenerated mock type for the IServer type
type IServer struct {
	mock.Mock
}

// Initialize provides a mock function with given fields:
func (_m *IServer) Initialize() {
	_m.Called()
}

type mockConstructorTestingTNewIServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewIServer creates a new instance of IServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIServer(t mockConstructorTestingTNewIServer) *IServer {
	mock := &IServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}