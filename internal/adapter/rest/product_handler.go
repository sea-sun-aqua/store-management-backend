package rest

import (
	"github.com/FLUKKIES/marketplace-backend/domain/exceptions"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	Register(c *fiber.Ctx) error
	GetAllProducts(c *fiber.Ctx) error
	FindByName(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
}

type productHandler struct {
	service usecases.ProductUseCase
}


func NewProductHandler(service usecases.ProductUseCase) ProductHandler {
	return &productHandler{
		service: service,
	}
}

// FindByID implements ProductHandler.
func (p *productHandler) FindByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FindByName implements ProductHandler.
func (p *productHandler) FindByName(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAllProducts implements ProductHandler.
func (p *productHandler) GetAllProducts(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Register implements ProductHandler.
func (p *productHandler) Register(c *fiber.Ctx) error {
	var req *requests.ProductRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := p.service.Register(c.Context(), req); err != nil {
		switch err {
		case exceptions.ErrDuplicatedEmail:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email already registered",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product registered successfully",
	})
}
