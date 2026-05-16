package main

import (
	"GoTracker/internal/middleware"
	"log"
	"net/http"

	db "GoTracker/internal"
	"GoTracker/internal/cache"
	apphttp "GoTracker/internal/http"
	"GoTracker/internal/queue"
	"GoTracker/internal/repository"
	"GoTracker/internal/service"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")

	dbase := db.MustConnect(dsn)

	defer dbase.Close()

	cacheTTL, _ := strconv.Atoi(os.Getenv("REDIS_TTL"))
	redisCache := cache.NewRedisCache(os.Getenv("REDIS_ADDR"), cacheTTL)

	queue.StartConsumer()

	repo := repository.NewPostgresOrderRepo(dbase)
	svc := service.NewOrderService(repo, redisCache)
	server := apphttp.NewHandler(svc)

	handler := middleware.RecoveryMiddleware(
		middleware.LoggingMiddleware(server),
	)

	log.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
