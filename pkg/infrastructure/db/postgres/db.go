package postgres

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	dataSourceName := os.ExpandEnv("postgres://${PG_USER}:${PG_PASS}@${PG_HOST}:${PG_PORT}/${PG_DB}?sslmode=disable&search_path=${PG_SCHEMA}")
	return sql.Open("postgres", dataSourceName)
}
