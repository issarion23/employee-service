package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/issarion23/employee-service/internal/handler"
	"github.com/issarion23/employee-service/internal/models"
)

type mockEmployeeService struct{}

func (m *mockEmployeeService) CreateEmployee(req *models.CreateEmployeeRequest) (*models.Employee, error) {
	return &models.Employee{
		ID:       1,
		FullName: req.FullName,
		Phone:    req.Phone,
		City:     req.City,
	}, nil
}

func (m *mockEmployeeService) GetEmployeeByID(id int) (*models.Employee, error) {
	return &models.Employee{
		ID:       id,
		FullName: "Tecтонов Тест",
		Phone:    "+77001234567",
		City:     "Алматы",
	}, nil
}

func (m *mockEmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return []models.Employee{
		{ID: 1, FullName: "Tecтонов Тест", Phone: "+77001234567", City: "Алматы"},
		{ID: 2, FullName: "Тестобеков ТТ", Phone: "+77007654321", City: "Астана"},
	}, nil
}

func TestCreateEmployee(t *testing.T) {
	mockService := &mockEmployeeService{}
	h := handler.NewEmployeeHandler(mockService)

	reqBody := models.CreateEmployeeRequest{
		FullName: "Тест Тест",
		Phone:    "+77007777777",
		City:     "Алматы",
	}
	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/employees", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateEmployee)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var employee models.Employee
	if err := json.NewDecoder(rr.Body).Decode(&employee); err != nil {
		t.Fatal(err)
	}

	if employee.FullName != reqBody.FullName {
		t.Errorf("handler returned unexpected body: got %v want %v", employee.FullName, reqBody.FullName)
	}
}

func TestGetAllEmployees(t *testing.T) {
	mockService := &mockEmployeeService{}
	h := handler.NewEmployeeHandler(mockService)

	req, err := http.NewRequest("GET", "/api/employees", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetAllEmployees)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var employees []models.Employee
	if err := json.NewDecoder(rr.Body).Decode(&employees); err != nil {
		t.Fatal(err)
	}

	if len(employees) != 2 {
		t.Errorf("handler returned unexpected number of employees: got %v want %v", len(employees), 2)
	}
}
