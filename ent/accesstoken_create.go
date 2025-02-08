// Code generated by ent, DO NOT EDIT.

package ent

import (
	"AltTube-Go/ent/accesstoken"
	"AltTube-Go/ent/refreshtoken"
	"AltTube-Go/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccessTokenCreate is the builder for creating a AccessToken entity.
type AccessTokenCreate struct {
	config
	mutation *AccessTokenMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (atc *AccessTokenCreate) SetCreateTime(t time.Time) *AccessTokenCreate {
	atc.mutation.SetCreateTime(t)
	return atc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (atc *AccessTokenCreate) SetNillableCreateTime(t *time.Time) *AccessTokenCreate {
	if t != nil {
		atc.SetCreateTime(*t)
	}
	return atc
}

// SetUpdateTime sets the "update_time" field.
func (atc *AccessTokenCreate) SetUpdateTime(t time.Time) *AccessTokenCreate {
	atc.mutation.SetUpdateTime(t)
	return atc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (atc *AccessTokenCreate) SetNillableUpdateTime(t *time.Time) *AccessTokenCreate {
	if t != nil {
		atc.SetUpdateTime(*t)
	}
	return atc
}

// SetToken sets the "token" field.
func (atc *AccessTokenCreate) SetToken(s string) *AccessTokenCreate {
	atc.mutation.SetToken(s)
	return atc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (atc *AccessTokenCreate) SetNillableToken(s *string) *AccessTokenCreate {
	if s != nil {
		atc.SetToken(*s)
	}
	return atc
}

// SetUserID sets the "user_id" field.
func (atc *AccessTokenCreate) SetUserID(s string) *AccessTokenCreate {
	atc.mutation.SetUserID(s)
	return atc
}

// SetRefreshTokenID sets the "refresh_token_id" field.
func (atc *AccessTokenCreate) SetRefreshTokenID(u uint) *AccessTokenCreate {
	atc.mutation.SetRefreshTokenID(u)
	return atc
}

// SetExpiry sets the "expiry" field.
func (atc *AccessTokenCreate) SetExpiry(t time.Time) *AccessTokenCreate {
	atc.mutation.SetExpiry(t)
	return atc
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (atc *AccessTokenCreate) SetNillableExpiry(t *time.Time) *AccessTokenCreate {
	if t != nil {
		atc.SetExpiry(*t)
	}
	return atc
}

// SetID sets the "id" field.
func (atc *AccessTokenCreate) SetID(u uint) *AccessTokenCreate {
	atc.mutation.SetID(u)
	return atc
}

// SetUser sets the "user" edge to the User entity.
func (atc *AccessTokenCreate) SetUser(u *User) *AccessTokenCreate {
	return atc.SetUserID(u.ID)
}

// SetRefreshToken sets the "refresh_token" edge to the RefreshToken entity.
func (atc *AccessTokenCreate) SetRefreshToken(r *RefreshToken) *AccessTokenCreate {
	return atc.SetRefreshTokenID(r.ID)
}

// Mutation returns the AccessTokenMutation object of the builder.
func (atc *AccessTokenCreate) Mutation() *AccessTokenMutation {
	return atc.mutation
}

// Save creates the AccessToken in the database.
func (atc *AccessTokenCreate) Save(ctx context.Context) (*AccessToken, error) {
	atc.defaults()
	return withHooks(ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *AccessTokenCreate) SaveX(ctx context.Context) *AccessToken {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *AccessTokenCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *AccessTokenCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *AccessTokenCreate) defaults() {
	if _, ok := atc.mutation.CreateTime(); !ok {
		v := accesstoken.DefaultCreateTime()
		atc.mutation.SetCreateTime(v)
	}
	if _, ok := atc.mutation.UpdateTime(); !ok {
		v := accesstoken.DefaultUpdateTime()
		atc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *AccessTokenCreate) check() error {
	if _, ok := atc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "AccessToken.create_time"`)}
	}
	if _, ok := atc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "AccessToken.update_time"`)}
	}
	if _, ok := atc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "AccessToken.user_id"`)}
	}
	if _, ok := atc.mutation.RefreshTokenID(); !ok {
		return &ValidationError{Name: "refresh_token_id", err: errors.New(`ent: missing required field "AccessToken.refresh_token_id"`)}
	}
	if len(atc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "AccessToken.user"`)}
	}
	if len(atc.mutation.RefreshTokenIDs()) == 0 {
		return &ValidationError{Name: "refresh_token", err: errors.New(`ent: missing required edge "AccessToken.refresh_token"`)}
	}
	return nil
}

func (atc *AccessTokenCreate) sqlSave(ctx context.Context) (*AccessToken, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint(id)
	}
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *AccessTokenCreate) createSpec() (*AccessToken, *sqlgraph.CreateSpec) {
	var (
		_node = &AccessToken{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(accesstoken.Table, sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeUint))
	)
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := atc.mutation.CreateTime(); ok {
		_spec.SetField(accesstoken.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := atc.mutation.UpdateTime(); ok {
		_spec.SetField(accesstoken.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := atc.mutation.Token(); ok {
		_spec.SetField(accesstoken.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := atc.mutation.Expiry(); ok {
		_spec.SetField(accesstoken.FieldExpiry, field.TypeTime, value)
		_node.Expiry = value
	}
	if nodes := atc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.UserTable,
			Columns: []string{accesstoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atc.mutation.RefreshTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.RefreshTokenTable,
			Columns: []string{accesstoken.RefreshTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RefreshTokenID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccessTokenCreateBulk is the builder for creating many AccessToken entities in bulk.
type AccessTokenCreateBulk struct {
	config
	err      error
	builders []*AccessTokenCreate
}

// Save creates the AccessToken entities in the database.
func (atcb *AccessTokenCreateBulk) Save(ctx context.Context) ([]*AccessToken, error) {
	if atcb.err != nil {
		return nil, atcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*AccessToken, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccessTokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint(id)
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *AccessTokenCreateBulk) SaveX(ctx context.Context) []*AccessToken {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *AccessTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *AccessTokenCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}
