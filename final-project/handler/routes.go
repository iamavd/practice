package handler

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Error string `json: "error"`
}

func SendError(err error, w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ResponseError{Error: err.Error()})
}

func SendResponse(w http.ResponseWriter, res interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func CreateMux(employeeService EmployeeService, departmentService DepartmentService) *http.ServeMux {

	employeeHandler := EmployeeHandler{EmployeeService: employeeService}
	departmentHandler := DepartmentHandler{DepartmentService: departmentService}

	mux := http.NewServeMux()

	mux.HandleFunc("/employee/add", employeeHandler.Add)
	mux.HandleFunc("/employee/get", employeeHandler.GetEmployee)
	mux.HandleFunc("/employee/getlist", employeeHandler.GetEmployeeList)
	mux.HandleFunc("/employee/remove", employeeHandler.RemoveEmployee)
	mux.HandleFunc("/employee/edit", employeeHandler.EditEmployee)

	mux.HandleFunc("/department/add", departmentHandler.Add)
	mux.HandleFunc("/department/getlist", departmentHandler.GetDepartmentList)
	mux.HandleFunc("/department/empadd", departmentHandler.AddToDepartment)
	mux.HandleFunc("/department/headedit", departmentHandler.EditHeadOfDepartment)
	mux.HandleFunc("/department/empremove", departmentHandler.RemoveEmloyeeFromDepartment)
	mux.HandleFunc("/department/edit", departmentHandler.EditDepartment)
	mux.HandleFunc("/department/remove", departmentHandler.RemoveDepartment)

	return mux
}

func CreateMuxEmp(employeeService EmployeeService) *http.ServeMux {

	employeeHandler := EmployeeHandler{EmployeeService: employeeService}
	//departmentHandler := DepartmentHandler{DepartmentService: departmentService}

	mux := http.NewServeMux()

	mux.HandleFunc("/employee/add", employeeHandler.Add)
	mux.HandleFunc("/employee/get", employeeHandler.GetEmployee)
	mux.HandleFunc("/employee/getlist", employeeHandler.GetEmployeeList)
	mux.HandleFunc("/employee/remove", employeeHandler.RemoveEmployee)
	mux.HandleFunc("/employee/edit", employeeHandler.EditEmployee)

	return mux
}
