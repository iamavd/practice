package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//Department contains the information about departments - name, workers and head of department
type Department struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"deptName"`
	Description string               `bson:"description"`
	Head        primitive.ObjectID   `bson:"head"`
	Employees   []primitive.ObjectID `bson:"employees"`
}
