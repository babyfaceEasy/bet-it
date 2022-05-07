package repositories_test

//https://faun.pub/how-to-test-database-repository-in-golang-771b59c8084e

import (
	"database/sql"
	"log"
	"path"
	"runtime"
	"testing"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PostgreSQLTestSuite struct {
	suite.Suite
	db *sql.DB
	m  *migrate.Migrate
}

const databaseTestName = "elivate9jago_test_db"

//Run before all test start
func (s *PostgreSQLTestSuite) SetUpSuite() {
	var err error

	s.db, err = sql.Open("postgres", "postgresql://postgres:p@ssw0rd1@127.0.0.1:5432/elivate9jago_test_db?sslmode=disable")
	require.NoError(s.T(), err)

	// Migrate up to create all tables
	// Find migration path
	_, filename, _, _ := runtime.Caller(0)
	migrationPath := "file://" + path.Join(path.Dir(filename), "migrations")
	// db driver
	driver, err := postgres.WithInstance(s.db, &postgres.Config{DatabaseName: databaseTestName})
	require.NoError(s.T(), err)

	// Init migrate
	s.m, err = migrate.NewWithDatabaseInstance(migrationPath, databaseTestName, driver)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), s.m)
	require.NoError(s.T(), s.m.Up())
}

// Rn after each tet finished
func (s *PostgreSQLTestSuite) TearDownTest() {
	// truncate all tables. empty all tables
	query := `SELECT TABLE_NAME FROM information_schema.tables WHERE tables_schema='` + databaseTestName + `'`
	rows, err := s.db.Query(query)
	require.NoError(s.T(), err)

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Printf("errorscanning tablename %s", err)
		}
		if tableName == "schema_migrations" {
			continue
		}
		queryTruncate := "TRUNCATE TABLE " + tableName
		_, err = s.db.Exec(queryTruncate)
		require.NoError(s.T(), err)
	}

	err = rows.Close()
	require.NoError(s.T(), err)
}

//Run after all test has finished
func (s *PostgreSQLTestSuite) TearDownSuite() {
	//Migrate down (drop all tables)
	require.NoError(s.T(), s.m.Down())
}

func (s *PostgreSQLTestSuite) TestPostgreSQLTestSuite(t *testing.T) {
	// Skip the test when using "-short" flag
	if testing.Short() {
		t.Skip("Skipping long-running test")
	}

	suite.Run(t, new(PostgreSQLTestSuite))
}
