package main

import (
	"context"
	"fmt"

	"github.com/DaniilStepanenko/database-communication/pkg/infrastructure/db/postgres"
	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository"
	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository/ent"
)

func main() {
	db, err := postgres.NewDB()
	check(err)

	repo := ent.NewCustomerRepository(db)
	ctx := context.Background()
	result, err := repo.List(
		ctx,
		&repository.ListOptions{
			Sort: []repository.SortOrder{{
				Property: "customer_id",
			}},
		},
		&repository.CustomerCriteria{
			StoreID: ptr(2),
			Query:   ptr("Alex"),
		},
	)
	check(err)

	for i := range result {
		fmt.Printf("%+v\n", result[i])
	}
	fmt.Printf("Found %d customers\n", len(result))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ptr[T any](val T) *T {
	return &val
}
