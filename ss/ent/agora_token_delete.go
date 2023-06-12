// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/agora_token"
	"CSBackendTmp/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AgoraTokenDelete is the builder for deleting a Agora_token entity.
type AgoraTokenDelete struct {
	config
	hooks    []Hook
	mutation *AgoraTokenMutation
}

// Where appends a list predicates to the AgoraTokenDelete builder.
func (atd *AgoraTokenDelete) Where(ps ...predicate.Agora_token) *AgoraTokenDelete {
	atd.mutation.Where(ps...)
	return atd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atd *AgoraTokenDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atd.hooks) == 0 {
		affected, err = atd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgoraTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atd.mutation = mutation
			affected, err = atd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atd.hooks) - 1; i >= 0; i-- {
			if atd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (atd *AgoraTokenDelete) ExecX(ctx context.Context) int {
	n, err := atd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atd *AgoraTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: agora_token.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: agora_token.FieldID,
			},
		},
	}
	if ps := atd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, atd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AgoraTokenDeleteOne is the builder for deleting a single Agora_token entity.
type AgoraTokenDeleteOne struct {
	atd *AgoraTokenDelete
}

// Exec executes the deletion query.
func (atdo *AgoraTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := atdo.atd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{agora_token.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atdo *AgoraTokenDeleteOne) ExecX(ctx context.Context) {
	atdo.atd.ExecX(ctx)
}