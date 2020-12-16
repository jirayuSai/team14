// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/playlist_video"
	"github.com/tanapon395/playlist-video/ent/predicate"
	"github.com/tanapon395/playlist-video/ent/resolution"
)

// ResolutionUpdate is the builder for updating Resolution entities.
type ResolutionUpdate struct {
	config
	hooks      []Hook
	mutation   *ResolutionMutation
	predicates []predicate.Resolution
}

// Where adds a new predicate for the builder.
func (ru *ResolutionUpdate) Where(ps ...predicate.Resolution) *ResolutionUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetValue sets the value field.
func (ru *ResolutionUpdate) SetValue(i int) *ResolutionUpdate {
	ru.mutation.ResetValue()
	ru.mutation.SetValue(i)
	return ru
}

// AddValue adds i to value.
func (ru *ResolutionUpdate) AddValue(i int) *ResolutionUpdate {
	ru.mutation.AddValue(i)
	return ru
}

// AddPlaylistVideoIDs adds the playlist_videos edge to Playlist_Video by ids.
func (ru *ResolutionUpdate) AddPlaylistVideoIDs(ids ...int) *ResolutionUpdate {
	ru.mutation.AddPlaylistVideoIDs(ids...)
	return ru
}

// AddPlaylistVideos adds the playlist_videos edges to Playlist_Video.
func (ru *ResolutionUpdate) AddPlaylistVideos(p ...*Playlist_Video) *ResolutionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPlaylistVideoIDs(ids...)
}

// Mutation returns the ResolutionMutation object of the builder.
func (ru *ResolutionUpdate) Mutation() *ResolutionMutation {
	return ru.mutation
}

// RemovePlaylistVideoIDs removes the playlist_videos edge to Playlist_Video by ids.
func (ru *ResolutionUpdate) RemovePlaylistVideoIDs(ids ...int) *ResolutionUpdate {
	ru.mutation.RemovePlaylistVideoIDs(ids...)
	return ru
}

// RemovePlaylistVideos removes playlist_videos edges to Playlist_Video.
func (ru *ResolutionUpdate) RemovePlaylistVideos(p ...*Playlist_Video) *ResolutionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePlaylistVideoIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *ResolutionUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ru.mutation.Value(); ok {
		if err := resolution.ValueValidator(v); err != nil {
			return 0, &ValidationError{Name: "value", err: fmt.Errorf("ent: validator failed for field \"value\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResolutionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResolutionUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResolutionUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResolutionUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ResolutionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resolution.Table,
			Columns: resolution.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resolution.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resolution.FieldValue,
		})
	}
	if value, ok := ru.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resolution.FieldValue,
		})
	}
	if nodes := ru.mutation.RemovedPlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resolution.PlaylistVideosTable,
			Columns: []string{resolution.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resolution.PlaylistVideosTable,
			Columns: []string{resolution.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resolution.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ResolutionUpdateOne is the builder for updating a single Resolution entity.
type ResolutionUpdateOne struct {
	config
	hooks    []Hook
	mutation *ResolutionMutation
}

// SetValue sets the value field.
func (ruo *ResolutionUpdateOne) SetValue(i int) *ResolutionUpdateOne {
	ruo.mutation.ResetValue()
	ruo.mutation.SetValue(i)
	return ruo
}

// AddValue adds i to value.
func (ruo *ResolutionUpdateOne) AddValue(i int) *ResolutionUpdateOne {
	ruo.mutation.AddValue(i)
	return ruo
}

// AddPlaylistVideoIDs adds the playlist_videos edge to Playlist_Video by ids.
func (ruo *ResolutionUpdateOne) AddPlaylistVideoIDs(ids ...int) *ResolutionUpdateOne {
	ruo.mutation.AddPlaylistVideoIDs(ids...)
	return ruo
}

// AddPlaylistVideos adds the playlist_videos edges to Playlist_Video.
func (ruo *ResolutionUpdateOne) AddPlaylistVideos(p ...*Playlist_Video) *ResolutionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPlaylistVideoIDs(ids...)
}

// Mutation returns the ResolutionMutation object of the builder.
func (ruo *ResolutionUpdateOne) Mutation() *ResolutionMutation {
	return ruo.mutation
}

// RemovePlaylistVideoIDs removes the playlist_videos edge to Playlist_Video by ids.
func (ruo *ResolutionUpdateOne) RemovePlaylistVideoIDs(ids ...int) *ResolutionUpdateOne {
	ruo.mutation.RemovePlaylistVideoIDs(ids...)
	return ruo
}

// RemovePlaylistVideos removes playlist_videos edges to Playlist_Video.
func (ruo *ResolutionUpdateOne) RemovePlaylistVideos(p ...*Playlist_Video) *ResolutionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePlaylistVideoIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ruo *ResolutionUpdateOne) Save(ctx context.Context) (*Resolution, error) {
	if v, ok := ruo.mutation.Value(); ok {
		if err := resolution.ValueValidator(v); err != nil {
			return nil, &ValidationError{Name: "value", err: fmt.Errorf("ent: validator failed for field \"value\": %w", err)}
		}
	}

	var (
		err  error
		node *Resolution
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResolutionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResolutionUpdateOne) SaveX(ctx context.Context) *Resolution {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *ResolutionUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResolutionUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ResolutionUpdateOne) sqlSave(ctx context.Context) (r *Resolution, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resolution.Table,
			Columns: resolution.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resolution.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Resolution.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resolution.FieldValue,
		})
	}
	if value, ok := ruo.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resolution.FieldValue,
		})
	}
	if nodes := ruo.mutation.RemovedPlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resolution.PlaylistVideosTable,
			Columns: []string{resolution.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resolution.PlaylistVideosTable,
			Columns: []string{resolution.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Resolution{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resolution.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
