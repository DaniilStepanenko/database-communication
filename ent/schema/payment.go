package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

func (Payment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "payment"},
	}
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().StorageKey("payment_id"),
		field.Int("customer_id"),
		field.Int("staff_id"),
		field.Int("rental_id"),
		field.Float("amount"),
		field.Time("payment_date"),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payer", Customer.Type).
			Ref("payments").
			Unique().
			Field("customer_id").
			Required(),
	}
}
