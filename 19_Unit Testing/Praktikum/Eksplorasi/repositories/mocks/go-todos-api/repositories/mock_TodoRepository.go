// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "go-todos-api/models"

	mock "github.com/stretchr/testify/mock"
)

// MockTodoRepository is an autogenerated mock type for the TodoRepository type
type MockTodoRepository struct {
	mock.Mock
}

type MockTodoRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTodoRepository) EXPECT() *MockTodoRepository_Expecter {
	return &MockTodoRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: categoryInput
func (_m *MockTodoRepository) Create(categoryInput models.TodoInput) (models.Todo, error) {
	ret := _m.Called(categoryInput)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(models.TodoInput) (models.Todo, error)); ok {
		return rf(categoryInput)
	}
	if rf, ok := ret.Get(0).(func(models.TodoInput) models.Todo); ok {
		r0 = rf(categoryInput)
	} else {
		r0 = ret.Get(0).(models.Todo)
	}

	if rf, ok := ret.Get(1).(func(models.TodoInput) error); ok {
		r1 = rf(categoryInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTodoRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockTodoRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - categoryInput models.TodoInput
func (_e *MockTodoRepository_Expecter) Create(categoryInput interface{}) *MockTodoRepository_Create_Call {
	return &MockTodoRepository_Create_Call{Call: _e.mock.On("Create", categoryInput)}
}

func (_c *MockTodoRepository_Create_Call) Run(run func(categoryInput models.TodoInput)) *MockTodoRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.TodoInput))
	})
	return _c
}

func (_c *MockTodoRepository_Create_Call) Return(_a0 models.Todo, _a1 error) *MockTodoRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTodoRepository_Create_Call) RunAndReturn(run func(models.TodoInput) (models.Todo, error)) *MockTodoRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *MockTodoRepository) Delete(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTodoRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockTodoRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id string
func (_e *MockTodoRepository_Expecter) Delete(id interface{}) *MockTodoRepository_Delete_Call {
	return &MockTodoRepository_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *MockTodoRepository_Delete_Call) Run(run func(id string)) *MockTodoRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTodoRepository_Delete_Call) Return(_a0 error) *MockTodoRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTodoRepository_Delete_Call) RunAndReturn(run func(string) error) *MockTodoRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *MockTodoRepository) GetAll() ([]models.Todo, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Todo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTodoRepository_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockTodoRepository_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *MockTodoRepository_Expecter) GetAll() *MockTodoRepository_GetAll_Call {
	return &MockTodoRepository_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *MockTodoRepository_GetAll_Call) Run(run func()) *MockTodoRepository_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTodoRepository_GetAll_Call) Return(_a0 []models.Todo, _a1 error) *MockTodoRepository_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTodoRepository_GetAll_Call) RunAndReturn(run func() ([]models.Todo, error)) *MockTodoRepository_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: id
func (_m *MockTodoRepository) GetByID(id string) (models.Todo, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Todo, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Todo); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Todo)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTodoRepository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type MockTodoRepository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - id string
func (_e *MockTodoRepository_Expecter) GetByID(id interface{}) *MockTodoRepository_GetByID_Call {
	return &MockTodoRepository_GetByID_Call{Call: _e.mock.On("GetByID", id)}
}

func (_c *MockTodoRepository_GetByID_Call) Run(run func(id string)) *MockTodoRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTodoRepository_GetByID_Call) Return(_a0 models.Todo, _a1 error) *MockTodoRepository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTodoRepository_GetByID_Call) RunAndReturn(run func(string) (models.Todo, error)) *MockTodoRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: categoryInput, id
func (_m *MockTodoRepository) Update(categoryInput models.TodoInput, id string) (models.Todo, error) {
	ret := _m.Called(categoryInput, id)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(models.TodoInput, string) (models.Todo, error)); ok {
		return rf(categoryInput, id)
	}
	if rf, ok := ret.Get(0).(func(models.TodoInput, string) models.Todo); ok {
		r0 = rf(categoryInput, id)
	} else {
		r0 = ret.Get(0).(models.Todo)
	}

	if rf, ok := ret.Get(1).(func(models.TodoInput, string) error); ok {
		r1 = rf(categoryInput, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTodoRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockTodoRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - categoryInput models.TodoInput
//   - id string
func (_e *MockTodoRepository_Expecter) Update(categoryInput interface{}, id interface{}) *MockTodoRepository_Update_Call {
	return &MockTodoRepository_Update_Call{Call: _e.mock.On("Update", categoryInput, id)}
}

func (_c *MockTodoRepository_Update_Call) Run(run func(categoryInput models.TodoInput, id string)) *MockTodoRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.TodoInput), args[1].(string))
	})
	return _c
}

func (_c *MockTodoRepository_Update_Call) Return(_a0 models.Todo, _a1 error) *MockTodoRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTodoRepository_Update_Call) RunAndReturn(run func(models.TodoInput, string) (models.Todo, error)) *MockTodoRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTodoRepository creates a new instance of MockTodoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTodoRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTodoRepository {
	mock := &MockTodoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}