package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CreatorId struct {
	mixin.Schema
}

func (CreatorId) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("creator_id").
			Comment("creator_id").
			Immutable().
			Optional(),
	}
}
