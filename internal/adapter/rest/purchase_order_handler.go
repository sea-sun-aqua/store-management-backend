package rest

import (
	"github.com/FLUKKIES/marketplace-backend/domain/exceptions"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

type PurchaseOrderHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	UpdateStatusOrder(c *fiber.Ctx) error
}

type purchaseOrderHandler struct {
	service usecases.PurchaseOrderUseCase
}

func NewPurchaseOrderHandler(service usecases.PurchaseOrderUseCase) PurchaseOrderHandler {
	return &purchaseOrderHandler{service: service}
}

func (p *purchaseOrderHandler) Create(c *fiber.Ctx) error {
	var req *requests.PurchaseOrderCreateRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := p.service.Create(c.Context(), req); err != nil {
		switch err {
		case exceptions.ErrDuplicatedPurchaseIDOrder:
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
		"message": "Purchase order created successfully",
	})
}

func (p *purchaseOrderHandler) GetAll(c *fiber.Ctx) error {
	purchaseOrders, err := p.service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(purchaseOrders)
}

func (p *purchaseOrderHandler) UpdateStatusOrder(c *fiber.Ctx) error {
	purchaseOrderID := c.Params("PurchaseOrderID")
	var req *requests.PurchaseOrderUpdateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	purchaseOrder, err := p.service.UpdateStatusByID(c.Context(), purchaseOrderID, req)
	if err != nil {
		switch err {
		case exceptions.ErrStatusInvalid:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Purchase status invalid",
			})
		case exceptions.ErrPurchaseOrderIDNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Purchase order not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(purchaseOrder)
}
