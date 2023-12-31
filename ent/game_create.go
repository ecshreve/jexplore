// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/jexplore/ent/clue"
	"github.com/ecshreve/jexplore/ent/game"
	"github.com/ecshreve/jexplore/ent/season"
)

// GameCreate is the builder for creating a Game entity.
type GameCreate struct {
	config
	mutation *GameMutation
	hooks    []Hook
}

// SetShow sets the "show" field.
func (gc *GameCreate) SetShow(i int) *GameCreate {
	gc.mutation.SetShow(i)
	return gc
}

// SetAirDate sets the "air_date" field.
func (gc *GameCreate) SetAirDate(t time.Time) *GameCreate {
	gc.mutation.SetAirDate(t)
	return gc
}

// SetTapeDate sets the "tape_date" field.
func (gc *GameCreate) SetTapeDate(t time.Time) *GameCreate {
	gc.mutation.SetTapeDate(t)
	return gc
}

// SetSeasonID sets the "season_id" field.
func (gc *GameCreate) SetSeasonID(i int) *GameCreate {
	gc.mutation.SetSeasonID(i)
	return gc
}

// SetNillableSeasonID sets the "season_id" field if the given value is not nil.
func (gc *GameCreate) SetNillableSeasonID(i *int) *GameCreate {
	if i != nil {
		gc.SetSeasonID(*i)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GameCreate) SetID(i int) *GameCreate {
	gc.mutation.SetID(i)
	return gc
}

// SetSeason sets the "season" edge to the Season entity.
func (gc *GameCreate) SetSeason(s *Season) *GameCreate {
	return gc.SetSeasonID(s.ID)
}

// AddClueIDs adds the "clues" edge to the Clue entity by IDs.
func (gc *GameCreate) AddClueIDs(ids ...int) *GameCreate {
	gc.mutation.AddClueIDs(ids...)
	return gc
}

// AddClues adds the "clues" edges to the Clue entity.
func (gc *GameCreate) AddClues(c ...*Clue) *GameCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return gc.AddClueIDs(ids...)
}

// Mutation returns the GameMutation object of the builder.
func (gc *GameCreate) Mutation() *GameMutation {
	return gc.mutation
}

// Save creates the Game in the database.
func (gc *GameCreate) Save(ctx context.Context) (*Game, error) {
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GameCreate) SaveX(ctx context.Context) *Game {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GameCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GameCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GameCreate) check() error {
	if _, ok := gc.mutation.Show(); !ok {
		return &ValidationError{Name: "show", err: errors.New(`ent: missing required field "Game.show"`)}
	}
	if _, ok := gc.mutation.AirDate(); !ok {
		return &ValidationError{Name: "air_date", err: errors.New(`ent: missing required field "Game.air_date"`)}
	}
	if _, ok := gc.mutation.TapeDate(); !ok {
		return &ValidationError{Name: "tape_date", err: errors.New(`ent: missing required field "Game.tape_date"`)}
	}
	return nil
}

func (gc *GameCreate) sqlSave(ctx context.Context) (*Game, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GameCreate) createSpec() (*Game, *sqlgraph.CreateSpec) {
	var (
		_node = &Game{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(game.Table, sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt))
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.Show(); ok {
		_spec.SetField(game.FieldShow, field.TypeInt, value)
		_node.Show = value
	}
	if value, ok := gc.mutation.AirDate(); ok {
		_spec.SetField(game.FieldAirDate, field.TypeTime, value)
		_node.AirDate = value
	}
	if value, ok := gc.mutation.TapeDate(); ok {
		_spec.SetField(game.FieldTapeDate, field.TypeTime, value)
		_node.TapeDate = value
	}
	if nodes := gc.mutation.SeasonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.SeasonTable,
			Columns: []string{game.SeasonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SeasonID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.CluesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.CluesTable,
			Columns: []string{game.CluesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GameCreateBulk is the builder for creating many Game entities in bulk.
type GameCreateBulk struct {
	config
	builders []*GameCreate
}

// Save creates the Game entities in the database.
func (gcb *GameCreateBulk) Save(ctx context.Context) ([]*Game, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Game, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GameMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GameCreateBulk) SaveX(ctx context.Context) []*Game {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GameCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GameCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}
