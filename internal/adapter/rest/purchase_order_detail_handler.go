package rest

import (
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

type PurchaseOrderDetailHandler interface {
	GetAll(c *fiber.Ctx) error
}

type purchaseOrderDetailHandler struct {
	service usecases.PurchaseOrderDetailUseCase
}

func NewPurchaseOrderDetailHandler(service usecases.PurchaseOrderDetailUseCase) PurchaseOrderDetailHandler {
	return &purchaseOrderDetailHandler{
		service: service,
	}
}

func (p *purchaseOrderDetailHandler) GetAll(c *fiber.Ctx) error {
	purchaseOrderDetail, err := p.service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(purchaseOrderDetail)
}
