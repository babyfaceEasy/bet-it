package routes

import (
	"github.com/gofiber/fiber/v2"
)


func GetAPIRoutes(app *fiber.App) {
	api := app.Group("api/")

	getAdminRoutes(api)
	getCustomerRoutes(api)

}
