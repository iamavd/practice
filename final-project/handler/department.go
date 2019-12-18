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
	AddHeadOfDepartment(ctx context.Context, departmentId string, employeeId string) error
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

func (dept DepartmentHandler) AddHeadOfDepartment(w http.ResponseWriter, r *http.Request) {
	deptId := r.URL.Query().Get("departmentId")
	empId := r.URL.Query().Get("employeeId")
	err := dept.DepartmentService.AddHeadOfDepartment(r.Context(), deptId, empId)
	if err != nil {
		SendError(err, w, http.StatusInternalServerError)
		return
	}
	SendResponse(w, nil)

	return
}
