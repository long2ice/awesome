package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TopicCategory holds the schema definition for the TopicCategory entity.
type TopicCategory struct {
	ent.Schema
}

// Fields of the TopicCategory.
func (TopicCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("icon"),
	}
}

// Edges of the TopicCategory.
func (TopicCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("topics", Topic.Type),
	}
}

// Annotations of the TopicCategory.
func (TopicCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "topic_category"},
	}
}
