// Code generated by mockery v2.40.1. DO NOT EDIT.

package client

import (
	context "context"

	client "github.com/openfga/go-sdk/client"

	mock "github.com/stretchr/testify/mock"
)

// MockSdkClientCheckRequestInterface is an autogenerated mock type for the SdkClientCheckRequestInterface type
type MockSdkClientCheckRequestInterface struct {
	mock.Mock
}

type MockSdkClientCheckRequestInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSdkClientCheckRequestInterface) EXPECT() *MockSdkClientCheckRequestInterface_Expecter {
	return &MockSdkClientCheckRequestInterface_Expecter{mock: &_m.Mock}
}

// Body provides a mock function with given fields: body
func (_m *MockSdkClientCheckRequestInterface) Body(body client.ClientCheckRequest) client.SdkClientCheckRequestInterface {
	ret := _m.Called(body)

	if len(ret) == 0 {
		panic("no return value specified for Body")
	}

	var r0 client.SdkClientCheckRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientCheckRequest) client.SdkClientCheckRequestInterface); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientCheckRequestInterface)
		}
	}

	return r0
}

// MockSdkClientCheckRequestInterface_Body_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Body'
type MockSdkClientCheckRequestInterface_Body_Call struct {
	*mock.Call
}

// Body is a helper method to define mock.On call
//   - body client.ClientCheckRequest
func (_e *MockSdkClientCheckRequestInterface_Expecter) Body(body interface{}) *MockSdkClientCheckRequestInterface_Body_Call {
	return &MockSdkClientCheckRequestInterface_Body_Call{Call: _e.mock.On("Body", body)}
}

func (_c *MockSdkClientCheckRequestInterface_Body_Call) Run(run func(body client.ClientCheckRequest)) *MockSdkClientCheckRequestInterface_Body_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientCheckRequest))
	})
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_Body_Call) Return(_a0 client.SdkClientCheckRequestInterface) *MockSdkClientCheckRequestInterface_Body_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_Body_Call) RunAndReturn(run func(client.ClientCheckRequest) client.SdkClientCheckRequestInterface) *MockSdkClientCheckRequestInterface_Body_Call {
	_c.Call.Return(run)
	return _c
}

// Execute provides a mock function with given fields:
func (_m *MockSdkClientCheckRequestInterface) Execute() (*client.ClientCheckResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *client.ClientCheckResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*client.ClientCheckResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *client.ClientCheckResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientCheckResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSdkClientCheckRequestInterface_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockSdkClientCheckRequestInterface_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockSdkClientCheckRequestInterface_Expecter) Execute() *MockSdkClientCheckRequestInterface_Execute_Call {
	return &MockSdkClientCheckRequestInterface_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockSdkClientCheckRequestInterface_Execute_Call) Run(run func()) *MockSdkClientCheckRequestInterface_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_Execute_Call) Return(_a0 *client.ClientCheckResponse, _a1 error) *MockSdkClientCheckRequestInterface_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_Execute_Call) RunAndReturn(run func() (*client.ClientCheckResponse, error)) *MockSdkClientCheckRequestInterface_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuthorizationModelIdOverride provides a mock function with given fields:
func (_m *MockSdkClientCheckRequestInterface) GetAuthorizationModelIdOverride() *string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorizationModelIdOverride")
	}

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// MockSdkClientCheckRequestInterface_GetAuthorizationModelIdOverride_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthorizationModelIdOverride'
type MockSdkClientCheckRequestInterface_GetAuthorizationModelIdOverride_Call struct {
	*mock.Call
}

// GetAuthorizationModelIdOverride is a helper method to define mock.On call
func (_e *MockSdkClientCheckRequestInterface_Expecter) GetAuthorizationModelIdOverride() *MockSdkClientCheckRequestInterface_GetAuthorizationModelIdOverride_Call {
	return &MockSdkClientCheckRequestInterface_GetAuthorizationModelIdOverride_Call{Call: _e.mock.On("GetAuthorizationModelIdOverride")}
}

func (_c *MockSdkClientCheckRequestInterface_GetAuthorizationModelIdOverride_Call) Run(run func()) *MockSdkClientCheckRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_GetAuthorizationModelIdOverride_Call) Return(_a0 *string) *MockSdkClientCheckRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_GetAuthorizationModelIdOverride_Call) RunAndReturn(run func() *string) *MockSdkClientCheckRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Return(run)
	return _c
}

// GetBody provides a mock function with given fields:
func (_m *MockSdkClientCheckRequestInterface) GetBody() *client.ClientCheckRequest {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBody")
	}

	var r0 *client.ClientCheckRequest
	if rf, ok := ret.Get(0).(func() *client.ClientCheckRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientCheckRequest)
		}
	}

	return r0
}

