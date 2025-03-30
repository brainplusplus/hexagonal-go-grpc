// Code generated by mockery v2.53.3. DO NOT EDIT.

package db

import (
	context "context"
	entity "hexagonal-go-grpc/internal/adapters/service/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockICustomerDbRepository is an autogenerated mock type for the ICustomerDbRepository type
type MockICustomerDbRepository struct {
	mock.Mock
}

type MockICustomerDbRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockICustomerDbRepository) EXPECT() *MockICustomerDbRepository_Expecter {
	return &MockICustomerDbRepository_Expecter{mock: &_m.Mock}
}

// ActivateById provides a mock function with given fields: ctx, request
func (_m *MockICustomerDbRepository) ActivateById(ctx context.Context, request entity.CustomerIdRequest) (entity.Customer, int, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ActivateById")
	}

	var r0 entity.Customer
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerIdRequest) (entity.Customer, int, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerIdRequest) entity.Customer); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(entity.Customer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.CustomerIdRequest) int); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, entity.CustomerIdRequest) error); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockICustomerDbRepository_ActivateById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ActivateById'
type MockICustomerDbRepository_ActivateById_Call struct {
	*mock.Call
}

// ActivateById is a helper method to define mock.On call
//   - ctx context.Context
//   - request entity.CustomerIdRequest
func (_e *MockICustomerDbRepository_Expecter) ActivateById(ctx interface{}, request interface{}) *MockICustomerDbRepository_ActivateById_Call {
	return &MockICustomerDbRepository_ActivateById_Call{Call: _e.mock.On("ActivateById", ctx, request)}
}

func (_c *MockICustomerDbRepository_ActivateById_Call) Run(run func(ctx context.Context, request entity.CustomerIdRequest)) *MockICustomerDbRepository_ActivateById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.CustomerIdRequest))
	})
	return _c
}

func (_c *MockICustomerDbRepository_ActivateById_Call) Return(_a0 entity.Customer, _a1 int, _a2 error) *MockICustomerDbRepository_ActivateById_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockICustomerDbRepository_ActivateById_Call) RunAndReturn(run func(context.Context, entity.CustomerIdRequest) (entity.Customer, int, error)) *MockICustomerDbRepository_ActivateById_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockICustomerDbRepository) Create(ctx context.Context, request entity.CustomerCreateRequest) (entity.Customer, int, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 entity.Customer
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerCreateRequest) (entity.Customer, int, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerCreateRequest) entity.Customer); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(entity.Customer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.CustomerCreateRequest) int); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, entity.CustomerCreateRequest) error); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockICustomerDbRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockICustomerDbRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request entity.CustomerCreateRequest
func (_e *MockICustomerDbRepository_Expecter) Create(ctx interface{}, request interface{}) *MockICustomerDbRepository_Create_Call {
	return &MockICustomerDbRepository_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockICustomerDbRepository_Create_Call) Run(run func(ctx context.Context, request entity.CustomerCreateRequest)) *MockICustomerDbRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.CustomerCreateRequest))
	})
	return _c
}

func (_c *MockICustomerDbRepository_Create_Call) Return(_a0 entity.Customer, _a1 int, _a2 error) *MockICustomerDbRepository_Create_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockICustomerDbRepository_Create_Call) RunAndReturn(run func(context.Context, entity.CustomerCreateRequest) (entity.Customer, int, error)) *MockICustomerDbRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// DeactivateById provides a mock function with given fields: ctx, request
func (_m *MockICustomerDbRepository) DeactivateById(ctx context.Context, request entity.CustomerIdRequest) (entity.Customer, int, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeactivateById")
	}

	var r0 entity.Customer
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerIdRequest) (entity.Customer, int, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerIdRequest) entity.Customer); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(entity.Customer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.CustomerIdRequest) int); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, entity.CustomerIdRequest) error); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockICustomerDbRepository_DeactivateById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeactivateById'
type MockICustomerDbRepository_DeactivateById_Call struct {
	*mock.Call
}

