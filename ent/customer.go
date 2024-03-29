// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DaniilStepanenko/database-communication/ent/customer"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StoreID holds the value of the "store_id" field.
	StoreID int `json:"store_id,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// AddressID holds the value of the "address_id" field.
	AddressID int `json:"address_id,omitempty"`
	// Activebool holds the value of the "activebool" field.
	Activebool bool `json:"activebool,omitempty"`
	// CreateDate holds the value of the "create_date" field.
	CreateDate time.Time `json:"create_date,omitempty"`
	// LastUpdate holds the value of the "last_update" field.
	LastUpdate time.Time `json:"last_update,omitempty"`
	// Active holds the value of the "active" field.
	Active int `json:"active,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges CustomerEdges `json:"edges"`
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Payments holds the value of the payments edge.
	Payments []*Payment `json:"payments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PaymentsOrErr returns the Payments value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) PaymentsOrErr() ([]*Payment, error) {
	if e.loadedTypes[0] {
		return e.Payments, nil
	}
	return nil, &NotLoadedError{edge: "payments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldActivebool:
			values[i] = new(sql.NullBool)
		case customer.FieldID, customer.FieldStoreID, customer.FieldAddressID, customer.FieldActive:
			values[i] = new(sql.NullInt64)
		case customer.FieldFirstName, customer.FieldLastName, customer.FieldEmail:
			values[i] = new(sql.NullString)
		case customer.FieldCreateDate, customer.FieldLastUpdate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Customer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case customer.FieldStoreID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field store_id", values[i])
			} else if value.Valid {
				c.StoreID = int(value.Int64)
			}
		case customer.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				c.FirstName = value.String
			}
		case customer.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				c.LastName = value.String
			}
		case customer.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case customer.FieldAddressID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field address_id", values[i])
			} else if value.Valid {
				c.AddressID = int(value.Int64)
			}
		case customer.FieldActivebool:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field activebool", values[i])
			} else if value.Valid {
				c.Activebool = value.Bool
			}
		case customer.FieldCreateDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_date", values[i])
			} else if value.Valid {
				c.CreateDate = value.Time
			}
		case customer.FieldLastUpdate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_update", values[i])
			} else if value.Valid {
				c.LastUpdate = value.Time
			}
		case customer.FieldActive:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				c.Active = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPayments queries the "payments" edge of the Customer entity.
func (c *Customer) QueryPayments() *PaymentQuery {
	return (&CustomerClient{config: c.config}).QueryPayments(c)
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("store_id=")
	builder.WriteString(fmt.Sprintf("%v", c.StoreID))
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(c.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(c.LastName)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("address_id=")
	builder.WriteString(fmt.Sprintf("%v", c.AddressID))
	builder.WriteString(", ")
	builder.WriteString("activebool=")
	builder.WriteString(fmt.Sprintf("%v", c.Activebool))
	builder.WriteString(", ")
	builder.WriteString("create_date=")
	builder.WriteString(c.CreateDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_update=")
	builder.WriteString(c.LastUpdate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", c.Active))
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer

func (c Customers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
