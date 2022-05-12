package controllers

import "github.com/gofiber/fiber/v2"

type AuthController struct {}

func NewAuthController() {}

func (ac *AuthController) CustomerLogin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "handler for customer login",
	})
}


func (ac *AuthController) AdminLogin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "handler for admin login",
	})
}