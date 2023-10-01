include .env

gen-sql:
	@mkdir -p ./internal/generated/db
	@echo "Generate sql..."
	@sqlc generate
# Migrations
migrations-create:
	mkdir -p migrations
	@read -p "Name of the migration: " migration \
	&& echo "Create migrations $$migration at postgres ${TILE_GENERATOR_POSTGRES_MASTER_DSN}" \
	&& goose -dir migrations postgres "${TILE_GENERATOR_POSTGRES_MASTER_DSN}" create $$migration sql

migrations-up:
	@goose -dir migrations postgres "${TILE_GENERATOR_POSTGRES_GOOSE_DSN}" up

migrations-down:
	@goose -dir migrations postgres "${TILE_GENERATOR_POSTGRES_GOOSE_DSN}" down

docker-run:
	@echo "Running docker-compose..."
	@docker compose -p tile-generator up -d

docker-stop:
	@echo "Stopping docker-compose..."
	@docker compose -p tile-generator down

tidy:
	@go mod tidy

vendor:
	@go mod vendor

clean-vendor:
	@echo "Removing vendor directory..."
	@rm -rf ./vendor

deps: tidy vendor

gen-sql:
	@mkdir -p ./internal/generated/db
	@rm ./internal/generated/db/*.*
	@echo "Generate sql..."
	@sqlc generate

run:
	go run -race cmd/cli/main.go