package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mfmahendr/car-rental/internal/application"
	"github.com/mfmahendr/car-rental/internal/application/input"
	"github.com/mfmahendr/car-rental/internal/delivery/http/request"
)

type BookingController struct {
	usecase application.BookingUsecase
}

func NewBookingController(u application.BookingUsecase) *BookingController {
	return &BookingController{usecase: u}
}

func (h *BookingController) List(c *fiber.Ctx) error {
	in := input.PaginationInput{
		Page: c.QueryInt("page", 1),
		Size: c.QueryInt("size", 10),
	}


	res, err := h.usecase.ListBookings(c.Context(), in)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *BookingController) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	res, err := h.usecase.FindByID(c.Context(), uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *BookingController) Create(c *fiber.Ctx) error {
	var req request.CreateBooking

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	in := input.CreateBookingInput{
		CustomerID: req.CustomerID,
		CarID:      req.CarID,
		StartRent:  req.StartRent,
		EndRent:    req.EndRent,
	}

	res, err := h.usecase.CreateBooking(c.Context(), in)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *BookingController) UpdateRentDate(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var req request.UpdateBookingRentDate
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	in := input.UpdateBookingRentDateInput{
		StartRent: req.StartRent,
		EndRent:   req.EndRent,
	}

	if err := h.usecase.UpdateRentDate(c.Context(), uint(id), in); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *BookingController) FinishBooking(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.usecase.FinishBooking(c.Context(), uint(id)); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *BookingController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.usecase.DeleteBooking(c.Context(), uint(id)); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *BookingController) CustomerHistory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	res, err := h.usecase.GetCustomerBookingHistory(c.Context(), uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
