package main

import (
	"GoTracker/internal/middleware"
	"log"
	"net/http"

	apphttp "GoTracker/internal/http"
	"GoTracker/internal/repository"
	"GoTracker/internal/service"
)

func main() {
	repo := repository.NewInMemoryOrderRepo()
	svc := service.NewOrderService(repo)
	server := apphttp.NewHandler(svc)

	handler := middleware.RecoveryMiddleware(
		middleware.LoggingMiddleware(server),
	)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
