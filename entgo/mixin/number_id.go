package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

var _ ent.Mixin = (*NumberId)(nil)

type NumberId struct{ mixin.Schema }

func (NumberId) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Comment("id").
			Positive().
			Immutable().
			Unique(),
	}
}

// Indexes of the NumberId.
func (NumberId) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
