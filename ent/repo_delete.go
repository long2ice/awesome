// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/long2ice/awesome/ent/predicate"
	"github.com/long2ice/awesome/ent/repo"
)

// RepoDelete is the builder for deleting a Repo entity.
type RepoDelete struct {
	config
	hooks    []Hook
	mutation *RepoMutation
}

// Where appends a list predicates to the RepoDelete builder.
func (rd *RepoDelete) Where(ps ...predicate.Repo) *RepoDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *RepoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, RepoMutation](ctx, rd.sqlExec, rd.mutation, rd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *RepoDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *RepoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(repo.Table, sqlgraph.NewFieldSpec(repo.FieldID, field.TypeInt))
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rd.mutation.done = true
	return affected, err
}

// RepoDeleteOne is the builder for deleting a single Repo entity.
type RepoDeleteOne struct {
	rd *RepoDelete
}

// Where appends a list predicates to the RepoDelete builder.
func (rdo *RepoDeleteOne) Where(ps ...predicate.Repo) *RepoDeleteOne {
	rdo.rd.mutation.Where(ps...)
	return rdo
}

// Exec executes the deletion query.
func (rdo *RepoDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{repo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *RepoDeleteOne) ExecX(ctx context.Context) {
	if err := rdo.Exec(ctx); err != nil {
		panic(err)
	}
}
