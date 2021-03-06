// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/nutzann/app/ent/area"
	"github.com/nutzann/app/ent/level"
	"github.com/nutzann/app/ent/predicate"
)

// LevelUpdate is the builder for updating Level entities.
type LevelUpdate struct {
	config
	hooks      []Hook
	mutation   *LevelMutation
	predicates []predicate.Level
}

// Where adds a new predicate for the builder.
func (lu *LevelUpdate) Where(ps ...predicate.Level) *LevelUpdate {
	lu.predicates = append(lu.predicates, ps...)
	return lu
}

// SetName sets the name field.
func (lu *LevelUpdate) SetName(s string) *LevelUpdate {
	lu.mutation.SetName(s)
	return lu
}

// AddAreaIDs adds the area edge to Area by ids.
func (lu *LevelUpdate) AddAreaIDs(ids ...int) *LevelUpdate {
	lu.mutation.AddAreaIDs(ids...)
	return lu
}

// AddArea adds the area edges to Area.
func (lu *LevelUpdate) AddArea(a ...*Area) *LevelUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lu.AddAreaIDs(ids...)
}

// Mutation returns the LevelMutation object of the builder.
func (lu *LevelUpdate) Mutation() *LevelMutation {
	return lu.mutation
}

// RemoveAreaIDs removes the area edge to Area by ids.
func (lu *LevelUpdate) RemoveAreaIDs(ids ...int) *LevelUpdate {
	lu.mutation.RemoveAreaIDs(ids...)
	return lu
}

// RemoveArea removes area edges to Area.
func (lu *LevelUpdate) RemoveArea(a ...*Area) *LevelUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lu.RemoveAreaIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (lu *LevelUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LevelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LevelUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LevelUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LevelUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LevelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   level.Table,
			Columns: level.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: level.FieldID,
			},
		},
	}
	if ps := lu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: level.FieldName,
		})
	}
	if nodes := lu.mutation.RemovedAreaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   level.AreaTable,
			Columns: []string{level.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.AreaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   level.AreaTable,
			Columns: []string{level.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{level.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// LevelUpdateOne is the builder for updating a single Level entity.
type LevelUpdateOne struct {
	config
	hooks    []Hook
	mutation *LevelMutation
}

// SetName sets the name field.
func (luo *LevelUpdateOne) SetName(s string) *LevelUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// AddAreaIDs adds the area edge to Area by ids.
func (luo *LevelUpdateOne) AddAreaIDs(ids ...int) *LevelUpdateOne {
	luo.mutation.AddAreaIDs(ids...)
	return luo
}

// AddArea adds the area edges to Area.
func (luo *LevelUpdateOne) AddArea(a ...*Area) *LevelUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return luo.AddAreaIDs(ids...)
}

// Mutation returns the LevelMutation object of the builder.
func (luo *LevelUpdateOne) Mutation() *LevelMutation {
	return luo.mutation
}

// RemoveAreaIDs removes the area edge to Area by ids.
func (luo *LevelUpdateOne) RemoveAreaIDs(ids ...int) *LevelUpdateOne {
	luo.mutation.RemoveAreaIDs(ids...)
	return luo
}

// RemoveArea removes area edges to Area.
func (luo *LevelUpdateOne) RemoveArea(a ...*Area) *LevelUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return luo.RemoveAreaIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (luo *LevelUpdateOne) Save(ctx context.Context) (*Level, error) {

	var (
		err  error
		node *Level
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LevelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LevelUpdateOne) SaveX(ctx context.Context) *Level {
	l, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return l
}

// Exec executes the query on the entity.
func (luo *LevelUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LevelUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LevelUpdateOne) sqlSave(ctx context.Context) (l *Level, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   level.Table,
			Columns: level.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: level.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Level.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := luo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: level.FieldName,
		})
	}
	if nodes := luo.mutation.RemovedAreaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   level.AreaTable,
			Columns: []string{level.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.AreaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   level.AreaTable,
			Columns: []string{level.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	l = &Level{config: luo.config}
	_spec.Assign = l.assignValues
	_spec.ScanValues = l.scanValues()
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{level.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return l, nil
}
