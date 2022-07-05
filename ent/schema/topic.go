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
		field.String("sub_name"),
		field.String("description"),
		field.String("url").Unique(),
		field.String("github_url"),
		field.Int("platform_id"),
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("platform", Platform.Type).Ref("topics").Required().Field("platform_id").Unique(),
		edge.To("repos", Repo.Type),
	}
}

// Annotations of the Topic.
func (Topic) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "topic"},
	}
}
