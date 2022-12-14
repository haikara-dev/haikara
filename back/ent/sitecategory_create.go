// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/haikara-dev/haikara/ent/site"
	"github.com/haikara-dev/haikara/ent/sitecategory"
)

// SiteCategoryCreate is the builder for creating a SiteCategory entity.
type SiteCategoryCreate struct {
	config
	mutation *SiteCategoryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (scc *SiteCategoryCreate) SetCreatedAt(t time.Time) *SiteCategoryCreate {
	scc.mutation.SetCreatedAt(t)
	return scc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (scc *SiteCategoryCreate) SetNillableCreatedAt(t *time.Time) *SiteCategoryCreate {
	if t != nil {
		scc.SetCreatedAt(*t)
	}
	return scc
}

// SetUpdatedAt sets the "updated_at" field.
func (scc *SiteCategoryCreate) SetUpdatedAt(t time.Time) *SiteCategoryCreate {
	scc.mutation.SetUpdatedAt(t)
	return scc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (scc *SiteCategoryCreate) SetNillableUpdatedAt(t *time.Time) *SiteCategoryCreate {
	if t != nil {
		scc.SetUpdatedAt(*t)
	}
	return scc
}

// SetLabel sets the "label" field.
func (scc *SiteCategoryCreate) SetLabel(s string) *SiteCategoryCreate {
	scc.mutation.SetLabel(s)
	return scc
}

// AddSiteIDs adds the "sites" edge to the Site entity by IDs.
func (scc *SiteCategoryCreate) AddSiteIDs(ids ...int) *SiteCategoryCreate {
	scc.mutation.AddSiteIDs(ids...)
	return scc
}

// AddSites adds the "sites" edges to the Site entity.
func (scc *SiteCategoryCreate) AddSites(s ...*Site) *SiteCategoryCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scc.AddSiteIDs(ids...)
}

// Mutation returns the SiteCategoryMutation object of the builder.
func (scc *SiteCategoryCreate) Mutation() *SiteCategoryMutation {
	return scc.mutation
}

// Save creates the SiteCategory in the database.
func (scc *SiteCategoryCreate) Save(ctx context.Context) (*SiteCategory, error) {
	var (
		err  error
		node *SiteCategory
	)
	scc.defaults()
	if len(scc.hooks) == 0 {
		if err = scc.check(); err != nil {
			return nil, err
		}
		node, err = scc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SiteCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = scc.check(); err != nil {
				return nil, err
			}
			scc.mutation = mutation
			if node, err = scc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(scc.hooks) - 1; i >= 0; i-- {
			if scc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, scc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SiteCategory)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SiteCategoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (scc *SiteCategoryCreate) SaveX(ctx context.Context) *SiteCategory {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *SiteCategoryCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *SiteCategoryCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *SiteCategoryCreate) defaults() {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		v := sitecategory.DefaultCreatedAt()
		scc.mutation.SetCreatedAt(v)
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		v := sitecategory.DefaultUpdatedAt()
		scc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *SiteCategoryCreate) check() error {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SiteCategory.created_at"`)}
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SiteCategory.updated_at"`)}
	}
	if _, ok := scc.mutation.Label(); !ok {
		return &ValidationError{Name: "label", err: errors.New(`ent: missing required field "SiteCategory.label"`)}
	}
	if v, ok := scc.mutation.Label(); ok {
		if err := sitecategory.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "SiteCategory.label": %w`, err)}
		}
	}
	return nil
}

func (scc *SiteCategoryCreate) sqlSave(ctx context.Context) (*SiteCategory, error) {
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (scc *SiteCategoryCreate) createSpec() (*SiteCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &SiteCategory{config: scc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sitecategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sitecategory.FieldID,
			},
		}
	)
	if value, ok := scc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sitecategory.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := scc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sitecategory.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := scc.mutation.Label(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sitecategory.FieldLabel,
		})
		_node.Label = value
	}
	if nodes := scc.mutation.SitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sitecategory.SitesTable,
			Columns: sitecategory.SitesPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SiteCategoryCreateBulk is the builder for creating many SiteCategory entities in bulk.
type SiteCategoryCreateBulk struct {
	config
	builders []*SiteCategoryCreate
}

// Save creates the SiteCategory entities in the database.
func (sccb *SiteCategoryCreateBulk) Save(ctx context.Context) ([]*SiteCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*SiteCategory, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SiteCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *SiteCategoryCreateBulk) SaveX(ctx context.Context) []*SiteCategory {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *SiteCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *SiteCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}
