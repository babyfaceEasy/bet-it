//go:build wireinject
// +build wireinject

package di

import (
	"elivate9ja-go/app/admin/controllers"
	"elivate9ja-go/app/admin/repositories"
	"elivate9ja-go/app/admin/services"
	authControllers "elivate9ja-go/app/auth/controllers"
	authServices "elivate9ja-go/app/auth/services"

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
	provideValidator)

var AdminServiceSet = wire.NewSet(repositories.NewAdminRepository, wire.Bind(new(repositories.IAdminRepository), new(*repositories.AdminRepository)))

var AuthServiceSet = wire.NewSet(
	authServices.NewAuthService, wire.Bind(new(authServices.IAuthService),new(*authServices.AuthService)),
	services.NewAdminService, wire.Bind(new(services.IAdminService), new(*services.AdminService)),
	repositories.NewAdminRepository, wire.Bind(new(repositories.IAdminRepository), new(*repositories.AdminRepository)),
)

func InitializeAdminController(db bun.IDB) (*controllers.AdminController, error) {
	wire.Build(controllers.NewAdminController, AdminControllerSet)
	return &controllers.AdminController{}, nil
}

func InitializeAdminService(db bun.IDB) (*services.AdminService, error) {
	wire.Build(services.NewAdminService, AdminServiceSet)
	return &services.AdminService{}, nil
}

func InitializeAuthController(db bun.IDB) (*authControllers.AuthController, error) {
	wire.Build(authControllers.NewAuthController, AuthServiceSet)
	return &authControllers.AuthController{}, nil
}
