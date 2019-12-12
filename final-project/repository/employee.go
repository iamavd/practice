package repository

import (
	"final-project/model"
	"time"
)

type EmployeeRepository interface {
	Add(employee *model.Employee) error
	GetEmployeeById(id objectid.ObjectID) (*model.Employee, error)
	GetEmployeesList() ([]*model.Employee, error)
}

type employeeRepository struct {

}

func (repo *employeeRepository) Add(employee *model.Employee) error {
	//TODO: add worker
	return nil
} 

func (repo *employeeRepository) GetEmployeeById(id objectid.ObjectID) (*model.Employee, error) {
	if id == nil {
		return nil,nil
	}
	return &model.Employee {
		ID: "12345",
		Name: "Ivan",
		Surname: "Ivanov",
		Sex: "male",
		DateOfBirth: time.Parse(shortForm, "1986-Feb-03"),
		Profession: "developer",
		Address: "any town"
	}, nil
}

func GetEmployeesList() ([]*model.Employee, error) {
	return *[]model.Employee{
				&model.Employee{
				ID: "1",
				Name: "Ivan",
				Surname: "Ivanov",
				Sex: "male",
				DateOfBirth: time.Parse(shortForm, "1986-Feb-03"),
				Profession: "developer",
				Address: "any town"
			},
			&model.Employee{
				ID: "2",
				Name: "John",
				Surname: "Johanson",
				Sex: "male",
				DateOfBirth: time.Parse(shortForm, "1996-Feb-04"),
				Profession: "developer",
				Address: "any town"
			},
	}, nil
}
	

