package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Repo holds the schema definition for the Repo entity.
type Repo struct {
	ent.Schema
}

// Fields of the Repo.
func (Repo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name"),
		field.Text("description"),
		field.Text("url"),
		field.Text("sub_topic"),
		field.Enum("type").Values("repo", "resource"),
		field.Int("star_count").Optional(),
		field.Int("fork_count").Optional(),
		field.Int("watch_count").Optional(),
		field.Time("updated_at").Optional(),
		field.Int("topic_id"),
	}
}

// Edges of the Repo.
func (Repo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("topics", Topic.Type).Ref("repos").
			Field("topic_id").Required().Unique(),
	}
}

// Annotations of the Repo.
func (Repo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "repo"},
		edge.Annotation{StructTag: `json:"-"`},
	}
}
