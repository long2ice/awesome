// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/long2ice/awesome/ent/platform"
	"github.com/long2ice/awesome/ent/predicate"
	"github.com/long2ice/awesome/ent/topic"
)

// PlatformUpdate is the builder for updating Platform entities.
type PlatformUpdate struct {
	config
	hooks     []Hook
	mutation  *PlatformMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PlatformUpdate builder.
func (pu *PlatformUpdate) Where(ps ...predicate.Platform) *PlatformUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PlatformUpdate) SetName(s string) *PlatformUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetIcon sets the "icon" field.
func (pu *PlatformUpdate) SetIcon(s string) *PlatformUpdate {
	pu.mutation.SetIcon(s)
	return pu
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (pu *PlatformUpdate) SetNillableIcon(s *string) *PlatformUpdate {
	if s != nil {
		pu.SetIcon(*s)
	}
	return pu
}

// ClearIcon clears the value of the "icon" field.
func (pu *PlatformUpdate) ClearIcon() *PlatformUpdate {
	pu.mutation.ClearIcon()
	return pu
}

// AddTopicIDs adds the "topics" edge to the Topic entity by IDs.
func (pu *PlatformUpdate) AddTopicIDs(ids ...int) *PlatformUpdate {
	pu.mutation.AddTopicIDs(ids...)
	return pu
}

// AddTopics adds the "topics" edges to the Topic entity.
func (pu *PlatformUpdate) AddTopics(t ...*Topic) *PlatformUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTopicIDs(ids...)
}

// Mutation returns the PlatformMutation object of the builder.
func (pu *PlatformUpdate) Mutation() *PlatformMutation {
	return pu.mutation
}

// ClearTopics clears all "topics" edges to the Topic entity.
func (pu *PlatformUpdate) ClearTopics() *PlatformUpdate {
	pu.mutation.ClearTopics()
	return pu
}

// RemoveTopicIDs removes the "topics" edge to Topic entities by IDs.
func (pu *PlatformUpdate) RemoveTopicIDs(ids ...int) *PlatformUpdate {
	pu.mutation.RemoveTopicIDs(ids...)
	return pu
}

// RemoveTopics removes "topics" edges to Topic entities.
func (pu *PlatformUpdate) RemoveTopics(t ...*Topic) *PlatformUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTopicIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlatformUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PlatformMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlatformUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlatformUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlatformUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PlatformUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PlatformUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PlatformUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(platform.Table, platform.Columns, sqlgraph.NewFieldSpec(platform.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(platform.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Icon(); ok {
		_spec.SetField(platform.FieldIcon, field.TypeString, value)
	}
	if pu.mutation.IconCleared() {
		_spec.ClearField(platform.FieldIcon, field.TypeString)
	}
	if pu.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.TopicsTable,
			Columns: []string{platform.TopicsColumn},
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
	if nodes := pu.mutation.RemovedTopicsIDs(); len(nodes) > 0 && !pu.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.TopicsTable,
			Columns: []string{platform.TopicsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.TopicsTable,
			Columns: []string{platform.TopicsColumn},
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
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{platform.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PlatformUpdateOne is the builder for updating a single Platform entity.
type PlatformUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PlatformMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (puo *PlatformUpdateOne) SetName(s string) *PlatformUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetIcon sets the "icon" field.
func (puo *PlatformUpdateOne) SetIcon(s string) *PlatformUpdateOne {
	puo.mutation.SetIcon(s)
	return puo
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (puo *PlatformUpdateOne) SetNillableIcon(s *string) *PlatformUpdateOne {
	if s != nil {
		puo.SetIcon(*s)
	}
	return puo
}

// ClearIcon clears the value of the "icon" field.
func (puo *PlatformUpdateOne) ClearIcon() *PlatformUpdateOne {
	puo.mutation.ClearIcon()
	return puo
}

// AddTopicIDs adds the "topics" edge to the Topic entity by IDs.
func (puo *PlatformUpdateOne) AddTopicIDs(ids ...int) *PlatformUpdateOne {
	puo.mutation.AddTopicIDs(ids...)
	return puo
}

// AddTopics adds the "topics" edges to the Topic entity.
func (puo *PlatformUpdateOne) AddTopics(t ...*Topic) *PlatformUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTopicIDs(ids...)
}

// Mutation returns the PlatformMutation object of the builder.
func (puo *PlatformUpdateOne) Mutation() *PlatformMutation {
	return puo.mutation
}

// ClearTopics clears all "topics" edges to the Topic entity.
func (puo *PlatformUpdateOne) ClearTopics() *PlatformUpdateOne {
	puo.mutation.ClearTopics()
	return puo
}

// RemoveTopicIDs removes the "topics" edge to Topic entities by IDs.
func (puo *PlatformUpdateOne) RemoveTopicIDs(ids ...int) *PlatformUpdateOne {
	puo.mutation.RemoveTopicIDs(ids...)
	return puo
}

// RemoveTopics removes "topics" edges to Topic entities.
func (puo *PlatformUpdateOne) RemoveTopics(t ...*Topic) *PlatformUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTopicIDs(ids...)
}

// Where appends a list predicates to the PlatformUpdate builder.
func (puo *PlatformUpdateOne) Where(ps ...predicate.Platform) *PlatformUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlatformUpdateOne) Select(field string, fields ...string) *PlatformUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Platform entity.
func (puo *PlatformUpdateOne) Save(ctx context.Context) (*Platform, error) {
	return withHooks[*Platform, PlatformMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlatformUpdateOne) SaveX(ctx context.Context) *Platform {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlatformUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlatformUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PlatformUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PlatformUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PlatformUpdateOne) sqlSave(ctx context.Context) (_node *Platform, err error) {
	_spec := sqlgraph.NewUpdateSpec(platform.Table, platform.Columns, sqlgraph.NewFieldSpec(platform.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Platform.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, platform.FieldID)
		for _, f := range fields {
			if !platform.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != platform.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(platform.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Icon(); ok {
		_spec.SetField(platform.FieldIcon, field.TypeString, value)
	}
	if puo.mutation.IconCleared() {
		_spec.ClearField(platform.FieldIcon, field.TypeString)
	}
	if puo.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.TopicsTable,
			Columns: []string{platform.TopicsColumn},
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
	if nodes := puo.mutation.RemovedTopicsIDs(); len(nodes) > 0 && !puo.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.TopicsTable,
			Columns: []string{platform.TopicsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.TopicsTable,
			Columns: []string{platform.TopicsColumn},
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
	_spec.AddModifiers(puo.modifiers...)
	_node = &Platform{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{platform.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
