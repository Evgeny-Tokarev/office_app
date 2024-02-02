// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	employee_repository "github.com/evgeny-tokarev/office_app/backend/internal/repositories/employee_repository"
	mock "github.com/stretchr/testify/mock"
)

// MockStore is an autogenerated mock type for the Store type
type MockStore struct {
	mock.Mock
}

type MockStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStore) EXPECT() *MockStore_Expecter {
	return &MockStore_Expecter{mock: &_m.Mock}
}

// AttachePhoto provides a mock function with given fields: ctx, arg
func (_m *MockStore) AttachePhoto(ctx context.Context, arg employee_repository.AttachePhotoParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for AttachePhoto")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, employee_repository.AttachePhotoParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_AttachePhoto_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AttachePhoto'
type MockStore_AttachePhoto_Call struct {
	*mock.Call
}

// AttachePhoto is a helper method to define mock.On call
//   - ctx context.Context
//   - arg employee_repository.AttachePhotoParams
func (_e *MockStore_Expecter) AttachePhoto(ctx interface{}, arg interface{}) *MockStore_AttachePhoto_Call {
	return &MockStore_AttachePhoto_Call{Call: _e.mock.On("AttachePhoto", ctx, arg)}
}

func (_c *MockStore_AttachePhoto_Call) Run(run func(ctx context.Context, arg employee_repository.AttachePhotoParams)) *MockStore_AttachePhoto_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(employee_repository.AttachePhotoParams))
	})
	return _c
}

func (_c *MockStore_AttachePhoto_Call) Return(_a0 error) *MockStore_AttachePhoto_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_AttachePhoto_Call) RunAndReturn(run func(context.Context, employee_repository.AttachePhotoParams) error) *MockStore_AttachePhoto_Call {
	_c.Call.Return(run)
	return _c
}

// CreateEmployee provides a mock function with given fields: ctx, arg
func (_m *MockStore) CreateEmployee(ctx context.Context, arg employee_repository.CreateEmployeeParams) (employee_repository.Employee, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateEmployee")
	}

	var r0 employee_repository.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, employee_repository.CreateEmployeeParams) (employee_repository.Employee, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, employee_repository.CreateEmployeeParams) employee_repository.Employee); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(employee_repository.Employee)
	}

	if rf, ok := ret.Get(1).(func(context.Context, employee_repository.CreateEmployeeParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_CreateEmployee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEmployee'
type MockStore_CreateEmployee_Call struct {
	*mock.Call
}

// CreateEmployee is a helper method to define mock.On call
//   - ctx context.Context
//   - arg employee_repository.CreateEmployeeParams
func (_e *MockStore_Expecter) CreateEmployee(ctx interface{}, arg interface{}) *MockStore_CreateEmployee_Call {
	return &MockStore_CreateEmployee_Call{Call: _e.mock.On("CreateEmployee", ctx, arg)}
}

func (_c *MockStore_CreateEmployee_Call) Run(run func(ctx context.Context, arg employee_repository.CreateEmployeeParams)) *MockStore_CreateEmployee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(employee_repository.CreateEmployeeParams))
	})
	return _c
}

func (_c *MockStore_CreateEmployee_Call) Return(_a0 employee_repository.Employee, _a1 error) *MockStore_CreateEmployee_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_CreateEmployee_Call) RunAndReturn(run func(context.Context, employee_repository.CreateEmployeeParams) (employee_repository.Employee, error)) *MockStore_CreateEmployee_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteEmployee provides a mock function with given fields: ctx, id
func (_m *MockStore) DeleteEmployee(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_DeleteEmployee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteEmployee'
type MockStore_DeleteEmployee_Call struct {
	*mock.Call
}

// DeleteEmployee is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockStore_Expecter) DeleteEmployee(ctx interface{}, id interface{}) *MockStore_DeleteEmployee_Call {
	return &MockStore_DeleteEmployee_Call{Call: _e.mock.On("DeleteEmployee", ctx, id)}
}

func (_c *MockStore_DeleteEmployee_Call) Run(run func(ctx context.Context, id int64)) *MockStore_DeleteEmployee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockStore_DeleteEmployee_Call) Return(_a0 error) *MockStore_DeleteEmployee_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_DeleteEmployee_Call) RunAndReturn(run func(context.Context, int64) error) *MockStore_DeleteEmployee_Call {
	_c.Call.Return(run)
	return _c
}

// GetEmployee provides a mock function with given fields: ctx, id
func (_m *MockStore) GetEmployee(ctx context.Context, id int64) (employee_repository.Employee, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetEmployee")
	}

	var r0 employee_repository.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (employee_repository.Employee, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) employee_repository.Employee); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(employee_repository.Employee)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_GetEmployee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEmployee'
type MockStore_GetEmployee_Call struct {
	*mock.Call
}

// GetEmployee is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockStore_Expecter) GetEmployee(ctx interface{}, id interface{}) *MockStore_GetEmployee_Call {
	return &MockStore_GetEmployee_Call{Call: _e.mock.On("GetEmployee", ctx, id)}
}

func (_c *MockStore_GetEmployee_Call) Run(run func(ctx context.Context, id int64)) *MockStore_GetEmployee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockStore_GetEmployee_Call) Return(_a0 employee_repository.Employee, _a1 error) *MockStore_GetEmployee_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_GetEmployee_Call) RunAndReturn(run func(context.Context, int64) (employee_repository.Employee, error)) *MockStore_GetEmployee_Call {
	_c.Call.Return(run)
	return _c
}

