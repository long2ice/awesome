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
	"github.com/long2ice/awesome/ent/predicate"
	"github.com/long2ice/awesome/ent/repo"
	"github.com/long2ice/awesome/ent/topic"
)

// RepoUpdate is the builder for updating Repo entities.
type RepoUpdate struct {
	config
	hooks    []Hook
	mutation *RepoMutation
}

// Where appends a list predicates to the RepoUpdate builder.
func (ru *RepoUpdate) Where(ps ...predicate.Repo) *RepoUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RepoUpdate) SetName(s string) *RepoUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetDescription sets the "description" field.
func (ru *RepoUpdate) SetDescription(s string) *RepoUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetURL sets the "url" field.
func (ru *RepoUpdate) SetURL(s string) *RepoUpdate {
	ru.mutation.SetURL(s)
	return ru
}

// SetSubTopic sets the "sub_topic" field.
func (ru *RepoUpdate) SetSubTopic(s string) *RepoUpdate {
	ru.mutation.SetSubTopic(s)
	return ru
}

// SetType sets the "type" field.
func (ru *RepoUpdate) SetType(r repo.Type) *RepoUpdate {
	ru.mutation.SetType(r)
	return ru
}

// SetStarCount sets the "star_count" field.
func (ru *RepoUpdate) SetStarCount(i int) *RepoUpdate {
	ru.mutation.ResetStarCount()
	ru.mutation.SetStarCount(i)
	return ru
}

// SetNillableStarCount sets the "star_count" field if the given value is not nil.
func (ru *RepoUpdate) SetNillableStarCount(i *int) *RepoUpdate {
	if i != nil {
		ru.SetStarCount(*i)
	}
	return ru
}

// AddStarCount adds i to the "star_count" field.
func (ru *RepoUpdate) AddStarCount(i int) *RepoUpdate {
	ru.mutation.AddStarCount(i)
	return ru
}

// ClearStarCount clears the value of the "star_count" field.
func (ru *RepoUpdate) ClearStarCount() *RepoUpdate {
	ru.mutation.ClearStarCount()
	return ru
}

// SetForkCount sets the "fork_count" field.
func (ru *RepoUpdate) SetForkCount(i int) *RepoUpdate {
	ru.mutation.ResetForkCount()
	ru.mutation.SetForkCount(i)
	return ru
}

// SetNillableForkCount sets the "fork_count" field if the given value is not nil.
func (ru *RepoUpdate) SetNillableForkCount(i *int) *RepoUpdate {
	if i != nil {
		ru.SetForkCount(*i)
	}
	return ru
}

// AddForkCount adds i to the "fork_count" field.
func (ru *RepoUpdate) AddForkCount(i int) *RepoUpdate {
	ru.mutation.AddForkCount(i)
	return ru
}

// ClearForkCount clears the value of the "fork_count" field.
func (ru *RepoUpdate) ClearForkCount() *RepoUpdate {
	ru.mutation.ClearForkCount()
	return ru
}

// SetWatchCount sets the "watch_count" field.
func (ru *RepoUpdate) SetWatchCount(i int) *RepoUpdate {
	ru.mutation.ResetWatchCount()
	ru.mutation.SetWatchCount(i)
	return ru
}

// SetNillableWatchCount sets the "watch_count" field if the given value is not nil.
func (ru *RepoUpdate) SetNillableWatchCount(i *int) *RepoUpdate {
	if i != nil {
		ru.SetWatchCount(*i)
	}
	return ru
}

// AddWatchCount adds i to the "watch_count" field.
func (ru *RepoUpdate) AddWatchCount(i int) *RepoUpdate {
	ru.mutation.AddWatchCount(i)
	return ru
}

// ClearWatchCount clears the value of the "watch_count" field.
func (ru *RepoUpdate) ClearWatchCount() *RepoUpdate {
	ru.mutation.ClearWatchCount()
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RepoUpdate) SetUpdatedAt(t time.Time) *RepoUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ru *RepoUpdate) SetNillableUpdatedAt(t *time.Time) *RepoUpdate {
	if t != nil {
		ru.SetUpdatedAt(*t)
	}
	return ru
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ru *RepoUpdate) ClearUpdatedAt() *RepoUpdate {
	ru.mutation.ClearUpdatedAt()
	return ru
}

// SetTopicID sets the "topic_id" field.
func (ru *RepoUpdate) SetTopicID(i int) *RepoUpdate {
	ru.mutation.SetTopicID(i)
	return ru
}

// SetTopicsID sets the "topics" edge to the Topic entity by ID.
func (ru *RepoUpdate) SetTopicsID(id int) *RepoUpdate {
	ru.mutation.SetTopicsID(id)
	return ru
}

