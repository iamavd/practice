package handler

import (
	"context"
	"final-project/model"

	"encoding/json"
	"net/http"
)

type DepartmentService interface {
	AddDepartment(ctx context.Context, m model.Department) (*model.IDresponse, error)
	AddEmployeeToDepartment(ctx context.Context, departmentId string, employeeId string) error
	EditHeadOfDepartment(ctx context.Context, departmentId string, employeeId string) error
	RemoveEmloyeeFromDepartment(ctx context.Context, departmentId string, employeeId string) error
	EditDepartment(ctx context.Context, departmentId string, m model.Department) error
	RemoveDepartment(ctx context.Context, departmentId string) error
}

type DepartmentHandler struct {
	DepartmentService DepartmentService
}

func (dept DepartmentHandler) Add(w http.ResponseWriter, r *http.Request) {
	var m model.Department
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		SendError(err, w, http.StatusBadRequest)
		return
	}
	res, err := dept.DepartmentService.AddDepartment(r.Context(), m)
	if err != nil {
		SendError(err, w, http.StatusInternalServerError)
		return
	}
	SendResponse(w, res)
}

func (dept DepartmentHandler) AddToDepartment(w http.ResponseWriter, r *http.Request) {
	deptId := r.URL.Query().Get("departmentId")
	empId := r.URL.Query().Get("employeeId")
	err := dept.DepartmentService.AddEmployeeToDepartment(r.Context(), deptId, empId)
	if err != nil {
		SendError(err, w, http.StatusInternalServerError)
		return
	}
	SendResponse(w, nil)

	return
}

func (dept DepartmentHandler) RemoveEmloyeeFromDepartment(w http.ResponseWriter, r *http.Request) {
	deptId := r.URL.Query().Get("departmentId")
	empId := r.URL.Query().Get("employeeId")
	err := dept.DepartmentService.RemoveEmloyeeFromDepartment(r.Context(), deptId, empId)
	if err != nil {
		SendError(err, w, http.StatusInternalServerError)
		return
	}
	SendResponse(w, nil)

	return
}

func (dept DepartmentHandler) EditHeadOfDepartment(w http.ResponseWriter, r *http.Request) {
	deptId := r.URL.Query().Get("departmentId")
	empId := r.URL.Query().Get("employeeId")
	err := dept.DepartmentService.EditHeadOfDepartment(r.Context(), deptId, empId)
	if err != nil {
		SendError(err, w, http.StatusInternalServerError)
		return
	}
	SendResponse(w, nil)

	return
}

func (dept DepartmentHandler) EditDepartment(w http.ResponseWriter, r *http.Request) {
	deptId := r.URL.Query().Get("departmentId")

	var m model.Department
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		SendError(err, w, http.StatusBadRequest)
		return
	}

	err := dept.DepartmentService.EditDepartment(r.Context(), deptId, m)
	if err != nil {
		SendError(err, w, http.StatusInternalServerError)
		return
	}
	SendResponse(w, nil)

	return
}

func (dept DepartmentHandler) RemoveDepartment(w http.ResponseWriter, r *http.Request) {
	deptId := r.URL.Query().Get("departmentId")
	err := dept.DepartmentService.RemoveDepartment(r.Context(), deptId)

	if err != nil {
		SendError(err, w, http.StatusInternalServerError)
	}

	SendResponse(w, nil)
}
