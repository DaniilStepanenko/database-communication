package repository

import (
	"context"
	"time"

	"github.com/DaniilStepanenko/database-communication/pkg/model"
)

type CustomerCriteria struct {
	StoreID *int
	Query   *string
}

type CustomerRepository interface {
	Create(ctx context.Context, customer *model.Customer) (id int, err error)
	Read(ctx context.Context, id int) (*model.Customer, error)
	Update(ctx context.Context, customer *model.Customer) error
	Delete(ctx context.Context, id int) error

	List(ctx context.Context, list *ListOptions, criteria *CustomerCriteria) ([]*model.Customer, error)
	FindActiveCustomers(ctx context.Context, lastPaymentNotLater time.Time) ([]*model.Customer, error)
}
