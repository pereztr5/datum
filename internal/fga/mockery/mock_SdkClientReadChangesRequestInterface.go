// Code generated by mockery v2.40.1. DO NOT EDIT.

package client

import (
	context "context"

	client "github.com/openfga/go-sdk/client"

	mock "github.com/stretchr/testify/mock"

	openfga "github.com/openfga/go-sdk"
)

// MockSdkClientReadChangesRequestInterface is an autogenerated mock type for the SdkClientReadChangesRequestInterface type
type MockSdkClientReadChangesRequestInterface struct {
	mock.Mock
}

type MockSdkClientReadChangesRequestInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSdkClientReadChangesRequestInterface) EXPECT() *MockSdkClientReadChangesRequestInterface_Expecter {
	return &MockSdkClientReadChangesRequestInterface_Expecter{mock: &_m.Mock}
}

// Body provides a mock function with given fields: body
func (_m *MockSdkClientReadChangesRequestInterface) Body(body client.ClientReadChangesRequest) client.SdkClientReadChangesRequestInterface {
	ret := _m.Called(body)

	if len(ret) == 0 {
		panic("no return value specified for Body")
	}

	var r0 client.SdkClientReadChangesRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientReadChangesRequest) client.SdkClientReadChangesRequestInterface); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientReadChangesRequestInterface)
		}
	}

	return r0
}

// MockSdkClientReadChangesRequestInterface_Body_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Body'
type MockSdkClientReadChangesRequestInterface_Body_Call struct {
	*mock.Call
}

// Body is a helper method to define mock.On call
//   - body client.ClientReadChangesRequest
func (_e *MockSdkClientReadChangesRequestInterface_Expecter) Body(body interface{}) *MockSdkClientReadChangesRequestInterface_Body_Call {
	return &MockSdkClientReadChangesRequestInterface_Body_Call{Call: _e.mock.On("Body", body)}
}

func (_c *MockSdkClientReadChangesRequestInterface_Body_Call) Run(run func(body client.ClientReadChangesRequest)) *MockSdkClientReadChangesRequestInterface_Body_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientReadChangesRequest))
	})
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_Body_Call) Return(_a0 client.SdkClientReadChangesRequestInterface) *MockSdkClientReadChangesRequestInterface_Body_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_Body_Call) RunAndReturn(run func(client.ClientReadChangesRequest) client.SdkClientReadChangesRequestInterface) *MockSdkClientReadChangesRequestInterface_Body_Call {
	_c.Call.Return(run)
	return _c
}

// Execute provides a mock function with given fields:
func (_m *MockSdkClientReadChangesRequestInterface) Execute() (*openfga.ReadChangesResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *openfga.ReadChangesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*openfga.ReadChangesResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *openfga.ReadChangesResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openfga.ReadChangesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSdkClientReadChangesRequestInterface_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockSdkClientReadChangesRequestInterface_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockSdkClientReadChangesRequestInterface_Expecter) Execute() *MockSdkClientReadChangesRequestInterface_Execute_Call {
	return &MockSdkClientReadChangesRequestInterface_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockSdkClientReadChangesRequestInterface_Execute_Call) Run(run func()) *MockSdkClientReadChangesRequestInterface_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_Execute_Call) Return(_a0 *openfga.ReadChangesResponse, _a1 error) *MockSdkClientReadChangesRequestInterface_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_Execute_Call) RunAndReturn(run func() (*openfga.ReadChangesResponse, error)) *MockSdkClientReadChangesRequestInterface_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// GetBody provides a mock function with given fields:
func (_m *MockSdkClientReadChangesRequestInterface) GetBody() *client.ClientReadChangesRequest {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBody")
	}

	var r0 *client.ClientReadChangesRequest
	if rf, ok := ret.Get(0).(func() *client.ClientReadChangesRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientReadChangesRequest)
		}
	}

	return r0
}

// MockSdkClientReadChangesRequestInterface_GetBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBody'
type MockSdkClientReadChangesRequestInterface_GetBody_Call struct {
	*mock.Call
}

// GetBody is a helper method to define mock.On call
func (_e *MockSdkClientReadChangesRequestInterface_Expecter) GetBody() *MockSdkClientReadChangesRequestInterface_GetBody_Call {
	return &MockSdkClientReadChangesRequestInterface_GetBody_Call{Call: _e.mock.On("GetBody")}
}

