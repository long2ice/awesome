// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/long2ice/awesome/ent/predicate"
	"github.com/long2ice/awesome/ent/project"
	"github.com/long2ice/awesome/ent/topic"
	"github.com/long2ice/awesome/ent/topiccategory"
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

// SetDescription sets the "description" field.
func (tu *TopicUpdate) SetDescription(s string) *TopicUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetTopicCategoryID sets the "topic_category_id" field.
func (tu *TopicUpdate) SetTopicCategoryID(i int) *TopicUpdate {
	tu.mutation.SetTopicCategoryID(i)
	return tu
}

// SetTopiccategoryID sets the "topiccategory" edge to the TopicCategory entity by ID.
func (tu *TopicUpdate) SetTopiccategoryID(id int) *TopicUpdate {
	tu.mutation.SetTopiccategoryID(id)
	return tu
}

// SetTopiccategory sets the "topiccategory" edge to the TopicCategory entity.
func (tu *TopicUpdate) SetTopiccategory(t *TopicCategory) *TopicUpdate {
	return tu.SetTopiccategoryID(t.ID)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (tu *TopicUpdate) AddProjectIDs(ids ...int) *TopicUpdate {
	tu.mutation.AddProjectIDs(ids...)
	return tu
}

// AddProjects adds the "projects" edges to the Project entity.
func (tu *TopicUpdate) AddProjects(p ...*Project) *TopicUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.AddProjectIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tu *TopicUpdate) Mutation() *TopicMutation {
	return tu.mutation
}

// ClearTopiccategory clears the "topiccategory" edge to the TopicCategory entity.
func (tu *TopicUpdate) ClearTopiccategory() *TopicUpdate {
	tu.mutation.ClearTopiccategory()
	return tu
}

// ClearProjects clears all "projects" edges to the Project entity.
func (tu *TopicUpdate) ClearProjects() *TopicUpdate {
	tu.mutation.ClearProjects()
	return tu
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (tu *TopicUpdate) RemoveProjectIDs(ids ...int) *TopicUpdate {
	tu.mutation.RemoveProjectIDs(ids...)
	return tu
}

// RemoveProjects removes "projects" edges to Project entities.
func (tu *TopicUpdate) RemoveProjects(p ...*Project) *TopicUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.RemoveProjectIDs(ids...)
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
	if _, ok := tu.mutation.TopiccategoryID(); tu.mutation.TopiccategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Topic.topiccategory"`)
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
	if value, ok := tu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldDescription,
		})
	}
	if tu.mutation.TopiccategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.TopiccategoryTable,
			Columns: []string{topic.TopiccategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topiccategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TopiccategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.TopiccategoryTable,
			Columns: []string{topic.TopiccategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topiccategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ProjectsTable,
			Columns: []string{topic.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !tu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ProjectsTable,
			Columns: []string{topic.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ProjectsTable,
			Columns: []string{topic.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
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

// SetDescription sets the "description" field.
func (tuo *TopicUpdateOne) SetDescription(s string) *TopicUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetTopicCategoryID sets the "topic_category_id" field.
func (tuo *TopicUpdateOne) SetTopicCategoryID(i int) *TopicUpdateOne {
	tuo.mutation.SetTopicCategoryID(i)
	return tuo
}

// SetTopiccategoryID sets the "topiccategory" edge to the TopicCategory entity by ID.
func (tuo *TopicUpdateOne) SetTopiccategoryID(id int) *TopicUpdateOne {
	tuo.mutation.SetTopiccategoryID(id)
	return tuo
}

// SetTopiccategory sets the "topiccategory" edge to the TopicCategory entity.
func (tuo *TopicUpdateOne) SetTopiccategory(t *TopicCategory) *TopicUpdateOne {
	return tuo.SetTopiccategoryID(t.ID)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (tuo *TopicUpdateOne) AddProjectIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.AddProjectIDs(ids...)
	return tuo
}

// AddProjects adds the "projects" edges to the Project entity.
func (tuo *TopicUpdateOne) AddProjects(p ...*Project) *TopicUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.AddProjectIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tuo *TopicUpdateOne) Mutation() *TopicMutation {
	return tuo.mutation
}

// ClearTopiccategory clears the "topiccategory" edge to the TopicCategory entity.
func (tuo *TopicUpdateOne) ClearTopiccategory() *TopicUpdateOne {
	tuo.mutation.ClearTopiccategory()
	return tuo
}

// ClearProjects clears all "projects" edges to the Project entity.
func (tuo *TopicUpdateOne) ClearProjects() *TopicUpdateOne {
	tuo.mutation.ClearProjects()
	return tuo
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (tuo *TopicUpdateOne) RemoveProjectIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.RemoveProjectIDs(ids...)
	return tuo
}

// RemoveProjects removes "projects" edges to Project entities.
func (tuo *TopicUpdateOne) RemoveProjects(p ...*Project) *TopicUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.RemoveProjectIDs(ids...)
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
	if _, ok := tuo.mutation.TopiccategoryID(); tuo.mutation.TopiccategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Topic.topiccategory"`)
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
	if value, ok := tuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldDescription,
		})
	}
	if tuo.mutation.TopiccategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.TopiccategoryTable,
			Columns: []string{topic.TopiccategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topiccategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TopiccategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.TopiccategoryTable,
			Columns: []string{topic.TopiccategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topiccategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ProjectsTable,
			Columns: []string{topic.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !tuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ProjectsTable,
			Columns: []string{topic.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.ProjectsTable,
			Columns: []string{topic.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
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
