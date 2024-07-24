include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

build:
	@go build -o ./bin/main ./cmd/main.go 

run: build
	@./bin/main

test:
	@go test -v ./...

migration:
	migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down