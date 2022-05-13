package services

import (
	"elivate9ja-go/app/admin/entities"

	"github.com/google/uuid"
)

type IAdminService interface {
	VerifyAdmin(email string, password string) (bool, error)
	GetAdminByEmail(email string) (entities.AdminEntity, error)
	GetAllAdminsCount() int
	CreateAdmin(entities.AdminEntity) (bool, error)
	GetAdmin(uuid.UUID) (entities.AdminEntity, error)
	GetAdmins() ([]entities.AdminEntity, error)
}
