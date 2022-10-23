package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Meeting holds the schema definition for the Meeting entity.
type Meeting struct {
	ent.Schema
}

// Fields of the Meeting.
func (Meeting) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.Time("starts_at"),
		field.Time("ends_at"),
		field.Text("description").
			Optional().
			Nillable(),
	}
}

// Mixin of the Meeting.
func (Meeting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Meeting.
func (Meeting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organizer", User.Type).
			Ref("meetings").
			Required(),
	}
}
