// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/long2ice/awesome/ent/repo"
	"github.com/long2ice/awesome/ent/topic"
)

// RepoCreate is the builder for creating a Repo entity.
type RepoCreate struct {
	config
	mutation *RepoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (rc *RepoCreate) SetName(s string) *RepoCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetDescription sets the "description" field.
func (rc *RepoCreate) SetDescription(s string) *RepoCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetURL sets the "url" field.
func (rc *RepoCreate) SetURL(s string) *RepoCreate {
	rc.mutation.SetURL(s)
	return rc
}

// SetSubTopic sets the "sub_topic" field.
func (rc *RepoCreate) SetSubTopic(s string) *RepoCreate {
	rc.mutation.SetSubTopic(s)
	return rc
}

// SetType sets the "type" field.
func (rc *RepoCreate) SetType(r repo.Type) *RepoCreate {
	rc.mutation.SetType(r)
	return rc
}

// SetStarCount sets the "star_count" field.
func (rc *RepoCreate) SetStarCount(i int) *RepoCreate {
	rc.mutation.SetStarCount(i)
	return rc
}

// SetNillableStarCount sets the "star_count" field if the given value is not nil.
func (rc *RepoCreate) SetNillableStarCount(i *int) *RepoCreate {
	if i != nil {
		rc.SetStarCount(*i)
	}
	return rc
}

// SetForkCount sets the "fork_count" field.
func (rc *RepoCreate) SetForkCount(i int) *RepoCreate {
	rc.mutation.SetForkCount(i)
	return rc
}

// SetNillableForkCount sets the "fork_count" field if the given value is not nil.
func (rc *RepoCreate) SetNillableForkCount(i *int) *RepoCreate {
	if i != nil {
		rc.SetForkCount(*i)
	}
	return rc
}

// SetWatchCount sets the "watch_count" field.
func (rc *RepoCreate) SetWatchCount(i int) *RepoCreate {
	rc.mutation.SetWatchCount(i)
	return rc
}

// SetNillableWatchCount sets the "watch_count" field if the given value is not nil.
func (rc *RepoCreate) SetNillableWatchCount(i *int) *RepoCreate {
	if i != nil {
		rc.SetWatchCount(*i)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RepoCreate) SetUpdatedAt(t time.Time) *RepoCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RepoCreate) SetNillableUpdatedAt(t *time.Time) *RepoCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetTopicID sets the "topic_id" field.
func (rc *RepoCreate) SetTopicID(i int) *RepoCreate {
	rc.mutation.SetTopicID(i)
	return rc
}

// SetTopicsID sets the "topics" edge to the Topic entity by ID.
func (rc *RepoCreate) SetTopicsID(id int) *RepoCreate {
	rc.mutation.SetTopicsID(id)
	return rc
}

// SetTopics sets the "topics" edge to the Topic entity.
func (rc *RepoCreate) SetTopics(t *Topic) *RepoCreate {
	return rc.SetTopicsID(t.ID)
}

// Mutation returns the RepoMutation object of the builder.
func (rc *RepoCreate) Mutation() *RepoMutation {
	return rc.mutation
}

// Save creates the Repo in the database.
func (rc *RepoCreate) Save(ctx context.Context) (*Repo, error) {
	var (
		err  error
		node *Repo
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RepoCreate) SaveX(ctx context.Context) *Repo {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RepoCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RepoCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RepoCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Repo.name"`)}
	}
	if _, ok := rc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Repo.description"`)}
	}
	if _, ok := rc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Repo.url"`)}
	}
	if _, ok := rc.mutation.SubTopic(); !ok {
		return &ValidationError{Name: "sub_topic", err: errors.New(`ent: missing required field "Repo.sub_topic"`)}
	}
	if _, ok := rc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Repo.type"`)}
	}
	if v, ok := rc.mutation.GetType(); ok {
		if err := repo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Repo.type": %w`, err)}
		}
	}
	if _, ok := rc.mutation.TopicID(); !ok {
		return &ValidationError{Name: "topic_id", err: errors.New(`ent: missing required field "Repo.topic_id"`)}
	}
	if _, ok := rc.mutation.TopicsID(); !ok {
		return &ValidationError{Name: "topics", err: errors.New(`ent: missing required edge "Repo.topics"`)}
	}
	return nil
}

