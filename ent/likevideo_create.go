// Code generated by ent, DO NOT EDIT.

package ent

import (
	"AltTube-Go/ent/likevideo"
	"AltTube-Go/ent/user"
	"AltTube-Go/ent/video"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LikeVideoCreate is the builder for creating a LikeVideo entity.
type LikeVideoCreate struct {
	config
	mutation *LikeVideoMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (lvc *LikeVideoCreate) SetCreateTime(t time.Time) *LikeVideoCreate {
	lvc.mutation.SetCreateTime(t)
	return lvc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (lvc *LikeVideoCreate) SetNillableCreateTime(t *time.Time) *LikeVideoCreate {
	if t != nil {
		lvc.SetCreateTime(*t)
	}
	return lvc
}

// SetUpdateTime sets the "update_time" field.
func (lvc *LikeVideoCreate) SetUpdateTime(t time.Time) *LikeVideoCreate {
	lvc.mutation.SetUpdateTime(t)
	return lvc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (lvc *LikeVideoCreate) SetNillableUpdateTime(t *time.Time) *LikeVideoCreate {
	if t != nil {
		lvc.SetUpdateTime(*t)
	}
	return lvc
}

// SetUserID sets the "user_id" field.
func (lvc *LikeVideoCreate) SetUserID(s string) *LikeVideoCreate {
	lvc.mutation.SetUserID(s)
	return lvc
}

// SetVideoID sets the "video_id" field.
func (lvc *LikeVideoCreate) SetVideoID(s string) *LikeVideoCreate {
	lvc.mutation.SetVideoID(s)
	return lvc
}

// SetUser sets the "user" edge to the User entity.
func (lvc *LikeVideoCreate) SetUser(u *User) *LikeVideoCreate {
	return lvc.SetUserID(u.ID)
}

// SetVideo sets the "video" edge to the Video entity.
func (lvc *LikeVideoCreate) SetVideo(v *Video) *LikeVideoCreate {
	return lvc.SetVideoID(v.ID)
}

// Mutation returns the LikeVideoMutation object of the builder.
func (lvc *LikeVideoCreate) Mutation() *LikeVideoMutation {
	return lvc.mutation
}

// Save creates the LikeVideo in the database.
func (lvc *LikeVideoCreate) Save(ctx context.Context) (*LikeVideo, error) {
	lvc.defaults()
	return withHooks(ctx, lvc.sqlSave, lvc.mutation, lvc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lvc *LikeVideoCreate) SaveX(ctx context.Context) *LikeVideo {
	v, err := lvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lvc *LikeVideoCreate) Exec(ctx context.Context) error {
	_, err := lvc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lvc *LikeVideoCreate) ExecX(ctx context.Context) {
	if err := lvc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lvc *LikeVideoCreate) defaults() {
	if _, ok := lvc.mutation.CreateTime(); !ok {
		v := likevideo.DefaultCreateTime()
		lvc.mutation.SetCreateTime(v)
	}
	if _, ok := lvc.mutation.UpdateTime(); !ok {
		v := likevideo.DefaultUpdateTime()
		lvc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lvc *LikeVideoCreate) check() error {
	if _, ok := lvc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "LikeVideo.create_time"`)}
	}
	if _, ok := lvc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "LikeVideo.update_time"`)}
	}
	if _, ok := lvc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "LikeVideo.user_id"`)}
	}
	if _, ok := lvc.mutation.VideoID(); !ok {
		return &ValidationError{Name: "video_id", err: errors.New(`ent: missing required field "LikeVideo.video_id"`)}
	}
	if len(lvc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "LikeVideo.user"`)}
	}
	if len(lvc.mutation.VideoIDs()) == 0 {
		return &ValidationError{Name: "video", err: errors.New(`ent: missing required edge "LikeVideo.video"`)}
	}
	return nil
}

func (lvc *LikeVideoCreate) sqlSave(ctx context.Context) (*LikeVideo, error) {
	if err := lvc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lvc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lvc.mutation.id = &_node.ID
	lvc.mutation.done = true
	return _node, nil
}

func (lvc *LikeVideoCreate) createSpec() (*LikeVideo, *sqlgraph.CreateSpec) {
	var (
		_node = &LikeVideo{config: lvc.config}
		_spec = sqlgraph.NewCreateSpec(likevideo.Table, sqlgraph.NewFieldSpec(likevideo.FieldID, field.TypeInt))
	)
	if value, ok := lvc.mutation.CreateTime(); ok {
		_spec.SetField(likevideo.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := lvc.mutation.UpdateTime(); ok {
		_spec.SetField(likevideo.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if nodes := lvc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   likevideo.UserTable,
			Columns: []string{likevideo.UserColumn},
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
	if nodes := lvc.mutation.VideoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   likevideo.VideoTable,
			Columns: []string{likevideo.VideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(video.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.VideoID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LikeVideoCreateBulk is the builder for creating many LikeVideo entities in bulk.
type LikeVideoCreateBulk struct {
	config
	err      error
	builders []*LikeVideoCreate
}

// Save creates the LikeVideo entities in the database.
func (lvcb *LikeVideoCreateBulk) Save(ctx context.Context) ([]*LikeVideo, error) {
	if lvcb.err != nil {
		return nil, lvcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lvcb.builders))
	nodes := make([]*LikeVideo, len(lvcb.builders))
	mutators := make([]Mutator, len(lvcb.builders))
	for i := range lvcb.builders {
		func(i int, root context.Context) {
			builder := lvcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LikeVideoMutation)
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
					_, err = mutators[i+1].Mutate(root, lvcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lvcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lvcb *LikeVideoCreateBulk) SaveX(ctx context.Context) []*LikeVideo {
	v, err := lvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lvcb *LikeVideoCreateBulk) Exec(ctx context.Context) error {
	_, err := lvcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lvcb *LikeVideoCreateBulk) ExecX(ctx context.Context) {
	if err := lvcb.Exec(ctx); err != nil {
		panic(err)
	}
}
