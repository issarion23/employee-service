package service

import (
	"github.com/issarion23/employee-service/internal/models"
	"github.com/issarion23/employee-service/internal/repo"

	"errors"
)

type employeeService struct {
	repo repo.EmployeeRepository
}

func NewEmployeeService(repo repo.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CreateEmployee(req *models.CreateEmployeeRequest) (*models.Employee, error) {
	if req.FullName == "" || req.Phone == "" || req.City == "" {
		return nil, errors.New("all fields are required")
	}

	employee := &models.Employee{
		FullName: req.FullName,
		Phone:    req.Phone,
		City:     req.City,
	}

	if err := s.repo.Create(employee); err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *employeeService) GetEmployeeByID(id int) (*models.Employee, error) {
	if id <= 0 {
		return nil, errors.New("invalid employee ID")
	}
	return s.repo.GetByID(id)
}

func (s *employeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.repo.GetAll()
}
