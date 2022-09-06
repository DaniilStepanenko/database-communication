PG_HOST := localhost
PG_PORT := 5432
PG_USER := postgres
PG_PASS := postgres
PG_DB := postgres
PG_SCHEMA := dvds

dev-db:
	docker-compose -f docker-compose.dev.yml up -d db adminer

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

playground:
	echo docker-compose -f docker-compose.dev.yml build app && \
 		docker-compose -f docker-compose.dev.yml run --rm --entrypoint=$(ENTRYPOINT) app

jet-playground:
	make playground ENTRYPOINT="/bin/jet"

gorm-playground:
	make playground ENTRYPOINT="/bin/gorm"

sqlc-playground:
	make playground ENTRYPOINT="/bin/sqlc"

ent-playground:
	make playground ENTRYPOINT="/bin/ent"


pre-commit: dependency generate

generate:
	go generate -x ./...

dependency:
	go mod tidy && go mod vendor
