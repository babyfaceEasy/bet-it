package main

import (
	"elivate9ja-go/configs"
	"elivate9ja-go/routes"
	"elivate9ja-go/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load env values
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading env values")
	}

	fiberConfig := configs.FiberConfig()

	// create fiber app
	app := fiber.New(fiberConfig)

	// routes
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "good"})
	})
	routes.GetAPIRoutes(app)
	/*

		// testing jwt middleware
		app.Get("/generate-token", func(c *fiber.Ctx) error {
			userClaim := &jwt.MapClaims{
				"email": "oodegbaro@gmail.com",
				"role": "admin",
				"id": 7,
			}
			token, err := middlewares.Encode(userClaim, 1000)
			if err != nil {
				return c.SendStatus(500)
			}

			return c.SendString(token)
		})

		app.Use(middlewares.New(middlewares.Config{}))

		app.Get("/protected", func(c *fiber.Ctx) error {
			claimData := c.Locals("jwtClaims")

			if claimData == nil {
				return c.SendString("JWT was bypassed")
			} else {
				return c.JSON(claimData)
			}
		})
	*/

	utils.StartServer(app)
}