// DeactivateById is a helper method to define mock.On call
//   - ctx context.Context
//   - request entity.CustomerIdRequest
func (_e *MockICustomerDbRepository_Expecter) DeactivateById(ctx interface{}, request interface{}) *MockICustomerDbRepository_DeactivateById_Call {
	return &MockICustomerDbRepository_DeactivateById_Call{Call: _e.mock.On("DeactivateById", ctx, request)}
}

func (_c *MockICustomerDbRepository_DeactivateById_Call) Run(run func(ctx context.Context, request entity.CustomerIdRequest)) *MockICustomerDbRepository_DeactivateById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.CustomerIdRequest))
	})
	return _c
}

func (_c *MockICustomerDbRepository_DeactivateById_Call) Return(_a0 entity.Customer, _a1 int, _a2 error) *MockICustomerDbRepository_DeactivateById_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockICustomerDbRepository_DeactivateById_Call) RunAndReturn(run func(context.Context, entity.CustomerIdRequest) (entity.Customer, int, error)) *MockICustomerDbRepository_DeactivateById_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, request
func (_m *MockICustomerDbRepository) DeleteById(ctx context.Context, request entity.CustomerIdRequest) (int, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteById")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerIdRequest) (int, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerIdRequest) int); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.CustomerIdRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockICustomerDbRepository_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockICustomerDbRepository_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - request entity.CustomerIdRequest
func (_e *MockICustomerDbRepository_Expecter) DeleteById(ctx interface{}, request interface{}) *MockICustomerDbRepository_DeleteById_Call {
	return &MockICustomerDbRepository_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, request)}
}

func (_c *MockICustomerDbRepository_DeleteById_Call) Run(run func(ctx context.Context, request entity.CustomerIdRequest)) *MockICustomerDbRepository_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.CustomerIdRequest))
	})
	return _c
}

func (_c *MockICustomerDbRepository_DeleteById_Call) Return(_a0 int, _a1 error) *MockICustomerDbRepository_DeleteById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockICustomerDbRepository_DeleteById_Call) RunAndReturn(run func(context.Context, entity.CustomerIdRequest) (int, error)) *MockICustomerDbRepository_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// GetByEmailOrPhone provides a mock function with given fields: ctx, request
func (_m *MockICustomerDbRepository) GetByEmailOrPhone(ctx context.Context, request entity.CustomerEmailOrPhoneRequest) (entity.Customer, int, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmailOrPhone")
	}

	var r0 entity.Customer
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerEmailOrPhoneRequest) (entity.Customer, int, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerEmailOrPhoneRequest) entity.Customer); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(entity.Customer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.CustomerEmailOrPhoneRequest) int); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, entity.CustomerEmailOrPhoneRequest) error); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockICustomerDbRepository_GetByEmailOrPhone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmailOrPhone'
type MockICustomerDbRepository_GetByEmailOrPhone_Call struct {
	*mock.Call
}

// GetByEmailOrPhone is a helper method to define mock.On call
//   - ctx context.Context
//   - request entity.CustomerEmailOrPhoneRequest
func (_e *MockICustomerDbRepository_Expecter) GetByEmailOrPhone(ctx interface{}, request interface{}) *MockICustomerDbRepository_GetByEmailOrPhone_Call {
	return &MockICustomerDbRepository_GetByEmailOrPhone_Call{Call: _e.mock.On("GetByEmailOrPhone", ctx, request)}
}

func (_c *MockICustomerDbRepository_GetByEmailOrPhone_Call) Run(run func(ctx context.Context, request entity.CustomerEmailOrPhoneRequest)) *MockICustomerDbRepository_GetByEmailOrPhone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.CustomerEmailOrPhoneRequest))
	})
	return _c
}

