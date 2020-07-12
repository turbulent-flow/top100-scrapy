PROJECTNAME=$(shell basename "${PWD}")

## migrate: migrate the database, if you want more fine-grained operations of the migration, you can look at the comments of the `MigrateDB()` method in the `pkg/automation/migration.go` file.
migrate: go-migrate

## migrate-test-db: migrate the test database, if you want more fine-grained operations of the migration, you can look at the comments of the `MigrateDB()` method in the `pkg/automation/migration.go` file.
migrate-test-db: go-migrate-test-db

## init: initialize the resouces of the project, e.g., initialize the database.
init: go-init-db

## populate: populate the canned data into the tables of the database. It's prepared for the workflow of the development.
populate: go-populate-seeds

## test: run all the test of the project.
test: go-migrate-test-db go-test

## compile: compile the most common instructions located the `./cmd` directory into the `./bin` directory.
compile: go-tidy go-compile-migration \
		 go-compile-producer-of-categories-insertion \
         go-compile-producer-of-products-insertion go-compile-subscriber

## compile-remaining-cmd: compile the remaining instructions except the above ones.
compile-remaining-cmd: go-compile-testdb-migration go-compile-db-initialization \
					   go-compile-test-db-initialization go-compile-seeds-population

go-migrate:
	@echo "  > Processing the migration..."
	@./bin/migrate up
	@echo "  > Done."

go-migrate-test-db:
	@echo "  > Processing the migration of the test database...."
	@./bin/migrate_test_db up
	@echo "  > Done."

go-init-db:
	@echo "  > Processing the initialization of the db..."
	@./bin/initialize_db
	@./bin/initialize_test_db
	@echo "  > Done."

go-populate-seeds:
	@echo "  > Populating the canned data into the db..."
	@./bin/populate_seeds
	@echo "  > Done."

go-test:
	@echo "  > Testing..."
	@gotest -v ./...

go-compile-migration:
	@echo "  > Compiling the instruction of migration..."
	@go build -o ./bin/ ./cmd/migrate
	@echo "  > Done."

go-compile-testdb-migration:
	@echo "  > Compiling the instruction of the migration of the test database..."
	@go build -o ./bin/ ./cmd/migrate_test_db
	@echo "  > Done."

go-compile-db-initialization:
	@echo "  > Compiling the insturction of the initialization of the db..."
	@go build -o ./bin ./cmd/initialize_db
	@echo "  > Done."

go-compile-test-db-initialization:
	@echo "  > Compiling the insturction of the initialization of the test db..."
	@go build -o ./bin ./cmd/initialize_test_db
	@echo "  > Done."

go-compile-producer-of-categories-insertion:
	@echo "  > Compiling the instruction of the insertion of the category queue..."
	@go build -o ./bin/ ./cmd/enqueue_categories_insertion
	@echo "  > Done."

go-compile-producer-of-products-insertion:
	@echo "  > Compiling the instruction of the insertion of the product queue..."
	@go build -o ./bin/ ./cmd/enqueue_products_insertion
	@echo "  > Done."

go-compile-subscriber: 
	@echo "  > Compiling the instruction of the workers which have subscribed the dedicated queue..."
	@go build -o ./bin/ ./cmd/consume
	@echo "  > Done."

go-compile-seeds-population:
	@echo "  > Compiling the instruction of the population of the seeds..."
	@go build -o ./bin/ ./cmd/populate_seeds
	@echo " > Done."

go-tidy:
	@echo "  > Tidying the dependencies from the \`go.mod\` file..."
	@go mod tidy
	@echo "  > Done."

.PHONY: help
help: Makefile
	@echo
	@echo "  Choose a command to run in "${PROJECTNAME}": "
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
