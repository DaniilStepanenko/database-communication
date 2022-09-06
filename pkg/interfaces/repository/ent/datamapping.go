package ent

import (
	"github.com/DaniilStepanenko/database-communication/ent"
	"github.com/DaniilStepanenko/database-communication/pkg/model"
)

func customerToDomainModel(c *ent.Customer) *model.Customer {
	return &model.Customer{
		Id:         c.ID,
		StoreID:    c.StoreID,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		Email:      c.Email,
		AddressID:  c.AddressID,
		ActiveBool: c.Activebool,
		CreateDate: c.CreateDate,
		LastUpdate: c.LastUpdate,
		Active:     c.Active,
	}
}

func customerFromDomainModel(c *model.Customer) *ent.Customer {
	return &ent.Customer{
		ID:         c.Id,
		StoreID:    c.StoreID,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		Email:      c.Email,
		AddressID:  c.AddressID,
		Activebool: c.ActiveBool,
		CreateDate: c.CreateDate,
		LastUpdate: c.LastUpdate,
		Active:     c.Active,
	}
}

func customersToDomainModel(rows []*ent.Customer) (result []*model.Customer) {
	result = make([]*model.Customer, len(rows))
	for i := range rows {
		result[i] = customerToDomainModel(rows[i])
	}

	return
}