// SetTopics sets the "topics" edge to the Topic entity.
func (ru *RepoUpdate) SetTopics(t *Topic) *RepoUpdate {
	return ru.SetTopicsID(t.ID)
}

// Mutation returns the RepoMutation object of the builder.
func (ru *RepoUpdate) Mutation() *RepoMutation {
	return ru.mutation
}

// ClearTopics clears the "topics" edge to the Topic entity.
func (ru *RepoUpdate) ClearTopics() *RepoUpdate {
	ru.mutation.ClearTopics()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RepoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RepoUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RepoUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RepoUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RepoUpdate) check() error {
	if v, ok := ru.mutation.GetType(); ok {
		if err := repo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Repo.type": %w`, err)}
		}
	}
	if _, ok := ru.mutation.TopicsID(); ru.mutation.TopicsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Repo.topics"`)
	}
	return nil
}

func (ru *RepoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repo.Table,
			Columns: repo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repo.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldName,
		})
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldDescription,
		})
	}
	if value, ok := ru.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldURL,
		})
	}
	if value, ok := ru.mutation.SubTopic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldSubTopic,
		})
	}
	if value, ok := ru.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: repo.FieldType,
		})
	}
	if value, ok := ru.mutation.StarCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldStarCount,
		})
	}
	if value, ok := ru.mutation.AddedStarCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldStarCount,
		})
	}
	if ru.mutation.StarCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: repo.FieldStarCount,
		})
	}
	if value, ok := ru.mutation.ForkCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldForkCount,
		})
	}
	if value, ok := ru.mutation.AddedForkCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldForkCount,
		})
	}
	if ru.mutation.ForkCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: repo.FieldForkCount,
		})
	}
	if value, ok := ru.mutation.WatchCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldWatchCount,
		})
	}
	if value, ok := ru.mutation.AddedWatchCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldWatchCount,
		})
	}
	if ru.mutation.WatchCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: repo.FieldWatchCount,
		})
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldUpdatedAt,
		})
	}
	if ru.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: repo.FieldUpdatedAt,
		})
	}
	if ru.mutation.TopicsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.TopicsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RepoUpdateOne is the builder for updating a single Repo entity.
type RepoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RepoMutation
}

// SetName sets the "name" field.
func (ruo *RepoUpdateOne) SetName(s string) *RepoUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RepoUpdateOne) SetDescription(s string) *RepoUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetURL sets the "url" field.
func (ruo *RepoUpdateOne) SetURL(s string) *RepoUpdateOne {
	ruo.mutation.SetURL(s)
	return ruo
}

// SetSubTopic sets the "sub_topic" field.
func (ruo *RepoUpdateOne) SetSubTopic(s string) *RepoUpdateOne {
	ruo.mutation.SetSubTopic(s)
	return ruo
}

// SetType sets the "type" field.
func (ruo *RepoUpdateOne) SetType(r repo.Type) *RepoUpdateOne {
	ruo.mutation.SetType(r)
	return ruo
}

// SetStarCount sets the "star_count" field.
func (ruo *RepoUpdateOne) SetStarCount(i int) *RepoUpdateOne {
	ruo.mutation.ResetStarCount()
	ruo.mutation.SetStarCount(i)
	return ruo
}

// SetNillableStarCount sets the "star_count" field if the given value is not nil.
func (ruo *RepoUpdateOne) SetNillableStarCount(i *int) *RepoUpdateOne {
	if i != nil {
		ruo.SetStarCount(*i)
	}
	return ruo
}

// AddStarCount adds i to the "star_count" field.
func (ruo *RepoUpdateOne) AddStarCount(i int) *RepoUpdateOne {
	ruo.mutation.AddStarCount(i)
	return ruo
}

// ClearStarCount clears the value of the "star_count" field.
func (ruo *RepoUpdateOne) ClearStarCount() *RepoUpdateOne {
	ruo.mutation.ClearStarCount()
	return ruo
}

// SetForkCount sets the "fork_count" field.
func (ruo *RepoUpdateOne) SetForkCount(i int) *RepoUpdateOne {
	ruo.mutation.ResetForkCount()
	ruo.mutation.SetForkCount(i)
	return ruo
}

// SetNillableForkCount sets the "fork_count" field if the given value is not nil.
func (ruo *RepoUpdateOne) SetNillableForkCount(i *int) *RepoUpdateOne {
	if i != nil {
		ruo.SetForkCount(*i)
	}
	return ruo
}

// AddForkCount adds i to the "fork_count" field.
func (ruo *RepoUpdateOne) AddForkCount(i int) *RepoUpdateOne {
	ruo.mutation.AddForkCount(i)
	return ruo
}

