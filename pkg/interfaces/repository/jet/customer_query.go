package jet

import (
	"time"

	. "github.com/go-jet/jet/v2/postgres"

	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository"
	"github.com/DaniilStepanenko/database-communication/pkg/interfaces/repository/jet/internal/postgres/dvds/table"
	"github.com/DaniilStepanenko/database-communication/pkg/model"
)

func createCustomerQuery(c *model.Customer) InsertStatement {
	return table.Customer.INSERT(
		table.Customer.StoreID,
		table.Customer.FirstName,
		table.Customer.LastName,
		table.Customer.Email,
		table.Customer.AddressID,
		table.Customer.Activebool,
		table.Customer.CreateDate,
		table.Customer.LastUpdate,
		table.Customer.Active,
	).VALUES(
		c.StoreID,
		c.FirstName,
		c.LastName,
		c.Email,
		c.AddressID,
		c.ActiveBool,
		c.CreateDate,
		c.LastUpdate,
		c.Active,
	).RETURNING(
		table.Customer.CustomerID,
	)
}

func readCustomerQuery(id int) SelectStatement {
	return SELECT(
		table.Customer.StoreID,
		table.Customer.FirstName,
		table.Customer.LastName,
		table.Customer.Email,
		table.Customer.AddressID,
		table.Customer.Activebool,
		table.Customer.CreateDate,
		table.Customer.LastUpdate,
		table.Customer.Active,
	).FROM(
		table.Customer,
	).WHERE(
		table.Customer.CustomerID.EQ(Int(int64(id))),
	)
}

func updateCustomerQuery(c *model.Customer) UpdateStatement {
	return table.Customer.UPDATE(
		table.Customer.StoreID,
		table.Customer.FirstName,
		table.Customer.LastName,
		table.Customer.Email,
		table.Customer.AddressID,
		table.Customer.Activebool,
		table.Customer.CreateDate,
		table.Customer.LastUpdate,
		table.Customer.Active,
	).SET(
		c.StoreID,
		c.FirstName,
		c.LastName,
		c.Email,
		c.AddressID,
		c.ActiveBool,
		c.CreateDate,
		c.LastUpdate,
		c.Active,
	).WHERE(
		table.Customer.CustomerID.EQ(Int(int64(c.Id))),
	)
}

func deleteCustomerQuery(id int) DeleteStatement {
	return table.Customer.DELETE().
		WHERE(
			table.Customer.CustomerID.EQ(Int(int64(id))),
		)
}

func findActiveCustomersQuery(lastPaymentNotEarlier time.Time) SelectStatement {
	const maxPaymentDate = "max_payment_date"

	payments := SELECT(
		asIs(table.Payment.CustomerID),
		MAX(table.Payment.PaymentDate).AS(maxPaymentDate),
	).FROM(
		table.Payment,
	).GROUP_BY(
		table.Payment.CustomerID,
	)

	return SELECT(
		table.Customer.AllColumns,
	).FROM(
		table.Customer.INNER_JOIN(
			payments.AsTable(table.Payment.TableName()), table.Customer.CustomerID.EQ(table.Payment.CustomerID),
		),
	).WHERE(
		TimestampColumn(maxPaymentDate).GT_EQ(TimestampT(lastPaymentNotEarlier)),
	)
}

func listCustomersQuery(listOptions *repository.ListOptions, criteria *repository.CustomerCriteria) SelectStatement {
	stmt := SELECT(
		table.Customer.AllColumns,
	).FROM(
		table.Customer,
	)

	var conditions = Bool(true)
	if criteria != nil && criteria.StoreID != nil {
		conditions = conditions.AND(table.Customer.StoreID.EQ(Int(int64(*criteria.StoreID))))
	}

	if criteria != nil && criteria.Query != nil {
		query := RawString("'%' || query || '%'", map[string]interface{}{"query": *criteria.Query})

		conditions = conditions.AND(
			table.Customer.FirstName.LIKE(query).
				OR(table.Customer.LastName.LIKE(query)).
				OR(table.Customer.Email.LIKE(query)),
		)
	}
	stmt.WHERE(conditions)

	if listOptions != nil {
		for _, sort := range listOptions.Sort {
			if sort.Direction == repository.SortDirectionASC {
				stmt.ORDER_BY(RawString(sort.Property).ASC())
			} else {
				stmt.ORDER_BY(RawString(sort.Property).DESC())
			}
		}
	}

	if listOptions != nil && listOptions.Offset > 0 {
		stmt.OFFSET(int64(listOptions.Offset))
	}

	if listOptions != nil && listOptions.Limit > 0 {
		stmt.LIMIT(int64(listOptions.Limit))
	}

	return stmt
}
