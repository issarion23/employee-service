package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/issarion23/employee-service/internal/service"
	"go.uber.org/zap"
)

type Handler struct {
	service *service.EmployeeService
	logger  *zap.Logger
}

func NewHandler(s *service.EmployeeService, l *zap.Logger) *Handler {
	return &Handler{service: s, logger: l}
}

type CreateEmployeeRequest struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	City     string `json:"city"`
}

func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var req CreateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	emp, err := h.service.CreateEmployee(r.Context(), req.FullName, req.Phone, req.City)
	if err == service.ErrValidation {
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	} else if err != nil {
		h.logger.Error("create employee failed", zap.Error(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/v1/employees/%s", emp.ID))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func (h *Handler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/employees/"):]
	emp, err := h.service.GetEmployee(r.Context(), id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(emp)
}
