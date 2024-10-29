package rest

import (
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

type OrderDetailHandler interface {
	GetAll(c *fiber.Ctx) error
}

type orderDetailHandler struct {
	service usecases.OrderDetailUseCase
}

func NewOrderDetailHandler(service usecases.OrderDetailUseCase) OrderDetailHandler {
	return &orderDetailHandler{
		service: service,
	}
}

func (o *orderDetailHandler) GetAll(c *fiber.Ctx) error {
	orderDetail, err := o.service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(orderDetail)
}

