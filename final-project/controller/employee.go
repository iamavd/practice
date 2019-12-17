package controller

import (
	"context"
	"final-project/model"
	mongodb "final-project/mongodb"
)

type Employee struct {
	DbEmployee mongodb.EmployeeCollection
}

func (emp Employee) AddEmployee(ctx context.Context, m model.Employee) (*model.IDresponse, error) {
	newEmp, err := emp.DbEmployee.Add(ctx, m)

	if err != nil {
		return nil, err
	}

	return &model.IDresponse{Id: newEmp.ID.Hex()}, nil
}

func (emp Employee) GetEmployeeByID(ctx context.Context, id string) (*model.Employee, error) {
	newEmp, err := emp.DbEmployee.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return newEmp, nil
}

func (emp Employee) GetEmployeeList(ctx context.Context) (*[]model.Employee, error) {
	newEmpList := &[]model.Employee{}
	newEmpList, err := emp.DbEmployee.GetList(ctx)

	if err != nil {
		return nil, err
	}

	return newEmpList, nil
}
