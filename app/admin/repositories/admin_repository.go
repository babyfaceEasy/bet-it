package repositories

import (
	"context"
	"elivate9ja-go/app/admin/entities"
	"elivate9ja-go/app/admin/models"
	"fmt"
	"log"

	"github.com/google/uuid"
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

func (ar *AdminRepository) GetAdmin(id uuid.UUID) (entities.AdminEntity, error) {
	adminModel := new(models.Admin)
	ctx := context.Background()
	err := ar.db.NewSelect().Model(adminModel).Where("id = ?", id).Scan(ctx)
	if err != nil {
		// Log error and then create new error to send back
		return entities.AdminEntity{}, err
	}

	return entities.AdminEntity{
		ID: adminModel.ID,
		Name: adminModel.FullName,
		Email: adminModel.Email,
		SuperAdmin: adminModel.SuperAdmin,
		CreatedAt: adminModel.CreatedAt,
		UpdatedAt: adminModel.UpdatedAt,
	}, nil
}

func (ar *AdminRepository) GetAdmins() ([]entities.AdminEntity, error) {
	var admins []entities.AdminEntity
	var modelAdmins []models.Admin
	ctx := context.Background()
	_, err := ar.db.NewSelect().Model(&modelAdmins).ScanAndCount(ctx)
	if err != nil {
		// Log error then create new error to send back
		return []entities.AdminEntity{}, err 
	}

	for _, adminModel := range modelAdmins {
		adminEntity := entities.AdminEntity{
			ID: adminModel.ID,
			Name: adminModel.FullName,
			Email: adminModel.Email,
			SuperAdmin: adminModel.SuperAdmin,
			CreatedAt: adminModel.CreatedAt,
			UpdatedAt: adminModel.UpdatedAt,
		}
		admins = append(admins, adminEntity)
	}

	return admins, nil

}

func (ar *AdminRepository) VerifyAdmin(email, password string) (bool, error) {
	ctx := context.Background()
	exists, err := ar.db.NewSelect().
		Model((*models.Admin)(nil)).
		Where("email = ?", email).
		Where("password = ?", password).
		Exists(ctx)
	if err != nil {
		// log error, inspect and create new error
		return false, err
	}

	if !exists {
		return false, nil
	}

	return true, nil
}

func (ar *AdminRepository) GetAdminByEmail(email string) (entities.AdminEntity, error) {
	adminModel := new(models.Admin)
	ctx := context.Background()
	err := ar.db.NewSelect().Model(adminModel).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		//log error and then create more friendly error to return
		return entities.AdminEntity{}, err
	}

	return entities.AdminEntity{
		ID: adminModel.ID,
		Name: adminModel.FullName,
		Email: adminModel.Email,
		SuperAdmin: adminModel.SuperAdmin,
		CreatedAt: adminModel.CreatedAt,
		UpdatedAt: adminModel.UpdatedAt,
	}, nil
}
