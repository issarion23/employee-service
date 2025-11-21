package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/issarion23/employee-service/internal/models"
	"github.com/issarion23/employee-service/internal/repo"
	"time"
)

var ErrValidation = errors.New("validation failed")

type EmployeeService struct {
	repo      *repo.EmployeeRepo
	validator *validator.Validate
}

func NewEmployeeService(r *repo.EmployeeRepo) *EmployeeService {
	return &EmployeeService{
		repo:      r,
		validator: validator.New(),
	}
}

type CreateEmployeeRequest struct {
	FullName string `validate:"required,max=256"`
	Phone    string `validate:"required,e164"`
	City     string `validate:"required,max=128"`
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, fullName, phone, city string) (*models.Employee, error) {
	req := CreateEmployeeRequest{FullName: fullName, Phone: phone, City: city}
	if err := s.validator.Struct(req); err != nil {
		return nil, ErrValidation
	}

	e := &models.Employee{
		ID:        uuid.New().String(),
		FullName:  fullName,
		Phone:     phone,
		City:      city,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}

	return e, nil
}

func (s *EmployeeService) GetEmployee(ctx context.Context, id string) (*models.Employee, error) {
	return s.repo.GetByID(ctx, id)
}
