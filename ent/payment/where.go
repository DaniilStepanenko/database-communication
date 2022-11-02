// Code generated by ent, DO NOT EDIT.

package payment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DaniilStepanenko/database-communication/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CustomerID applies equality check predicate on the "customer_id" field. It's identical to CustomerIDEQ.
func CustomerID(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomerID), v))
	})
}

// StaffID applies equality check predicate on the "staff_id" field. It's identical to StaffIDEQ.
func StaffID(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStaffID), v))
	})
}

// RentalID applies equality check predicate on the "rental_id" field. It's identical to RentalIDEQ.
func RentalID(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRentalID), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// PaymentDate applies equality check predicate on the "payment_date" field. It's identical to PaymentDateEQ.
func PaymentDate(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentDate), v))
	})
}

// CustomerIDEQ applies the EQ predicate on the "customer_id" field.
func CustomerIDEQ(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomerID), v))
	})
}

// CustomerIDNEQ applies the NEQ predicate on the "customer_id" field.
func CustomerIDNEQ(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCustomerID), v))
	})
}

// CustomerIDIn applies the In predicate on the "customer_id" field.
func CustomerIDIn(vs ...int) predicate.Payment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCustomerID), v...))
	})
}

// CustomerIDNotIn applies the NotIn predicate on the "customer_id" field.
func CustomerIDNotIn(vs ...int) predicate.Payment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCustomerID), v...))
	})
}

// StaffIDEQ applies the EQ predicate on the "staff_id" field.
func StaffIDEQ(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStaffID), v))
	})
}

// StaffIDNEQ applies the NEQ predicate on the "staff_id" field.
func StaffIDNEQ(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStaffID), v))
	})
}

// StaffIDIn applies the In predicate on the "staff_id" field.
func StaffIDIn(vs ...int) predicate.Payment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStaffID), v...))
	})
}

// StaffIDNotIn applies the NotIn predicate on the "staff_id" field.
func StaffIDNotIn(vs ...int) predicate.Payment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStaffID), v...))
	})
}

// StaffIDGT applies the GT predicate on the "staff_id" field.
func StaffIDGT(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStaffID), v))
	})
}

// StaffIDGTE applies the GTE predicate on the "staff_id" field.
func StaffIDGTE(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStaffID), v))
	})
}

// StaffIDLT applies the LT predicate on the "staff_id" field.
func StaffIDLT(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStaffID), v))
	})
}

// StaffIDLTE applies the LTE predicate on the "staff_id" field.
func StaffIDLTE(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStaffID), v))
	})
}

// RentalIDEQ applies the EQ predicate on the "rental_id" field.
func RentalIDEQ(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRentalID), v))
	})
}

// RentalIDNEQ applies the NEQ predicate on the "rental_id" field.
func RentalIDNEQ(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRentalID), v))
	})
}

// RentalIDIn applies the In predicate on the "rental_id" field.
func RentalIDIn(vs ...int) predicate.Payment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRentalID), v...))
	})
}

// RentalIDNotIn applies the NotIn predicate on the "rental_id" field.
func RentalIDNotIn(vs ...int) predicate.Payment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRentalID), v...))
	})
}

// RentalIDGT applies the GT predicate on the "rental_id" field.
func RentalIDGT(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRentalID), v))
	})
}

// RentalIDGTE applies the GTE predicate on the "rental_id" field.
func RentalIDGTE(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRentalID), v))
	})
}

// RentalIDLT applies the LT predicate on the "rental_id" field.
func RentalIDLT(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRentalID), v))
	})
}

// RentalIDLTE applies the LTE predicate on the "rental_id" field.
func RentalIDLTE(v int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRentalID), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Payment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Payment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// PaymentDateEQ applies the EQ predicate on the "payment_date" field.
func PaymentDateEQ(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentDate), v))
	})
}

// PaymentDateNEQ applies the NEQ predicate on the "payment_date" field.
func PaymentDateNEQ(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaymentDate), v))
	})
}

// PaymentDateIn applies the In predicate on the "payment_date" field.
func PaymentDateIn(vs ...time.Time) predicate.Payment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPaymentDate), v...))
	})
}

// PaymentDateNotIn applies the NotIn predicate on the "payment_date" field.
func PaymentDateNotIn(vs ...time.Time) predicate.Payment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPaymentDate), v...))
	})
}

// PaymentDateGT applies the GT predicate on the "payment_date" field.
func PaymentDateGT(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaymentDate), v))
	})
}

// PaymentDateGTE applies the GTE predicate on the "payment_date" field.
func PaymentDateGTE(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaymentDate), v))
	})
}

// PaymentDateLT applies the LT predicate on the "payment_date" field.
func PaymentDateLT(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaymentDate), v))
	})
}

// PaymentDateLTE applies the LTE predicate on the "payment_date" field.
func PaymentDateLTE(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaymentDate), v))
	})
}

// HasPayer applies the HasEdge predicate on the "payer" edge.
func HasPayer() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PayerTable, CustomerFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PayerTable, PayerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPayerWith applies the HasEdge predicate on the "payer" edge with a given conditions (other predicates).
func HasPayerWith(preds ...predicate.Customer) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PayerInverseTable, CustomerFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PayerTable, PayerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		p(s.Not())
	})
}