// MockSdkClientCheckRequestInterface_GetBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBody'
type MockSdkClientCheckRequestInterface_GetBody_Call struct {
	*mock.Call
}

// GetBody is a helper method to define mock.On call
func (_e *MockSdkClientCheckRequestInterface_Expecter) GetBody() *MockSdkClientCheckRequestInterface_GetBody_Call {
	return &MockSdkClientCheckRequestInterface_GetBody_Call{Call: _e.mock.On("GetBody")}
}

func (_c *MockSdkClientCheckRequestInterface_GetBody_Call) Run(run func()) *MockSdkClientCheckRequestInterface_GetBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_GetBody_Call) Return(_a0 *client.ClientCheckRequest) *MockSdkClientCheckRequestInterface_GetBody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_GetBody_Call) RunAndReturn(run func() *client.ClientCheckRequest) *MockSdkClientCheckRequestInterface_GetBody_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields:
func (_m *MockSdkClientCheckRequestInterface) GetContext() context.Context {
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

// MockSdkClientCheckRequestInterface_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockSdkClientCheckRequestInterface_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
func (_e *MockSdkClientCheckRequestInterface_Expecter) GetContext() *MockSdkClientCheckRequestInterface_GetContext_Call {
	return &MockSdkClientCheckRequestInterface_GetContext_Call{Call: _e.mock.On("GetContext")}
}

func (_c *MockSdkClientCheckRequestInterface_GetContext_Call) Run(run func()) *MockSdkClientCheckRequestInterface_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_GetContext_Call) Return(_a0 context.Context) *MockSdkClientCheckRequestInterface_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_GetContext_Call) RunAndReturn(run func() context.Context) *MockSdkClientCheckRequestInterface_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetOptions provides a mock function with given fields:
func (_m *MockSdkClientCheckRequestInterface) GetOptions() *client.ClientCheckOptions {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOptions")
	}

	var r0 *client.ClientCheckOptions
	if rf, ok := ret.Get(0).(func() *client.ClientCheckOptions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientCheckOptions)
		}
	}

	return r0
}

// MockSdkClientCheckRequestInterface_GetOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOptions'
type MockSdkClientCheckRequestInterface_GetOptions_Call struct {
	*mock.Call
}

// GetOptions is a helper method to define mock.On call
func (_e *MockSdkClientCheckRequestInterface_Expecter) GetOptions() *MockSdkClientCheckRequestInterface_GetOptions_Call {
	return &MockSdkClientCheckRequestInterface_GetOptions_Call{Call: _e.mock.On("GetOptions")}
}

func (_c *MockSdkClientCheckRequestInterface_GetOptions_Call) Run(run func()) *MockSdkClientCheckRequestInterface_GetOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_GetOptions_Call) Return(_a0 *client.ClientCheckOptions) *MockSdkClientCheckRequestInterface_GetOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_GetOptions_Call) RunAndReturn(run func() *client.ClientCheckOptions) *MockSdkClientCheckRequestInterface_GetOptions_Call {
	_c.Call.Return(run)
	return _c
}

// Options provides a mock function with given fields: options
func (_m *MockSdkClientCheckRequestInterface) Options(options client.ClientCheckOptions) client.SdkClientCheckRequestInterface {
	ret := _m.Called(options)

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 client.SdkClientCheckRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientCheckOptions) client.SdkClientCheckRequestInterface); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientCheckRequestInterface)
		}
	}

	return r0
}

// MockSdkClientCheckRequestInterface_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type MockSdkClientCheckRequestInterface_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
//   - options client.ClientCheckOptions
func (_e *MockSdkClientCheckRequestInterface_Expecter) Options(options interface{}) *MockSdkClientCheckRequestInterface_Options_Call {
	return &MockSdkClientCheckRequestInterface_Options_Call{Call: _e.mock.On("Options", options)}
}

func (_c *MockSdkClientCheckRequestInterface_Options_Call) Run(run func(options client.ClientCheckOptions)) *MockSdkClientCheckRequestInterface_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientCheckOptions))
	})
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_Options_Call) Return(_a0 client.SdkClientCheckRequestInterface) *MockSdkClientCheckRequestInterface_Options_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientCheckRequestInterface_Options_Call) RunAndReturn(run func(client.ClientCheckOptions) client.SdkClientCheckRequestInterface) *MockSdkClientCheckRequestInterface_Options_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSdkClientCheckRequestInterface creates a new instance of MockSdkClientCheckRequestInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSdkClientCheckRequestInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSdkClientCheckRequestInterface {
	mock := &MockSdkClientCheckRequestInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
