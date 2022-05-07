package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Admin struct {
	bun.BaseModel `bun:"table:admins,alias:a"`
	ID        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	CreatedAt time.Time ``
	UpdatedAt time.Time ``
	FullName  string    ``
	Email     string    ``
	Password  string    ``
}
