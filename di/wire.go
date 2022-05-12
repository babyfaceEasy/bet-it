//go:build wireinject
// +build wireinject

package di

import (
	"elivate9ja-go/app/admin/controllers"
	"elivate9ja-go/app/admin/repositories"
	"elivate9ja-go/app/admin/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/uptrace/bun"
)

func provideValidator() *validator.Validate {
	return validator.New()
}

var AdminControllerSet = wire.NewSet(
	services.NewAdminService, wire.Bind(new(services.IAdminService), new(*services.AdminService)),
	repositories.NewAdminRepository, wire.Bind(new(repositories.IAdminRepository), new(*repositories.AdminRepository)),
	provideValidator )

func InitializeAdminController(db bun.IDB) (*controllers.AdminController, error) {
	wire.Build(controllers.NewAdminController, AdminControllerSet)
	return &controllers.AdminController{}, nil
}

func InitializeAdminService()
