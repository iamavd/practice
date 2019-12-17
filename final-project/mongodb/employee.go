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
	GetList(ctx context.Context) (*[]model.Employee, error)
	Remove(ctx context.Context, id string) error
	Update(ctx context.Context, id string, m model.Employee) error
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

func (emp Employee) GetList(ctx context.Context) (*[]model.Employee, error) {
	cursor, err := emp.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	var m []model.Employee

	if err := cursor.All(ctx, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

func (emp Employee) Remove(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = emp.DeleteOne(ctx, bson.M{
		"_id": _id,
	})

	if err != nil {
		return err
	}
	return nil
}

func (emp Employee) Update(ctx context.Context, id string, m model.Employee) error {
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	m.ID = _id

	_, err = emp.ReplaceOne(ctx, bson.M{
		"_id": _id,
	}, m)

	if err != nil {
		return err
	}
	return nil
}
