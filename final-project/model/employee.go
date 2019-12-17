package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Employee contains the personal information about employees
type Employee struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Surname     string             `bson:"surname"`
	Sex         string             `bson:"sex"`
	DateOfBirth time.Time          `bson:"birth"`
	Profession  string             `bson:"profession"`
	Address     string             `bson:"address"`
}

type IDresponse struct {
	Id string `json:"id"`
}
