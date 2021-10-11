package model

import "time"

type Customer struct {
	Id         int
	StoreID    int
	FirstName  string
	LastName   string
	Email      string
	AddressID  int
	ActiveBool bool
	CreateDate time.Time
	LastUpdate time.Time
	Active     int
}