func (_c *MockICustomerDbRepository_GetByEmailOrPhone_Call) Return(_a0 entity.Customer, _a1 int, _a2 error) *MockICustomerDbRepository_GetByEmailOrPhone_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockICustomerDbRepository_GetByEmailOrPhone_Call) RunAndReturn(run func(context.Context, entity.CustomerEmailOrPhoneRequest) (entity.Customer, int, error)) *MockICustomerDbRepository_GetByEmailOrPhone_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, request
func (_m *MockICustomerDbRepository) GetById(ctx context.Context, request entity.CustomerIdRequest) (entity.Customer, int, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 entity.Customer
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerIdRequest) (entity.Customer, int, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerIdRequest) entity.Customer); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(entity.Customer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.CustomerIdRequest) int); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, entity.CustomerIdRequest) error); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockICustomerDbRepository_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockICustomerDbRepository_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - request entity.CustomerIdRequest
func (_e *MockICustomerDbRepository_Expecter) GetById(ctx interface{}, request interface{}) *MockICustomerDbRepository_GetById_Call {
	return &MockICustomerDbRepository_GetById_Call{Call: _e.mock.On("GetById", ctx, request)}
}

func (_c *MockICustomerDbRepository_GetById_Call) Run(run func(ctx context.Context, request entity.CustomerIdRequest)) *MockICustomerDbRepository_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.CustomerIdRequest))
	})
	return _c
}

func (_c *MockICustomerDbRepository_GetById_Call) Return(_a0 entity.Customer, _a1 int, _a2 error) *MockICustomerDbRepository_GetById_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockICustomerDbRepository_GetById_Call) RunAndReturn(run func(context.Context, entity.CustomerIdRequest) (entity.Customer, int, error)) *MockICustomerDbRepository_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: ctx, request
func (_m *MockICustomerDbRepository) GetList(ctx context.Context, request entity.CustomerListRequest) ([]entity.Customer, int, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetList")
	}

	var r0 []entity.Customer
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerListRequest) ([]entity.Customer, int, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerListRequest) []entity.Customer); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.CustomerListRequest) int); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, entity.CustomerListRequest) error); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockICustomerDbRepository_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockICustomerDbRepository_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - ctx context.Context
//   - request entity.CustomerListRequest
func (_e *MockICustomerDbRepository_Expecter) GetList(ctx interface{}, request interface{}) *MockICustomerDbRepository_GetList_Call {
	return &MockICustomerDbRepository_GetList_Call{Call: _e.mock.On("GetList", ctx, request)}
}

func (_c *MockICustomerDbRepository_GetList_Call) Run(run func(ctx context.Context, request entity.CustomerListRequest)) *MockICustomerDbRepository_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.CustomerListRequest))
	})
	return _c
}

func (_c *MockICustomerDbRepository_GetList_Call) Return(_a0 []entity.Customer, _a1 int, _a2 error) *MockICustomerDbRepository_GetList_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockICustomerDbRepository_GetList_Call) RunAndReturn(run func(context.Context, entity.CustomerListRequest) ([]entity.Customer, int, error)) *MockICustomerDbRepository_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockICustomerDbRepository) Update(ctx context.Context, request entity.CustomerUpdateRequest) (entity.Customer, int, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 entity.Customer
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerUpdateRequest) (entity.Customer, int, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.CustomerUpdateRequest) entity.Customer); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(entity.Customer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.CustomerUpdateRequest) int); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, entity.CustomerUpdateRequest) error); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockICustomerDbRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockICustomerDbRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request entity.CustomerUpdateRequest
func (_e *MockICustomerDbRepository_Expecter) Update(ctx interface{}, request interface{}) *MockICustomerDbRepository_Update_Call {
	return &MockICustomerDbRepository_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockICustomerDbRepository_Update_Call) Run(run func(ctx context.Context, request entity.CustomerUpdateRequest)) *MockICustomerDbRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.CustomerUpdateRequest))
	})
	return _c
}

func (_c *MockICustomerDbRepository_Update_Call) Return(_a0 entity.Customer, _a1 int, _a2 error) *MockICustomerDbRepository_Update_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockICustomerDbRepository_Update_Call) RunAndReturn(run func(context.Context, entity.CustomerUpdateRequest) (entity.Customer, int, error)) *MockICustomerDbRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockICustomerDbRepository creates a new instance of MockICustomerDbRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockICustomerDbRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockICustomerDbRepository {
	mock := &MockICustomerDbRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
