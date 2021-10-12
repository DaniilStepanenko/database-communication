package gorm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"

	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository"
	"github.com/DaniilStepanenko/database-communication/pkg/model"
)

func NewCustomerRepository(db *sql.DB) (*CustomerRepository, error) {
	pgConf := postgres.Config{Conn: db}
	gormConf := &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	pg := postgres.New(pgConf)

	gormDb, err := gorm.Open(pg, gormConf)
	if err != nil {
		return nil, fmt.Errorf("gorm: open: %w", err)
	}

	return &CustomerRepository{gormDb}, nil
}

type CustomerRepository struct {
	db *gorm.DB
}

func (c *CustomerRepository) Create(ctx context.Context, customer *model.Customer) (int, error) {
	gormCustomer := toGORMCustomer(customer)

	err := c.db.
		WithContext(ctx).
		Create(gormCustomer).
		Error

	if err != nil {
		return 0, fmt.Errorf("gorm: create: %w", err)
	}

	return gormCustomer.ID, nil
}

func (c *CustomerRepository) Read(ctx context.Context, id int) (*model.Customer, error) {
	customer := &gormCustomer{ID: id}

	err := c.db.
		WithContext(ctx).
		First(customer).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repository.NoData
	}

	if err != nil {
		return nil, fmt.Errorf("gorm: read: %w", err)
	}

	return customer.toDomain(), nil
}

func (c CustomerRepository) Update(ctx context.Context, customer *model.Customer) error {
	err := c.db.
		WithContext(ctx).
		Updates(toGORMCustomer(customer)).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NoData
	}

	if err != nil {
		return fmt.Errorf("gorm: update: %w", err)
	}

	return nil
}

func (c CustomerRepository) Delete(ctx context.Context, id int) error {
	err := c.db.
		WithContext(ctx).
		Delete(gormCustomer{ID: id}).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.NoData
	}

	if err != nil {
		return fmt.Errorf("gorm: update: %w", err)
	}

	return nil
}

func (c *CustomerRepository) FindActiveCustomers(ctx context.Context, lastPaymentNotEarlier time.Time) (result []*model.Customer, err error) {
	payments := c.db.
		Select([]string{
			"customer_id",
			"MAX(payment_date) AS \"max_payment_date\"",
		}).
		Model(&gormPayment{}).
		Group("customer_id")

	err = c.db.
		WithContext(ctx).
		Model(&gormCustomer{}).
		Joins("INNER JOIN (?) AS payment ON customer.customer_id = payment.customer_id", payments).
		Where("max_payment_date >= ?::timestamp without time zone", lastPaymentNotEarlier).
		Scan(&result).
		Error

	if err != nil {
		return nil, fmt.Errorf("gorm: find: %w", err)
	}

	return result, nil
}

func (c *CustomerRepository) List(ctx context.Context, listOptions *repository.ListOptions, criteria *repository.CustomerCriteria) (customers []*model.Customer, err error) {
	var result []gormCustomer

	stmt := c.db.WithContext(ctx)

	var conditions = clause.AndConditions{}
	if criteria != nil && criteria.StoreID != nil {
		conditions.Exprs = append(
			conditions.Exprs,
			clause.Eq{
				Column: "store_id",
				Value:  *criteria.StoreID,
			},
		)
	}

	if criteria != nil && criteria.Query != nil {
		query := strings.Join([]string{"%", *criteria.Query, "%"}, "")
		or := clause.Or(
			clause.Like{
				Column: "first_name",
				Value:  query,
			},
			clause.Like{
				Column: "last_name",
				Value:  query,
			},
			clause.Like{
				Column: "email",
				Value:  query,
			})

		conditions.Exprs = append(conditions.Exprs, or)
	}

	if len(conditions.Exprs) > 0 {
		stmt = stmt.Where(conditions)
	}

	if listOptions != nil {
		for _, sort := range listOptions.Sort {
			orderByColumn := clause.OrderByColumn{
				Column: clause.Column{Name: sort.Property},
				Desc:   sort.Direction != repository.SortDirectionASC,
			}
			stmt.Order(orderByColumn)
		}
	}

	if listOptions != nil && listOptions.Offset > 0 {
		stmt.Offset(listOptions.Offset)
	}

	if listOptions != nil && listOptions.Limit > 0 {
		stmt.Limit(listOptions.Limit)
	}

	err = stmt.Find(&result).Error
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	customers = make([]*model.Customer, len(result))
	for i := range result {
		customers[i] = result[i].toDomain()
	}
	return customers, nil
}

var _ repository.CustomerRepository = (*CustomerRepository)(nil)
