// Code generated by mockery v2.38.0. DO NOT EDIT.

package mock_objectstore

import mock "github.com/stretchr/testify/mock"

// MockVirtualSpacesStore is an autogenerated mock type for the VirtualSpacesStore type
type MockVirtualSpacesStore struct {
	mock.Mock
}

type MockVirtualSpacesStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVirtualSpacesStore) EXPECT() *MockVirtualSpacesStore_Expecter {
	return &MockVirtualSpacesStore_Expecter{mock: &_m.Mock}
}

// DeleteVirtualSpace provides a mock function with given fields: spaceID
func (_m *MockVirtualSpacesStore) DeleteVirtualSpace(spaceID string) error {
	ret := _m.Called(spaceID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteVirtualSpace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(spaceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockVirtualSpacesStore_DeleteVirtualSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteVirtualSpace'
type MockVirtualSpacesStore_DeleteVirtualSpace_Call struct {
	*mock.Call
}

// DeleteVirtualSpace is a helper method to define mock.On call
//   - spaceID string
func (_e *MockVirtualSpacesStore_Expecter) DeleteVirtualSpace(spaceID interface{}) *MockVirtualSpacesStore_DeleteVirtualSpace_Call {
	return &MockVirtualSpacesStore_DeleteVirtualSpace_Call{Call: _e.mock.On("DeleteVirtualSpace", spaceID)}
}

func (_c *MockVirtualSpacesStore_DeleteVirtualSpace_Call) Run(run func(spaceID string)) *MockVirtualSpacesStore_DeleteVirtualSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockVirtualSpacesStore_DeleteVirtualSpace_Call) Return(_a0 error) *MockVirtualSpacesStore_DeleteVirtualSpace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVirtualSpacesStore_DeleteVirtualSpace_Call) RunAndReturn(run func(string) error) *MockVirtualSpacesStore_DeleteVirtualSpace_Call {
	_c.Call.Return(run)
	return _c
}

// ListVirtualSpaces provides a mock function with given fields:
func (_m *MockVirtualSpacesStore) ListVirtualSpaces() ([]string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListVirtualSpaces")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVirtualSpacesStore_ListVirtualSpaces_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListVirtualSpaces'
type MockVirtualSpacesStore_ListVirtualSpaces_Call struct {
	*mock.Call
}

// ListVirtualSpaces is a helper method to define mock.On call
func (_e *MockVirtualSpacesStore_Expecter) ListVirtualSpaces() *MockVirtualSpacesStore_ListVirtualSpaces_Call {
	return &MockVirtualSpacesStore_ListVirtualSpaces_Call{Call: _e.mock.On("ListVirtualSpaces")}
}

func (_c *MockVirtualSpacesStore_ListVirtualSpaces_Call) Run(run func()) *MockVirtualSpacesStore_ListVirtualSpaces_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockVirtualSpacesStore_ListVirtualSpaces_Call) Return(_a0 []string, _a1 error) *MockVirtualSpacesStore_ListVirtualSpaces_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVirtualSpacesStore_ListVirtualSpaces_Call) RunAndReturn(run func() ([]string, error)) *MockVirtualSpacesStore_ListVirtualSpaces_Call {
	_c.Call.Return(run)
	return _c
}

// SaveVirtualSpace provides a mock function with given fields: id
func (_m *MockVirtualSpacesStore) SaveVirtualSpace(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for SaveVirtualSpace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockVirtualSpacesStore_SaveVirtualSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveVirtualSpace'
type MockVirtualSpacesStore_SaveVirtualSpace_Call struct {
	*mock.Call
}

// SaveVirtualSpace is a helper method to define mock.On call
//   - id string
func (_e *MockVirtualSpacesStore_Expecter) SaveVirtualSpace(id interface{}) *MockVirtualSpacesStore_SaveVirtualSpace_Call {
	return &MockVirtualSpacesStore_SaveVirtualSpace_Call{Call: _e.mock.On("SaveVirtualSpace", id)}
}

func (_c *MockVirtualSpacesStore_SaveVirtualSpace_Call) Run(run func(id string)) *MockVirtualSpacesStore_SaveVirtualSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockVirtualSpacesStore_SaveVirtualSpace_Call) Return(_a0 error) *MockVirtualSpacesStore_SaveVirtualSpace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVirtualSpacesStore_SaveVirtualSpace_Call) RunAndReturn(run func(string) error) *MockVirtualSpacesStore_SaveVirtualSpace_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVirtualSpacesStore creates a new instance of MockVirtualSpacesStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVirtualSpacesStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVirtualSpacesStore {
	mock := &MockVirtualSpacesStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
