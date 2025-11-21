package test

import (
	"bytes"
	"context"
	"github.com/issarion23/employee-service/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/issarion23/employee-service/internal/handler"
	"github.com/issarion23/employee-service/internal/service"
	"go.uber.org/zap"
)

type mockService struct{}

func (m *service.EmployeeService) CreateEmployee(_ context.Context, fullName, phone, city string) (*models.Employee, error) {
	return &models.Employee{
		ID:       "uuid",
		FullName: fullName,
		Phone:    phone,
		City:     city,
	}, nil
}

func TestCreateEmployeeHandler(t *testing.T) {
	svc := &mockService{}
	h := handler.NewHandler(svc, zap.NewNop())
	reqBody := `{"full_name":"Иван","phone":"+77771234567","city":"Алматы"}`
	req := httptest.NewRequest(http.MethodPost, "/v1/employees", bytes.NewReader([]byte(reqBody)))
	w := httptest.NewRecorder()
	h.CreateEmployee(w, req)
	if w.Result().StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Result().StatusCode)
	}
}
