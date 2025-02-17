// Code generated by mockery v2.38.0. DO NOT EDIT.

package mock_wallet

import (
	app "github.com/anyproto/any-sync/app"
	accountdata "github.com/anyproto/any-sync/commonspace/object/accountdata"

	crypto "github.com/anyproto/any-sync/util/crypto"

	mock "github.com/stretchr/testify/mock"
)

// MockWallet is an autogenerated mock type for the Wallet type
type MockWallet struct {
	mock.Mock
}

type MockWallet_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWallet) EXPECT() *MockWallet_Expecter {
	return &MockWallet_Expecter{mock: &_m.Mock}
}

// Account provides a mock function with given fields:
func (_m *MockWallet) Account() *accountdata.AccountKeys {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Account")
	}

	var r0 *accountdata.AccountKeys
	if rf, ok := ret.Get(0).(func() *accountdata.AccountKeys); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*accountdata.AccountKeys)
		}
	}

	return r0
}

// MockWallet_Account_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Account'
type MockWallet_Account_Call struct {
	*mock.Call
}

// Account is a helper method to define mock.On call
func (_e *MockWallet_Expecter) Account() *MockWallet_Account_Call {
	return &MockWallet_Account_Call{Call: _e.mock.On("Account")}
}

func (_c *MockWallet_Account_Call) Run(run func()) *MockWallet_Account_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWallet_Account_Call) Return(_a0 *accountdata.AccountKeys) *MockWallet_Account_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWallet_Account_Call) RunAndReturn(run func() *accountdata.AccountKeys) *MockWallet_Account_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccountPrivkey provides a mock function with given fields:
func (_m *MockWallet) GetAccountPrivkey() crypto.PrivKey {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAccountPrivkey")
	}

	var r0 crypto.PrivKey
	if rf, ok := ret.Get(0).(func() crypto.PrivKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivKey)
		}
	}

	return r0
}

// MockWallet_GetAccountPrivkey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountPrivkey'
type MockWallet_GetAccountPrivkey_Call struct {
	*mock.Call
}

// GetAccountPrivkey is a helper method to define mock.On call
func (_e *MockWallet_Expecter) GetAccountPrivkey() *MockWallet_GetAccountPrivkey_Call {
	return &MockWallet_GetAccountPrivkey_Call{Call: _e.mock.On("GetAccountPrivkey")}
}

func (_c *MockWallet_GetAccountPrivkey_Call) Run(run func()) *MockWallet_GetAccountPrivkey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWallet_GetAccountPrivkey_Call) Return(_a0 crypto.PrivKey) *MockWallet_GetAccountPrivkey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWallet_GetAccountPrivkey_Call) RunAndReturn(run func() crypto.PrivKey) *MockWallet_GetAccountPrivkey_Call {
	_c.Call.Return(run)
	return _c
}

// GetDevicePrivkey provides a mock function with given fields:
func (_m *MockWallet) GetDevicePrivkey() crypto.PrivKey {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDevicePrivkey")
	}

	var r0 crypto.PrivKey
	if rf, ok := ret.Get(0).(func() crypto.PrivKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivKey)
		}
	}

	return r0
}

// MockWallet_GetDevicePrivkey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDevicePrivkey'
type MockWallet_GetDevicePrivkey_Call struct {
	*mock.Call
}

// GetDevicePrivkey is a helper method to define mock.On call
func (_e *MockWallet_Expecter) GetDevicePrivkey() *MockWallet_GetDevicePrivkey_Call {
	return &MockWallet_GetDevicePrivkey_Call{Call: _e.mock.On("GetDevicePrivkey")}
}

func (_c *MockWallet_GetDevicePrivkey_Call) Run(run func()) *MockWallet_GetDevicePrivkey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWallet_GetDevicePrivkey_Call) Return(_a0 crypto.PrivKey) *MockWallet_GetDevicePrivkey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWallet_GetDevicePrivkey_Call) RunAndReturn(run func() crypto.PrivKey) *MockWallet_GetDevicePrivkey_Call {
	_c.Call.Return(run)
	return _c
}

// GetMasterKey provides a mock function with given fields:
func (_m *MockWallet) GetMasterKey() crypto.PrivKey {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMasterKey")
	}

	var r0 crypto.PrivKey
	if rf, ok := ret.Get(0).(func() crypto.PrivKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivKey)
		}
	}

	return r0
}

