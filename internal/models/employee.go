package models

import "time"

type Employee struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}
