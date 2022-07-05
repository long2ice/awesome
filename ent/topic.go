// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/long2ice/awesome/ent/platform"
	"github.com/long2ice/awesome/ent/topic"
)

// Topic is the model entity for the Topic schema.
type Topic struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// SubName holds the value of the "sub_name" field.
	SubName string `json:"sub_name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// GithubURL holds the value of the "github_url" field.
	GithubURL string `json:"github_url,omitempty"`
	// PlatformID holds the value of the "platform_id" field.
	PlatformID int `json:"platform_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TopicQuery when eager-loading is set.
	Edges TopicEdges `json:"-"`
}

// TopicEdges holds the relations/edges for other nodes in the graph.
type TopicEdges struct {
	// Platform holds the value of the platform edge.
	Platform *Platform `json:"platform,omitempty"`
	// Repos holds the value of the repos edge.
	Repos []*Repo `json:"repos,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PlatformOrErr returns the Platform value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TopicEdges) PlatformOrErr() (*Platform, error) {
	if e.loadedTypes[0] {
		if e.Platform == nil {
			// The edge platform was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: platform.Label}
		}
		return e.Platform, nil
	}
	return nil, &NotLoadedError{edge: "platform"}
}

// ReposOrErr returns the Repos value or an error if the edge
// was not loaded in eager-loading.
func (e TopicEdges) ReposOrErr() ([]*Repo, error) {
	if e.loadedTypes[1] {
		return e.Repos, nil
	}
	return nil, &NotLoadedError{edge: "repos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Topic) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case topic.FieldID, topic.FieldPlatformID:
			values[i] = new(sql.NullInt64)
		case topic.FieldName, topic.FieldSubName, topic.FieldDescription, topic.FieldURL, topic.FieldGithubURL:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Topic", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Topic fields.
func (t *Topic) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case topic.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case topic.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case topic.FieldSubName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sub_name", values[i])
			} else if value.Valid {
				t.SubName = value.String
			}
		case topic.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case topic.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				t.URL = value.String
			}
		case topic.FieldGithubURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field github_url", values[i])
			} else if value.Valid {
				t.GithubURL = value.String
			}
		case topic.FieldPlatformID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field platform_id", values[i])
			} else if value.Valid {
				t.PlatformID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPlatform queries the "platform" edge of the Topic entity.
func (t *Topic) QueryPlatform() *PlatformQuery {
	return (&TopicClient{config: t.config}).QueryPlatform(t)
}

// QueryRepos queries the "repos" edge of the Topic entity.
func (t *Topic) QueryRepos() *RepoQuery {
	return (&TopicClient{config: t.config}).QueryRepos(t)
}

// Update returns a builder for updating this Topic.
// Note that you need to call Topic.Unwrap() before calling this method if this Topic
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Topic) Update() *TopicUpdateOne {
	return (&TopicClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Topic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Topic) Unwrap() *Topic {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Topic is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Topic) String() string {
	var builder strings.Builder
	builder.WriteString("Topic(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", sub_name=")
	builder.WriteString(t.SubName)
	builder.WriteString(", description=")
	builder.WriteString(t.Description)
	builder.WriteString(", url=")
	builder.WriteString(t.URL)
	builder.WriteString(", github_url=")
	builder.WriteString(t.GithubURL)
	builder.WriteString(", platform_id=")
	builder.WriteString(fmt.Sprintf("%v", t.PlatformID))
	builder.WriteByte(')')
	return builder.String()
}

// Topics is a parsable slice of Topic.
type Topics []*Topic

func (t Topics) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
