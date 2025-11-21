package repo

import (
	"context"
	"github.com/issarion23/employee-service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepo struct {
	db *pgxpool.Pool
}

func NewEmployeeRepo(db *pgxpool.Pool) *EmployeeRepo {
	return &EmployeeRepo{db: db}
}

func (r *EmployeeRepo) Create(ctx context.Context, e *models.Employee) error {
	q := `INSERT INTO employees (id, full_name, phone, city, created_at) VALUES ($1,$2,$3,$4,$5)`
	_, err := r.db.Exec(ctx, q, e.ID, e.FullName, e.Phone, e.City, e.CreatedAt)
	return err
}

func (r *EmployeeRepo) GetByID(ctx context.Context, id string) (*models.Employee, error) {
	e := &models.Employee{}
	q := `SELECT id, full_name, phone, city, created_at FROM employees WHERE id=$1`
	err := r.db.QueryRow(ctx, q, id).Scan(&e.ID, &e.FullName, &e.Phone, &e.City, &e.CreatedAt)
	if err != nil {
		return nil, err
	}
	return e, nil
}
