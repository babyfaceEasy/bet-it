package repositories

import (
	"elivate9ja-go/app/admin/entities"

	"github.com/google/uuid"
)

type IAdminRepository interface {
	CountAdmins() int
	SaveAdmin(entities.AdminEntity) (bool, error)
	GetAdmin(id uuid.UUID) (entities.AdminEntity, error)
	GetAdmins()([]entities.AdminEntity, error)
}
