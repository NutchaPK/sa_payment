package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// BillStatus holds the schema definition for the BillStatus entity.
type BillStatus struct {
	ent.Schema
}

// Fields of the BillStatus.
func (BillStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("BillStatus").Unique(),
	}
}

// Edges of the BillStatus.
func (BillStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bills", Bill.Type).StorageKey(edge.Column("billstatus_id")),
	}
}
