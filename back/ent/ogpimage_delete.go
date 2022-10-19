// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/haikara-dev/haikara/ent/ogpimage"
	"github.com/haikara-dev/haikara/ent/predicate"
)

// OGPImageDelete is the builder for deleting a OGPImage entity.
type OGPImageDelete struct {
	config
	hooks    []Hook
	mutation *OGPImageMutation
}

// Where appends a list predicates to the OGPImageDelete builder.
func (oid *OGPImageDelete) Where(ps ...predicate.OGPImage) *OGPImageDelete {
	oid.mutation.Where(ps...)
	return oid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oid *OGPImageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oid.hooks) == 0 {
		affected, err = oid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OGPImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oid.mutation = mutation
			affected, err = oid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oid.hooks) - 1; i >= 0; i-- {
			if oid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (oid *OGPImageDelete) ExecX(ctx context.Context) int {
	n, err := oid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oid *OGPImageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: ogpimage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ogpimage.FieldID,
			},
		},
	}
	if ps := oid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, oid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// OGPImageDeleteOne is the builder for deleting a single OGPImage entity.
type OGPImageDeleteOne struct {
	oid *OGPImageDelete
}

// Exec executes the deletion query.
func (oido *OGPImageDeleteOne) Exec(ctx context.Context) error {
	n, err := oido.oid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ogpimage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oido *OGPImageDeleteOne) ExecX(ctx context.Context) {
	oido.oid.ExecX(ctx)
}