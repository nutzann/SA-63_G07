// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/nutzann/app/ent/employee"
	"github.com/nutzann/app/ent/predicate"
)

// EmployeeDelete is the builder for deleting a Employee entity.
type EmployeeDelete struct {
	config
	hooks      []Hook
	mutation   *EmployeeMutation
	predicates []predicate.Employee
}

// Where adds a new predicate to the delete builder.
func (ed *EmployeeDelete) Where(ps ...predicate.Employee) *EmployeeDelete {
	ed.predicates = append(ed.predicates, ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *EmployeeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ed.hooks) == 0 {
		affected, err = ed.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ed.mutation = mutation
			affected, err = ed.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ed.hooks) - 1; i >= 0; i-- {
			mut = ed.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ed.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *EmployeeDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *EmployeeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: employee.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employee.FieldID,
			},
		},
	}
	if ps := ed.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
}

// EmployeeDeleteOne is the builder for deleting a single Employee entity.
type EmployeeDeleteOne struct {
	ed *EmployeeDelete
}

// Exec executes the deletion query.
func (edo *EmployeeDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{employee.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *EmployeeDeleteOne) ExecX(ctx context.Context) {
	edo.ed.ExecX(ctx)
}
