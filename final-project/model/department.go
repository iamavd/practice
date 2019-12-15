package model

//import "github.com/mongodb/mongo-go-driver/bson"
import "labix.org/v2/mgo/bson"

//Department contains the information about departments - name, workers and head of department
type Department struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"deptName"`
	Description string        `bson:"description"`
	Head        bson.ObjectId     `bson:"head"`
	Employees   []bson.ObjectId   `bson:"employees"`
}
