package handler

import (
	"errors"
	"final-project/handler/mocks"
	"final-project/model"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/mock"
)

func Test_AddDeptStatusOK(t *testing.T) {
	var m model.Department
	ds := &mocks.DepartmentService{}
	resp := model.IDresponse{Id: "response_id"}

	ds.On("AddDepartment", mock.Anything, m).Return(&resp, nil)

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/add").
		WithJSON(m).
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("{\"id\":\"response_id\"}\n")
}

func Test_AddDeptBadRequest(t *testing.T) {
	var m model.Department
	ds := &mocks.DepartmentService{}
	resp := model.IDresponse{Id: "response_id"}

	ds.On("AddDepartment", mock.Anything, m).Return(&resp, nil)

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/add").
		WithJSON("").
		Expect().
		Status(http.StatusBadRequest).
		Body().
		Equal("{\"Error\":\"json: cannot unmarshal string into Go value of type model.Department\"}\n")
}

func Test_AddDeptError(t *testing.T) {
	var m model.Department
	ds := &mocks.DepartmentService{}

	ds.On("AddDepartment", mock.Anything, m).Return(nil, errors.New("this mock returns error"))

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/add").
		WithJSON(m).
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"this mock returns error\"}\n")
}

func Test_GetDeptListStatusOK(t *testing.T) {
	var m []model.Department
	ds := &mocks.DepartmentService{}

	ds.On("GetDepartmentList", mock.Anything).Return(&m, nil)

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/getlist").
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("null\n")
}

func Test_GetDeptListError(t *testing.T) {
	ds := &mocks.DepartmentService{}

	ds.On("GetDepartmentList", mock.Anything).Return(nil, errors.New("this mock returns error"))

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/getlist").
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"this mock returns error\"}\n")
}

func Test_AddEmpToDeptStatusOK(t *testing.T) {
	ds := &mocks.DepartmentService{}

	ds.On("AddEmployeeToDepartment", mock.Anything, "someDept", "someEmp").Return(nil)

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/empadd").WithQuery("departmentId", "someDept").
		WithQuery("employeeId", "someEmp").
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("null\n")
}

func Test_AddEmpToDeptError(t *testing.T) {
	ds := &mocks.DepartmentService{}

	ds.On("AddEmployeeToDepartment", mock.Anything, "someDept", "someEmp").Return(errors.New("this mock returns error"))

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/empadd").WithQuery("departmentId", "someDept").
		WithQuery("employeeId", "someEmp").
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"this mock returns error\"}\n")
}

func Test_RemoveEmpFromDeptStatusOK(t *testing.T) {
	ds := &mocks.DepartmentService{}

	ds.On("RemoveEmloyeeFromDepartment", mock.Anything, "someDept", "someEmp").Return(nil)

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/empremove").WithQuery("departmentId", "someDept").
		WithQuery("employeeId", "someEmp").
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("null\n")
}

func Test_RemoveEmpFromDeptError(t *testing.T) {
	ds := &mocks.DepartmentService{}

	ds.On("RemoveEmloyeeFromDepartment", mock.Anything, "someDept", "someEmp").Return(errors.New("this mock returns error"))

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/empremove").WithQuery("departmentId", "someDept").
		WithQuery("employeeId", "someEmp").
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"this mock returns error\"}\n")
}

func Test_EditDeptHeadStatusOK(t *testing.T) {
	ds := &mocks.DepartmentService{}

	ds.On("EditHeadOfDepartment", mock.Anything, "someDept", "someEmp").Return(nil)

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/headedit").WithQuery("departmentId", "someDept").
		WithQuery("employeeId", "someEmp").
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("null\n")
}

func Test_EditDeptHeadError(t *testing.T) {
	ds := &mocks.DepartmentService{}

	ds.On("EditHeadOfDepartment", mock.Anything, "someDept", "someEmp").Return(errors.New("this mock returns error"))

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/headedit").WithQuery("departmentId", "someDept").
		WithQuery("employeeId", "someEmp").
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"this mock returns error\"}\n")
}

func Test_EditDepartmentStatusOK(t *testing.T) {
	ds := &mocks.DepartmentService{}
	var m model.Department

	ds.On("EditDepartment", mock.Anything, "someDeptId", m).Return(nil)

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/edit").WithQuery("departmentId", "someDeptId").
		WithJSON(m).
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("null\n")
}

func Test_EditDepartmentInternalError(t *testing.T) {
	ds := &mocks.DepartmentService{}
	var m model.Department

	ds.On("EditDepartment", mock.Anything, "someDeptId", m).Return(errors.New("this mock returns error"))

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/edit").WithQuery("departmentId", "someDeptId").
		WithJSON(m).
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"this mock returns error\"}\n")
}

func Test_EditDepartmentBadRequest(t *testing.T) {
	ds := &mocks.DepartmentService{}
	var m model.Department

	ds.On("EditDepartment", mock.Anything, "someDeptId", m).Return(errors.New("this mock returns error"))

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/edit").WithQuery("departmentId", "someDeptId").
		WithJSON("").
		Expect().
		Status(http.StatusBadRequest).
		Body().
		Equal("{\"Error\":\"json: cannot unmarshal string into Go value of type model.Department\"}\n")
}

func Test_RemoveDepartmentStatusOK(t *testing.T) {
	ds := &mocks.DepartmentService{}

	ds.On("RemoveDepartment", mock.Anything, "someDeptId").Return(nil)

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/remove").WithQuery("departmentId", "someDeptId").
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("null\n")
}

func Test_RemoveDepartmentError(t *testing.T) {
	ds := &mocks.DepartmentService{}

	ds.On("RemoveDepartment", mock.Anything, "someDeptId").Return(errors.New("this mock returns error"))

	mux := CreateMux(nil, ds)

	exp := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})

	exp.GET("/department/remove").WithQuery("departmentId", "someDeptId").
		Expect().
		Status(http.StatusInternalServerError).
		Body().
		Equal("{\"Error\":\"this mock returns error\"}\nnull\n")
}
