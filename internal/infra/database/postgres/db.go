package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mfmahendr/car-rental/internal/config"
)

func NewPool(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.DBName,
	)

	pgCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse postgres config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres pool connection: %w", err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	return pool, nil
}