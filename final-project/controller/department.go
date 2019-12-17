package controller

import (
	"context"
	"final-project/model"
	mongodb "final-project/mongodb"
)

type Department struct {
	DbDepartmnet mongodb.DepartmentCollection
}

func (dept Department) AddDepartment(ctx context.Context, m model.Department) (*model.IDresponse, error) {
	newDept, err := dept.DbDepartmnet.AddDept(ctx, m)

	if err != nil {
		return nil, err
	}

	return &model.IDresponse{Id: newDept.ID.Hex()}, nil
}

func (dept Department) AddEmployeeToDepartment(ctx context.Context, departmentId string, employeeId string) error {
	return nil
}
