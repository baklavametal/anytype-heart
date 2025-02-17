// Code generated by mockery v2.38.0. DO NOT EDIT.

package mock_database

import (
	model "github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	mock "github.com/stretchr/testify/mock"
)

// MockSchema is an autogenerated mock type for the Schema type
type MockSchema struct {
	mock.Mock
}

type MockSchema_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSchema) EXPECT() *MockSchema_Expecter {
	return &MockSchema_Expecter{mock: &_m.Mock}
}

// ListRelations provides a mock function with given fields:
func (_m *MockSchema) ListRelations() []*model.RelationLink {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListRelations")
	}

	var r0 []*model.RelationLink
	if rf, ok := ret.Get(0).(func() []*model.RelationLink); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RelationLink)
		}
	}

	return r0
}

// MockSchema_ListRelations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListRelations'
type MockSchema_ListRelations_Call struct {
	*mock.Call
}

// ListRelations is a helper method to define mock.On call
func (_e *MockSchema_Expecter) ListRelations() *MockSchema_ListRelations_Call {
	return &MockSchema_ListRelations_Call{Call: _e.mock.On("ListRelations")}
}

func (_c *MockSchema_ListRelations_Call) Run(run func()) *MockSchema_ListRelations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSchema_ListRelations_Call) Return(_a0 []*model.RelationLink) *MockSchema_ListRelations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSchema_ListRelations_Call) RunAndReturn(run func() []*model.RelationLink) *MockSchema_ListRelations_Call {
	_c.Call.Return(run)
	return _c
}

// RequiredRelations provides a mock function with given fields:
func (_m *MockSchema) RequiredRelations() []*model.RelationLink {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequiredRelations")
	}

	var r0 []*model.RelationLink
	if rf, ok := ret.Get(0).(func() []*model.RelationLink); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RelationLink)
		}
	}

	return r0
}

// MockSchema_RequiredRelations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequiredRelations'
type MockSchema_RequiredRelations_Call struct {
	*mock.Call
}

// RequiredRelations is a helper method to define mock.On call
func (_e *MockSchema_Expecter) RequiredRelations() *MockSchema_RequiredRelations_Call {
	return &MockSchema_RequiredRelations_Call{Call: _e.mock.On("RequiredRelations")}
}

func (_c *MockSchema_RequiredRelations_Call) Run(run func()) *MockSchema_RequiredRelations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSchema_RequiredRelations_Call) Return(_a0 []*model.RelationLink) *MockSchema_RequiredRelations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSchema_RequiredRelations_Call) RunAndReturn(run func() []*model.RelationLink) *MockSchema_RequiredRelations_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSchema creates a new instance of MockSchema. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSchema(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSchema {
	mock := &MockSchema{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
