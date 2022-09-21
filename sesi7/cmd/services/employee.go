package services

import (
	"sesi7/cmd/entities"
	"sesi7/cmd/repositories"
	"sesi7/cmd/request"
)

type EmployeeService interface {
	CreateEmployee(input request.RegisterEmployee) (entities.Employee, error)
	GetAllEmployee() ([]entities.Employee, error)
}

type services struct {
	repository repositories.RepositoryEmployee
}

func EmployeeNewService(repository repositories.RepositoryEmployee) *services {
	return &services{repository}
}

func (s *services) CreateEmployee(input request.RegisterEmployee) (entities.Employee, error) {
	// fmt.Printf("%# v", input)
	employee := entities.Employee{}
	employee.Full_name = input.Full_name
	employee.Email = input.Email
	employee.Age = input.Age
	employee.Division = input.Division

	// fmt.Printf("%#v \n", employee)

	newEmployee, err := s.repository.CreateEmployee(employee)

	if err != nil {
		return newEmployee, err
	}

	return newEmployee, nil
}

func (s *services) GetAllEmployee() ([]entities.Employee, error) {
	employees, err := s.repository.GetAllEmployee()

	if err != nil {
		return employees, err
	}

	return employees, nil
}