// MockWallet_GetMasterKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMasterKey'
type MockWallet_GetMasterKey_Call struct {
	*mock.Call
}

// GetMasterKey is a helper method to define mock.On call
func (_e *MockWallet_Expecter) GetMasterKey() *MockWallet_GetMasterKey_Call {
	return &MockWallet_GetMasterKey_Call{Call: _e.mock.On("GetMasterKey")}
}

func (_c *MockWallet_GetMasterKey_Call) Run(run func()) *MockWallet_GetMasterKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWallet_GetMasterKey_Call) Return(_a0 crypto.PrivKey) *MockWallet_GetMasterKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWallet_GetMasterKey_Call) RunAndReturn(run func() crypto.PrivKey) *MockWallet_GetMasterKey_Call {
	_c.Call.Return(run)
	return _c
}

// GetOldAccountKey provides a mock function with given fields:
func (_m *MockWallet) GetOldAccountKey() crypto.PrivKey {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOldAccountKey")
	}

	var r0 crypto.PrivKey
	if rf, ok := ret.Get(0).(func() crypto.PrivKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivKey)
		}
	}

	return r0
}

// MockWallet_GetOldAccountKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOldAccountKey'
type MockWallet_GetOldAccountKey_Call struct {
	*mock.Call
}

// GetOldAccountKey is a helper method to define mock.On call
func (_e *MockWallet_Expecter) GetOldAccountKey() *MockWallet_GetOldAccountKey_Call {
	return &MockWallet_GetOldAccountKey_Call{Call: _e.mock.On("GetOldAccountKey")}
}

func (_c *MockWallet_GetOldAccountKey_Call) Run(run func()) *MockWallet_GetOldAccountKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWallet_GetOldAccountKey_Call) Return(_a0 crypto.PrivKey) *MockWallet_GetOldAccountKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWallet_GetOldAccountKey_Call) RunAndReturn(run func() crypto.PrivKey) *MockWallet_GetOldAccountKey_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockWallet) Init(a *app.App) error {
	ret := _m.Called(a)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWallet_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockWallet_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockWallet_Expecter) Init(a interface{}) *MockWallet_Init_Call {
	return &MockWallet_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockWallet_Init_Call) Run(run func(a *app.App)) *MockWallet_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockWallet_Init_Call) Return(err error) *MockWallet_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockWallet_Init_Call) RunAndReturn(run func(*app.App) error) *MockWallet_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockWallet) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockWallet_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockWallet_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockWallet_Expecter) Name() *MockWallet_Name_Call {
	return &MockWallet_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockWallet_Name_Call) Run(run func()) *MockWallet_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWallet_Name_Call) Return(name string) *MockWallet_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockWallet_Name_Call) RunAndReturn(run func() string) *MockWallet_Name_Call {
	_c.Call.Return(run)
	return _c
}

// RepoPath provides a mock function with given fields:
func (_m *MockWallet) RepoPath() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RepoPath")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockWallet_RepoPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RepoPath'
type MockWallet_RepoPath_Call struct {
	*mock.Call
}

// RepoPath is a helper method to define mock.On call
func (_e *MockWallet_Expecter) RepoPath() *MockWallet_RepoPath_Call {
	return &MockWallet_RepoPath_Call{Call: _e.mock.On("RepoPath")}
}

func (_c *MockWallet_RepoPath_Call) Run(run func()) *MockWallet_RepoPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWallet_RepoPath_Call) Return(_a0 string) *MockWallet_RepoPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWallet_RepoPath_Call) RunAndReturn(run func() string) *MockWallet_RepoPath_Call {
	_c.Call.Return(run)
	return _c
}

// RootPath provides a mock function with given fields:
func (_m *MockWallet) RootPath() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RootPath")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockWallet_RootPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RootPath'
type MockWallet_RootPath_Call struct {
	*mock.Call
}

// RootPath is a helper method to define mock.On call
func (_e *MockWallet_Expecter) RootPath() *MockWallet_RootPath_Call {
	return &MockWallet_RootPath_Call{Call: _e.mock.On("RootPath")}
}

func (_c *MockWallet_RootPath_Call) Run(run func()) *MockWallet_RootPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWallet_RootPath_Call) Return(_a0 string) *MockWallet_RootPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWallet_RootPath_Call) RunAndReturn(run func() string) *MockWallet_RootPath_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWallet creates a new instance of MockWallet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWallet(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWallet {
	mock := &MockWallet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
