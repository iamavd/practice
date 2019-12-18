package handler

import (
	//"context"
	"errors"
	"final-project/handler/mocks"
	"final-project/model"
	"github.com/gavv/httpexpect"
	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func Test_Add(t *testing.T) {
	var m model.Employee
	es := &mocks.EmployeeService{}
	es.On("AddEmployee", mock.Anything, m).Return(nil, errors.New("this mock return error"))

	mux := CreateMux(es, nil)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/employee/add").
		WithJSON(m).
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"this mock return error\"}\n")

}