func (_c *MockSdkClientReadChangesRequestInterface_GetBody_Call) Run(run func()) *MockSdkClientReadChangesRequestInterface_GetBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_GetBody_Call) Return(_a0 *client.ClientReadChangesRequest) *MockSdkClientReadChangesRequestInterface_GetBody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_GetBody_Call) RunAndReturn(run func() *client.ClientReadChangesRequest) *MockSdkClientReadChangesRequestInterface_GetBody_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields:
func (_m *MockSdkClientReadChangesRequestInterface) GetContext() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetContext")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockSdkClientReadChangesRequestInterface_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockSdkClientReadChangesRequestInterface_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
func (_e *MockSdkClientReadChangesRequestInterface_Expecter) GetContext() *MockSdkClientReadChangesRequestInterface_GetContext_Call {
	return &MockSdkClientReadChangesRequestInterface_GetContext_Call{Call: _e.mock.On("GetContext")}
}

func (_c *MockSdkClientReadChangesRequestInterface_GetContext_Call) Run(run func()) *MockSdkClientReadChangesRequestInterface_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_GetContext_Call) Return(_a0 context.Context) *MockSdkClientReadChangesRequestInterface_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_GetContext_Call) RunAndReturn(run func() context.Context) *MockSdkClientReadChangesRequestInterface_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetOptions provides a mock function with given fields:
func (_m *MockSdkClientReadChangesRequestInterface) GetOptions() *client.ClientReadChangesOptions {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOptions")
	}

	var r0 *client.ClientReadChangesOptions
	if rf, ok := ret.Get(0).(func() *client.ClientReadChangesOptions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientReadChangesOptions)
		}
	}

	return r0
}

// MockSdkClientReadChangesRequestInterface_GetOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOptions'
type MockSdkClientReadChangesRequestInterface_GetOptions_Call struct {
	*mock.Call
}

// GetOptions is a helper method to define mock.On call
func (_e *MockSdkClientReadChangesRequestInterface_Expecter) GetOptions() *MockSdkClientReadChangesRequestInterface_GetOptions_Call {
	return &MockSdkClientReadChangesRequestInterface_GetOptions_Call{Call: _e.mock.On("GetOptions")}
}

func (_c *MockSdkClientReadChangesRequestInterface_GetOptions_Call) Run(run func()) *MockSdkClientReadChangesRequestInterface_GetOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_GetOptions_Call) Return(_a0 *client.ClientReadChangesOptions) *MockSdkClientReadChangesRequestInterface_GetOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_GetOptions_Call) RunAndReturn(run func() *client.ClientReadChangesOptions) *MockSdkClientReadChangesRequestInterface_GetOptions_Call {
	_c.Call.Return(run)
	return _c
}

// Options provides a mock function with given fields: options
func (_m *MockSdkClientReadChangesRequestInterface) Options(options client.ClientReadChangesOptions) client.SdkClientReadChangesRequestInterface {
	ret := _m.Called(options)

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 client.SdkClientReadChangesRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientReadChangesOptions) client.SdkClientReadChangesRequestInterface); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientReadChangesRequestInterface)
		}
	}

	return r0
}

// MockSdkClientReadChangesRequestInterface_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type MockSdkClientReadChangesRequestInterface_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
//   - options client.ClientReadChangesOptions
func (_e *MockSdkClientReadChangesRequestInterface_Expecter) Options(options interface{}) *MockSdkClientReadChangesRequestInterface_Options_Call {
	return &MockSdkClientReadChangesRequestInterface_Options_Call{Call: _e.mock.On("Options", options)}
}

func (_c *MockSdkClientReadChangesRequestInterface_Options_Call) Run(run func(options client.ClientReadChangesOptions)) *MockSdkClientReadChangesRequestInterface_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientReadChangesOptions))
	})
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_Options_Call) Return(_a0 client.SdkClientReadChangesRequestInterface) *MockSdkClientReadChangesRequestInterface_Options_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientReadChangesRequestInterface_Options_Call) RunAndReturn(run func(client.ClientReadChangesOptions) client.SdkClientReadChangesRequestInterface) *MockSdkClientReadChangesRequestInterface_Options_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSdkClientReadChangesRequestInterface creates a new instance of MockSdkClientReadChangesRequestInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSdkClientReadChangesRequestInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSdkClientReadChangesRequestInterface {
	mock := &MockSdkClientReadChangesRequestInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
