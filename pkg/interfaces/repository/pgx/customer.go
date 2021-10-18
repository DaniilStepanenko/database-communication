package pgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository"
	"github.com/DaniilStepanenko/database-communication/pkg/model"
)

func NewCustomerRepository(conn *pgx.Conn) *CustomerRepository {
	return &CustomerRepository{conn}
}

type CustomerRepository struct {
	conn *pgx.Conn
}

func (c *CustomerRepository) List(ctx context.Context, list *repository.ListOptions, criteria *repository.CustomerCriteria) ([]*model.Customer, error) {
	query, args := listCustomersQuery(list, criteria)

	rows, err := c.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query %s: %w", query, err)
	}
	defer rows.Close()

	var res []*model.Customer
	for rows.Next() {
		customer := &model.Customer{}

		err = rows.Scan(
			&customer.Id,
			&customer.StoreID,
			&customer.FirstName,
			&customer.LastName,
			&customer.Email,
			&customer.AddressID,
			&customer.ActiveBool,
			&customer.CreateDate,
			&customer.LastUpdate,
			&customer.Active,
		)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		res = append(res, customer)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return res, nil
}

func listCustomersQuery(listOptions *repository.ListOptions, criteria *repository.CustomerCriteria) (string, []interface{}) {
	query := "SELECT * FROM dvds.customer WHERE TRUE"
	args := make([]interface{}, 0)

	if criteria != nil && criteria.StoreID != nil {
		args = append(args, *criteria.StoreID)
		query += fmt.Sprintf(" AND customer.store_id = $%d", len(args))
	}

	if criteria != nil && criteria.Query != nil {
		args = append(args, *criteria.Query)
		query += fmt.Sprintf(" AND customer.first_name LIKE '%%' || $%d || '%%'", len(args))

		args = append(args, *criteria.Query)
		query += fmt.Sprintf(" AND customer.last_name LIKE '%%' || $%d || '%%'", len(args))

		args = append(args, *criteria.Query)
		query += fmt.Sprintf(" AND customer.email LIKE '%%' || $%d || '%%'", len(args))

	}

	if listOptions != nil {
		for _, sort := range listOptions.Sort {
			args = append(args, sort.Property)
			query += fmt.Sprintf(" ORDER BY $%d", len(args))

			if sort.Direction == repository.SortDirectionDESC {
				query += " DESC"
			} else {
				query += " ASC"
			}
		}
	}

	if listOptions != nil && listOptions.Offset > 0 {
		args = append(args, listOptions.Offset)
		query += fmt.Sprintf(" OFFSET $%d", len(args))
	}

	if listOptions != nil && listOptions.Limit > 0 {
		args = append(args, listOptions.Limit)
		query += fmt.Sprintf(" OFFSET $%d", len(args))
	}

	return query, args
}

// TODO: remove next line
var ListCustomersQuery = listCustomersQuery
