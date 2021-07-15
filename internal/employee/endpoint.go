package employee

import (
	"basic_golang_echo/internal/config"
	"basic_golang_echo/internal/model"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	cv  *config.Configs
	srv employeeService //Service Interface Tier
}

type employeeService interface {
	GetEmployeeById(ctx context.Context, empId string) ([]model.Employee, error)
	GetEmployeeByFirstName(ctx context.Context, firstName string) ([]model.Employee, error)
	AddEmployee(ctx context.Context, emp model.Employee) error
	UpdateSalary(ctx context.Context, emp model.Employee) error
	DeleteEmployeeById(ctx context.Context, empId string) error
}

func NewEndpoint(cv *config.Configs) *Endpoint {
	return &Endpoint{
		cv:  cv,
		srv: NewService(cv),
	}
}

func (e Endpoint) GetEmployeeById(c echo.Context) error {
	var request struct {
		EmployeeId string `json:"id"`
	}

	//check ว่า request ตรงกับที่ระบุเอาไว้ไหม
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.EmployeeResponse{
			ServerStatus: http.StatusBadRequest,
			Message:      "get employee by id failed",
		})
	}
	//เรียกใช้งาน service แล้วได้สถานะ bad request กลับมา
	data, err := e.srv.GetEmployeeById(c.Request().Context(), request.EmployeeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.EmployeeResponse{
			ServerStatus: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}
	//เรียกใช้งาน service แล้วได้สถานะ ok กลับมา
	return c.JSON(http.StatusOK, model.EmployeeResponse{
		ServerStatus: http.StatusOK,
		Message:      "get employee by id success",
		Employee:     data,
	})
}

func (e Endpoint) AddEmployee(c echo.Context) error {
	var newEmployee model.Employee
	//check ว่า request ตรงกับที่ระบุเอาไว้ไหม
	if err := c.Bind(&newEmployee); err != nil {
		return c.JSON(http.StatusBadRequest, model.StatusMessageResponse{
			ServerStatus: http.StatusBadRequest,
			Message:      "add employee by id failed",
		})
	}
	//เรียกใช้งาน service แล้วได้สถานะ bad request กลับมา
	err := e.srv.AddEmployee(c.Request().Context(), newEmployee)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.StatusMessageResponse{
			ServerStatus: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}
	//เรียกใช้งาน service แล้วได้สถานะ ok กลับมา
	return c.JSON(http.StatusOK, model.StatusMessageResponse{
		ServerStatus: http.StatusOK,
		Message:      "add employee by id success",
	})
}

func (e Endpoint) UpdateSalary(c echo.Context) error {
	var em model.Employee
	//check ว่า request ตรงกับที่ระบุเอาไว้ไหม
	if err := c.Bind(&em); err != nil {
		return c.JSON(http.StatusBadRequest, model.StatusMessageResponse{
			ServerStatus: http.StatusBadRequest,
			Message:      "update employee by id failed",
		})
	}
	//เรียกใช้งาน service แล้วได้สถานะ bad request กลับมา
	err := e.srv.UpdateSalary(c.Request().Context(), em)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.StatusMessageResponse{
			ServerStatus: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}
	//เรียกใช้งาน service แล้วได้สถานะ ok กลับมา
	return c.JSON(http.StatusOK, model.StatusMessageResponse{
		ServerStatus: http.StatusOK,
		Message:      "update employee by id success",
	})
}

func (e Endpoint) DeleteEmployeeById(c echo.Context) error {
	var em model.Employee
	//check ว่า request ตรงกับที่ระบุเอาไว้ไหม
	if err := c.Bind(&em); err != nil {
		return c.JSON(http.StatusBadRequest, model.StatusMessageResponse{
			ServerStatus: http.StatusBadRequest,
			Message:      "delete employee by id failed",
		})
	}
	//เรียกใช้งาน service แล้วได้สถานะ bad request กลับมา
	err := e.srv.DeleteEmployeeById(c.Request().Context(), em.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.StatusMessageResponse{
			ServerStatus: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}
	//เรียกใช้งาน service แล้วได้สถานะ ok กลับมา
	return c.JSON(http.StatusOK, model.StatusMessageResponse{
		ServerStatus: http.StatusOK,
		Message:      "delete employee by id success",
	})
}

func (e Endpoint) GetEmployeeByFirstName() error {
	return nil
}
