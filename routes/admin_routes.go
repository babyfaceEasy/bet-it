package routes

import (
	"elivate9ja-go/di"
	"elivate9ja-go/utils"

	"github.com/gofiber/fiber/v2"
)

func getAdminRoutes(router fiber.Router) {

	db := utils.GetDBConnection()

	//controllers
	//adminController := controllers.NewAdminController()
	adminController, err := di.InitializeAdminController(db)
	if err != nil {
		panic(err)
	}

	// group admin routes
	admin := router.Group("/admin")

	admin.Get("/", adminController.GetAdminsCount)
	admin.Post("/", adminController.NewAdmin)
	admin.Get("/:id", adminController.ViewAdmin)

}
