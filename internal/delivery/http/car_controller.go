package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mfmahendr/car-rental/internal/application"
	"github.com/mfmahendr/car-rental/internal/application/input"
	"github.com/mfmahendr/car-rental/internal/delivery/http/request"
)

type CarController struct {
	usecase application.CarUsecase
}

func NewCarController(u application.CarUsecase) *CarController {
	return &CarController{u}
}

func (h *CarController) List(c *fiber.Ctx) error {
	in := input.PaginationInput{
		Page: c.QueryInt("page", 1),
		Size: c.QueryInt("size", 10),
	}

	res, err := h.usecase.ListCars(c.Context(), in)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *CarController) Get(c *fiber.Ctx) error {
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

func (h *CarController) Create(c *fiber.Ctx) error {
	var in input.CreateCarInput
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.usecase.CreateCar(c.Context(), in)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *CarController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id format"})
	}

	var req request.UpdateCar
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	in := input.UpdateCarInput{
		Name: req.Name,
		Stock: req.Stock,
		DailyRent: req.DailyRent,
	}

	if err := h.usecase.UpdateCar(c.Context(), uint(id), in); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *CarController) UpdateStock(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id format"})
	}

	var req request.UpdateCarStock
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err = h.usecase.UpdateStock(c.Context(), int64(id), req.Stock); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *CarController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id format"})
	}

	if err = h.usecase.DeleteCar(c.Context(), uint(id)); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