// ClearForkCount clears the value of the "fork_count" field.
func (ruo *RepoUpdateOne) ClearForkCount() *RepoUpdateOne {
	ruo.mutation.ClearForkCount()
	return ruo
}

// SetWatchCount sets the "watch_count" field.
func (ruo *RepoUpdateOne) SetWatchCount(i int) *RepoUpdateOne {
	ruo.mutation.ResetWatchCount()
	ruo.mutation.SetWatchCount(i)
	return ruo
}

// SetNillableWatchCount sets the "watch_count" field if the given value is not nil.
func (ruo *RepoUpdateOne) SetNillableWatchCount(i *int) *RepoUpdateOne {
	if i != nil {
		ruo.SetWatchCount(*i)
	}
	return ruo
}

// AddWatchCount adds i to the "watch_count" field.
func (ruo *RepoUpdateOne) AddWatchCount(i int) *RepoUpdateOne {
	ruo.mutation.AddWatchCount(i)
	return ruo
}

// ClearWatchCount clears the value of the "watch_count" field.
func (ruo *RepoUpdateOne) ClearWatchCount() *RepoUpdateOne {
	ruo.mutation.ClearWatchCount()
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RepoUpdateOne) SetUpdatedAt(t time.Time) *RepoUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ruo *RepoUpdateOne) SetNillableUpdatedAt(t *time.Time) *RepoUpdateOne {
	if t != nil {
		ruo.SetUpdatedAt(*t)
	}
	return ruo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ruo *RepoUpdateOne) ClearUpdatedAt() *RepoUpdateOne {
	ruo.mutation.ClearUpdatedAt()
	return ruo
}

// SetTopicID sets the "topic_id" field.
func (ruo *RepoUpdateOne) SetTopicID(i int) *RepoUpdateOne {
	ruo.mutation.SetTopicID(i)
	return ruo
}

// SetTopicsID sets the "topics" edge to the Topic entity by ID.
func (ruo *RepoUpdateOne) SetTopicsID(id int) *RepoUpdateOne {
	ruo.mutation.SetTopicsID(id)
	return ruo
}

// SetTopics sets the "topics" edge to the Topic entity.
func (ruo *RepoUpdateOne) SetTopics(t *Topic) *RepoUpdateOne {
	return ruo.SetTopicsID(t.ID)
}

// Mutation returns the RepoMutation object of the builder.
func (ruo *RepoUpdateOne) Mutation() *RepoMutation {
	return ruo.mutation
}

// ClearTopics clears the "topics" edge to the Topic entity.
func (ruo *RepoUpdateOne) ClearTopics() *RepoUpdateOne {
	ruo.mutation.ClearTopics()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RepoUpdateOne) Select(field string, fields ...string) *RepoUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Repo entity.
func (ruo *RepoUpdateOne) Save(ctx context.Context) (*Repo, error) {
	var (
		err  error
		node *Repo
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RepoUpdateOne) SaveX(ctx context.Context) *Repo {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RepoUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RepoUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RepoUpdateOne) check() error {
	if v, ok := ruo.mutation.GetType(); ok {
		if err := repo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Repo.type": %w`, err)}
		}
	}
	if _, ok := ruo.mutation.TopicsID(); ruo.mutation.TopicsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Repo.topics"`)
	}
	return nil
}

func (ruo *RepoUpdateOne) sqlSave(ctx context.Context) (_node *Repo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repo.Table,
			Columns: repo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repo.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Repo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repo.FieldID)
		for _, f := range fields {
			if !repo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != repo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldName,
		})
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldDescription,
		})
	}
	if value, ok := ruo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldURL,
		})
	}
	if value, ok := ruo.mutation.SubTopic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldSubTopic,
		})
	}
	if value, ok := ruo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: repo.FieldType,
		})
	}
	if value, ok := ruo.mutation.StarCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldStarCount,
		})
	}
	if value, ok := ruo.mutation.AddedStarCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldStarCount,
		})
	}
	if ruo.mutation.StarCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: repo.FieldStarCount,
		})
	}
	if value, ok := ruo.mutation.ForkCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldForkCount,
		})
	}
	if value, ok := ruo.mutation.AddedForkCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldForkCount,
		})
	}
	if ruo.mutation.ForkCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: repo.FieldForkCount,
		})
	}
	if value, ok := ruo.mutation.WatchCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldWatchCount,
		})
	}
	if value, ok := ruo.mutation.AddedWatchCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repo.FieldWatchCount,
		})
	}
	if ruo.mutation.WatchCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: repo.FieldWatchCount,
		})
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldUpdatedAt,
		})
	}
	if ruo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: repo.FieldUpdatedAt,
		})
	}
	if ruo.mutation.TopicsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.TopicsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Repo{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
