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
	"github.com/long2ice/awesome/ent/topic"
)

// PlatformCreate is the builder for creating a Platform entity.
type PlatformCreate struct {
	config
	mutation *PlatformMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (pc *PlatformCreate) SetName(s string) *PlatformCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetIcon sets the "icon" field.
func (pc *PlatformCreate) SetIcon(s string) *PlatformCreate {
	pc.mutation.SetIcon(s)
	return pc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (pc *PlatformCreate) SetNillableIcon(s *string) *PlatformCreate {
	if s != nil {
		pc.SetIcon(*s)
	}
	return pc
}

// AddTopicIDs adds the "topics" edge to the Topic entity by IDs.
func (pc *PlatformCreate) AddTopicIDs(ids ...int) *PlatformCreate {
	pc.mutation.AddTopicIDs(ids...)
	return pc
}

// AddTopics adds the "topics" edges to the Topic entity.
func (pc *PlatformCreate) AddTopics(t ...*Topic) *PlatformCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTopicIDs(ids...)
}

// Mutation returns the PlatformMutation object of the builder.
func (pc *PlatformCreate) Mutation() *PlatformMutation {
	return pc.mutation
}

// Save creates the Platform in the database.
func (pc *PlatformCreate) Save(ctx context.Context) (*Platform, error) {
	return withHooks[*Platform, PlatformMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlatformCreate) SaveX(ctx context.Context) *Platform {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlatformCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlatformCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlatformCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Platform.name"`)}
	}
	return nil
}

func (pc *PlatformCreate) sqlSave(ctx context.Context) (*Platform, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PlatformCreate) createSpec() (*Platform, *sqlgraph.CreateSpec) {
	var (
		_node = &Platform{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(platform.Table, sqlgraph.NewFieldSpec(platform.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(platform.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Icon(); ok {
		_spec.SetField(platform.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if nodes := pc.mutation.TopicsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Platform.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlatformUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pc *PlatformCreate) OnConflict(opts ...sql.ConflictOption) *PlatformUpsertOne {
	pc.conflict = opts
	return &PlatformUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Platform.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PlatformCreate) OnConflictColumns(columns ...string) *PlatformUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PlatformUpsertOne{
		create: pc,
	}
}

type (
	// PlatformUpsertOne is the builder for "upsert"-ing
	//  one Platform node.
	PlatformUpsertOne struct {
		create *PlatformCreate
	}

	// PlatformUpsert is the "OnConflict" setter.
	PlatformUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *PlatformUpsert) SetName(v string) *PlatformUpsert {
	u.Set(platform.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlatformUpsert) UpdateName() *PlatformUpsert {
	u.SetExcluded(platform.FieldName)
	return u
}

// SetIcon sets the "icon" field.
func (u *PlatformUpsert) SetIcon(v string) *PlatformUpsert {
	u.Set(platform.FieldIcon, v)
	return u
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *PlatformUpsert) UpdateIcon() *PlatformUpsert {
	u.SetExcluded(platform.FieldIcon)
	return u
}

// ClearIcon clears the value of the "icon" field.
func (u *PlatformUpsert) ClearIcon() *PlatformUpsert {
	u.SetNull(platform.FieldIcon)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Platform.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PlatformUpsertOne) UpdateNewValues() *PlatformUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Platform.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PlatformUpsertOne) Ignore() *PlatformUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlatformUpsertOne) DoNothing() *PlatformUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlatformCreate.OnConflict
// documentation for more info.
func (u *PlatformUpsertOne) Update(set func(*PlatformUpsert)) *PlatformUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlatformUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PlatformUpsertOne) SetName(v string) *PlatformUpsertOne {
	return u.Update(func(s *PlatformUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlatformUpsertOne) UpdateName() *PlatformUpsertOne {
	return u.Update(func(s *PlatformUpsert) {
		s.UpdateName()
	})
}

// SetIcon sets the "icon" field.
func (u *PlatformUpsertOne) SetIcon(v string) *PlatformUpsertOne {
	return u.Update(func(s *PlatformUpsert) {
		s.SetIcon(v)
	})
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *PlatformUpsertOne) UpdateIcon() *PlatformUpsertOne {
	return u.Update(func(s *PlatformUpsert) {
		s.UpdateIcon()
	})
}

// ClearIcon clears the value of the "icon" field.
func (u *PlatformUpsertOne) ClearIcon() *PlatformUpsertOne {
	return u.Update(func(s *PlatformUpsert) {
		s.ClearIcon()
	})
}

// Exec executes the query.
func (u *PlatformUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlatformCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlatformUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PlatformUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PlatformUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PlatformCreateBulk is the builder for creating many Platform entities in bulk.
type PlatformCreateBulk struct {
	config
	builders []*PlatformCreate
	conflict []sql.ConflictOption
}

// Save creates the Platform entities in the database.
func (pcb *PlatformCreateBulk) Save(ctx context.Context) ([]*Platform, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Platform, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlatformMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlatformCreateBulk) SaveX(ctx context.Context) []*Platform {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlatformCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlatformCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Platform.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlatformUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pcb *PlatformCreateBulk) OnConflict(opts ...sql.ConflictOption) *PlatformUpsertBulk {
	pcb.conflict = opts
	return &PlatformUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Platform.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PlatformCreateBulk) OnConflictColumns(columns ...string) *PlatformUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PlatformUpsertBulk{
		create: pcb,
	}
}

// PlatformUpsertBulk is the builder for "upsert"-ing
// a bulk of Platform nodes.
type PlatformUpsertBulk struct {
	create *PlatformCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Platform.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PlatformUpsertBulk) UpdateNewValues() *PlatformUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Platform.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PlatformUpsertBulk) Ignore() *PlatformUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlatformUpsertBulk) DoNothing() *PlatformUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlatformCreateBulk.OnConflict
// documentation for more info.
func (u *PlatformUpsertBulk) Update(set func(*PlatformUpsert)) *PlatformUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlatformUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PlatformUpsertBulk) SetName(v string) *PlatformUpsertBulk {
	return u.Update(func(s *PlatformUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlatformUpsertBulk) UpdateName() *PlatformUpsertBulk {
	return u.Update(func(s *PlatformUpsert) {
		s.UpdateName()
	})
}

// SetIcon sets the "icon" field.
func (u *PlatformUpsertBulk) SetIcon(v string) *PlatformUpsertBulk {
	return u.Update(func(s *PlatformUpsert) {
		s.SetIcon(v)
	})
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *PlatformUpsertBulk) UpdateIcon() *PlatformUpsertBulk {
	return u.Update(func(s *PlatformUpsert) {
		s.UpdateIcon()
	})
}

// ClearIcon clears the value of the "icon" field.
func (u *PlatformUpsertBulk) ClearIcon() *PlatformUpsertBulk {
	return u.Update(func(s *PlatformUpsert) {
		s.ClearIcon()
	})
}

// Exec executes the query.
func (u *PlatformUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PlatformCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlatformCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlatformUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
