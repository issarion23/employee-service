package repo

import "github.com/issarion23/employee-service/internal/models"

type EmployeeRepository interface {
	Create(employee *models.Employee) error
	GetByID(id int) (*models.Employee, error)
	GetAll() ([]models.Employee, error)
}
