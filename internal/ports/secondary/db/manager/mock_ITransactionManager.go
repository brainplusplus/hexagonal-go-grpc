// Code generated by mockery v2.53.3. DO NOT EDIT.

package manager

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockITransactionManager is an autogenerated mock type for the ITransactionManager type
type MockITransactionManager struct {
	mock.Mock
}

type MockITransactionManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockITransactionManager) EXPECT() *MockITransactionManager_Expecter {
	return &MockITransactionManager_Expecter{mock: &_m.Mock}
}

// WithTransaction provides a mock function with given fields: fn, ctx, args
func (_m *MockITransactionManager) WithTransaction(fn func(context.Context, ...interface{}) ([]interface{}, int, error), ctx context.Context, args ...interface{}) ([]interface{}, int, error) {
	var _ca []interface{}
	_ca = append(_ca, fn, ctx)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for WithTransaction")
	}

	var r0 []interface{}
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(func(context.Context, ...interface{}) ([]interface{}, int, error), context.Context, ...interface{}) ([]interface{}, int, error)); ok {
		return rf(fn, ctx, args...)
	}
	if rf, ok := ret.Get(0).(func(func(context.Context, ...interface{}) ([]interface{}, int, error), context.Context, ...interface{}) []interface{}); ok {
		r0 = rf(fn, ctx, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(func(context.Context, ...interface{}) ([]interface{}, int, error), context.Context, ...interface{}) int); ok {
		r1 = rf(fn, ctx, args...)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(func(context.Context, ...interface{}) ([]interface{}, int, error), context.Context, ...interface{}) error); ok {
		r2 = rf(fn, ctx, args...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockITransactionManager_WithTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTransaction'
type MockITransactionManager_WithTransaction_Call struct {
	*mock.Call
}

// WithTransaction is a helper method to define mock.On call
//   - fn func(context.Context , ...interface{})([]interface{} , int , error)
//   - ctx context.Context
//   - args ...interface{}
func (_e *MockITransactionManager_Expecter) WithTransaction(fn interface{}, ctx interface{}, args ...interface{}) *MockITransactionManager_WithTransaction_Call {
	return &MockITransactionManager_WithTransaction_Call{Call: _e.mock.On("WithTransaction",
		append([]interface{}{fn, ctx}, args...)...)}
}

func (_c *MockITransactionManager_WithTransaction_Call) Run(run func(fn func(context.Context, ...interface{}) ([]interface{}, int, error), ctx context.Context, args ...interface{})) *MockITransactionManager_WithTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(func(context.Context, ...interface{}) ([]interface{}, int, error)), args[1].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockITransactionManager_WithTransaction_Call) Return(_a0 []interface{}, _a1 int, _a2 error) *MockITransactionManager_WithTransaction_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockITransactionManager_WithTransaction_Call) RunAndReturn(run func(func(context.Context, ...interface{}) ([]interface{}, int, error), context.Context, ...interface{}) ([]interface{}, int, error)) *MockITransactionManager_WithTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockITransactionManager creates a new instance of MockITransactionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockITransactionManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockITransactionManager {
	mock := &MockITransactionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
