package main

import (
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

	handler := apphttp.RecoveryMiddleware(
		apphttp.LoggingMiddleware(server),
	)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
