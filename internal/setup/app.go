package setup

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mfmahendr/car-rental/internal/config"
)

type Application struct {
	config     config.AppConfig
	httpServer *fiber.App
	db         *pgxpool.Pool
	val        *validator.Validate
}

func NewApplication(cfg config.AppConfig, httpSrv *fiber.App, db *pgxpool.Pool, v *validator.Validate) *Application {
	return &Application{
		config:     cfg,
		httpServer: httpSrv,
		db:         db,
		val:        v,
	}
}
