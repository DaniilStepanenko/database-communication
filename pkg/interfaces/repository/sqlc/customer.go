package sqlc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository"
	"github.com/DaniilStepanenko/database-communication/pkg/model"
)

func NewCustomerRepository(db DBTX) *CustomerRepository {
	return &CustomerRepository{New(db)}
}

type CustomerRepository struct {
	query *Queries
}

func (c *CustomerRepository) Create(ctx context.Context, customer *model.Customer) (int, error) {
	id, err := c.query.CreateCustomer(ctx, CreateCustomerParams{
		StoreID:    int16(customer.StoreID),
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		Email:      sql.NullString{Valid: customer.Email != "", String: customer.Email},
		AddressID:  int16(customer.AddressID),
		Activebool: customer.ActiveBool,
		CreateDate: customer.CreateDate,
		LastUpdate: sql.NullTime{Valid: !customer.LastUpdate.IsZero(), Time: customer.LastUpdate},
		Active:     sql.NullInt32{Valid: customer.Active != 0, Int32: int32(customer.Active)},
	})
	if err != nil {
		return 0, fmt.Errorf("query create customer: %w", err)
	}

	return int(id), nil
}

func (c *CustomerRepository) Read(ctx context.Context, id int) (*model.Customer, error) {
	row, err := c.query.ReadCustomer(ctx, int32(id))
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, repository.NoData
	}

	customer := &model.Customer{
		Id:         id,
		StoreID:    int(row.StoreID),
		FirstName:  row.FirstName,
		LastName:   row.LastName,
		Email:      row.Email.String,
		AddressID:  int(row.AddressID),
		ActiveBool: row.Activebool,
		CreateDate: row.CreateDate,
		LastUpdate: row.LastUpdate.Time,
		Active:     int(row.Active.Int32),
	}
	return customer, nil
}

func (c *CustomerRepository) Update(ctx context.Context, customer *model.Customer) error {
	params := UpdateCustomerParams{
		CustomerID: int32(customer.Id),
		StoreID:    int16(customer.StoreID),
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		Email:      sql.NullString{Valid: customer.Email != "", String: customer.Email},
		AddressID:  int16(customer.AddressID),
		Activebool: customer.ActiveBool,
		CreateDate: customer.CreateDate,
		LastUpdate: sql.NullTime{Valid: !customer.LastUpdate.IsZero(), Time: customer.LastUpdate},
		Active:     sql.NullInt32{Valid: customer.Active != 0, Int32: int32(customer.Active)},
	}

	err := c.query.UpdateCustomer(ctx, params)
	if err != nil {
		return fmt.Errorf("query update customer: %w", err)
	}

	return nil
}

func (c *CustomerRepository) Delete(ctx context.Context, id int) error {
	err := c.query.DeleteCustomer(ctx, int32(id))
	if err != nil {
		return fmt.Errorf("query delete %d: %w", id, err)
	}

	return nil
}

func (c *CustomerRepository) FindActiveCustomers(ctx context.Context, lastPaymentNotEarlier time.Time) ([]*model.Customer, error) {
	rows, err := c.query.FindActiveCustomers(ctx, lastPaymentNotEarlier)
	if err != nil {
		return nil, fmt.Errorf("query active customers: %w", err)
	}

	customers := make([]*model.Customer, len(rows))
	for i, row := range rows {
		customers[i] = &model.Customer{
			Id:         int(row.CustomerID),
			StoreID:    int(row.StoreID),
			FirstName:  row.FirstName,
			LastName:   row.LastName,
			Email:      row.Email.String,
			AddressID:  int(row.AddressID),
			ActiveBool: row.Activebool,
			CreateDate: row.CreateDate,
			LastUpdate: row.LastUpdate.Time,
			Active:     int(row.Active.Int32),
		}
	}
	return customers, nil
}

func (c *CustomerRepository) List(ctx context.Context, list *repository.ListOptions, criteria *repository.CustomerCriteria) ([]*model.Customer, error) {
	panic("implement me")
}

var _ repository.CustomerRepository = (*CustomerRepository)(nil)
