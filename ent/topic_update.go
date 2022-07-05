// Code generated by entc, DO NOT EDIT.

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
	"github.com/long2ice/awesome/ent/repo"
	"github.com/long2ice/awesome/ent/topic"
)

// TopicUpdate is the builder for updating Topic entities.
type TopicUpdate struct {
	config
	hooks    []Hook
	mutation *TopicMutation
}

// Where appends a list predicates to the TopicUpdate builder.
func (tu *TopicUpdate) Where(ps ...predicate.Topic) *TopicUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TopicUpdate) SetName(s string) *TopicUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetSubName sets the "sub_name" field.
func (tu *TopicUpdate) SetSubName(s string) *TopicUpdate {
	tu.mutation.SetSubName(s)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TopicUpdate) SetDescription(s string) *TopicUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetURL sets the "url" field.
func (tu *TopicUpdate) SetURL(s string) *TopicUpdate {
	tu.mutation.SetURL(s)
	return tu
}

// SetGithubURL sets the "github_url" field.
func (tu *TopicUpdate) SetGithubURL(s string) *TopicUpdate {
	tu.mutation.SetGithubURL(s)
	return tu
}

// SetPlatformID sets the "platform_id" field.
func (tu *TopicUpdate) SetPlatformID(i int) *TopicUpdate {
	tu.mutation.SetPlatformID(i)
	return tu
}

// SetPlatform sets the "platform" edge to the Platform entity.
func (tu *TopicUpdate) SetPlatform(p *Platform) *TopicUpdate {
	return tu.SetPlatformID(p.ID)
}

// AddRepoIDs adds the "repos" edge to the Repo entity by IDs.
func (tu *TopicUpdate) AddRepoIDs(ids ...int) *TopicUpdate {
	tu.mutation.AddRepoIDs(ids...)
	return tu
}

// AddRepos adds the "repos" edges to the Repo entity.
func (tu *TopicUpdate) AddRepos(r ...*Repo) *TopicUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.AddRepoIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tu *TopicUpdate) Mutation() *TopicMutation {
	return tu.mutation
}

// ClearPlatform clears the "platform" edge to the Platform entity.
func (tu *TopicUpdate) ClearPlatform() *TopicUpdate {
	tu.mutation.ClearPlatform()
	return tu
}

// ClearRepos clears all "repos" edges to the Repo entity.
func (tu *TopicUpdate) ClearRepos() *TopicUpdate {
	tu.mutation.ClearRepos()
	return tu
}

// RemoveRepoIDs removes the "repos" edge to Repo entities by IDs.
func (tu *TopicUpdate) RemoveRepoIDs(ids ...int) *TopicUpdate {
	tu.mutation.RemoveRepoIDs(ids...)
	return tu
}

// RemoveRepos removes "repos" edges to Repo entities.
func (tu *TopicUpdate) RemoveRepos(r ...*Repo) *TopicUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.RemoveRepoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TopicUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TopicUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TopicUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TopicUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TopicUpdate) check() error {
	if _, ok := tu.mutation.PlatformID(); tu.mutation.PlatformCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Topic.platform"`)
	}
	return nil
}

