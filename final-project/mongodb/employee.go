package mongodb

import (
	"context"
	"final-project/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Employee struct {
	*mongo.Collection
}

type EmployeeCollection interface {
	Add(ctx context.Context, m model.Employee) (*model.Employee, error)
	GetByID(ctx context.Context, id string) (*model.Employee, error)
}

func (emp Employee) Add(ctx context.Context, m model.Employee) (*model.Employee, error) {
	m.ID = primitive.NewObjectID()
	_, err := emp.InsertOne(ctx, m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (emp Employee) GetByID(ctx context.Context, id string) (*model.Employee, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	cursor, err := emp.Find(ctx, bson.M{
		"_id": _id,
	})
	cursor.Next(ctx)
	var m model.Employee
	if err := cursor.Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}
