package handler

import (
	"context"
	"final-project/model"

	"encoding/json"
	"net/http"
)

type EmployeeService interface {
	AddEmployee(ctx context.Context, m model.Employee) (*model.IDresponse, error)
	GetEmployeeByID(ctx context.Context, id string) (*model.Employee, error)
	GetEmployeeList(ctx context.Context) (*[]model.Employee, error)
	DeleteEmployee(ctx context.Context, id string) error
	ModifyEmployee(ctx context.Context, id string, m model.Employee) error
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

func (emp EmployeeHandler) GetEmployeeList(w http.ResponseWriter, r *http.Request) {
	res, err := emp.EmployeeService.GetEmployeeList(r.Context())

	if err != nil {
		SendError(err, w, http.StatusInternalServerError)
		return
	}
	SendResponse(w, res)
}

func (emp EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("employeeId")
	err := emp.EmployeeService.DeleteEmployee(r.Context(), id)

	if err != nil {
		SendError(err, w, http.StatusInternalServerError)
	}

	SendResponse(w, nil)
}

func (emp EmployeeHandler) EditEmployee(w http.ResponseWriter, r *http.Request) {
	var m model.Employee
	id := r.URL.Query().Get("employeeId")

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		SendError(err, w, http.StatusBadRequest)
		return
	}
	err1 := emp.EmployeeService.ModifyEmployee(r.Context(), id, m)
	if err1 != nil {
		SendError(err1, w, http.StatusInternalServerError)
		return
	}
	SendResponse(w, nil)
}
