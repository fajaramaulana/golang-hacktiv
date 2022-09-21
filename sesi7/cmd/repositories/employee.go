package repositories

import (
	"database/sql"
	"sesi7/cmd/entities"
)

type RepositoryEmployee interface {
	CreateEmployee(employee entities.Employee) (entities.Employee, error)
	GetAllEmployee() ([]entities.Employee, error)
}

type repositories struct {
	db *sql.DB
}

func NewRepositoryEmployee(db *sql.DB) RepositoryEmployee {
	return &repositories{
		db: db,
	}
}

func (r *repositories) CreateEmployee(employee entities.Employee) (entities.Employee, error) {
	var newEmployee entities.Employee
	sqlStatement := `INSERT INTO employees (full_name, email, age, division) VALUES ($1, $2, $3, $4) RETURNING id, full_name, email, age, division`

	stmt, err := r.db.Prepare(sqlStatement)

	if err != nil {
		return newEmployee, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(employee.Full_name, employee.Email, employee.Age, employee.Division)
	// err = r.db.QueryRow(sqlStatement, employee.Full_name, employee.Email, employee.Age, employee.Division).Scan(&newEmployee.Id, &newEmployee.Full_name, &newEmployee.Email, &newEmployee.Age, &newEmployee.Division)

	if err != nil {
		return newEmployee, err
	}

	return newEmployee, nil
}

func (r *repositories) GetAllEmployee() ([]entities.Employee, error) {
	var employees []entities.Employee
	sqlStatement := `SELECT id, full_name, email, age, division FROM employees`

	stmt, err := r.db.Prepare(sqlStatement)

	if err != nil {
		return employees, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return employees, err
	}

	if err != nil {
		return employees, err
	}

	defer rows.Close()

	for rows.Next() {
		var employee entities.Employee
		err = rows.Scan(&employee.Id, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			return employees, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}
