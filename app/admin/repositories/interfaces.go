package repositories

import "elivate9ja-go/app/admin/entities"

type IAdminRepository interface {
	CountAdmins() int
	SaveAdmin(entities.AdminEntity) (bool, error)
}