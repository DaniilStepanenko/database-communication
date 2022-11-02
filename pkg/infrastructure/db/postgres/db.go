package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() (db *sql.DB, err error) {
	dataSourceName := os.ExpandEnv("postgres://${PG_USER}:${PG_PASS}@${PG_HOST}:${PG_PORT}/${PG_DB}?sslmode=disable&search_path=${PG_SCHEMA}")
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return
}
