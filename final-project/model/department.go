package model

import "github.com/mongodb/mongo-go-driver/bson/objectid"

//Department contains the information about departments - name, workers and head of department
type Department struct {
	ID          objectid.ObjectID   `bson:"_id"`
	Name        string              `bson:"deptName"`
	Description string              `bson:"description"`
	Head        objectid.ObjectID   `bson:"head"`
	Employees   []objectid.ObjectID `bson:"employees"`
}
