package handler

import (
	"final-project/repository"
	"fmt"
	"net/http"
)

type EmployeeHandler interface {
	GetEmployeeList(w http.ResponseWriter, r *http.Request)
}

type employeeHandler struct {
	employeeRepository repository.EmployeeRepository
}

/*
func Add(employeeRepository repository.EmployeeRepository) EmployeeHandler {
	return &EmployeeHandler{
		usersRepository: employeeRepository,
	}
}
*/
func (h *employeeHandler) GetEmployeeList(w http.ResponseWriter, r *http.Request) {

	emps, _ := h.employeeRepository.GetEmployeesList()

	fmt.Fprintln(w, *emps[0], *emps[1])

}
