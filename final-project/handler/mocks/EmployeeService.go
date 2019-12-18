// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "final-project/model"
)

// EmployeeService is an autogenerated mock type for the EmployeeService type
type EmployeeService struct {
	mock.Mock
}

// AddEmployee provides a mock function with given fields: ctx, m
func (_m *EmployeeService) AddEmployee(ctx context.Context, m model.Employee) (*model.IDresponse, error) {
	ret := _m.Called(ctx, m)

	var r0 *model.IDresponse
	if rf, ok := ret.Get(0).(func(context.Context, model.Employee) *model.IDresponse); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.IDresponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Employee) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEmployeeByID provides a mock function with given fields: ctx, id
func (_m *EmployeeService) GetEmployeeByID(ctx context.Context, id string) (*model.Employee, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Employee
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Employee); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Employee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEmployeeList provides a mock function with given fields: ctx
func (_m *EmployeeService) GetEmployeeList(ctx context.Context) (*[]model.Employee, error) {
	ret := _m.Called(ctx)

	var r0 *[]model.Employee
	if rf, ok := ret.Get(0).(func(context.Context) *[]model.Employee); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.Employee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyEmployee provides a mock function with given fields: ctx, id, m
func (_m *EmployeeService) ModifyEmployee(ctx context.Context, id string, m model.Employee) error {
	ret := _m.Called(ctx, id, m)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Employee) error); ok {
		r0 = rf(ctx, id, m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveEmployee provides a mock function with given fields: ctx, id
func (_m *EmployeeService) RemoveEmployee(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
