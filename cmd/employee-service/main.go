package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/issarion23/employee-service/internal/handler"
	"github.com/issarion23/employee-service/internal/repo"
	"github.com/issarion23/employee-service/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		logger.Fatal("failed to connect to db", zap.Error(err))
	}
	defer dbpool.Close()

	employeeRepo := repo.NewEmployeeRepo(dbpool)
	employeeService := service.NewEmployeeService(employeeRepo)
	h := handler.NewHandler(employeeService, logger)

	http.HandleFunc("/v1/employees", h.CreateEmployee)
	http.HandleFunc("/v1/employees/", h.GetEmployee)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Info("starting server", zap.String("port", port))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
