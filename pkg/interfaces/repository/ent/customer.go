package ent

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	entsql "entgo.io/ent/dialect/sql"

	"github.com/DaniilStepanenko/database-communication/ent"
	"github.com/DaniilStepanenko/database-communication/ent/customer"
	"github.com/DaniilStepanenko/database-communication/ent/payment"
	"github.com/DaniilStepanenko/database-communication/ent/predicate"
	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository"
	"github.com/DaniilStepanenko/database-communication/pkg/model"
)

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	drv := entsql.OpenDB("postgres", db)
	return &CustomerRepository{ent.NewClient(ent.Driver(drv)), true}
}

type CustomerRepository struct {
	client *ent.Client
	debug  bool
}

func (c *CustomerRepository) Create(ctx context.Context, model *model.Customer) (int, error) {
	created, err := c.
		getClient().
		Customer.
		Create().
		SetStoreID(model.StoreID).
		SetFirstName(model.FirstName).
		SetLastUpdate(model.LastUpdate).
		SetEmail(model.Email).
		SetAddressID(model.AddressID).
		SetActivebool(model.ActiveBool).
		SetCreateDate(model.CreateDate).
		SetLastUpdate(model.LastUpdate).
		SetActive(model.Active).
		Save(ctx)

	if err != nil {
		return 0, fmt.Errorf("ent save customer: %w", err)
	}

	return created.ID, nil
}

func (c *CustomerRepository) Read(ctx context.Context, id int) (*model.Customer, error) {
	row, err := c.
		getClient().
		Customer.
		Get(ctx, id)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, repository.NoData
	}
	if err != nil {
		return nil, fmt.Errorf("get by id = %d: %w", id, err)
	}

	return customerToDomainModel(row), nil
}

func (c *CustomerRepository) Update(ctx context.Context, model *model.Customer) error {
	_, err := c.
		getClient().
		Customer.
		UpdateOne(customerFromDomainModel(model)).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("update by id = %d: %w", model.Id, err)
	}

	return nil
}

func (c *CustomerRepository) Delete(ctx context.Context, id int) error {
	err := c.
		getClient().
		Customer.
		DeleteOne(&ent.Customer{ID: id}).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("delete by id = %d: %w", id, err)
	}

	return nil
}

func (c *CustomerRepository) List(ctx context.Context, listOptions *repository.ListOptions, criteria *repository.CustomerCriteria) ([]*model.Customer, error) {
	stmt := c.
		getClient().
		Customer.
		Query()

	var conditions []predicate.Customer
	if criteria != nil && criteria.StoreID != nil {
		conditions = append(conditions, customer.StoreID(*criteria.StoreID))
	}

	if criteria != nil && criteria.Query != nil {
		or := customer.Or(
			customer.FirstNameContains(*criteria.Query),
			customer.LastNameContains(*criteria.Query),
			customer.EmailContains(*criteria.Query),
		)
		conditions = append(conditions, or)
	}
	stmt.Where(conditions...)

	if listOptions != nil {
		for _, sort := range listOptions.Sort {
			if sort.Direction == repository.SortDirectionDESC {
				stmt.Order(ent.Desc(sort.Property))
			} else {
				stmt.Order(ent.Asc(sort.Property))
			}
		}
	}

	if listOptions != nil && listOptions.Offset > 0 {
		stmt.Offset(listOptions.Offset)
	}

	if listOptions != nil && listOptions.Limit > 0 {
		stmt.Limit(listOptions.Limit)
	}

	rows, err := stmt.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("list query failed: %w", err)
	}

	return customersToDomainModel(rows), nil
}

func (c *CustomerRepository) FindActiveCustomers(ctx context.Context, lastPaymentNotEarlier time.Time) ([]*model.Customer, error) {
	const fieldMaxPaymentDate = "max_payment_date"

	paymentTable := entsql.Table(payment.Table)
	paymentSubQuery := entsql.
		Select(
			payment.FieldCustomerID,
			entsql.As(entsql.Max(payment.FieldPaymentDate), fieldMaxPaymentDate),
		).
		From(paymentTable).
		GroupBy(
			payment.FieldCustomerID,
		).
		As(payment.Table)

	rows, err := c.
		getClient().
		Customer.
		Query().
		Where(
			func(s *entsql.Selector) {
				s.Join(paymentSubQuery).On(s.C(customer.FieldID), paymentTable.C(payment.FieldCustomerID))

				s.Where(entsql.GTE(paymentTable.C(fieldMaxPaymentDate), lastPaymentNotEarlier))
			},
		).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return customersToDomainModel(rows), nil
}

func (c *CustomerRepository) getClient() *ent.Client {
	if c.debug {
		return c.client.Debug()
	}

	return c.client
}

var _ repository.CustomerRepository = (*CustomerRepository)(nil)
