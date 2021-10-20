// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DaniilStepanenko/database-communication/ent/payment"
)

// PaymentCreate is the builder for creating a Payment entity.
type PaymentCreate struct {
	config
	mutation *PaymentMutation
	hooks    []Hook
}

// SetCustomerID sets the "customer_id" field.
func (pc *PaymentCreate) SetCustomerID(i int) *PaymentCreate {
	pc.mutation.SetCustomerID(i)
	return pc
}

// SetStaffID sets the "staff_id" field.
func (pc *PaymentCreate) SetStaffID(i int) *PaymentCreate {
	pc.mutation.SetStaffID(i)
	return pc
}

// SetRentalID sets the "rental_id" field.
func (pc *PaymentCreate) SetRentalID(i int) *PaymentCreate {
	pc.mutation.SetRentalID(i)
	return pc
}

// SetAmount sets the "amount" field.
func (pc *PaymentCreate) SetAmount(i int) *PaymentCreate {
	pc.mutation.SetAmount(i)
	return pc
}

// SetPaymentDate sets the "payment_date" field.
func (pc *PaymentCreate) SetPaymentDate(t time.Time) *PaymentCreate {
	pc.mutation.SetPaymentDate(t)
	return pc
}

// SetID sets the "id" field.
func (pc *PaymentCreate) SetID(i int) *PaymentCreate {
	pc.mutation.SetID(i)
	return pc
}

// Mutation returns the PaymentMutation object of the builder.
func (pc *PaymentCreate) Mutation() *PaymentMutation {
	return pc.mutation
}

// Save creates the Payment in the database.
func (pc *PaymentCreate) Save(ctx context.Context) (*Payment, error) {
	var (
		err  error
		node *Payment
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PaymentCreate) SaveX(ctx context.Context) *Payment {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PaymentCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PaymentCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PaymentCreate) check() error {
	if _, ok := pc.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer_id", err: errors.New(`ent: missing required field "customer_id"`)}
	}
	if _, ok := pc.mutation.StaffID(); !ok {
		return &ValidationError{Name: "staff_id", err: errors.New(`ent: missing required field "staff_id"`)}
	}
	if _, ok := pc.mutation.RentalID(); !ok {
		return &ValidationError{Name: "rental_id", err: errors.New(`ent: missing required field "rental_id"`)}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "amount"`)}
	}
	if _, ok := pc.mutation.PaymentDate(); !ok {
		return &ValidationError{Name: "payment_date", err: errors.New(`ent: missing required field "payment_date"`)}
	}
	return nil
}

func (pc *PaymentCreate) sqlSave(ctx context.Context) (*Payment, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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

func (pc *PaymentCreate) createSpec() (*Payment, *sqlgraph.CreateSpec) {
	var (
		_node = &Payment{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: payment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payment.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CustomerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldCustomerID,
		})
		_node.CustomerID = value
	}
	if value, ok := pc.mutation.StaffID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldStaffID,
		})
		_node.StaffID = value
	}
	if value, ok := pc.mutation.RentalID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldRentalID,
		})
		_node.RentalID = value
	}
	if value, ok := pc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := pc.mutation.PaymentDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payment.FieldPaymentDate,
		})
		_node.PaymentDate = value
	}
	return _node, _spec
}

// PaymentCreateBulk is the builder for creating many Payment entities in bulk.
type PaymentCreateBulk struct {
	config
	builders []*PaymentCreate
}

// Save creates the Payment entities in the database.
func (pcb *PaymentCreateBulk) Save(ctx context.Context) ([]*Payment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Payment, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PaymentCreateBulk) SaveX(ctx context.Context) []*Payment {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PaymentCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PaymentCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
