package services

import (
	"elivate9ja-go/app/admin/entities"
	"elivate9ja-go/app/admin/repositories"

	"github.com/google/uuid"
)


type AdminService struct {
	repo repositories.IAdminRepository
}

func NewAdminService(repo repositories.IAdminRepository) (*AdminService, error) {
	return &AdminService{repo}, nil
}

func (as *AdminService) GetAllAdminsCount() int {
	totalAdmins :=  as.repo.CountAdmins()

	return totalAdmins
}

func (as *AdminService) CreateAdmin(adminEntity entities.AdminEntity) (bool, error) {
	response, err :=  as.repo.SaveAdmin(adminEntity)
	if err != nil {
		// plan, format a service like error based on the db type error pased back.
		return response, err
	}

	return response, err
}

func (as * AdminService) GetAdmin(identifier uuid.UUID) (entities.AdminEntity, error) {
	response, err := as.repo.GetAdmin(identifier)
	if err != nil {
		return response, err
	}

	return response, err
}

func (as *AdminService) GetAdmins() ([]entities.AdminEntity, error) {
	response, err := as.repo.GetAdmins()
	if err != nil {
		return response, err
	}

	return response, err
}