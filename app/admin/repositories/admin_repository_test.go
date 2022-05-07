package repositories

import (
	"elivate9ja-go/app/admin/entities"
	"elivate9ja-go/utils"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)
type IntegrationTestSuite struct {
	suite.Suite
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, &IntegrationTestSuite{})
}

func (its *IntegrationTestSuite) SetupSuite() {
	
}

func TestSaveAdmin(t *testing.T) {
	requires := require.New(t)
	db := utils.GetTestDBConnection()
	defer db.Close()

	newAdmin := entities.AdminEntity{
		ID:         uuid.New(),
		Name:       "TestAdmin",
		Email:      "test@example.com",
		SuperAdmin: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	repo, err := NewAdminRepository(db)
	if err != nil {
		panic(err)
	}
	result, err := repo.SaveAdmin(newAdmin)

	requires.Nil(err)
	requires.Equal(true, result)
}
