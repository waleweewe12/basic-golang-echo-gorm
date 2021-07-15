package handler

import (
	"basic_golang_echo/internal/config"
	"basic_golang_echo/internal/employee"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CloseFunc func() error

type route struct {
	Group          string
	Path           string
	HttpMethod     string
	HandlerFunc    echo.HandlerFunc
	MiddlewareFunc []echo.MiddlewareFunc
}

func NewRoutes(e *echo.Echo, cv *config.Configs) ([]CloseFunc, error) {
	closeFuncs := make([]CloseFunc, 0)

	emp := employee.NewEndpoint(cv)
	routes := []route{
		{
			Group:          "employee",
			Path:           "/byId",
			HttpMethod:     http.MethodPost,
			HandlerFunc:    emp.GetEmployeeById,
			MiddlewareFunc: nil,
		},
		{
			Group:          "employee",
			Path:           "/add",
			HttpMethod:     http.MethodPost,
			HandlerFunc:    emp.AddEmployee,
			MiddlewareFunc: nil,
		},
		{
			Group:          "employee",
			Path:           "/salaryById",
			HttpMethod:     http.MethodPatch,
			HandlerFunc:    emp.UpdateSalary,
			MiddlewareFunc: nil,
		},
		{
			Group:          "employee",
			Path:           "/byId",
			HttpMethod:     http.MethodDelete,
			HandlerFunc:    emp.DeleteEmployeeById,
			MiddlewareFunc: nil,
		},
	}

	// http connection
	for _, rt := range routes {
		mw := []echo.MiddlewareFunc{
			middleware.BodyDumpWithConfig(BodyDumpConfig()),
		}
		mw = append(mw, rt.MiddlewareFunc...)
		e.Group(rt.Group).Add(rt.HttpMethod, rt.Path, rt.HandlerFunc, mw...)
	}

	return closeFuncs, nil
}
