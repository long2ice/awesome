package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.String("author"),
		field.String("author_email"),
		field.String("url"),
		field.Int("star"),
		field.Int("topic_id"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("topic", Topic.Type).Ref("projects").
			Field("topic_id").Required().Unique(),
	}
}

// Annotations of the Project.
func (Project) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "project"},
	}
}
