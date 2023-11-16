package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var _ ent.Mixin = (*CreatedAt)(nil)

type CreatedAt struct{ mixin.Schema }

func (CreatedAt) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Comment("created_at").
			Immutable().
			Optional().
			Nillable(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*UpdatedAt)(nil)

type UpdatedAt struct{ mixin.Schema }

func (UpdatedAt) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").
			Comment("updated_at").
			Optional().
			Nillable(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*DeletedAt)(nil)

type DeletedAt struct{ mixin.Schema }

func (DeletedAt) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Comment("deleted_at").
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*TimeAt)(nil)

type TimeAt struct{ mixin.Schema }

func (TimeAt) Fields() []ent.Field {
	var fields []ent.Field
	fields = append(fields, CreatedAt{}.Fields()...)
	fields = append(fields, UpdatedAt{}.Fields()...)
	fields = append(fields, DeletedAt{}.Fields()...)
	return fields
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*CreateTime)(nil)

type CreateTime struct{ mixin.Schema }

func (CreateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").
			Comment("create_time").
			Immutable().
			Optional().
			Nillable(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*UpdateTime)(nil)

type UpdateTime struct{ mixin.Schema }

func (UpdateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("update_time").
			Comment("update_time").
			Optional().
			Nillable(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*DeleteTime)(nil)

type DeleteTime struct{ mixin.Schema }

func (DeleteTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("delete_time").
			Comment("delete_time").
			Optional().
			Nillable(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*Time)(nil)

type Time struct{ mixin.Schema }

func (Time) Fields() []ent.Field {
	var fields []ent.Field
	fields = append(fields, CreateTime{}.Fields()...)
	fields = append(fields, UpdateTime{}.Fields()...)
	fields = append(fields, DeleteTime{}.Fields()...)
	return fields
}
