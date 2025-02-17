// Code generated by mockery v2.38.0. DO NOT EDIT.

package mock_subscription

import mock "github.com/stretchr/testify/mock"

// MockCollectionService is an autogenerated mock type for the CollectionService type
type MockCollectionService struct {
	mock.Mock
}

type MockCollectionService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCollectionService) EXPECT() *MockCollectionService_Expecter {
	return &MockCollectionService_Expecter{mock: &_m.Mock}
}

// SubscribeForCollection provides a mock function with given fields: collectionID, subscriptionID
func (_m *MockCollectionService) SubscribeForCollection(collectionID string, subscriptionID string) ([]string, <-chan []string, error) {
	ret := _m.Called(collectionID, subscriptionID)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeForCollection")
	}

	var r0 []string
	var r1 <-chan []string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) ([]string, <-chan []string, error)); ok {
		return rf(collectionID, subscriptionID)
	}
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(collectionID, subscriptionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) <-chan []string); ok {
		r1 = rf(collectionID, subscriptionID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan []string)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(collectionID, subscriptionID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockCollectionService_SubscribeForCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeForCollection'
type MockCollectionService_SubscribeForCollection_Call struct {
	*mock.Call
}

// SubscribeForCollection is a helper method to define mock.On call
//   - collectionID string
//   - subscriptionID string
func (_e *MockCollectionService_Expecter) SubscribeForCollection(collectionID interface{}, subscriptionID interface{}) *MockCollectionService_SubscribeForCollection_Call {
	return &MockCollectionService_SubscribeForCollection_Call{Call: _e.mock.On("SubscribeForCollection", collectionID, subscriptionID)}
}

func (_c *MockCollectionService_SubscribeForCollection_Call) Run(run func(collectionID string, subscriptionID string)) *MockCollectionService_SubscribeForCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockCollectionService_SubscribeForCollection_Call) Return(_a0 []string, _a1 <-chan []string, _a2 error) *MockCollectionService_SubscribeForCollection_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockCollectionService_SubscribeForCollection_Call) RunAndReturn(run func(string, string) ([]string, <-chan []string, error)) *MockCollectionService_SubscribeForCollection_Call {
	_c.Call.Return(run)
	return _c
}

// UnsubscribeFromCollection provides a mock function with given fields: collectionID, subscriptionID
func (_m *MockCollectionService) UnsubscribeFromCollection(collectionID string, subscriptionID string) {
	_m.Called(collectionID, subscriptionID)
}

// MockCollectionService_UnsubscribeFromCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnsubscribeFromCollection'
type MockCollectionService_UnsubscribeFromCollection_Call struct {
	*mock.Call
}

// UnsubscribeFromCollection is a helper method to define mock.On call
//   - collectionID string
//   - subscriptionID string
func (_e *MockCollectionService_Expecter) UnsubscribeFromCollection(collectionID interface{}, subscriptionID interface{}) *MockCollectionService_UnsubscribeFromCollection_Call {
	return &MockCollectionService_UnsubscribeFromCollection_Call{Call: _e.mock.On("UnsubscribeFromCollection", collectionID, subscriptionID)}
}

func (_c *MockCollectionService_UnsubscribeFromCollection_Call) Run(run func(collectionID string, subscriptionID string)) *MockCollectionService_UnsubscribeFromCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockCollectionService_UnsubscribeFromCollection_Call) Return() *MockCollectionService_UnsubscribeFromCollection_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCollectionService_UnsubscribeFromCollection_Call) RunAndReturn(run func(string, string)) *MockCollectionService_UnsubscribeFromCollection_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCollectionService creates a new instance of MockCollectionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCollectionService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCollectionService {
	mock := &MockCollectionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