func (rc *RepoCreate) sqlSave(ctx context.Context) (*Repo, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RepoCreate) createSpec() (*Repo, *sqlgraph.CreateSpec) {
	var (
		_node = &Repo{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: repo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repo.FieldID,
			},
		}
	)
	_spec.OnConflict = rc.conflict
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := rc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := rc.mutation.SubTopic(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldSubTopic,
		})
		_node.SubTopic = value
	}
	if value, ok := rc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: repo.FieldType,
		})
		_node.Type = value
	}
	if value, ok := rc.mutation.StarCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldStarCount,
		})
		_node.StarCount = value
	}
	if value, ok := rc.mutation.ForkCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldForkCount,
		})
		_node.ForkCount = value
	}
	if value, ok := rc.mutation.WatchCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldWatchCount,
		})
		_node.WatchCount = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := rc.mutation.TopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repo.TopicsTable,
			Columns: []string{repo.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TopicID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Repo.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RepoUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (rc *RepoCreate) OnConflict(opts ...sql.ConflictOption) *RepoUpsertOne {
	rc.conflict = opts
	return &RepoUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Repo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (rc *RepoCreate) OnConflictColumns(columns ...string) *RepoUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RepoUpsertOne{
		create: rc,
	}
}

type (
	// RepoUpsertOne is the builder for "upsert"-ing
	//  one Repo node.
	RepoUpsertOne struct {
		create *RepoCreate
	}

	// RepoUpsert is the "OnConflict" setter.
	RepoUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *RepoUpsert) SetName(v string) *RepoUpsert {
	u.Set(repo.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RepoUpsert) UpdateName() *RepoUpsert {
	u.SetExcluded(repo.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *RepoUpsert) SetDescription(v string) *RepoUpsert {
	u.Set(repo.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RepoUpsert) UpdateDescription() *RepoUpsert {
	u.SetExcluded(repo.FieldDescription)
	return u
}

// SetURL sets the "url" field.
func (u *RepoUpsert) SetURL(v string) *RepoUpsert {
	u.Set(repo.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *RepoUpsert) UpdateURL() *RepoUpsert {
	u.SetExcluded(repo.FieldURL)
	return u
}

// SetSubTopic sets the "sub_topic" field.
func (u *RepoUpsert) SetSubTopic(v string) *RepoUpsert {
	u.Set(repo.FieldSubTopic, v)
	return u
}

// UpdateSubTopic sets the "sub_topic" field to the value that was provided on create.
func (u *RepoUpsert) UpdateSubTopic() *RepoUpsert {
	u.SetExcluded(repo.FieldSubTopic)
	return u
}

// SetType sets the "type" field.
func (u *RepoUpsert) SetType(v repo.Type) *RepoUpsert {
	u.Set(repo.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *RepoUpsert) UpdateType() *RepoUpsert {
	u.SetExcluded(repo.FieldType)
	return u
}

// SetStarCount sets the "star_count" field.
func (u *RepoUpsert) SetStarCount(v int) *RepoUpsert {
	u.Set(repo.FieldStarCount, v)
	return u
}

// UpdateStarCount sets the "star_count" field to the value that was provided on create.
func (u *RepoUpsert) UpdateStarCount() *RepoUpsert {
	u.SetExcluded(repo.FieldStarCount)
	return u
}

// AddStarCount adds v to the "star_count" field.
func (u *RepoUpsert) AddStarCount(v int) *RepoUpsert {
	u.Add(repo.FieldStarCount, v)
	return u
}

// ClearStarCount clears the value of the "star_count" field.
func (u *RepoUpsert) ClearStarCount() *RepoUpsert {
	u.SetNull(repo.FieldStarCount)
	return u
}

// SetForkCount sets the "fork_count" field.
func (u *RepoUpsert) SetForkCount(v int) *RepoUpsert {
	u.Set(repo.FieldForkCount, v)
	return u
}

// UpdateForkCount sets the "fork_count" field to the value that was provided on create.
func (u *RepoUpsert) UpdateForkCount() *RepoUpsert {
	u.SetExcluded(repo.FieldForkCount)
	return u
}

// AddForkCount adds v to the "fork_count" field.
func (u *RepoUpsert) AddForkCount(v int) *RepoUpsert {
	u.Add(repo.FieldForkCount, v)
	return u
}

// ClearForkCount clears the value of the "fork_count" field.
func (u *RepoUpsert) ClearForkCount() *RepoUpsert {
	u.SetNull(repo.FieldForkCount)
	return u
}

// SetWatchCount sets the "watch_count" field.
func (u *RepoUpsert) SetWatchCount(v int) *RepoUpsert {
	u.Set(repo.FieldWatchCount, v)
	return u
}

// UpdateWatchCount sets the "watch_count" field to the value that was provided on create.
func (u *RepoUpsert) UpdateWatchCount() *RepoUpsert {
	u.SetExcluded(repo.FieldWatchCount)
	return u
}

// AddWatchCount adds v to the "watch_count" field.
func (u *RepoUpsert) AddWatchCount(v int) *RepoUpsert {
	u.Add(repo.FieldWatchCount, v)
	return u
}

// ClearWatchCount clears the value of the "watch_count" field.
func (u *RepoUpsert) ClearWatchCount() *RepoUpsert {
	u.SetNull(repo.FieldWatchCount)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RepoUpsert) SetUpdatedAt(v time.Time) *RepoUpsert {
	u.Set(repo.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RepoUpsert) UpdateUpdatedAt() *RepoUpsert {
	u.SetExcluded(repo.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *RepoUpsert) ClearUpdatedAt() *RepoUpsert {
	u.SetNull(repo.FieldUpdatedAt)
	return u
}

// SetTopicID sets the "topic_id" field.
func (u *RepoUpsert) SetTopicID(v int) *RepoUpsert {
	u.Set(repo.FieldTopicID, v)
	return u
}

// UpdateTopicID sets the "topic_id" field to the value that was provided on create.
func (u *RepoUpsert) UpdateTopicID() *RepoUpsert {
	u.SetExcluded(repo.FieldTopicID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Repo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *RepoUpsertOne) UpdateNewValues() *RepoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Repo.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *RepoUpsertOne) Ignore() *RepoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RepoUpsertOne) DoNothing() *RepoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RepoCreate.OnConflict
// documentation for more info.
func (u *RepoUpsertOne) Update(set func(*RepoUpsert)) *RepoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RepoUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *RepoUpsertOne) SetName(v string) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RepoUpsertOne) UpdateName() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *RepoUpsertOne) SetDescription(v string) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RepoUpsertOne) UpdateDescription() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateDescription()
	})
}

