package routes

import "github.com/gofiber/fiber/v2"

func getCustomerRoutes(router fiber.Router) {
	customer :=  router.Group("/customer")

	customer.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "good to go",
		})
	})
}