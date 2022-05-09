package controllers

import (
	"elivate9ja-go/app/admin/dtos"
	"elivate9ja-go/app/admin/entities"
	"elivate9ja-go/app/admin/requests"
	"elivate9ja-go/app/admin/services"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type AdminController struct {
	adminService services.IAdminService
	validate     *validator.Validate
}

func NewAdminController(adminService services.IAdminService, validate *validator.Validate) (*AdminController, error) {
	return &AdminController{adminService, validate}, nil
}

func (ac *AdminController) GetAdminsCount(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"count":   ac.adminService.GetAllAdminsCount(),
	})
}

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}

	validationErrs := err.(validator.ValidationErrors)
	for _, e := range validationErrs{
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr) 
	}

	return errs
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

		//https://raw.githubusercontent.com/go-playground/validator/master/_examples/translations/main.go
		english := en.New()
		uni := ut.New(english, english)
		trans, _ := uni.GetTranslator("en")
		_ = en_translations.RegisterDefaultTranslations(ac.validate, trans)

		errs := translateError(err, trans)
		fmt.Printf("%v\n",errs)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": errs[0].Error(),
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

func (ac *AdminController) ViewAdmin(c *fiber.Ctx) error {
	var err error
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	admin, err := ac.adminService.GetAdmin(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	viewAdminDTO := dtos.GetViewAdminDTO(admin)

	return c.JSON(fiber.Map{
		"success": true,
		"data": viewAdminDTO,
	})
}
