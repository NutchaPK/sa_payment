package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// PayType holds the schema definition for the PayType entity.
type PayType struct {
	ent.Schema
}

// Fields of the PayType.
func (PayType) Fields() []ent.Field {
	return []ent.Field {
		field.String("TypeInfo").NotEmpty().Unique(),
	}
}

// Edges of the PayType.
func (PayType) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("payments",Payment.Type).StorageKey(edge.Column("paytype_id")),
	}
}
