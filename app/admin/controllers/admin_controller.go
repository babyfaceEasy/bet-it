package controllers

import (
	"elivate9ja-go/app/admin/entities"
	"elivate9ja-go/app/admin/requests"
	"elivate9ja-go/app/admin/services"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AdminController struct {
	adminService services.IAdminService
	validate     *validator.Validate
}

func NewAdminController(adminService services.IAdminService, validate *validator.Validate) (*AdminController, error) {
	return &AdminController{adminService, validate}, nil
}

func (ac *AdminController) GetAdmins(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"count":   ac.adminService.GetAllAdminsCount(),
	})
}

func (ac *AdminController) NewAdmin(c *fiber.Ctx) error {

	createAdminRequest := &requests.CreateAdminRequest{}
	// check to see if the value recieved is good
	if err := c.BodyParser(createAdminRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	fmt.Printf("Create Admin Request: %v\n", createAdminRequest.SuperAdmin)

	// validate here
	err := ac.validate.Struct(createAdminRequest)
	if err != nil {
		/*
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
		*/

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	adminEntity := entities.AdminEntity{
		ID:        uuid.New(),
		Name:      createAdminRequest.Name,
		Email:     createAdminRequest.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	response, err := ac.adminService.CreateAdmin(adminEntity)
	if err != nil {
		// plan format a http response and send back
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	if !response {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "An error occurred please try again later.",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Admin creation was successful.",
	})
}
