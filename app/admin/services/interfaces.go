package services

import "elivate9ja-go/app/admin/entities"

type IAdminService interface {
	GetAllAdminsCount() int
	CreateAdmin(entities.AdminEntity) (bool, error)
}