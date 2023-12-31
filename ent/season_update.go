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
	"github.com/ecshreve/jexplore/ent/game"
	"github.com/ecshreve/jexplore/ent/predicate"
	"github.com/ecshreve/jexplore/ent/season"
)

// SeasonUpdate is the builder for updating Season entities.
type SeasonUpdate struct {
	config
	hooks    []Hook
	mutation *SeasonMutation
}

// Where appends a list predicates to the SeasonUpdate builder.
func (su *SeasonUpdate) Where(ps ...predicate.Season) *SeasonUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetNumber sets the "number" field.
func (su *SeasonUpdate) SetNumber(i int) *SeasonUpdate {
	su.mutation.ResetNumber()
	su.mutation.SetNumber(i)
	return su
}

// AddNumber adds i to the "number" field.
func (su *SeasonUpdate) AddNumber(i int) *SeasonUpdate {
	su.mutation.AddNumber(i)
	return su
}

// SetStartDate sets the "start_date" field.
func (su *SeasonUpdate) SetStartDate(t time.Time) *SeasonUpdate {
	su.mutation.SetStartDate(t)
	return su
}

// SetEndDate sets the "end_date" field.
func (su *SeasonUpdate) SetEndDate(t time.Time) *SeasonUpdate {
	su.mutation.SetEndDate(t)
	return su
}

// AddGameIDs adds the "games" edge to the Game entity by IDs.
func (su *SeasonUpdate) AddGameIDs(ids ...int) *SeasonUpdate {
	su.mutation.AddGameIDs(ids...)
	return su
}

// AddGames adds the "games" edges to the Game entity.
func (su *SeasonUpdate) AddGames(g ...*Game) *SeasonUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return su.AddGameIDs(ids...)
}

// Mutation returns the SeasonMutation object of the builder.
func (su *SeasonUpdate) Mutation() *SeasonMutation {
	return su.mutation
}

// ClearGames clears all "games" edges to the Game entity.
func (su *SeasonUpdate) ClearGames() *SeasonUpdate {
	su.mutation.ClearGames()
	return su
}

// RemoveGameIDs removes the "games" edge to Game entities by IDs.
func (su *SeasonUpdate) RemoveGameIDs(ids ...int) *SeasonUpdate {
	su.mutation.RemoveGameIDs(ids...)
	return su
}

// RemoveGames removes "games" edges to Game entities.
func (su *SeasonUpdate) RemoveGames(g ...*Game) *SeasonUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return su.RemoveGameIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SeasonUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SeasonUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SeasonUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SeasonUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SeasonUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(season.Table, season.Columns, sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Number(); ok {
		_spec.SetField(season.FieldNumber, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedNumber(); ok {
		_spec.AddField(season.FieldNumber, field.TypeInt, value)
	}
	if value, ok := su.mutation.StartDate(); ok {
		_spec.SetField(season.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := su.mutation.EndDate(); ok {
		_spec.SetField(season.FieldEndDate, field.TypeTime, value)
	}
	if su.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.GamesTable,
			Columns: []string{season.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedGamesIDs(); len(nodes) > 0 && !su.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.GamesTable,
			Columns: []string{season.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.GamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.GamesTable,
			Columns: []string{season.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{season.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SeasonUpdateOne is the builder for updating a single Season entity.
type SeasonUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SeasonMutation
}

// SetNumber sets the "number" field.
func (suo *SeasonUpdateOne) SetNumber(i int) *SeasonUpdateOne {
	suo.mutation.ResetNumber()
	suo.mutation.SetNumber(i)
	return suo
}

// AddNumber adds i to the "number" field.
func (suo *SeasonUpdateOne) AddNumber(i int) *SeasonUpdateOne {
	suo.mutation.AddNumber(i)
	return suo
}

// SetStartDate sets the "start_date" field.
func (suo *SeasonUpdateOne) SetStartDate(t time.Time) *SeasonUpdateOne {
	suo.mutation.SetStartDate(t)
	return suo
}

// SetEndDate sets the "end_date" field.
func (suo *SeasonUpdateOne) SetEndDate(t time.Time) *SeasonUpdateOne {
	suo.mutation.SetEndDate(t)
	return suo
}

// AddGameIDs adds the "games" edge to the Game entity by IDs.
func (suo *SeasonUpdateOne) AddGameIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.AddGameIDs(ids...)
	return suo
}

// AddGames adds the "games" edges to the Game entity.
func (suo *SeasonUpdateOne) AddGames(g ...*Game) *SeasonUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return suo.AddGameIDs(ids...)
}

// Mutation returns the SeasonMutation object of the builder.
func (suo *SeasonUpdateOne) Mutation() *SeasonMutation {
	return suo.mutation
}

// ClearGames clears all "games" edges to the Game entity.
func (suo *SeasonUpdateOne) ClearGames() *SeasonUpdateOne {
	suo.mutation.ClearGames()
	return suo
}

// RemoveGameIDs removes the "games" edge to Game entities by IDs.
func (suo *SeasonUpdateOne) RemoveGameIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.RemoveGameIDs(ids...)
	return suo
}

// RemoveGames removes "games" edges to Game entities.
func (suo *SeasonUpdateOne) RemoveGames(g ...*Game) *SeasonUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return suo.RemoveGameIDs(ids...)
}

// Where appends a list predicates to the SeasonUpdate builder.
func (suo *SeasonUpdateOne) Where(ps ...predicate.Season) *SeasonUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SeasonUpdateOne) Select(field string, fields ...string) *SeasonUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Season entity.
func (suo *SeasonUpdateOne) Save(ctx context.Context) (*Season, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SeasonUpdateOne) SaveX(ctx context.Context) *Season {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SeasonUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SeasonUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SeasonUpdateOne) sqlSave(ctx context.Context) (_node *Season, err error) {
	_spec := sqlgraph.NewUpdateSpec(season.Table, season.Columns, sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Season.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, season.FieldID)
		for _, f := range fields {
			if !season.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != season.FieldID {
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
	if value, ok := suo.mutation.Number(); ok {
		_spec.SetField(season.FieldNumber, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedNumber(); ok {
		_spec.AddField(season.FieldNumber, field.TypeInt, value)
	}
	if value, ok := suo.mutation.StartDate(); ok {
		_spec.SetField(season.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := suo.mutation.EndDate(); ok {
		_spec.SetField(season.FieldEndDate, field.TypeTime, value)
	}
	if suo.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.GamesTable,
			Columns: []string{season.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedGamesIDs(); len(nodes) > 0 && !suo.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.GamesTable,
			Columns: []string{season.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.GamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.GamesTable,
			Columns: []string{season.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Season{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{season.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
