// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/long2ice/awesome/ent/platform"
)

// Platform is the model entity for the Platform schema.
type Platform struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon string `json:"icon,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlatformQuery when eager-loading is set.
	Edges PlatformEdges `json:"-"`
}

// PlatformEdges holds the relations/edges for other nodes in the graph.
type PlatformEdges struct {
	// Topics holds the value of the topics edge.
	Topics []*Topic `json:"topics,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TopicsOrErr returns the Topics value or an error if the edge
// was not loaded in eager-loading.
func (e PlatformEdges) TopicsOrErr() ([]*Topic, error) {
	if e.loadedTypes[0] {
		return e.Topics, nil
	}
	return nil, &NotLoadedError{edge: "topics"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Platform) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case platform.FieldID:
			values[i] = new(sql.NullInt64)
		case platform.FieldName, platform.FieldIcon:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Platform", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Platform fields.
func (pl *Platform) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case platform.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pl.ID = int(value.Int64)
		case platform.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case platform.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				pl.Icon = value.String
			}
		}
	}
	return nil
}

// QueryTopics queries the "topics" edge of the Platform entity.
func (pl *Platform) QueryTopics() *TopicQuery {
	return (&PlatformClient{config: pl.config}).QueryTopics(pl)
}

// Update returns a builder for updating this Platform.
// Note that you need to call Platform.Unwrap() before calling this method if this Platform
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Platform) Update() *PlatformUpdateOne {
	return (&PlatformClient{config: pl.config}).UpdateOne(pl)
}

// Unwrap unwraps the Platform entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Platform) Unwrap() *Platform {
	tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Platform is not a transactional entity")
	}
	pl.config.driver = tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Platform) String() string {
	var builder strings.Builder
	builder.WriteString("Platform(")
	builder.WriteString(fmt.Sprintf("id=%v", pl.ID))
	builder.WriteString(", name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", icon=")
	builder.WriteString(pl.Icon)
	builder.WriteByte(')')
	return builder.String()
}

// Platforms is a parsable slice of Platform.
type Platforms []*Platform

func (pl Platforms) config(cfg config) {
	for _i := range pl {
		pl[_i].config = cfg
	}
}