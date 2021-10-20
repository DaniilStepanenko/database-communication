package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "customer"},
	}
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().StorageKey("customer_id"),
		field.Int("store_id"),
		field.String("first_name"),
		field.String("last_name"),
		field.String("email").Optional(),
		field.Int("address_id"),
		field.Bool("activebool"),
		field.Time("create_date"),
		field.Time("last_update").Optional(),
		field.Int("active").Optional(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("payments", Payment.Type),
	}
}
