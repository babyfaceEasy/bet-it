.PHONY:

APP_NAME = Elivate9ja
MIGRATIONS_FOLDER = $(PWD)/database/migrations 
DATABASE_URL = postgresql://postgres:p@ssw0rd1@127.0.0.1:5432/elivate9jago_db?sslmode=disable
DATABASE_TEST_URL = postgresql://postgres:p@ssw0rd1@127.0.0.1:5432/elivate9jago_test_db?sslmode=disable

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

migrate.version:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" version

migrate.create:
	migrate create -ext sql -dir $(MIGRATIONS_FOLDER) $(migration_name)

test.migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_TEST_URL)" up

test.migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_TEST_URL)" down