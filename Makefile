PG_HOST := localhost
PG_PORT := 5432
PG_USER := postgres
PG_PASS := postgres
PG_DB := postgres
PG_SCHEMA := dvds

dev-db:
	docker-compose -f docker-compose.dev.yml up -d

jet-pg:
	jet -source=PostgreSQL \
		-host=$(PG_HOST) \
		-port=$(PG_PORT) \
		-user=$(PG_USER) \
		-password=$(PG_PASS) \
		-dbname=$(PG_DB) \
		-schema=$(PG_SCHEMA) \
		-path=./pkg/interfaces/repository/jet/internal

sqlc:
	sqlc generate

jet-playground:
	PG_HOST=$(PG_HOST) PG_PORT=$(PG_PORT) PG_USER=$(PG_USER) PG_PASS=$(PG_PASS) PG_DB=$(PG_DB) PG_SCHEMA=$(PG_SCHEMA) go run ./cmd/jet

gorm-playground:
	PG_HOST=$(PG_HOST) PG_PORT=$(PG_PORT) PG_USER=$(PG_USER) PG_PASS=$(PG_PASS) PG_DB=$(PG_DB) PG_SCHEMA=$(PG_SCHEMA) go run ./cmd/gorm

sqlc-playground:
	PG_HOST=$(PG_HOST) PG_PORT=$(PG_PORT) PG_USER=$(PG_USER) PG_PASS=$(PG_PASS) PG_DB=$(PG_DB) PG_SCHEMA=$(PG_SCHEMA) go run ./cmd/sqlc

ent-playground:
	PG_HOST=$(PG_HOST) PG_PORT=$(PG_PORT) PG_USER=$(PG_USER) PG_PASS=$(PG_PASS) PG_DB=$(PG_DB) PG_SCHEMA=$(PG_SCHEMA) go run ./cmd/ent

pre-commit: dependency generate

generate:
	go generate -x ./...

dependency:
	go mod tidy && go mod vendor