// SetURL sets the "url" field.
func (u *RepoUpsertOne) SetURL(v string) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *RepoUpsertOne) UpdateURL() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateURL()
	})
}

// SetSubTopic sets the "sub_topic" field.
func (u *RepoUpsertOne) SetSubTopic(v string) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.SetSubTopic(v)
	})
}

// UpdateSubTopic sets the "sub_topic" field to the value that was provided on create.
func (u *RepoUpsertOne) UpdateSubTopic() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateSubTopic()
	})
}

// SetType sets the "type" field.
func (u *RepoUpsertOne) SetType(v repo.Type) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *RepoUpsertOne) UpdateType() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateType()
	})
}

// SetStarCount sets the "star_count" field.
func (u *RepoUpsertOne) SetStarCount(v int) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.SetStarCount(v)
	})
}

// AddStarCount adds v to the "star_count" field.
func (u *RepoUpsertOne) AddStarCount(v int) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.AddStarCount(v)
	})
}

// UpdateStarCount sets the "star_count" field to the value that was provided on create.
func (u *RepoUpsertOne) UpdateStarCount() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateStarCount()
	})
}

// ClearStarCount clears the value of the "star_count" field.
func (u *RepoUpsertOne) ClearStarCount() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.ClearStarCount()
	})
}

// SetForkCount sets the "fork_count" field.
func (u *RepoUpsertOne) SetForkCount(v int) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.SetForkCount(v)
	})
}

// AddForkCount adds v to the "fork_count" field.
func (u *RepoUpsertOne) AddForkCount(v int) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.AddForkCount(v)
	})
}

// UpdateForkCount sets the "fork_count" field to the value that was provided on create.
func (u *RepoUpsertOne) UpdateForkCount() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateForkCount()
	})
}

// ClearForkCount clears the value of the "fork_count" field.
func (u *RepoUpsertOne) ClearForkCount() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.ClearForkCount()
	})
}

// SetWatchCount sets the "watch_count" field.
func (u *RepoUpsertOne) SetWatchCount(v int) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.SetWatchCount(v)
	})
}

// AddWatchCount adds v to the "watch_count" field.
func (u *RepoUpsertOne) AddWatchCount(v int) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.AddWatchCount(v)
	})
}

// UpdateWatchCount sets the "watch_count" field to the value that was provided on create.
func (u *RepoUpsertOne) UpdateWatchCount() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateWatchCount()
	})
}

// ClearWatchCount clears the value of the "watch_count" field.
func (u *RepoUpsertOne) ClearWatchCount() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.ClearWatchCount()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RepoUpsertOne) SetUpdatedAt(v time.Time) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RepoUpsertOne) UpdateUpdatedAt() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *RepoUpsertOne) ClearUpdatedAt() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetTopicID sets the "topic_id" field.
func (u *RepoUpsertOne) SetTopicID(v int) *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.SetTopicID(v)
	})
}

// UpdateTopicID sets the "topic_id" field to the value that was provided on create.
func (u *RepoUpsertOne) UpdateTopicID() *RepoUpsertOne {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateTopicID()
	})
}

