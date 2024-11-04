package rest

import (

	"github.com/FLUKKIES/marketplace-backend/domain/exceptions"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

type StaffHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type staffHandler struct {
	service usecases.StaffUseCase
}

func NewStaffHandler(service usecases.StaffUseCase) StaffHandler {
	return &staffHandler{
		service: service,
	}
}

func (s *staffHandler) Login(c *fiber.Ctx) error {
	var req *requests.StaffLoginRequest
		if err := c.BodyParser(&req); err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	

	// Login staff
	staff, err := s.service.Login(c.Context(), req)

	if err != nil {
		switch err {
		case exceptions.ErrLoginFailed:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Login failed",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(staff)

}

func (s *staffHandler) Register(c *fiber.Ctx) error {
	var req *requests.StaffRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	// Register staff
	if err := s.service.Register(c.Context(), req); err != nil {
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
		"message": "Staff registered successfully",
	})
}
