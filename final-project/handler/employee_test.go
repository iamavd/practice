package handler

import (
	"errors"
	"final-project/handler/mocks"
	"final-project/model"
	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func Test_AddError(t *testing.T) {
	var m model.Employee
	es := &mocks.EmployeeService{}
	es.On("AddEmployee", mock.Anything, m).Return(nil, errors.New("this mock returns error"))

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
		Equal("{\"Error\":\"this mock returns error\"}\n")

}

func Test_AddStatusOK(t *testing.T) {
	var m model.Employee
	es := &mocks.EmployeeService{}
	resp := model.IDresponse{Id: "response_id"}
	es.On("AddEmployee", mock.Anything, m).Return(&resp, nil)

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
		Status(http.StatusOK).
		Body().
		Equal("{\"id\":\"response_id\"}\n")
}

func Test_AddBadRequest(t *testing.T) {
	var m model.Employee
	es := &mocks.EmployeeService{}
	resp := model.IDresponse{Id: "response_id"}
	es.On("AddEmployee", mock.Anything, m).Return(&resp, nil)

	mux := CreateMux(es, nil)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/employee/add").
		WithJSON("").
		Expect().
		Status(http.StatusBadRequest).
		Body().
		Equal("{\"Error\":\"json: cannot unmarshal string into Go value of type model.Employee\"}\n")
}

func Test_GetByIDStatusOK(t *testing.T) {
	var m model.Employee
	es := &mocks.EmployeeService{}
	es.On("GetEmployeeByID", mock.Anything, "someid").Return(&m, nil)

	mux := CreateMux(es, nil)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/employee/get").
		WithQuery("employeeId", "someid").
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("{\"ID\":\"000000000000000000000000\",\"Name\":\"\",\"Surname\":\"\",\"Sex\":\"\",\"DateOfBirth\":\"0001-01-01T00:00:00Z\",\"Profession\":\"\",\"Address\":\"\"}\n")
}

func Test_GetByIDError(t *testing.T) {
	es := &mocks.EmployeeService{}
	es.On("GetEmployeeByID", mock.Anything, "").Return(nil, errors.New("this mock returns error"))

	mux := CreateMux(es, nil)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/employee/get").
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"this mock returns error\"}\n")

}

func Test_GetEmloyeeListStatusOK(t *testing.T) {
	es := &mocks.EmployeeService{}
	var m []model.Employee

	es.On("GetEmployeeList", mock.Anything).Return(&m, nil)

	mux := CreateMux(es, nil)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/employee/getlist").
		WithJSON(m).
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("null\n")
}

func Test_GetEmloyeeListError(t *testing.T) {
	es := &mocks.EmployeeService{}
	var m []model.Employee

	es.On("GetEmployeeList", mock.Anything).Return(nil, errors.New("this mock returns error"))

	mux := CreateMux(es, nil)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/employee/getlist").
		WithJSON(m).
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"this mock returns error\"}\n")
}

func Test_RemoveStatusOK(t *testing.T) {
	es := &mocks.EmployeeService{}
	es.On("RemoveEmployee", mock.Anything, "").Return(nil)

	mux := CreateMux(es, nil)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/employee/remove").WithQuery("id", "someid").
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("null\n")
}

func Test_RemoveError(t *testing.T) {
	es := &mocks.EmployeeService{}
	es.On("RemoveEmployee", mock.Anything, "").Return(errors.New("remove_error"))

	mux := CreateMux(es, nil)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/employee/remove").WithQuery("id", "someid").
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"remove_error\"}\nnull\n")
}

func Test_EditEmployeeStatusOK(t *testing.T) {
	es := &mocks.EmployeeService{}
	var m model.Employee

	es.On("ModifyEmployee", mock.Anything, "someid", m).Return(nil)

	mux := CreateMux(es, nil)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/employee/edit").WithQuery("employeeId", "someid").
		WithJSON(m).
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("null\n")
}

func Test_EditEmployeeError(t *testing.T) {
	es := &mocks.EmployeeService{}
	var m model.Employee

	es.On("ModifyEmployee", mock.Anything, "someid", m).Return(errors.New("this mock returns error"))

	mux := CreateMux(es, nil)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/employee/edit").WithQuery("employeeId", "someid").
		WithJSON(m).
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"this mock returns error\"}\n")
}

func Test_EditEmployeeBadRequest(t *testing.T) {
	es := &mocks.EmployeeService{}
	var m model.Employee

	es.On("ModifyEmployee", mock.Anything, "someid", m).Return(nil)

	mux := CreateMux(es, nil)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/employee/edit").WithQuery("employeeId", "someid").
		WithJSON("").
		Expect().
		Status(http.StatusBadRequest).
		Body().
		Equal("{\"Error\":\"json: cannot unmarshal string into Go value of type model.Employee\"}\n")
}
