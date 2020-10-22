package schema

import (

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Time("AddDatetime"),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("paytype", PayType.Type).Ref("payments").Unique(),
		edge.From("paymentstatus",PaymentStatus.Type).Ref("payments").Unique(),
		edge.From("bill", Bill.Type).Ref("payments").Unique(),
	}
}
