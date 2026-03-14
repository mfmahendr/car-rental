package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mfmahendr/car-rental/internal/application"
	"github.com/mfmahendr/car-rental/internal/application/input"
	"github.com/mfmahendr/car-rental/internal/delivery/http/request"
)

type CustomerController struct {
	usecase application.CustomerUsecase
}

func NewCustomerController(u application.CustomerUsecase) *CustomerController {
	return &CustomerController{u}
}

func (h *CustomerController) List(c *fiber.Ctx) error {
	in := input.PaginationInput{
		Page: c.QueryInt("page", 1),
		Size: c.QueryInt("size", 10),
	}

	res, err := h.usecase.ListCustomers(c.Context(), in)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *CustomerController) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id format"})
	}

	res, err := h.usecase.GetByID(c.Context(), uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *CustomerController) Create(c *fiber.Ctx) error {
	var in input.CreateCustomerInput
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.usecase.CreateCustomer(c.Context(), in)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *CustomerController) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var req request.UpdateCustomer
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	in := input.UpdateCustomerInput{
		Name:        req.Name,
		NIK:         req.NIK,
		PhoneNumber: req.PhoneNumber,
	}

	if err := h.usecase.UpdateCustomer(c.Context(), uint(id), in); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}


func (h *CustomerController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err := h.usecase.DeleteCustomer(c.Context(), uint(id))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
