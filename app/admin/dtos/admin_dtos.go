package dtos

import (
	"elivate9ja-go/app/admin/entities"
	"time"

	"github.com/google/uuid"
)

type ViewAdminDTO struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetViewAdminDTO(adminEntity entities.AdminEntity) ViewAdminDTO {
	return ViewAdminDTO{
		ID: adminEntity.ID,
		Name: adminEntity.Name,
		Email: adminEntity.Email,
		CreatedAt: adminEntity.CreatedAt,
		UpdatedAt: adminEntity.UpdatedAt,
	}
}
