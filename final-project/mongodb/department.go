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
	AddDept(ctx context.Context, m model.Department) (*model.Department, error)
	GetList(ctx context.Context) (*[]model.Department, error)
	AddEmployeeToDept(ctx context.Context, departmentId string, employeeId string) error
	EditDeptHead(ctx context.Context, departmentId string, employeeId string) error
	RemoveEmployee(ctx context.Context, departmentId string, employeeId string) error
	EditDept(ctx context.Context, departmentId string, m model.Department) error
	Remove(ctx context.Context, departmentId string) error
}

func (dept Department) AddDept(ctx context.Context, m model.Department) (*model.Department, error) {
	m.ID = primitive.NewObjectID()
	m.Employees = []primitive.ObjectID{}
	_, err := dept.InsertOne(ctx, m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (dept Department) GetList(ctx context.Context) (*[]model.Department, error) {
	cursor, err := dept.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	var m []model.Department

	if err := cursor.All(ctx, &m); err != nil {
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

	filter := bson.M{
		"_id": _idDept,
	}

	update := bson.M{
		"$push": bson.M{
			"employees": _idEmp,
		},
	}
	_, err = dept.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (dept Department) RemoveEmployee(ctx context.Context, departmentId string, employeeId string) error {
	_idDept, err := primitive.ObjectIDFromHex(departmentId)

	if err != nil {
		return err
	}

	_idEmp, err := primitive.ObjectIDFromHex(employeeId)

	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": _idDept,
	}

	update := bson.M{
		"$pull": bson.M{
			"employees": _idEmp,
		},
	}
	_, err = dept.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (dept Department) EditDeptHead(ctx context.Context, departmentId string, employeeId string) error {
	_idDept, err := primitive.ObjectIDFromHex(departmentId)
	if err != nil {
		return err
	}
	_idEmp, err := primitive.ObjectIDFromHex(employeeId)

	if err != nil {
		return err
	}
	filter := bson.M{"_id": _idDept}
	update := bson.M{
		"$set": bson.M{
			"head": _idEmp,
		},
	}

	_, err = dept.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (dept Department) EditDept(ctx context.Context, departmentId string, m model.Department) error {
	_idDept, err := primitive.ObjectIDFromHex(departmentId)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": _idDept}
	update := bson.M{
		"$set": bson.M{
			"deptName":    m.Name,
			"description": m.Description,
		},
	}

	_, err = dept.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (dept Department) Remove(ctx context.Context, departmentId string) error {
	_idDept, err := primitive.ObjectIDFromHex(departmentId)

	if err != nil {
		return err
	}

	_, err = dept.DeleteOne(ctx, bson.M{
		"_id": _idDept,
	})

	if err != nil {
		return err
	}
	return nil
}
