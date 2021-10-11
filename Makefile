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