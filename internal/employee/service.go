package employee

import (
	"basic_golang_echo/internal/config"
	"basic_golang_echo/internal/model"
	"context"
)

type employeeRepo interface {
	GetEmployeeById(ctx context.Context, empId string) ([]model.Employee, error)
	GetEmployeeByFirstName(ctx context.Context, firstName string) ([]model.Employee, error)
	AddEmployee(ctx context.Context, emp model.Employee) error
	UpdateSalary(ctx context.Context, emp model.Employee) error
	DeleteEmployeeById(ctx context.Context, empId string) error
}

type Service struct {
	cv   *config.Configs
	repo employeeRepo
}

func NewService(cv *config.Configs) *Service {
	return &Service{
		cv:   cv,
		repo: NewRepo(cv),
	}
}

func (srv *Service) GetEmployeeById(ctx context.Context, empId string) ([]model.Employee, error) {
	return srv.repo.GetEmployeeById(ctx, empId)
}

func (srv *Service) AddEmployee(ctx context.Context, emp model.Employee) error {
	return srv.repo.AddEmployee(ctx, emp)
}

func (srv *Service) UpdateSalary(ctx context.Context, emp model.Employee) error {
	return srv.repo.UpdateSalary(ctx, emp)
}

func (srv *Service) DeleteEmployeeById(ctx context.Context, empId string) error {
	return srv.repo.DeleteEmployeeById(ctx, empId)
}

func (srv *Service) GetEmployeeByFirstName(ctx context.Context, firstName string) ([]model.Employee, error) {
	return nil, nil
}
