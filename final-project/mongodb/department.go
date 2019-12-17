package mongodb

import (
	"context"
	"final-project/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Department struct {
	*mongo.Collection
}

type DepartmentCollection interface {
	AddDept(ctx context.Context, m model.Employee) (*model.Department, error)
	AddEmployeeToDept(ctx context.Context, departmentId string, employeeId string) error
}

func (dept Department) AddDept(ctx context.Context, m model.Department) (*model.Department, error) {
	m.ID = primitive.NewObjectID()
	_, err := dept.InsertOne(ctx, m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (dept Department) AddEmployeeToDept(ctx context.Context, departmentId string, employeeId string) error {
	_idDept, err := primitive.ObjectIDFromHex(departmentId)

	if err != nil {
		return err
	}

	_idEmp, err := primitive.ObjectIDFromHex(employeeId)

	if err != nil {
		return err
	}
	
	_, err = dept.UpdateOne(ctx, bson.M{
		"_id" : _idDept,
	}, bson.M{
		"$push" : {
			"employees" : _idEmp,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
