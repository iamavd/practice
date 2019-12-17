package handler

import (
	"context"
	"final-project/model"

	"encoding/json"
	"net/http"
)

type EmployeeService interface {
	//GetEmployeeList(w http.ResponseWriter, r *http.Request)
	AddEmployee(ctx context.Context, m model.Employee) (*model.IDresponse, error)
	GetEmployeeByID(ctx context.Context, id string) (*model.Employee, error)
}

type EmployeeHandler struct {
	EmployeeService EmployeeService
}

func (emp EmployeeHandler) Add(w http.ResponseWriter, r *http.Request) {
	var m model.Employee
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		SendError(err, w, http.StatusBadRequest)
		return
	}
	res, err := emp.EmployeeService.AddEmployee(r.Context(), m)
	if err != nil {
		SendError(err, w, http.StatusInternalServerError)
		return
	}
	SendResponse(w, res)
}

func (emp EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("employeeId")
	res, err := emp.EmployeeService.GetEmployeeByID(r.Context(), id)
	if err != nil {
		SendError(err, w, http.StatusInternalServerError)
		return
	}
	SendResponse(w, res)
}

/*
func (h *employeeHandler) GetEmployeeList(w http.ResponseWriter, r *http.Request) {

		emps, _ := h.employeeRepository.GetEmployeesList()

		fmt.Fprintln(w, *emps[0], *emps[1])

}
*/
