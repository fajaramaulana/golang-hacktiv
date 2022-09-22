package repositories

import (
	"sesi7/cmd/entities"

	"gorm.io/gorm"
)

type RepositoryEmployee interface {
	CreateEmployee(employee entities.Employee) (entities.Employee, error)
	GetAllEmployee() ([]entities.Employee, error)
}

type repositories struct {
	db *gorm.DB
}

func NewRepositoryEmployee(db *gorm.DB) RepositoryEmployee {
	return &repositories{
		db: db,
	}
}

func (r *repositories) CreateEmployee(employee entities.Employee) (entities.Employee, error) {
	err := r.db.Create(&employee).Error

	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repositories) GetAllEmployee() ([]entities.Employee, error) {
	var employees []entities.Employee
	err := r.db.Find(&employees).Error
	if err != nil {
		return nil, err
	}

	return employees, nil
}
