package routes

import (
	"elivate9ja-go/di"
	"elivate9ja-go/utils"

	"github.com/gofiber/fiber/v2"
)

func getAuthRoutes(router fiber.Router) {

	db := utils.GetDBConnection()
	authController, err := di.InitializeAuthController(db)
	if err != nil {
		panic(err)
	}

	auth := router.Group("/auth")

	auth.Post("/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Customer's login",
		})
	})

	auth.Post("/register", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Customer's registration endpoint.",
		})
	})

	auth.Post("/admin-login", authController.AdminLogin)

}
