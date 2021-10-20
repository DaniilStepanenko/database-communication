// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DaniilStepanenko/database-communication/ent/payment"
	"github.com/DaniilStepanenko/database-communication/ent/predicate"
)

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCustomerID sets the "customer_id" field.
func (pu *PaymentUpdate) SetCustomerID(i int) *PaymentUpdate {
	pu.mutation.ResetCustomerID()
	pu.mutation.SetCustomerID(i)
	return pu
}

// AddCustomerID adds i to the "customer_id" field.
func (pu *PaymentUpdate) AddCustomerID(i int) *PaymentUpdate {
	pu.mutation.AddCustomerID(i)
	return pu
}

// SetStaffID sets the "staff_id" field.
func (pu *PaymentUpdate) SetStaffID(i int) *PaymentUpdate {
	pu.mutation.ResetStaffID()
	pu.mutation.SetStaffID(i)
	return pu
}

// AddStaffID adds i to the "staff_id" field.
func (pu *PaymentUpdate) AddStaffID(i int) *PaymentUpdate {
	pu.mutation.AddStaffID(i)
	return pu
}

// SetRentalID sets the "rental_id" field.
func (pu *PaymentUpdate) SetRentalID(i int) *PaymentUpdate {
	pu.mutation.ResetRentalID()
	pu.mutation.SetRentalID(i)
	return pu
}

// AddRentalID adds i to the "rental_id" field.
func (pu *PaymentUpdate) AddRentalID(i int) *PaymentUpdate {
	pu.mutation.AddRentalID(i)
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PaymentUpdate) SetAmount(i int) *PaymentUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(i)
	return pu
}

// AddAmount adds i to the "amount" field.
func (pu *PaymentUpdate) AddAmount(i int) *PaymentUpdate {
	pu.mutation.AddAmount(i)
	return pu
}

// SetPaymentDate sets the "payment_date" field.
func (pu *PaymentUpdate) SetPaymentDate(t time.Time) *PaymentUpdate {
	pu.mutation.SetPaymentDate(t)
	return pu
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payment.Table,
			Columns: payment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payment.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CustomerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldCustomerID,
		})
	}
	if value, ok := pu.mutation.AddedCustomerID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldCustomerID,
		})
	}
	if value, ok := pu.mutation.StaffID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldStaffID,
		})
	}
	if value, ok := pu.mutation.AddedStaffID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldStaffID,
		})
	}
	if value, ok := pu.mutation.RentalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldRentalID,
		})
	}
	if value, ok := pu.mutation.AddedRentalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldRentalID,
		})
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := pu.mutation.PaymentDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payment.FieldPaymentDate,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentMutation
}

// SetCustomerID sets the "customer_id" field.
func (puo *PaymentUpdateOne) SetCustomerID(i int) *PaymentUpdateOne {
	puo.mutation.ResetCustomerID()
	puo.mutation.SetCustomerID(i)
	return puo
}

// AddCustomerID adds i to the "customer_id" field.
func (puo *PaymentUpdateOne) AddCustomerID(i int) *PaymentUpdateOne {
	puo.mutation.AddCustomerID(i)
	return puo
}

// SetStaffID sets the "staff_id" field.
func (puo *PaymentUpdateOne) SetStaffID(i int) *PaymentUpdateOne {
	puo.mutation.ResetStaffID()
	puo.mutation.SetStaffID(i)
	return puo
}

// AddStaffID adds i to the "staff_id" field.
func (puo *PaymentUpdateOne) AddStaffID(i int) *PaymentUpdateOne {
	puo.mutation.AddStaffID(i)
	return puo
}

// SetRentalID sets the "rental_id" field.
func (puo *PaymentUpdateOne) SetRentalID(i int) *PaymentUpdateOne {
	puo.mutation.ResetRentalID()
	puo.mutation.SetRentalID(i)
	return puo
}

// AddRentalID adds i to the "rental_id" field.
func (puo *PaymentUpdateOne) AddRentalID(i int) *PaymentUpdateOne {
	puo.mutation.AddRentalID(i)
	return puo
}

// SetAmount sets the "amount" field.
func (puo *PaymentUpdateOne) SetAmount(i int) *PaymentUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(i)
	return puo
}

// AddAmount adds i to the "amount" field.
func (puo *PaymentUpdateOne) AddAmount(i int) *PaymentUpdateOne {
	puo.mutation.AddAmount(i)
	return puo
}

// SetPaymentDate sets the "payment_date" field.
func (puo *PaymentUpdateOne) SetPaymentDate(t time.Time) *PaymentUpdateOne {
	puo.mutation.SetPaymentDate(t)
	return puo
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PaymentUpdateOne) Select(field string, fields ...string) *PaymentUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	var (
		err  error
		node *Payment
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payment.Table,
			Columns: payment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payment.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Payment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for _, f := range fields {
			if !payment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.CustomerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldCustomerID,
		})
	}
	if value, ok := puo.mutation.AddedCustomerID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldCustomerID,
		})
	}
	if value, ok := puo.mutation.StaffID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldStaffID,
		})
	}
	if value, ok := puo.mutation.AddedStaffID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldStaffID,
		})
	}
	if value, ok := puo.mutation.RentalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldRentalID,
		})
	}
	if value, ok := puo.mutation.AddedRentalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldRentalID,
		})
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := puo.mutation.PaymentDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payment.FieldPaymentDate,
		})
	}
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}