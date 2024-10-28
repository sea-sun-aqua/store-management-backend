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
	FindByID(c *fiber.Ctx) error
	UpdateProductByID(c *fiber.Ctx) error
}

type productHandler struct {
	service usecases.ProductUseCase
}



func NewProductHandler(service usecases.ProductUseCase) ProductHandler {
	return &productHandler{
		service: service,
	}
}

func (p *productHandler) UpdateProductByID(c *fiber.Ctx) error {

	ProductID := c.Params("ProductID")

	var req *requests.ProductUpdateAmountRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product, err := p.service.UpdateProductByID(c.Context(), ProductID, req)
	if err != nil {
		switch err {
		case exceptions.ErrProductNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (p *productHandler) FindByID(c *fiber.Ctx) error {
	ProductID := c.Params("ProductID")
	println(ProductID)
	product, err := p.service.FindByID(c.Context(), ProductID)
	if err != nil {
		switch err {
		case exceptions.ErrProductNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(product)
}

func (p *productHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := p.service.GetAllProducts(c.Context())
	if err != nil {
		switch err {
		case exceptions.ErrProductNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(products)
}

func (p *productHandler) Register(c *fiber.Ctx) error {
	var req *requests.ProductRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := p.service.Register(c.Context(), req); err != nil {
		switch err {
		case exceptions.ErrDuplicatedIDProduct:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "ID already registered",
			})
		case exceptions.ErrDuplicatedNameProduct:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Name already registered",
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
