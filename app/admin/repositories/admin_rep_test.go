package repositories_test

import (
	"elivate9ja-go/app/admin/entities"
	"elivate9ja-go/app/admin/repositories"
	"elivate9ja-go/utils"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func (s *PostgreSQLTestSuite) TestCountAdmins() {
	admins := []entities.AdminEntity{
		{
			ID:         uuid.New(),
			Name:       "Admin One",
			Email:      "admin@example.com",
			SuperAdmin: true,
			CreatedAt:  time.Now().UTC(),
			UpdatedAt:  time.Now().UTC(),
		},
	}
	s.seedAdmins(admins)

	// get bun connection
	dbConn := utils.GetTestDBConnection()
	repo, err := repositories.NewAdminRepository(dbConn)
	if err != nil {
		require.FailNow(s.T(), err.Error())
	}
	result := repo.CountAdmins()
	//require.NoError(s.T(), err)
	require.Equal(s.T(), len(admins), result)
}
