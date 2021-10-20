// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DaniilStepanenko/database-communication/ent/customer"
	"github.com/DaniilStepanenko/database-communication/ent/payment"
)

// CustomerCreate is the builder for creating a Customer entity.
type CustomerCreate struct {
	config
	mutation *CustomerMutation
	hooks    []Hook
}

// SetStoreID sets the "store_id" field.
func (cc *CustomerCreate) SetStoreID(i int) *CustomerCreate {
	cc.mutation.SetStoreID(i)
	return cc
}

// SetFirstName sets the "first_name" field.
func (cc *CustomerCreate) SetFirstName(s string) *CustomerCreate {
	cc.mutation.SetFirstName(s)
	return cc
}

// SetLastName sets the "last_name" field.
func (cc *CustomerCreate) SetLastName(s string) *CustomerCreate {
	cc.mutation.SetLastName(s)
	return cc
}

// SetEmail sets the "email" field.
func (cc *CustomerCreate) SetEmail(s string) *CustomerCreate {
	cc.mutation.SetEmail(s)
	return cc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableEmail(s *string) *CustomerCreate {
	if s != nil {
		cc.SetEmail(*s)
	}
	return cc
}

// SetAddressID sets the "address_id" field.
func (cc *CustomerCreate) SetAddressID(i int) *CustomerCreate {
	cc.mutation.SetAddressID(i)
	return cc
}

// SetActivebool sets the "activebool" field.
func (cc *CustomerCreate) SetActivebool(b bool) *CustomerCreate {
	cc.mutation.SetActivebool(b)
	return cc
}

// SetCreateDate sets the "create_date" field.
func (cc *CustomerCreate) SetCreateDate(t time.Time) *CustomerCreate {
	cc.mutation.SetCreateDate(t)
	return cc
}

// SetLastUpdate sets the "last_update" field.
func (cc *CustomerCreate) SetLastUpdate(t time.Time) *CustomerCreate {
	cc.mutation.SetLastUpdate(t)
	return cc
}

// SetNillableLastUpdate sets the "last_update" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableLastUpdate(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetLastUpdate(*t)
	}
	return cc
}

// SetActive sets the "active" field.
func (cc *CustomerCreate) SetActive(i int) *CustomerCreate {
	cc.mutation.SetActive(i)
	return cc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableActive(i *int) *CustomerCreate {
	if i != nil {
		cc.SetActive(*i)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CustomerCreate) SetID(i int) *CustomerCreate {
	cc.mutation.SetID(i)
	return cc
}

// AddPaymentIDs adds the "payments" edge to the Payment entity by IDs.
func (cc *CustomerCreate) AddPaymentIDs(ids ...int) *CustomerCreate {
	cc.mutation.AddPaymentIDs(ids...)
	return cc
}

// AddPayments adds the "payments" edges to the Payment entity.
func (cc *CustomerCreate) AddPayments(p ...*Payment) *CustomerCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddPaymentIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cc *CustomerCreate) Mutation() *CustomerMutation {
	return cc.mutation
}

// Save creates the Customer in the database.
func (cc *CustomerCreate) Save(ctx context.Context) (*Customer, error) {
	var (
		err  error
		node *Customer
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CustomerCreate) SaveX(ctx context.Context) *Customer {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CustomerCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CustomerCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CustomerCreate) check() error {
	if _, ok := cc.mutation.StoreID(); !ok {
		return &ValidationError{Name: "store_id", err: errors.New(`ent: missing required field "store_id"`)}
	}
	if _, ok := cc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "first_name"`)}
	}
	if _, ok := cc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "last_name"`)}
	}
	if _, ok := cc.mutation.AddressID(); !ok {
		return &ValidationError{Name: "address_id", err: errors.New(`ent: missing required field "address_id"`)}
	}
	if _, ok := cc.mutation.Activebool(); !ok {
		return &ValidationError{Name: "activebool", err: errors.New(`ent: missing required field "activebool"`)}
	}
	if _, ok := cc.mutation.CreateDate(); !ok {
		return &ValidationError{Name: "create_date", err: errors.New(`ent: missing required field "create_date"`)}
	}
	return nil
}

func (cc *CustomerCreate) sqlSave(ctx context.Context) (*Customer, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (cc *CustomerCreate) createSpec() (*Customer, *sqlgraph.CreateSpec) {
	var (
		_node = &Customer{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: customer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.StoreID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: customer.FieldStoreID,
		})
		_node.StoreID = value
	}
	if value, ok := cc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := cc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := cc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := cc.mutation.AddressID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: customer.FieldAddressID,
		})
		_node.AddressID = value
	}
	if value, ok := cc.mutation.Activebool(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: customer.FieldActivebool,
		})
		_node.Activebool = value
	}
	if value, ok := cc.mutation.CreateDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldCreateDate,
		})
		_node.CreateDate = value
	}
	if value, ok := cc.mutation.LastUpdate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldLastUpdate,
		})
		_node.LastUpdate = value
	}
	if value, ok := cc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: customer.FieldActive,
		})
		_node.Active = value
	}
	if nodes := cc.mutation.PaymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PaymentsTable,
			Columns: []string{customer.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CustomerCreateBulk is the builder for creating many Customer entities in bulk.
type CustomerCreateBulk struct {
	config
	builders []*CustomerCreate
}

// Save creates the Customer entities in the database.
func (ccb *CustomerCreateBulk) Save(ctx context.Context) ([]*Customer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Customer, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CustomerCreateBulk) SaveX(ctx context.Context) []*Customer {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CustomerCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CustomerCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
