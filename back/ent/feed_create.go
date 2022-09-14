// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cubdesign/dailyfj/ent/feed"
	"github.com/cubdesign/dailyfj/ent/site"
)

// FeedCreate is the builder for creating a Feed entity.
type FeedCreate struct {
	config
	mutation *FeedMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (fc *FeedCreate) SetCreatedAt(t time.Time) *FeedCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FeedCreate) SetNillableCreatedAt(t *time.Time) *FeedCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FeedCreate) SetUpdatedAt(t time.Time) *FeedCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FeedCreate) SetNillableUpdatedAt(t *time.Time) *FeedCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetContents sets the "contents" field.
func (fc *FeedCreate) SetContents(s string) *FeedCreate {
	fc.mutation.SetContents(s)
	return fc
}

// SetSiteID sets the "site" edge to the Site entity by ID.
func (fc *FeedCreate) SetSiteID(id int) *FeedCreate {
	fc.mutation.SetSiteID(id)
	return fc
}

// SetSite sets the "site" edge to the Site entity.
func (fc *FeedCreate) SetSite(s *Site) *FeedCreate {
	return fc.SetSiteID(s.ID)
}

// Mutation returns the FeedMutation object of the builder.
func (fc *FeedCreate) Mutation() *FeedMutation {
	return fc.mutation
}

// Save creates the Feed in the database.
func (fc *FeedCreate) Save(ctx context.Context) (*Feed, error) {
	var (
		err  error
		node *Feed
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			if node, err = fc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			if fc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Feed)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FeedMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FeedCreate) SaveX(ctx context.Context) *Feed {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FeedCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FeedCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FeedCreate) defaults() {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := feed.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := feed.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FeedCreate) check() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Feed.created_at"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Feed.updated_at"`)}
	}
	if _, ok := fc.mutation.Contents(); !ok {
		return &ValidationError{Name: "contents", err: errors.New(`ent: missing required field "Feed.contents"`)}
	}
	if v, ok := fc.mutation.Contents(); ok {
		if err := feed.ContentsValidator(v); err != nil {
			return &ValidationError{Name: "contents", err: fmt.Errorf(`ent: validator failed for field "Feed.contents": %w`, err)}
		}
	}
	if _, ok := fc.mutation.SiteID(); !ok {
		return &ValidationError{Name: "site", err: errors.New(`ent: missing required edge "Feed.site"`)}
	}
	return nil
}

func (fc *FeedCreate) sqlSave(ctx context.Context) (*Feed, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fc *FeedCreate) createSpec() (*Feed, *sqlgraph.CreateSpec) {
	var (
		_node = &Feed{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: feed.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: feed.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.Contents(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldContents,
		})
		_node.Contents = value
	}
	if nodes := fc.mutation.SiteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feed.SiteTable,
			Columns: []string{feed.SiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.site_feeds = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FeedCreateBulk is the builder for creating many Feed entities in bulk.
type FeedCreateBulk struct {
	config
	builders []*FeedCreate
}

// Save creates the Feed entities in the database.
func (fcb *FeedCreateBulk) Save(ctx context.Context) ([]*Feed, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Feed, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeedMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FeedCreateBulk) SaveX(ctx context.Context) []*Feed {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FeedCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FeedCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}