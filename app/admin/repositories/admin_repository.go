package repositories

import (
	"context"
	"elivate9ja-go/app/admin/entities"
	"elivate9ja-go/app/admin/models"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

type AdminRepository struct {
	db bun.IDB
}

func NewAdminRepository(db bun.IDB) (*AdminRepository, error) {
	return &AdminRepository{db: db}, nil
}

func (ar *AdminRepository) CountAdmins() int {
	totalAdmins := 0
	ctx := context.Background()

	totalAdmins, err := ar.db.NewSelect().Model((*models.Admin)(nil)).Count(ctx)
	if err != nil {
		log.Printf("Error occured in getting total count. Reason: %v", err)
	}

	return totalAdmins
}

func (ar * AdminRepository) SaveAdmin(adminEntity entities.AdminEntity) (bool, error) {
	adminModel := &models.Admin{
		ID: adminEntity.ID,
		FullName: adminEntity.Name,
		Email: adminEntity.Email,
	}

	ctx := context.Background()

	res, err := ar.db.NewInsert().Model(adminModel).Exec(ctx)
	fmt.Printf("%v\n", res)
	if err != nil {
		// plan: log the error here then create your error and send back
		log.Println(err.Error())
		return false, nil
	}

	return true, nil
}
