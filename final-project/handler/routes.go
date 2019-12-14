package handler

import (
	"fmt"
	"net/http"
)

type Server struct {
	port            string
	employeeHandler handler.EmployeeHandler
}

func NewServer(port string, employeeHandler handler.EmployeeHandler) *Server {
	return &Server{
		port:            port,
		employeeHandler: employeeHandler,
	}
}

func (server *Server) ConfigureAndRun() {

	employeeMux := http.NewServeMux()
	employeeMux.HandleFunc("/employee", server.employeeHandler.GetEmployeeList)

	fmt.Printf("listening at %s", server.port)
	http.ListenAndServe(server.port, employeeMux)
}
