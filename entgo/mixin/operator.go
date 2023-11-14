package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var _ ent.Mixin = (*CreateBy)(nil)

type CreateBy struct{ mixin.Schema }

func (CreateBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("create_by").
			Comment("create_by").
			Optional(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*UpdateBy)(nil)

type UpdateBy struct{ mixin.Schema }

func (UpdateBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("update_by").
			Comment("update_by").
			Optional(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*DeleteBy)(nil)

type DeleteBy struct{ mixin.Schema }

func (DeleteBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("delete_by").
			Comment("delete_by").
			Optional(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*CreatedBy)(nil)

type CreatedBy struct{ mixin.Schema }

func (CreatedBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("created_by").
			Comment("created_by").
			Optional(),
	}
}

var _ ent.Mixin = (*UpdatedBy)(nil)

type UpdatedBy struct{ mixin.Schema }

func (UpdatedBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("updated_by").
			Comment("updated_by").
			Optional(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*DeletedBy)(nil)

type DeletedBy struct{ mixin.Schema }

func (DeletedBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("deleted_by").
			Comment("deleted_by").
			Optional(),
	}
}
