package controllers

import (
	"elivate9ja-go/app/auth/requests"
	"elivate9ja-go/app/auth/services"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService services.IAuthService
}

func NewAuthController(authService services.IAuthService) (*AuthController, error) {
	return &AuthController{authService}, nil
}

func (ac *AuthController) CustomerLogin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "handler for customer login",
	})
}

func (ac *AuthController) AdminLogin(c *fiber.Ctx) error {

	logAdminInRequest := &requests.LogAdminInRequest{}
	if err := c.BodyParser(logAdminInRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	err := ac.authService.LogAdminIn(logAdminInRequest.Email, logAdminInRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "handler for admin login",
	})
}
