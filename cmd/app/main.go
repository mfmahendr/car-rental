package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mfmahendr/car-rental/internal/config"
	"github.com/mfmahendr/car-rental/internal/delivery/http"
	"github.com/mfmahendr/car-rental/internal/infra/database/postgres"
	"github.com/mfmahendr/car-rental/internal/infra/validator"
	"github.com/mfmahendr/car-rental/internal/setup"
)

func main() {
	cfg := config.Load()

	httpServer := fiber.New(fiber.Config{
		AppName:      "Car Rental App",
		ErrorHandler: http.ErrorHandler,
		// Prefork: true,
	})
	validate := validator.NewValidator()
	pool, err := postgres.NewPool(cfg.Db)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return
	}
	appSetup := setup.NewApplication(*cfg, httpServer, pool, validate)
	appSetup.Setup()

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Println("server starting on", addr)
	if err := httpServer.Listen(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
