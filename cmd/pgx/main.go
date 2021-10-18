package main

import (
	"fmt"

	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository"
	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository/pgx"
)

func main() {
	fmt.Println(pgx.ListCustomersQuery(
		&repository.ListOptions{
			Sort: []repository.SortOrder{{
				Property: "customer_id",
			}},
		},
		&repository.CustomerCriteria{
			StoreID: intPtr(2),
			Query:   stringPtr("Alex"),
		},
	))
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
