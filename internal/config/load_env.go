package config

import (
	"time"

	"github.com/joho/godotenv"
)

func Load() *AppConfig {
	_ = godotenv.Load()

	cfg := &AppConfig{
		App: AppMeta{
			AppName:    Get("APP_NAME", "app"),
			AppVersion: Get("APP_VERSION", "0.0.1"),
		},

		Server: ServerConfig{
			Host: Get("SERVER_HOST", "0.0.0.0"),
			Port: Get("SERVER_PORT", 8080),
		},

		Db: DatabaseConfig{
			Host:     Get("DB_HOST", "localhost"),
			Port:     Get("DB_PORT", 5432),
			UserName: Get("DB_USERNAME", "postgres"),
			Password: Get("DB_PASSWORD", ""),
			DBName:   Get("DB_NAME", "postgres"),

			MaxConns: int32(Get("DB_MAX_CONNS", 10)),
			MinConns: int32(Get("DB_MIN_CONNS", 2)),

			MaxConnLifeTime: time.Duration(Get("DB_MAX_CONN_LIFETIME", 3600)) * time.Second,
			MaxConnIdleTime: time.Duration(Get("DB_MAX_CONN_IDLE_TIME", 300)) * time.Second,
		},
	}

	return cfg
}