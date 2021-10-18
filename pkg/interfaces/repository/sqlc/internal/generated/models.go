// Code generated by sqlc. DO NOT EDIT.

package generated

import (
	"database/sql"
	"time"
)

type DvdsCustomer struct {
	CustomerID int32
	StoreID    int16
	FirstName  string
	LastName   string
	Email      sql.NullString
	AddressID  int16
	Activebool bool
	CreateDate time.Time
	LastUpdate sql.NullTime
	Active     sql.NullInt32
}

type DvdsPayment struct {
	PaymentID   int32
	CustomerID  int16
	StaffID     int16
	RentalID    int32
	Amount      string
	PaymentDate time.Time
}