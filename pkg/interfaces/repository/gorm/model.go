package gorm

import (
	"time"

	"github.com/DaniilStepanenko/database-communication/pkg/model"
)

type gormCustomer struct {
	ID         int `gorm:"column:customer_id"`
	StoreId    int
	FirstName  string
	LastName   string
	Email      string
	AddressId  int
	ActiveBool bool `gorm:"column:activebool"`
	CreateDate time.Time
	LastUpdate time.Time
	Active     int
}

func (gormCustomer) TableName() string {
	return "dvds.customer"
}

func (c gormCustomer) toDomain() *model.Customer {
	return &model.Customer{
		Id:         c.ID,
		StoreID:    c.StoreId,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		Email:      c.Email,
		AddressID:  c.AddressId,
		ActiveBool: c.ActiveBool,
		CreateDate: c.CreateDate,
		LastUpdate: c.LastUpdate,
		Active:     c.Active,
	}
}

func toGORMCustomer(c *model.Customer) *gormCustomer {
	return &gormCustomer{
		ID:         c.Id,
		StoreId:    c.StoreID,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		Email:      c.Email,
		AddressId:  c.AddressID,
		ActiveBool: c.ActiveBool,
		CreateDate: c.CreateDate,
		LastUpdate: c.LastUpdate,
		Active:     c.Active,
	}
}

type gormPayment struct {
}

func (gormPayment) TableName() string {
	return "dvds.payment"
}
