// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationshipinfo"
	"entgo.io/ent/schema/field"
)

// RelationshipInfoUpdate is the builder for updating RelationshipInfo entities.
type RelationshipInfoUpdate struct {
	config
	hooks    []Hook
	mutation *RelationshipInfoMutation
}

// Where appends a list predicates to the RelationshipInfoUpdate builder.
func (riu *RelationshipInfoUpdate) Where(ps ...predicate.RelationshipInfo) *RelationshipInfoUpdate {
	riu.mutation.Where(ps...)
	return riu
}

// SetText sets the "text" field.
func (riu *RelationshipInfoUpdate) SetText(s string) *RelationshipInfoUpdate {
	riu.mutation.SetText(s)
	return riu
}

// Mutation returns the RelationshipInfoMutation object of the builder.
func (riu *RelationshipInfoUpdate) Mutation() *RelationshipInfoMutation {
	return riu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (riu *RelationshipInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(riu.hooks) == 0 {
		affected, err = riu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RelationshipInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			riu.mutation = mutation
			affected, err = riu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(riu.hooks) - 1; i >= 0; i-- {
			if riu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = riu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, riu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (riu *RelationshipInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := riu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (riu *RelationshipInfoUpdate) Exec(ctx context.Context) error {
	_, err := riu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riu *RelationshipInfoUpdate) ExecX(ctx context.Context) {
	if err := riu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (riu *RelationshipInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   relationshipinfo.Table,
			Columns: relationshipinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: relationshipinfo.FieldID,
			},
		},
	}
	if ps := riu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := riu.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relationshipinfo.FieldText,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, riu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{relationshipinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RelationshipInfoUpdateOne is the builder for updating a single RelationshipInfo entity.
type RelationshipInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RelationshipInfoMutation
}

// SetText sets the "text" field.
func (riuo *RelationshipInfoUpdateOne) SetText(s string) *RelationshipInfoUpdateOne {
	riuo.mutation.SetText(s)
	return riuo
}

// Mutation returns the RelationshipInfoMutation object of the builder.
func (riuo *RelationshipInfoUpdateOne) Mutation() *RelationshipInfoMutation {
	return riuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (riuo *RelationshipInfoUpdateOne) Select(field string, fields ...string) *RelationshipInfoUpdateOne {
	riuo.fields = append([]string{field}, fields...)
	return riuo
}

// Save executes the query and returns the updated RelationshipInfo entity.
func (riuo *RelationshipInfoUpdateOne) Save(ctx context.Context) (*RelationshipInfo, error) {
	var (
		err  error
		node *RelationshipInfo
	)
	if len(riuo.hooks) == 0 {
		node, err = riuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RelationshipInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			riuo.mutation = mutation
			node, err = riuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(riuo.hooks) - 1; i >= 0; i-- {
			if riuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = riuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, riuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RelationshipInfo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RelationshipInfoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (riuo *RelationshipInfoUpdateOne) SaveX(ctx context.Context) *RelationshipInfo {
	node, err := riuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (riuo *RelationshipInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := riuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riuo *RelationshipInfoUpdateOne) ExecX(ctx context.Context) {
	if err := riuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (riuo *RelationshipInfoUpdateOne) sqlSave(ctx context.Context) (_node *RelationshipInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   relationshipinfo.Table,
			Columns: relationshipinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: relationshipinfo.FieldID,
			},
		},
	}
	id, ok := riuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RelationshipInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := riuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, relationshipinfo.FieldID)
		for _, f := range fields {
			if !relationshipinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != relationshipinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := riuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := riuo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relationshipinfo.FieldText,
		})
	}
	_node = &RelationshipInfo{config: riuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, riuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{relationshipinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}