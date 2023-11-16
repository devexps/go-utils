package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var _ ent.Mixin = (*CreateTimestamp)(nil)

type CreateTimestamp struct{ mixin.Schema }

func (CreateTimestamp) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("create_time").
			Comment("create_time").
			Immutable().
			Optional().
			Nillable().
			DefaultFunc(time.Now().UnixMilli),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*UpdateTimestamp)(nil)

type UpdateTimestamp struct{ mixin.Schema }

func (UpdateTimestamp) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("update_time").
			Comment("update_time").
			Optional().
			Nillable().
			UpdateDefault(time.Now().UnixMilli),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*DeleteTimestamp)(nil)

type DeleteTimestamp struct{ mixin.Schema }

func (DeleteTimestamp) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("delete_time").
			Comment("delete_time").
			Optional().
			Nillable(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*Timestamp)(nil)

type Timestamp struct{ mixin.Schema }

func (Timestamp) Fields() []ent.Field {
	var fields []ent.Field
	fields = append(fields, CreateTimestamp{}.Fields()...)
	fields = append(fields, UpdateTimestamp{}.Fields()...)
	fields = append(fields, DeleteTimestamp{}.Fields()...)
	return fields
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*CreatedAtTimestamp)(nil)

type CreatedAtTimestamp struct{ mixin.Schema }

func (CreatedAtTimestamp) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("created_at").
			Comment("created_at").
			Immutable().
			Optional().
			Nillable().
			DefaultFunc(time.Now().UnixMilli),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*UpdatedAtTimestamp)(nil)

type UpdatedAtTimestamp struct{ mixin.Schema }

func (UpdatedAtTimestamp) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("updated_at").
			Comment("updated_at").
			Optional().
			Nillable().
			UpdateDefault(time.Now().UnixMilli),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*DeletedAtTimestamp)(nil)

type DeletedAtTimestamp struct{ mixin.Schema }

func (DeletedAtTimestamp) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("deleted_at").
			Comment("deleted_at").
			Optional().
			Nillable(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*TimestampAt)(nil)

type TimestampAt struct{ mixin.Schema }

func (TimestampAt) Fields() []ent.Field {
	var fields []ent.Field
	fields = append(fields, CreatedAtTimestamp{}.Fields()...)
	fields = append(fields, UpdatedAtTimestamp{}.Fields()...)
	fields = append(fields, DeletedAtTimestamp{}.Fields()...)
	return fields
}
