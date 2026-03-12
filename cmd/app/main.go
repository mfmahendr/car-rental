package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var port = flag.Int("port", 8080, "Port to listen on")
var address = flag.String("address", "localhost", "Address for the API to bind to")

func main() {
	flag.Parse()
	app := fiber.New(fiber.Config{
		AppName: "Car Rental App",
		// Prefork: true,
	})

	app.Get("/", func (c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	addr := fmt.Sprintf("%s:%d", *address, *port)
	log.Println("server starting on", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf(err.Error())
	}
}
