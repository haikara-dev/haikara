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
	"github.com/cubdesign/dailyfj/ent/article"
	"github.com/cubdesign/dailyfj/ent/predicate"
	"github.com/cubdesign/dailyfj/ent/site"
)

// SiteUpdate is the builder for updating Site entities.
type SiteUpdate struct {
	config
	hooks    []Hook
	mutation *SiteMutation
}

// Where appends a list predicates to the SiteUpdate builder.
func (su *SiteUpdate) Where(ps ...predicate.Site) *SiteUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SiteUpdate) SetUpdatedAt(t time.Time) *SiteUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetName sets the "name" field.
func (su *SiteUpdate) SetName(s string) *SiteUpdate {
	su.mutation.SetName(s)
	return su
}

// SetURL sets the "url" field.
func (su *SiteUpdate) SetURL(s string) *SiteUpdate {
	su.mutation.SetURL(s)
	return su
}

// SetActive sets the "active" field.
func (su *SiteUpdate) SetActive(b bool) *SiteUpdate {
	su.mutation.SetActive(b)
	return su
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (su *SiteUpdate) SetNillableActive(b *bool) *SiteUpdate {
	if b != nil {
		su.SetActive(*b)
	}
	return su
}

// AddArticleIDs adds the "articles" edge to the Article entity by IDs.
func (su *SiteUpdate) AddArticleIDs(ids ...int) *SiteUpdate {
	su.mutation.AddArticleIDs(ids...)
	return su
}

// AddArticles adds the "articles" edges to the Article entity.
func (su *SiteUpdate) AddArticles(a ...*Article) *SiteUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.AddArticleIDs(ids...)
}

// Mutation returns the SiteMutation object of the builder.
func (su *SiteUpdate) Mutation() *SiteMutation {
	return su.mutation
}

// ClearArticles clears all "articles" edges to the Article entity.
func (su *SiteUpdate) ClearArticles() *SiteUpdate {
	su.mutation.ClearArticles()
	return su
}

// RemoveArticleIDs removes the "articles" edge to Article entities by IDs.
func (su *SiteUpdate) RemoveArticleIDs(ids ...int) *SiteUpdate {
	su.mutation.RemoveArticleIDs(ids...)
	return su
}

// RemoveArticles removes "articles" edges to Article entities.
func (su *SiteUpdate) RemoveArticles(a ...*Article) *SiteUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.RemoveArticleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SiteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SiteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SiteUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SiteUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SiteUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SiteUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := site.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SiteUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := site.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Site.name": %w`, err)}
		}
	}
	if v, ok := su.mutation.URL(); ok {
		if err := site.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Site.url": %w`, err)}
		}
	}
	return nil
}

func (su *SiteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   site.Table,
			Columns: site.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: site.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: site.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: site.FieldName,
		})
	}
	if value, ok := su.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: site.FieldURL,
		})
	}
	if value, ok := su.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: site.FieldActive,
		})
	}
	if su.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   site.ArticlesTable,
			Columns: []string{site.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedArticlesIDs(); len(nodes) > 0 && !su.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   site.ArticlesTable,
			Columns: []string{site.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ArticlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   site.ArticlesTable,
			Columns: []string{site.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{site.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SiteUpdateOne is the builder for updating a single Site entity.
type SiteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SiteMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SiteUpdateOne) SetUpdatedAt(t time.Time) *SiteUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetName sets the "name" field.
func (suo *SiteUpdateOne) SetName(s string) *SiteUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetURL sets the "url" field.
func (suo *SiteUpdateOne) SetURL(s string) *SiteUpdateOne {
	suo.mutation.SetURL(s)
	return suo
}

// SetActive sets the "active" field.
func (suo *SiteUpdateOne) SetActive(b bool) *SiteUpdateOne {
	suo.mutation.SetActive(b)
	return suo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (suo *SiteUpdateOne) SetNillableActive(b *bool) *SiteUpdateOne {
	if b != nil {
		suo.SetActive(*b)
	}
	return suo
}

// AddArticleIDs adds the "articles" edge to the Article entity by IDs.
func (suo *SiteUpdateOne) AddArticleIDs(ids ...int) *SiteUpdateOne {
	suo.mutation.AddArticleIDs(ids...)
	return suo
}

// AddArticles adds the "articles" edges to the Article entity.
func (suo *SiteUpdateOne) AddArticles(a ...*Article) *SiteUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.AddArticleIDs(ids...)
}

// Mutation returns the SiteMutation object of the builder.
func (suo *SiteUpdateOne) Mutation() *SiteMutation {
	return suo.mutation
}

// ClearArticles clears all "articles" edges to the Article entity.
func (suo *SiteUpdateOne) ClearArticles() *SiteUpdateOne {
	suo.mutation.ClearArticles()
	return suo
}

// RemoveArticleIDs removes the "articles" edge to Article entities by IDs.
func (suo *SiteUpdateOne) RemoveArticleIDs(ids ...int) *SiteUpdateOne {
	suo.mutation.RemoveArticleIDs(ids...)
	return suo
}

// RemoveArticles removes "articles" edges to Article entities.
func (suo *SiteUpdateOne) RemoveArticles(a ...*Article) *SiteUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.RemoveArticleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SiteUpdateOne) Select(field string, fields ...string) *SiteUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Site entity.
func (suo *SiteUpdateOne) Save(ctx context.Context) (*Site, error) {
	var (
		err  error
		node *Site
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SiteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Site)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SiteMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SiteUpdateOne) SaveX(ctx context.Context) *Site {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SiteUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SiteUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SiteUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := site.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SiteUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := site.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Site.name": %w`, err)}
		}
	}
	if v, ok := suo.mutation.URL(); ok {
		if err := site.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Site.url": %w`, err)}
		}
	}
	return nil
}

func (suo *SiteUpdateOne) sqlSave(ctx context.Context) (_node *Site, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   site.Table,
			Columns: site.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: site.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Site.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, site.FieldID)
		for _, f := range fields {
			if !site.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != site.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: site.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: site.FieldName,
		})
	}
	if value, ok := suo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: site.FieldURL,
		})
	}
	if value, ok := suo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: site.FieldActive,
		})
	}
	if suo.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   site.ArticlesTable,
			Columns: []string{site.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedArticlesIDs(); len(nodes) > 0 && !suo.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   site.ArticlesTable,
			Columns: []string{site.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ArticlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   site.ArticlesTable,
			Columns: []string{site.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Site{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{site.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}