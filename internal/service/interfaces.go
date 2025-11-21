package service

import "github.com/issarion23/employee-service/internal/models"

type EmployeeService interface {
	CreateEmployee(req *models.CreateEmployeeRequest) (*models.Employee, error)
	GetEmployeeByID(id int) (*models.Employee, error)
	GetAllEmployees() ([]models.Employee, error)
}
