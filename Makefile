PROJECTNAME=$(shell basename "${PWD}")

## migrate: database migrations which offer the entire circle that has reversibility.
migrate: go-migrate

## init: initialize the resouces of the project, e.g., initialize the database.
init: go-init-db

## test: run all the test of the project.
test: go-migrate-test go-test

## compile: compile the instructions located the `./cmd` directory into the `./bin` directory.
compile: go-tidy go-compile-migration go-compile-testdb-migration \
		 go-compile-producer-of-categories-insertion \
         go-compile-producer-of-products-insertion go-compile-subscriber

go-migrate:
	@echo "  > Processing the migration..."
	@./bin/migrate up

go-migrate-test:
	@echo "  > Processing the migration of the test database...."
	@./bin/migrate_test_db

go-init-db:
	@echo "  > Processing the initialization of the db..."
	@./bin/initialize_db
	@./bin/initialize_test_db
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
