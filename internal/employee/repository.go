package employee

import (
	"basic_golang_echo/internal/config"
	"basic_golang_echo/internal/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Repo struct {
	cv *config.Configs
	db *gorm.DB
}

func NewRepo(cv *config.Configs) *Repo {
	return &Repo{
		cv: cv,
		db: config.GetDBInstance(),
	}
}

func (repo *Repo) GetEmployeeById(ctx context.Context, empId string) ([]model.Employee, error) {
	var em []model.Employee
	// if err := repo.db.Table("employee").Where("id = ?", empId).Find(&em).Error; err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	if err := repo.db.Raw("SELECT * FROM Employee WHERE ID = ?", empId).Scan(&em).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return em, nil
}

func (repo *Repo) AddEmployee(ctx context.Context, emp model.Employee) error {
	// if err := repo.db.Exec("INSERT INTO Employee VALUES(?, ?, ?, ?)", emp.ID, emp.Firstname, emp.Lastname, emp.Salary).
	// 	Error; err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	if err := repo.db.Table("employee").Create(&emp).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *Repo) UpdateSalary(ctx context.Context, emp model.Employee) error {
	// if err := repo.db.Exec("UPDATE Employee SET Salary = ? WHERE ID = ?", emp.Salary, emp.ID).Error; err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	if err := repo.db.Table("employee").Model(&model.Employee{}).
		Where("ID = ?", emp.ID).
		Update("salary", emp.Salary).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *Repo) DeleteEmployeeById(ctx context.Context, empId string) error {
	// if err := repo.db.Table("employee").Exec("DELETE FROM Employee WHERE ID = ?", empId).Error; err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	if err := repo.db.Table("employee").Where("ID = ?", empId).Delete(model.Employee{}).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *Repo) GetEmployeeByFirstName(ctx context.Context, firstName string) ([]model.Employee, error) {
	var em []model.Employee
	return em, nil
}
