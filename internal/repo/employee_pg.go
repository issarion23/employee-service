package repo

import (
	"github.com/issarion23/employee-service/internal/models"

	"github.com/jmoiron/sqlx"
)

type employeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) Create(employee *models.Employee) error {
	query := `
		INSERT INTO employees (full_name, phone, city)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	return r.db.QueryRow(query, employee.FullName, employee.Phone, employee.City).
		Scan(&employee.ID, &employee.CreatedAt, &employee.UpdatedAt)
}

func (r *employeeRepository) GetByID(id int) (*models.Employee, error) {
	var employee models.Employee
	query := `SELECT id, full_name, phone, city, created_at, updated_at FROM employees WHERE id = $1`
	err := r.db.Get(&employee, query, id)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	query := `SELECT id, full_name, phone, city, created_at, updated_at FROM employees ORDER BY id DESC`
	err := r.db.Select(&employees, query)
	return employees, err
}