func (tu *TopicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topic.Table,
			Columns: topic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: topic.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldName,
		})
	}
	if value, ok := tu.mutation.SubName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldSubName,
		})
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldDescription,
		})
	}
	if value, ok := tu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldURL,
		})
	}
	if value, ok := tu.mutation.GithubURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldGithubURL,
		})
	}
	if tu.mutation.PlatformCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.PlatformTable,
			Columns: []string{topic.PlatformColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: platform.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.PlatformIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.PlatformTable,
			Columns: []string{topic.PlatformColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: platform.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ReposCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ReposTable,
			Columns: []string{topic.ReposColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedReposIDs(); len(nodes) > 0 && !tu.mutation.ReposCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ReposTable,
			Columns: []string{topic.ReposColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ReposIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ReposTable,
			Columns: []string{topic.ReposColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TopicUpdateOne is the builder for updating a single Topic entity.
type TopicUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TopicMutation
}

// SetName sets the "name" field.
func (tuo *TopicUpdateOne) SetName(s string) *TopicUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetSubName sets the "sub_name" field.
func (tuo *TopicUpdateOne) SetSubName(s string) *TopicUpdateOne {
	tuo.mutation.SetSubName(s)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TopicUpdateOne) SetDescription(s string) *TopicUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetURL sets the "url" field.
func (tuo *TopicUpdateOne) SetURL(s string) *TopicUpdateOne {
	tuo.mutation.SetURL(s)
	return tuo
}

// SetGithubURL sets the "github_url" field.
func (tuo *TopicUpdateOne) SetGithubURL(s string) *TopicUpdateOne {
	tuo.mutation.SetGithubURL(s)
	return tuo
}

// SetPlatformID sets the "platform_id" field.
func (tuo *TopicUpdateOne) SetPlatformID(i int) *TopicUpdateOne {
	tuo.mutation.SetPlatformID(i)
	return tuo
}

// SetPlatform sets the "platform" edge to the Platform entity.
func (tuo *TopicUpdateOne) SetPlatform(p *Platform) *TopicUpdateOne {
	return tuo.SetPlatformID(p.ID)
}

// AddRepoIDs adds the "repos" edge to the Repo entity by IDs.
func (tuo *TopicUpdateOne) AddRepoIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.AddRepoIDs(ids...)
	return tuo
}

// AddRepos adds the "repos" edges to the Repo entity.
func (tuo *TopicUpdateOne) AddRepos(r ...*Repo) *TopicUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.AddRepoIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tuo *TopicUpdateOne) Mutation() *TopicMutation {
	return tuo.mutation
}

// ClearPlatform clears the "platform" edge to the Platform entity.
func (tuo *TopicUpdateOne) ClearPlatform() *TopicUpdateOne {
	tuo.mutation.ClearPlatform()
	return tuo
}

// ClearRepos clears all "repos" edges to the Repo entity.
func (tuo *TopicUpdateOne) ClearRepos() *TopicUpdateOne {
	tuo.mutation.ClearRepos()
	return tuo
}

// RemoveRepoIDs removes the "repos" edge to Repo entities by IDs.
func (tuo *TopicUpdateOne) RemoveRepoIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.RemoveRepoIDs(ids...)
	return tuo
}

// RemoveRepos removes "repos" edges to Repo entities.
func (tuo *TopicUpdateOne) RemoveRepos(r ...*Repo) *TopicUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.RemoveRepoIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TopicUpdateOne) Select(field string, fields ...string) *TopicUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Topic entity.
func (tuo *TopicUpdateOne) Save(ctx context.Context) (*Topic, error) {
	var (
		err  error
		node *Topic
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TopicUpdateOne) SaveX(ctx context.Context) *Topic {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TopicUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TopicUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TopicUpdateOne) check() error {
	if _, ok := tuo.mutation.PlatformID(); tuo.mutation.PlatformCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Topic.platform"`)
	}
	return nil
}

func (tuo *TopicUpdateOne) sqlSave(ctx context.Context) (_node *Topic, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topic.Table,
			Columns: topic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: topic.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Topic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topic.FieldID)
		for _, f := range fields {
			if !topic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != topic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldName,
		})
	}
	if value, ok := tuo.mutation.SubName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldSubName,
		})
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldDescription,
		})
	}
	if value, ok := tuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldURL,
		})
	}
	if value, ok := tuo.mutation.GithubURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldGithubURL,
		})
	}
	if tuo.mutation.PlatformCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.PlatformTable,
			Columns: []string{topic.PlatformColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: platform.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.PlatformIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.PlatformTable,
			Columns: []string{topic.PlatformColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: platform.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ReposCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ReposTable,
			Columns: []string{topic.ReposColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedReposIDs(); len(nodes) > 0 && !tuo.mutation.ReposCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ReposTable,
			Columns: []string{topic.ReposColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ReposIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ReposTable,
			Columns: []string{topic.ReposColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Topic{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
