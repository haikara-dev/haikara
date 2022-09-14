// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cubdesign/dailyfj/ent/feed"
	"github.com/cubdesign/dailyfj/ent/predicate"
	"github.com/cubdesign/dailyfj/ent/site"
)

// FeedUpdate is the builder for updating Feed entities.
type FeedUpdate struct {
	config
	hooks    []Hook
	mutation *FeedMutation
}

// Where appends a list predicates to the FeedUpdate builder.
func (fu *FeedUpdate) Where(ps ...predicate.Feed) *FeedUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FeedUpdate) SetUpdatedAt(t time.Time) *FeedUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetContents sets the "contents" field.
func (fu *FeedUpdate) SetContents(s string) *FeedUpdate {
	fu.mutation.SetContents(s)
	return fu
}

// SetSiteID sets the "site" edge to the Site entity by ID.
func (fu *FeedUpdate) SetSiteID(id int) *FeedUpdate {
	fu.mutation.SetSiteID(id)
	return fu
}

// SetSite sets the "site" edge to the Site entity.
func (fu *FeedUpdate) SetSite(s *Site) *FeedUpdate {
	return fu.SetSiteID(s.ID)
}

// Mutation returns the FeedMutation object of the builder.
func (fu *FeedUpdate) Mutation() *FeedMutation {
	return fu.mutation
}

// ClearSite clears the "site" edge to the Site entity.
func (fu *FeedUpdate) ClearSite() *FeedUpdate {
	fu.mutation.ClearSite()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FeedUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	fu.defaults()
	if len(fu.hooks) == 0 {
		if err = fu.check(); err != nil {
			return 0, err
		}
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fu.check(); err != nil {
				return 0, err
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FeedUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FeedUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FeedUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FeedUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := feed.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FeedUpdate) check() error {
	if v, ok := fu.mutation.Contents(); ok {
		if err := feed.ContentsValidator(v); err != nil {
			return &ValidationError{Name: "contents", err: fmt.Errorf(`ent: validator failed for field "Feed.contents": %w`, err)}
		}
	}
	if _, ok := fu.mutation.SiteID(); fu.mutation.SiteCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Feed.site"`)
	}
	return nil
}

func (fu *FeedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   feed.Table,
			Columns: feed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: feed.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldUpdatedAt,
		})
	}
	if value, ok := fu.mutation.Contents(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldContents,
		})
	}
	if fu.mutation.SiteCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.SiteIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FeedUpdateOne is the builder for updating a single Feed entity.
type FeedUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FeedMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FeedUpdateOne) SetUpdatedAt(t time.Time) *FeedUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetContents sets the "contents" field.
func (fuo *FeedUpdateOne) SetContents(s string) *FeedUpdateOne {
	fuo.mutation.SetContents(s)
	return fuo
}

// SetSiteID sets the "site" edge to the Site entity by ID.
func (fuo *FeedUpdateOne) SetSiteID(id int) *FeedUpdateOne {
	fuo.mutation.SetSiteID(id)
	return fuo
}

// SetSite sets the "site" edge to the Site entity.
func (fuo *FeedUpdateOne) SetSite(s *Site) *FeedUpdateOne {
	return fuo.SetSiteID(s.ID)
}

// Mutation returns the FeedMutation object of the builder.
func (fuo *FeedUpdateOne) Mutation() *FeedMutation {
	return fuo.mutation
}

// ClearSite clears the "site" edge to the Site entity.
func (fuo *FeedUpdateOne) ClearSite() *FeedUpdateOne {
	fuo.mutation.ClearSite()
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FeedUpdateOne) Select(field string, fields ...string) *FeedUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Feed entity.
func (fuo *FeedUpdateOne) Save(ctx context.Context) (*Feed, error) {
	var (
		err  error
		node *Feed
	)
	fuo.defaults()
	if len(fuo.hooks) == 0 {
		if err = fuo.check(); err != nil {
			return nil, err
		}
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fuo.check(); err != nil {
				return nil, err
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (fuo *FeedUpdateOne) SaveX(ctx context.Context) *Feed {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FeedUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FeedUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FeedUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := feed.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FeedUpdateOne) check() error {
	if v, ok := fuo.mutation.Contents(); ok {
		if err := feed.ContentsValidator(v); err != nil {
			return &ValidationError{Name: "contents", err: fmt.Errorf(`ent: validator failed for field "Feed.contents": %w`, err)}
		}
	}
	if _, ok := fuo.mutation.SiteID(); fuo.mutation.SiteCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Feed.site"`)
	}
	return nil
}

func (fuo *FeedUpdateOne) sqlSave(ctx context.Context) (_node *Feed, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   feed.Table,
			Columns: feed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: feed.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Feed.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feed.FieldID)
		for _, f := range fields {
			if !feed.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != feed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldUpdatedAt,
		})
	}
	if value, ok := fuo.mutation.Contents(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldContents,
		})
	}
	if fuo.mutation.SiteCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.SiteIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Feed{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}