// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/nutzann/app/ent/area"
	"github.com/nutzann/app/ent/employee"
)

// EmployeeCreate is the builder for creating a Employee entity.
type EmployeeCreate struct {
	config
	mutation *EmployeeMutation
	hooks    []Hook
}

// SetEmail sets the Email field.
func (ec *EmployeeCreate) SetEmail(s string) *EmployeeCreate {
	ec.mutation.SetEmail(s)
	return ec
}

// SetName sets the Name field.
func (ec *EmployeeCreate) SetName(s string) *EmployeeCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetUserID sets the User_id field.
func (ec *EmployeeCreate) SetUserID(s string) *EmployeeCreate {
	ec.mutation.SetUserID(s)
	return ec
}

// AddAreaIDs adds the area edge to Area by ids.
func (ec *EmployeeCreate) AddAreaIDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddAreaIDs(ids...)
	return ec
}

// AddArea adds the area edges to Area.
func (ec *EmployeeCreate) AddArea(a ...*Area) *EmployeeCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ec.AddAreaIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (ec *EmployeeCreate) Mutation() *EmployeeMutation {
	return ec.mutation
}

// Save creates the Employee in the database.
func (ec *EmployeeCreate) Save(ctx context.Context) (*Employee, error) {
	if _, ok := ec.mutation.Email(); !ok {
		return nil, &ValidationError{Name: "Email", err: errors.New("ent: missing required field \"Email\"")}
	}
	if _, ok := ec.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "Name", err: errors.New("ent: missing required field \"Name\"")}
	}
	if _, ok := ec.mutation.UserID(); !ok {
		return nil, &ValidationError{Name: "User_id", err: errors.New("ent: missing required field \"User_id\"")}
	}
	var (
		err  error
		node *Employee
	)
	if len(ec.hooks) == 0 {
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ec.mutation = mutation
			node, err = ec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EmployeeCreate) SaveX(ctx context.Context) *Employee {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ec *EmployeeCreate) sqlSave(ctx context.Context) (*Employee, error) {
	e, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	e.ID = int(id)
	return e, nil
}

func (ec *EmployeeCreate) createSpec() (*Employee, *sqlgraph.CreateSpec) {
	var (
		e     = &Employee{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: employee.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employee.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmail,
		})
		e.Email = value
	}
	if value, ok := ec.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldName,
		})
		e.Name = value
	}
	if value, ok := ec.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldUserID,
		})
		e.UserID = value
	}
	if nodes := ec.mutation.AreaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.AreaTable,
			Columns: []string{employee.AreaColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return e, _spec
}
