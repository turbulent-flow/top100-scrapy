PROJECTNAME=$(shell basename "${PWD}")

## migrate: database migrations which offer the entire circle that has reversibility.
migrate: go-migrate

## test: run all the test of the project.
test: go-migrate-test go-test

## compile: compile the instructions located the `./cmd` directory into the `./bin` directory.
compile: go-tidy go-compile-migration go-compile-producer-of-categories-insertion \
         go-compile-producer-of-products-insertion go-compile-subscriber

go-migrate:
	@echo "  > Processing the migration..."
	@./bin/migrate

go-migrate-test:
	@echo "  > Processing the migration of the test database...."
	@./bin/migrate_t

go-test:
	@echo "  > Testing..."
	@gotest -v ./...

go-compile-migration:
	@echo "  > Compiling the instruction of migration..."
	@go build -o ./bin/ ./cmd/migration/migrate.go
	@echo "  > Done."

go-compile-migration-test:
	@echo "  > Compiling the instruction of the migration of the test database..."
	@go build -o ./bin/ ./cmd/migration_test/migrate_t.go
	@echo "  > Done."

go-compile-producer-of-categories-insertion:
	@echo "  > Compiling the instruction of the insertion of the category queue..."
	@go build -o ./bin/ ./cmd/pub/categories_insertion/enqueue_categories_insertion.go
	@echo "  > Done."

go-compile-producer-of-products-insertion:
	@echo "  > Compiling the instruction of the insertion of the product queue..."
	@go build -o ./bin/ ./cmd/pub/products_inserttion/enqueue_products_insertion.go
	@echo "  > Done."

go-compile-subscriber: 
	@echo "  > Compiling the instruction of the workers which have subscribed the dedicated queue..."
	@go build -o ./bin/ ./cmd/sub/consume.go
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
