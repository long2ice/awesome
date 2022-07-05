package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Platform holds the schema definition for the Platform entity.
type Platform struct {
	ent.Schema
}

// Fields of the Platform.
func (Platform) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("icon").Optional(),
	}
}

// Edges of the Platform.
func (Platform) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("topics", Topic.Type),
	}
}

// Annotations of the Platform.
func (Platform) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "platform"},
	}
}