// Exec executes the query.
func (u *RepoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RepoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RepoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RepoUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RepoUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RepoCreateBulk is the builder for creating many Repo entities in bulk.
type RepoCreateBulk struct {
	config
	builders []*RepoCreate
	conflict []sql.ConflictOption
}

// Save creates the Repo entities in the database.
func (rcb *RepoCreateBulk) Save(ctx context.Context) ([]*Repo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Repo, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RepoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RepoCreateBulk) SaveX(ctx context.Context) []*Repo {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RepoCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RepoCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Repo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RepoUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (rcb *RepoCreateBulk) OnConflict(opts ...sql.ConflictOption) *RepoUpsertBulk {
	rcb.conflict = opts
	return &RepoUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Repo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (rcb *RepoCreateBulk) OnConflictColumns(columns ...string) *RepoUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RepoUpsertBulk{
		create: rcb,
	}
}

// RepoUpsertBulk is the builder for "upsert"-ing
// a bulk of Repo nodes.
type RepoUpsertBulk struct {
	create *RepoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Repo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *RepoUpsertBulk) UpdateNewValues() *RepoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Repo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *RepoUpsertBulk) Ignore() *RepoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RepoUpsertBulk) DoNothing() *RepoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RepoCreateBulk.OnConflict
// documentation for more info.
func (u *RepoUpsertBulk) Update(set func(*RepoUpsert)) *RepoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RepoUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *RepoUpsertBulk) SetName(v string) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RepoUpsertBulk) UpdateName() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *RepoUpsertBulk) SetDescription(v string) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RepoUpsertBulk) UpdateDescription() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateDescription()
	})
}

// SetURL sets the "url" field.
func (u *RepoUpsertBulk) SetURL(v string) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *RepoUpsertBulk) UpdateURL() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateURL()
	})
}

// SetSubTopic sets the "sub_topic" field.
func (u *RepoUpsertBulk) SetSubTopic(v string) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.SetSubTopic(v)
	})
}

// UpdateSubTopic sets the "sub_topic" field to the value that was provided on create.
func (u *RepoUpsertBulk) UpdateSubTopic() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateSubTopic()
	})
}

// SetType sets the "type" field.
func (u *RepoUpsertBulk) SetType(v repo.Type) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *RepoUpsertBulk) UpdateType() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateType()
	})
}

// SetStarCount sets the "star_count" field.
func (u *RepoUpsertBulk) SetStarCount(v int) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.SetStarCount(v)
	})
}

// AddStarCount adds v to the "star_count" field.
func (u *RepoUpsertBulk) AddStarCount(v int) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.AddStarCount(v)
	})
}

// UpdateStarCount sets the "star_count" field to the value that was provided on create.
func (u *RepoUpsertBulk) UpdateStarCount() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateStarCount()
	})
}

// ClearStarCount clears the value of the "star_count" field.
func (u *RepoUpsertBulk) ClearStarCount() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.ClearStarCount()
	})
}

// SetForkCount sets the "fork_count" field.
func (u *RepoUpsertBulk) SetForkCount(v int) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.SetForkCount(v)
	})
}

// AddForkCount adds v to the "fork_count" field.
func (u *RepoUpsertBulk) AddForkCount(v int) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.AddForkCount(v)
	})
}

// UpdateForkCount sets the "fork_count" field to the value that was provided on create.
func (u *RepoUpsertBulk) UpdateForkCount() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateForkCount()
	})
}

// ClearForkCount clears the value of the "fork_count" field.
func (u *RepoUpsertBulk) ClearForkCount() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.ClearForkCount()
	})
}

// SetWatchCount sets the "watch_count" field.
func (u *RepoUpsertBulk) SetWatchCount(v int) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.SetWatchCount(v)
	})
}

// AddWatchCount adds v to the "watch_count" field.
func (u *RepoUpsertBulk) AddWatchCount(v int) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.AddWatchCount(v)
	})
}

// UpdateWatchCount sets the "watch_count" field to the value that was provided on create.
func (u *RepoUpsertBulk) UpdateWatchCount() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateWatchCount()
	})
}

// ClearWatchCount clears the value of the "watch_count" field.
func (u *RepoUpsertBulk) ClearWatchCount() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.ClearWatchCount()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RepoUpsertBulk) SetUpdatedAt(v time.Time) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RepoUpsertBulk) UpdateUpdatedAt() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *RepoUpsertBulk) ClearUpdatedAt() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetTopicID sets the "topic_id" field.
func (u *RepoUpsertBulk) SetTopicID(v int) *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.SetTopicID(v)
	})
}

// UpdateTopicID sets the "topic_id" field to the value that was provided on create.
func (u *RepoUpsertBulk) UpdateTopicID() *RepoUpsertBulk {
	return u.Update(func(s *RepoUpsert) {
		s.UpdateTopicID()
	})
}

// Exec executes the query.
func (u *RepoUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RepoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RepoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RepoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
