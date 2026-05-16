package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	RedisAddr   string
	RedisTTL    int
	KafkaBroker string
	AppEnv      string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		AppEnv:      os.Getenv("APP_ENV"),
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	if cfg.RedisAddr == "" {
		cfg.RedisAddr = "localhost:6379"
	}

	cfg.RedisTTL = 60
	if ttl := os.Getenv("REDIS_TTL"); ttl != "" {
		value, err := strconv.Atoi(ttl)
		if err != nil {
			return nil, fmt.Errorf("REDIS_TTL must be a valid integer: %w", err)
		}
		cfg.RedisTTL = value
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	if cfg.KafkaBroker == "" {
		return nil, fmt.Errorf("KAFKA_BROKER is required")
	}

	return cfg, nil
}
