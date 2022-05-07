package entities

import (
	"time"

	"github.com/google/uuid"
)

type AdminEntity struct {
	ID         uuid.UUID
	Name       string
	Email      string
	SuperAdmin bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
