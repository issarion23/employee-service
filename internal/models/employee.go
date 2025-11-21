package models

import "time"

type Employee struct {
	ID        int       `json:"id" db:"id"`
	FullName  string    `json:"full_name" db:"full_name"`
	Phone     string    `json:"phone" db:"phone"`
	City      string    `json:"city" db:"city"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateEmployeeRequest struct {
	FullName string `json:"full_name" validate:"required,min=2,max=255"`
	Phone    string `json:"phone" validate:"required,min=10,max=20"`
	City     string `json:"city" validate:"required,min=2,max=100"`
}
