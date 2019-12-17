package handler

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	port            string
	employeeHandler EmployeeHandler
}

type ResponseError struct {
	Error string `json: "error"`
}

func NewServer(port string, employeeHandler EmployeeHandler) *Server {
	return &Server{
		port:            port,
		employeeHandler: employeeHandler,
	}
}

/*
func (server *Server) ConfigureAndRun() {

	employeeMux := http.NewServeMux()
	employeeMux.HandleFunc("/employee", server.employeeHandler.GetEmployeeList)

	fmt.Printf("listening at %s", server.port)
	http.ListenAndServe(server.port, employeeMux)
}
*/
func SendError(err error, w http.ResponseWriter, code int) {
	json.NewEncoder(w).Encode(ResponseError{Error: err.Error()})
	w.WriteHeader(code)
}

func SendResponse(w http.ResponseWriter, res interface{}) {
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)

}
