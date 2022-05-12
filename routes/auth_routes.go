package routes

import "github.com/gofiber/fiber/v2"


func getAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/login", func(c *fiber.Ctx)error {
		return c.JSON(fiber.Map{
			"message": "Customer's login",
		})
	})

	auth.Post("/register", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Customer's registration endpoint.",
		})
	})

	auth.Post("/admin-login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Admin's login",
		})
	})

}