package http

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mfmahendr/car-rental/internal/application"
	"github.com/mfmahendr/car-rental/internal/domain"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// application Errors
	if errors.Is(err, application.ErrCarNotFound) || errors.Is(err, application.ErrBookingNotFound) || errors.Is(err, domain.ErrNotFound) {
		code = fiber.StatusNotFound
		message = err.Error()
	} else if errors.Is(err, application.ErrCarOutOfStock) || errors.Is(err, domain.ErrDuplicate) {
		code = fiber.StatusConflict
		message = err.Error()
	} else if errors.Is(err, application.ErrBookingFinished) || errors.Is(err, application.ErrInvalidRentDate) {
		code = fiber.StatusBadRequest
		message = err.Error()
	}

	// validation errors
	if _, ok := err.(validator.ValidationErrors); ok {
		code = fiber.StatusBadRequest
		message = "Validation error: " + err.Error()
	}

	log.Println("Error:", err)

	return c.Status(code).JSON(fiber.Map{
		"error": message,
	})
}
