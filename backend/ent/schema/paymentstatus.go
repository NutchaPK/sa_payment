package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// PaymentStatus holds the schema definition for the PaymentStatus entity.
type PaymentStatus struct {
	ent.Schema
}

// Fields of the PaymentStatus.
func (PaymentStatus) Fields() []ent.Field {
	return []ent.Field {
		field.String("PaymentStatus").NotEmpty().Unique(),
	}
}

// Edges of the PaymentStatus.
func (PaymentStatus) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("payments",Payment.Type).StorageKey(edge.Column("paymentstatus_id")),
	}
}
