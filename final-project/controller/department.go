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

func (dept Department) AddHeadOfDepartment(ctx context.Context, departmentId string, employeeId string) error {
	err := dept.DbDepartment.AddDeptHead(ctx, departmentId, employeeId)
	if err != nil {
		return err
	}

	return nil
}
