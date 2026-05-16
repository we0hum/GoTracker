package main

import (
	db "GoTracker/internal"
	"GoTracker/internal/cache"
	"GoTracker/internal/config"
	apphttp "GoTracker/internal/http"
	"GoTracker/internal/middleware"
	"GoTracker/internal/queue"
	"GoTracker/internal/repository"
	"GoTracker/internal/service"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbase := db.MustConnect(cfg.DatabaseURL)
	defer dbase.Close()

	redisCache := cache.NewRedisCache(cfg.RedisAddr, cfg.RedisTTL)
	defer redisCache.Close()

	queue.StartConsumer(cfg.KafkaBroker)

	repo := repository.NewPostgresOrderRepo(dbase)
	orderService := service.NewOrderService(repo, redisCache)
	router := apphttp.NewRouter(orderService)

	handler := middleware.RecoveryMiddleware(
		middleware.LoggingMiddleware(router),
	)

	log.Printf("Сервер запущен на http://localhost:%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatal(err)
	}
}
