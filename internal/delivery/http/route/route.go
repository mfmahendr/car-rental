package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mfmahendr/car-rental/internal/delivery/http"
)

type Router struct {
	App                *fiber.App
	CarController      *http.CarController
	CustomerController *http.CustomerController
	BookingController  *http.BookingController
}

func (r *Router) RegisterCustomerRoutes() {
	customer := r.App.Group("/api/customers")

	customer.Get("/", r.CustomerController.List)
	customer.Get("/:id", r.CustomerController.Get)
	customer.Post("/", r.CustomerController.Create)
	customer.Put("/:id", r.CustomerController.Update)
	customer.Delete("/:id", r.CustomerController.Delete)

	customer.Get("/:id/bookings", r.BookingController.CustomerHistory)
}

func (r *Router) RegisterCarRoutes() {
	car := r.App.Group("/api/cars")

	car.Get("/", r.CarController.List)
	car.Get("/:id", r.CarController.Get)
	car.Post("/", r.CarController.Create)
	car.Put("/:id", r.CarController.Update)
	car.Patch("/:id/stock", r.CarController.UpdateStock)
	car.Delete("/:id", r.CarController.Delete)
}

func (r *Router) RegisterBookingRoutes() {
	booking := r.App.Group("/api/bookings")

	booking.Get("/", r.BookingController.List)
	booking.Get("/:id", r.BookingController.Get)
	booking.Post("/", r.BookingController.Create)

	booking.Patch("/:id/rent-date", r.BookingController.UpdateRentDate)
	booking.Patch("/:id/finish", r.BookingController.FinishBooking)

	booking.Delete("/:id", r.BookingController.Delete)
}
