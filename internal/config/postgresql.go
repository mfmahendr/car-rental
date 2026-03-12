package config

import (
	"time"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	UserName string
	Password string
	DBName   string

	MaxConns        int32
	MinConns        int32
	MaxConnLifeTime time.Duration
	MaxConnIdleTime time.Duration
}
