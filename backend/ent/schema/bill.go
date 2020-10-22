package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Amount"),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("billstatus", BillStatus.Type).Ref("bills").Unique(),
		edge.To("payments", Payment.Type).StorageKey(edge.Column("bill_id")).Unique(),
		edge.From("resident", User.Type).Ref("bills").Unique(),
	}
}
