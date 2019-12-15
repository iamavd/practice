package repository

import (
	"final-project/model"
	//"final-project/mongo"
	"time"
	//	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"labix.org/v2/mgo/bson"
)

type EmployeeRepository interface {
	Add(employee *model.Employee) error
	GetEmployeeById(id bson.ObjectId) (*model.Employee, error)
	GetEmployeesList() ([]*model.Employee, error)
}

type employeeRepository struct {
}

//TODO: need in dbclient
/*
func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{

	}
}*/

func (repo *employeeRepository) Add(employee *model.Employee) error {
	//TODO: add worker
	return nil
}

func (repo *employeeRepository) GetEmployeeById(id bson.ObjectId) (*model.Employee, error) {
	if id.String() == "" {
		return nil, nil
	}
	t, _ := time.Parse("2006-01-02", "1986-Feb-03")
	return &model.Employee{
		ID:          "12345",
		Name:        "Ivan",
		Surname:     "Ivanov",
		Sex:         "male",
		DateOfBirth: t,
		Profession:  "developer",
		Address:     "any town",
	}, nil
}

func GetEmployeesList() ([]*model.Employee, error) {
	t, _ := time.Parse("2006-01-02", "1986-Feb-03")

	return []*model.Employee{
		&model.Employee{
			ID:          "1",
			Name:        "Ivan",
			Surname:     "Ivanov",
			Sex:         "male",
			DateOfBirth: t,
			Profession:  "developer",
			Address:     "any town",
		},
		&model.Employee{
			ID:          "2",
			Name:        "John",
			Surname:     "Johanson",
			Sex:         "male",
			DateOfBirth: t,
			Profession:  "developer",
			Address:     "any town",
		},
	}, nil
}
