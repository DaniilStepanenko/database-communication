package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository/sqlc"

	"github.com/DaniilStepanenko/database-communication/pkg/infrastructure/db/postgres"
)

func main() {
	dataSourceName := os.ExpandEnv("postgres://${PG_USER}:${PG_PASS}@${PG_HOST}:${PG_PORT}/${PG_DB}?sslmode=disable")
	db, err := postgres.NewDB(dataSourceName)
	check(err)

	repo := sqlc.NewCustomerRepository(db)

	ctx := context.Background()
	t := time.Date(2007, 05, 14, 13, 44, 29, 0, time.UTC)
	customers, err := repo.FindActiveCustomers(ctx, t)
	check(err)

	for i := range customers {
		fmt.Printf("%+v\n", customers[i])
	}
	fmt.Printf("Found %d customers\n", len(customers))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
