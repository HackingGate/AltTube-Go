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
	"github.com/hackinggate/alttube-go/ent/likevideo"
	"github.com/hackinggate/alttube-go/ent/predicate"
	"github.com/hackinggate/alttube-go/ent/video"
)

// VideoUpdate is the builder for updating Video entities.
type VideoUpdate struct {
	config
	hooks    []Hook
	mutation *VideoMutation
}

// Where appends a list predicates to the VideoUpdate builder.
func (vu *VideoUpdate) Where(ps ...predicate.Video) *VideoUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetUpdateTime sets the "update_time" field.
func (vu *VideoUpdate) SetUpdateTime(t time.Time) *VideoUpdate {
	vu.mutation.SetUpdateTime(t)
	return vu
}

// SetTitle sets the "title" field.
func (vu *VideoUpdate) SetTitle(s string) *VideoUpdate {
	vu.mutation.SetTitle(s)
	return vu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableTitle(s *string) *VideoUpdate {
	if s != nil {
		vu.SetTitle(*s)
	}
	return vu
}

// SetDescription sets the "description" field.
func (vu *VideoUpdate) SetDescription(s string) *VideoUpdate {
	vu.mutation.SetDescription(s)
	return vu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableDescription(s *string) *VideoUpdate {
	if s != nil {
		vu.SetDescription(*s)
	}
	return vu
}

// SetUploadDate sets the "uploadDate" field.
func (vu *VideoUpdate) SetUploadDate(t time.Time) *VideoUpdate {
	vu.mutation.SetUploadDate(t)
	return vu
}

// SetNillableUploadDate sets the "uploadDate" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableUploadDate(t *time.Time) *VideoUpdate {
	if t != nil {
		vu.SetUploadDate(*t)
	}
	return vu
}

// SetUploader sets the "uploader" field.
func (vu *VideoUpdate) SetUploader(s string) *VideoUpdate {
	vu.mutation.SetUploader(s)
	return vu
}

// SetNillableUploader sets the "uploader" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableUploader(s *string) *VideoUpdate {
	if s != nil {
		vu.SetUploader(*s)
	}
	return vu
}

// SetUploaderUrl sets the "uploaderUrl" field.
func (vu *VideoUpdate) SetUploaderUrl(s string) *VideoUpdate {
	vu.mutation.SetUploaderUrl(s)
	return vu
}

// SetNillableUploaderUrl sets the "uploaderUrl" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableUploaderUrl(s *string) *VideoUpdate {
	if s != nil {
		vu.SetUploaderUrl(*s)
	}
	return vu
}

// SetThumbnailUrl sets the "thumbnailUrl" field.
func (vu *VideoUpdate) SetThumbnailUrl(s string) *VideoUpdate {
	vu.mutation.SetThumbnailUrl(s)
	return vu
}

// SetNillableThumbnailUrl sets the "thumbnailUrl" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableThumbnailUrl(s *string) *VideoUpdate {
	if s != nil {
		vu.SetThumbnailUrl(*s)
	}
	return vu
}

// AddLikeVideoIDs adds the "like_videos" edge to the LikeVideo entity by IDs.
func (vu *VideoUpdate) AddLikeVideoIDs(ids ...int) *VideoUpdate {
	vu.mutation.AddLikeVideoIDs(ids...)
	return vu
}

// AddLikeVideos adds the "like_videos" edges to the LikeVideo entity.
func (vu *VideoUpdate) AddLikeVideos(l ...*LikeVideo) *VideoUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return vu.AddLikeVideoIDs(ids...)
}

// Mutation returns the VideoMutation object of the builder.
func (vu *VideoUpdate) Mutation() *VideoMutation {
	return vu.mutation
}

// ClearLikeVideos clears all "like_videos" edges to the LikeVideo entity.
func (vu *VideoUpdate) ClearLikeVideos() *VideoUpdate {
	vu.mutation.ClearLikeVideos()
	return vu
}

// RemoveLikeVideoIDs removes the "like_videos" edge to LikeVideo entities by IDs.
func (vu *VideoUpdate) RemoveLikeVideoIDs(ids ...int) *VideoUpdate {
	vu.mutation.RemoveLikeVideoIDs(ids...)
	return vu
}

