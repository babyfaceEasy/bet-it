package repositories_test

import (
	"context"
	"elivate9ja-go/app/admin/entities"
	"elivate9ja-go/utils"
)

func (s *PostgreSQLTestSuite) seedAdmins(admins []entities.AdminEntity) {
	db := utils.GetDBConnection()
	db.NewInsert().Model(&admins).Column("id", "name", "email", "created_at", "updated_at").Exec(context.Background())
}
