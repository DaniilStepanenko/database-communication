package jet

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository"
	"github.com/DaniilStepanenko/database-communication/pkg/model"
)

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db}
}

type CustomerRepository struct {
	db *sql.DB
}

func (c *CustomerRepository) Create(ctx context.Context, customer *model.Customer) (id int, err error) {
	stmt := createCustomerQuery(customer)
	query, args := stmt.Sql()

	row := c.db.QueryRowContext(ctx, query, args...)

	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return
}

func (c *CustomerRepository) Read(ctx context.Context, id int) (*model.Customer, error) {
	stmt := readCustomerQuery(id)
	query, args := stmt.Sql()
	row := c.db.QueryRowContext(ctx, query, args...)

	customer := &model.Customer{Id: id}
	err := row.Scan(
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
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, repository.NoData
	}
	if err != nil {
		return nil, fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return customer, nil
}

func (c CustomerRepository) Update(ctx context.Context, customer *model.Customer) error {
	stmt := updateCustomerQuery(customer)
	query, args := stmt.Sql()
	_, err := c.db.ExecContext(ctx, query, args...)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return repository.NoData
	}
	if err != nil {
		return fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return nil
}

func (c CustomerRepository) Delete(ctx context.Context, id int) error {
	stmt := deleteCustomerQuery(id)
	query, args := stmt.Sql()
	_, err := c.db.ExecContext(ctx, query, args...)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return repository.NoData
	}
	if err != nil {
		return fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return nil
}

func (c *CustomerRepository) FindActiveCustomers(ctx context.Context, lastPaymentNotEarlier time.Time) ([]*model.Customer, error) {
	stmt := findActiveCustomersQuery(lastPaymentNotEarlier)
	query, args := stmt.Sql()

	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query %s: %w", stmt.DebugSql(), err)
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

func (c *CustomerRepository) List(ctx context.Context, list *repository.ListOptions, criteria *repository.CustomerCriteria) ([]*model.Customer, error) {
	stmt := listCustomersQuery(list, criteria)
	query, args := stmt.Sql()

	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query %s: %w", stmt.DebugSql(), err)
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

var _ repository.CustomerRepository = (*CustomerRepository)(nil)
