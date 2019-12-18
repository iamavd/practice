package controller

import (
	"context"
	"final-project/model"
	mongodb "final-project/mongodb"
)

type Department struct {
	DbDepartment mongodb.DepartmentCollection
}

func (dept Department) AddDepartment(ctx context.Context, m model.Department) (*model.IDresponse, error) {
	newDept, err := dept.DbDepartment.AddDept(ctx, m)

	if err != nil {
		return nil, err
	}

	return &model.IDresponse{Id: newDept.ID.Hex()}, nil
}

func (dept Department) GetDepartmentList(ctx context.Context) (*[]model.Department, error) {
	newDeptList := &[]model.Department{}
	newDeptList, err := dept.DbDepartment.GetList(ctx)

	if err != nil {
		return nil, err
	}

	return newDeptList, nil
}

func (dept Department) AddEmployeeToDepartment(ctx context.Context, departmentId string, employeeId string) error {
	err := dept.DbDepartment.AddEmployeeToDept(ctx, departmentId, employeeId)
	if err != nil {
		return err
	}

	return nil
}

func (dept Department) RemoveEmloyeeFromDepartment(ctx context.Context, departmentId string, employeeId string) error {
	err := dept.DbDepartment.RemoveEmployee(ctx, departmentId, employeeId)
	if err != nil {
		return err
	}

	return nil
}

func (dept Department) EditHeadOfDepartment(ctx context.Context, departmentId string, employeeId string) error {
	err := dept.DbDepartment.EditDeptHead(ctx, departmentId, employeeId)
	if err != nil {
		return err
	}

	return nil
}

func (dept Department) EditDepartment(ctx context.Context, departmentId string, m model.Department) error {
	err := dept.DbDepartment.EditDept(ctx, departmentId, m)
	if err != nil {
		return err
	}

	return nil
}

func (dept Department) RemoveDepartment(ctx context.Context, departmentId string) error {
	err := dept.DbDepartment.Remove(ctx, departmentId)

	if err != nil {
		return err
	}

	return nil
}