// RemoveLikeVideos removes "like_videos" edges to LikeVideo entities.
func (vu *VideoUpdate) RemoveLikeVideos(l ...*LikeVideo) *VideoUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return vu.RemoveLikeVideoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VideoUpdate) Save(ctx context.Context) (int, error) {
	vu.defaults()
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VideoUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VideoUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VideoUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vu *VideoUpdate) defaults() {
	if _, ok := vu.mutation.UpdateTime(); !ok {
		v := video.UpdateDefaultUpdateTime()
		vu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VideoUpdate) check() error {
	if v, ok := vu.mutation.Title(); ok {
		if err := video.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Video.title": %w`, err)}
		}
	}
	if v, ok := vu.mutation.Uploader(); ok {
		if err := video.UploaderValidator(v); err != nil {
			return &ValidationError{Name: "uploader", err: fmt.Errorf(`ent: validator failed for field "Video.uploader": %w`, err)}
		}
	}
	return nil
}

func (vu *VideoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(video.Table, video.Columns, sqlgraph.NewFieldSpec(video.FieldID, field.TypeString))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.UpdateTime(); ok {
		_spec.SetField(video.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := vu.mutation.Title(); ok {
		_spec.SetField(video.FieldTitle, field.TypeString, value)
	}
	if value, ok := vu.mutation.Description(); ok {
		_spec.SetField(video.FieldDescription, field.TypeString, value)
	}
	if value, ok := vu.mutation.UploadDate(); ok {
		_spec.SetField(video.FieldUploadDate, field.TypeTime, value)
	}
	if value, ok := vu.mutation.Uploader(); ok {
		_spec.SetField(video.FieldUploader, field.TypeString, value)
	}
	if value, ok := vu.mutation.UploaderUrl(); ok {
		_spec.SetField(video.FieldUploaderUrl, field.TypeString, value)
	}
	if value, ok := vu.mutation.ThumbnailUrl(); ok {
		_spec.SetField(video.FieldThumbnailUrl, field.TypeString, value)
	}
	if vu.mutation.LikeVideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.LikeVideosTable,
			Columns: []string{video.LikeVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(likevideo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedLikeVideosIDs(); len(nodes) > 0 && !vu.mutation.LikeVideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.LikeVideosTable,
			Columns: []string{video.LikeVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(likevideo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.LikeVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.LikeVideosTable,
			Columns: []string{video.LikeVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(likevideo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{video.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VideoUpdateOne is the builder for updating a single Video entity.
type VideoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VideoMutation
}

// SetUpdateTime sets the "update_time" field.
func (vuo *VideoUpdateOne) SetUpdateTime(t time.Time) *VideoUpdateOne {
	vuo.mutation.SetUpdateTime(t)
	return vuo
}

// SetTitle sets the "title" field.
func (vuo *VideoUpdateOne) SetTitle(s string) *VideoUpdateOne {
	vuo.mutation.SetTitle(s)
	return vuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableTitle(s *string) *VideoUpdateOne {
	if s != nil {
		vuo.SetTitle(*s)
	}
	return vuo
}

// SetDescription sets the "description" field.
func (vuo *VideoUpdateOne) SetDescription(s string) *VideoUpdateOne {
	vuo.mutation.SetDescription(s)
	return vuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableDescription(s *string) *VideoUpdateOne {
	if s != nil {
		vuo.SetDescription(*s)
	}
	return vuo
}

// SetUploadDate sets the "uploadDate" field.
func (vuo *VideoUpdateOne) SetUploadDate(t time.Time) *VideoUpdateOne {
	vuo.mutation.SetUploadDate(t)
	return vuo
}

// SetNillableUploadDate sets the "uploadDate" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableUploadDate(t *time.Time) *VideoUpdateOne {
	if t != nil {
		vuo.SetUploadDate(*t)
	}
	return vuo
}

// SetUploader sets the "uploader" field.
func (vuo *VideoUpdateOne) SetUploader(s string) *VideoUpdateOne {
	vuo.mutation.SetUploader(s)
	return vuo
}

// SetNillableUploader sets the "uploader" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableUploader(s *string) *VideoUpdateOne {
	if s != nil {
		vuo.SetUploader(*s)
	}
	return vuo
}

// SetUploaderUrl sets the "uploaderUrl" field.
func (vuo *VideoUpdateOne) SetUploaderUrl(s string) *VideoUpdateOne {
	vuo.mutation.SetUploaderUrl(s)
	return vuo
}

// SetNillableUploaderUrl sets the "uploaderUrl" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableUploaderUrl(s *string) *VideoUpdateOne {
	if s != nil {
		vuo.SetUploaderUrl(*s)
	}
	return vuo
}

// SetThumbnailUrl sets the "thumbnailUrl" field.
func (vuo *VideoUpdateOne) SetThumbnailUrl(s string) *VideoUpdateOne {
	vuo.mutation.SetThumbnailUrl(s)
	return vuo
}

// SetNillableThumbnailUrl sets the "thumbnailUrl" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableThumbnailUrl(s *string) *VideoUpdateOne {
	if s != nil {
		vuo.SetThumbnailUrl(*s)
	}
	return vuo
}

// AddLikeVideoIDs adds the "like_videos" edge to the LikeVideo entity by IDs.
func (vuo *VideoUpdateOne) AddLikeVideoIDs(ids ...int) *VideoUpdateOne {
	vuo.mutation.AddLikeVideoIDs(ids...)
	return vuo
}

// AddLikeVideos adds the "like_videos" edges to the LikeVideo entity.
func (vuo *VideoUpdateOne) AddLikeVideos(l ...*LikeVideo) *VideoUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return vuo.AddLikeVideoIDs(ids...)
}

// Mutation returns the VideoMutation object of the builder.
func (vuo *VideoUpdateOne) Mutation() *VideoMutation {
	return vuo.mutation
}

// ClearLikeVideos clears all "like_videos" edges to the LikeVideo entity.
func (vuo *VideoUpdateOne) ClearLikeVideos() *VideoUpdateOne {
	vuo.mutation.ClearLikeVideos()
	return vuo
}

// RemoveLikeVideoIDs removes the "like_videos" edge to LikeVideo entities by IDs.
func (vuo *VideoUpdateOne) RemoveLikeVideoIDs(ids ...int) *VideoUpdateOne {
	vuo.mutation.RemoveLikeVideoIDs(ids...)
	return vuo
}

// RemoveLikeVideos removes "like_videos" edges to LikeVideo entities.
func (vuo *VideoUpdateOne) RemoveLikeVideos(l ...*LikeVideo) *VideoUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return vuo.RemoveLikeVideoIDs(ids...)
}

// Where appends a list predicates to the VideoUpdate builder.
func (vuo *VideoUpdateOne) Where(ps ...predicate.Video) *VideoUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VideoUpdateOne) Select(field string, fields ...string) *VideoUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Video entity.
func (vuo *VideoUpdateOne) Save(ctx context.Context) (*Video, error) {
	vuo.defaults()
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VideoUpdateOne) SaveX(ctx context.Context) *Video {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VideoUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VideoUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vuo *VideoUpdateOne) defaults() {
	if _, ok := vuo.mutation.UpdateTime(); !ok {
		v := video.UpdateDefaultUpdateTime()
		vuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VideoUpdateOne) check() error {
	if v, ok := vuo.mutation.Title(); ok {
		if err := video.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Video.title": %w`, err)}
		}
	}
	if v, ok := vuo.mutation.Uploader(); ok {
		if err := video.UploaderValidator(v); err != nil {
			return &ValidationError{Name: "uploader", err: fmt.Errorf(`ent: validator failed for field "Video.uploader": %w`, err)}
		}
	}
	return nil
}

func (vuo *VideoUpdateOne) sqlSave(ctx context.Context) (_node *Video, err error) {
	if err := vuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(video.Table, video.Columns, sqlgraph.NewFieldSpec(video.FieldID, field.TypeString))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Video.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, video.FieldID)
		for _, f := range fields {
			if !video.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != video.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.UpdateTime(); ok {
		_spec.SetField(video.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := vuo.mutation.Title(); ok {
		_spec.SetField(video.FieldTitle, field.TypeString, value)
	}
	if value, ok := vuo.mutation.Description(); ok {
		_spec.SetField(video.FieldDescription, field.TypeString, value)
	}
	if value, ok := vuo.mutation.UploadDate(); ok {
		_spec.SetField(video.FieldUploadDate, field.TypeTime, value)
	}
	if value, ok := vuo.mutation.Uploader(); ok {
		_spec.SetField(video.FieldUploader, field.TypeString, value)
	}
	if value, ok := vuo.mutation.UploaderUrl(); ok {
		_spec.SetField(video.FieldUploaderUrl, field.TypeString, value)
	}
	if value, ok := vuo.mutation.ThumbnailUrl(); ok {
		_spec.SetField(video.FieldThumbnailUrl, field.TypeString, value)
	}
	if vuo.mutation.LikeVideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.LikeVideosTable,
			Columns: []string{video.LikeVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(likevideo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedLikeVideosIDs(); len(nodes) > 0 && !vuo.mutation.LikeVideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.LikeVideosTable,
			Columns: []string{video.LikeVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(likevideo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.LikeVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.LikeVideosTable,
			Columns: []string{video.LikeVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(likevideo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Video{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{video.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}