// GetImagePath provides a mock function with given fields: ctx, id
func (_m *MockStore) GetImagePath(ctx context.Context, id int64) (string, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetImagePath")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (string, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_GetImagePath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetImagePath'
type MockStore_GetImagePath_Call struct {
	*mock.Call
}

// GetImagePath is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockStore_Expecter) GetImagePath(ctx interface{}, id interface{}) *MockStore_GetImagePath_Call {
	return &MockStore_GetImagePath_Call{Call: _e.mock.On("GetImagePath", ctx, id)}
}

func (_c *MockStore_GetImagePath_Call) Run(run func(ctx context.Context, id int64)) *MockStore_GetImagePath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockStore_GetImagePath_Call) Return(_a0 string, _a1 error) *MockStore_GetImagePath_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_GetImagePath_Call) RunAndReturn(run func(context.Context, int64) (string, error)) *MockStore_GetImagePath_Call {
	_c.Call.Return(run)
	return _c
}

// ListEmployees provides a mock function with given fields: ctx, officeID
func (_m *MockStore) ListEmployees(ctx context.Context, officeID int64) ([]employee_repository.Employee, error) {
	ret := _m.Called(ctx, officeID)

	if len(ret) == 0 {
		panic("no return value specified for ListEmployees")
	}

	var r0 []employee_repository.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]employee_repository.Employee, error)); ok {
		return rf(ctx, officeID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []employee_repository.Employee); ok {
		r0 = rf(ctx, officeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]employee_repository.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, officeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_ListEmployees_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListEmployees'
type MockStore_ListEmployees_Call struct {
	*mock.Call
}

// ListEmployees is a helper method to define mock.On call
//   - ctx context.Context
//   - officeID int64
func (_e *MockStore_Expecter) ListEmployees(ctx interface{}, officeID interface{}) *MockStore_ListEmployees_Call {
	return &MockStore_ListEmployees_Call{Call: _e.mock.On("ListEmployees", ctx, officeID)}
}

func (_c *MockStore_ListEmployees_Call) Run(run func(ctx context.Context, officeID int64)) *MockStore_ListEmployees_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockStore_ListEmployees_Call) Return(_a0 []employee_repository.Employee, _a1 error) *MockStore_ListEmployees_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_ListEmployees_Call) RunAndReturn(run func(context.Context, int64) ([]employee_repository.Employee, error)) *MockStore_ListEmployees_Call {
	_c.Call.Return(run)
	return _c
}

// TransferEmployeeTx provides a mock function with given fields: ctx, arg
func (_m *MockStore) TransferEmployeeTx(ctx context.Context, arg employee_repository.EmployeeTransferTxParams) (employee_repository.EmployeeTransferTxResult, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for TransferEmployeeTx")
	}

	var r0 employee_repository.EmployeeTransferTxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, employee_repository.EmployeeTransferTxParams) (employee_repository.EmployeeTransferTxResult, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, employee_repository.EmployeeTransferTxParams) employee_repository.EmployeeTransferTxResult); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(employee_repository.EmployeeTransferTxResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, employee_repository.EmployeeTransferTxParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_TransferEmployeeTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransferEmployeeTx'
type MockStore_TransferEmployeeTx_Call struct {
	*mock.Call
}

// TransferEmployeeTx is a helper method to define mock.On call
//   - ctx context.Context
//   - arg employee_repository.EmployeeTransferTxParams
func (_e *MockStore_Expecter) TransferEmployeeTx(ctx interface{}, arg interface{}) *MockStore_TransferEmployeeTx_Call {
	return &MockStore_TransferEmployeeTx_Call{Call: _e.mock.On("TransferEmployeeTx", ctx, arg)}
}

func (_c *MockStore_TransferEmployeeTx_Call) Run(run func(ctx context.Context, arg employee_repository.EmployeeTransferTxParams)) *MockStore_TransferEmployeeTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(employee_repository.EmployeeTransferTxParams))
	})
	return _c
}

func (_c *MockStore_TransferEmployeeTx_Call) Return(_a0 employee_repository.EmployeeTransferTxResult, _a1 error) *MockStore_TransferEmployeeTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_TransferEmployeeTx_Call) RunAndReturn(run func(context.Context, employee_repository.EmployeeTransferTxParams) (employee_repository.EmployeeTransferTxResult, error)) *MockStore_TransferEmployeeTx_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateEmployee provides a mock function with given fields: ctx, arg
func (_m *MockStore) UpdateEmployee(ctx context.Context, arg employee_repository.UpdateEmployeeParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, employee_repository.UpdateEmployeeParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_UpdateEmployee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEmployee'
type MockStore_UpdateEmployee_Call struct {
	*mock.Call
}

// UpdateEmployee is a helper method to define mock.On call
//   - ctx context.Context
//   - arg employee_repository.UpdateEmployeeParams
func (_e *MockStore_Expecter) UpdateEmployee(ctx interface{}, arg interface{}) *MockStore_UpdateEmployee_Call {
	return &MockStore_UpdateEmployee_Call{Call: _e.mock.On("UpdateEmployee", ctx, arg)}
}

func (_c *MockStore_UpdateEmployee_Call) Run(run func(ctx context.Context, arg employee_repository.UpdateEmployeeParams)) *MockStore_UpdateEmployee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(employee_repository.UpdateEmployeeParams))
	})
	return _c
}

func (_c *MockStore_UpdateEmployee_Call) Return(_a0 error) *MockStore_UpdateEmployee_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_UpdateEmployee_Call) RunAndReturn(run func(context.Context, employee_repository.UpdateEmployeeParams) error) *MockStore_UpdateEmployee_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStore creates a new instance of MockStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStore {
	mock := &MockStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
