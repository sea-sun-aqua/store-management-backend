package rest

import (
	"github.com/FLUKKIES/marketplace-backend/domain/exceptions"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	UpdateStatusOrder(c *fiber.Ctx) error
}

type orderHandler struct {
	service usecases.OrderUseCase
}

func NewOrderHandler(service usecases.OrderUseCase) OrderHandler {
	return &orderHandler{
		service: service,
	}
}

func (o *orderHandler) Create(c *fiber.Ctx) error {
	var req *requests.OrderCreateRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	

	if err := o.service.Create(c.Context(), req); err != nil {
		switch err {
		case exceptions.ErrDuplicatedIDOrder:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "ID already registered",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
	})
}

func (o *orderHandler) GetAll(c *fiber.Ctx) error {
	orders, err := o.service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(orders)
}

func (o *orderHandler) UpdateStatusOrder(c *fiber.Ctx) error {
	orderID := c.Params("OrderID")
	var req *requests.OrderUpdateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	order, err := o.service.UpdateStatusByID(c.Context(), orderID, req)
	if err != nil {
		switch err {
		case exceptions.ErrStatusInvalid:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order status invalid",
			})
		case exceptions.ErrOrderIDNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(order)
}
