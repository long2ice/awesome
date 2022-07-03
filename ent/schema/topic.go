package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
	ent.Schema
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.Int("topic_category_id"),
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("topiccategory", TopicCategory.Type).Ref("topics").Required().Field("topic_category_id").Unique(),
		edge.To("projects", Project.Type),
	}
}

// Annotations of the Topic.
func (Topic) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "topic"},
	}
}
