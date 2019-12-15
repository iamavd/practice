package model

import (
	"time"

	//	"github.com/mongodb/mongo-go-driver/bson"
	"labix.org/v2/mgo/bson"
)

//Employee contains the personal information about employees
type Employee struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	Surname     string        `bson:"surname"`
	Sex         string        `bson:"sex"`
	DateOfBirth time.Time     `bson:"birth"`
	Profession  string        `bson:"profession"`
	Address     string        `bson:"address"`
}
