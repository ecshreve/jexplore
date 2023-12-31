// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/jexplore/ent/category"
	"github.com/ecshreve/jexplore/ent/clue"
	"github.com/ecshreve/jexplore/ent/game"
)

// ClueCreate is the builder for creating a Clue entity.
type ClueCreate struct {
	config
	mutation *ClueMutation
	hooks    []Hook
}

// SetQuestion sets the "question" field.
func (cc *ClueCreate) SetQuestion(s string) *ClueCreate {
	cc.mutation.SetQuestion(s)
	return cc
}

// SetAnswer sets the "answer" field.
func (cc *ClueCreate) SetAnswer(s string) *ClueCreate {
	cc.mutation.SetAnswer(s)
	return cc
}

// SetCategoryID sets the "category_id" field.
func (cc *ClueCreate) SetCategoryID(i int) *ClueCreate {
	cc.mutation.SetCategoryID(i)
	return cc
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (cc *ClueCreate) SetNillableCategoryID(i *int) *ClueCreate {
	if i != nil {
		cc.SetCategoryID(*i)
	}
	return cc
}

// SetGameID sets the "game_id" field.
func (cc *ClueCreate) SetGameID(i int) *ClueCreate {
	cc.mutation.SetGameID(i)
	return cc
}

// SetNillableGameID sets the "game_id" field if the given value is not nil.
func (cc *ClueCreate) SetNillableGameID(i *int) *ClueCreate {
	if i != nil {
		cc.SetGameID(*i)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ClueCreate) SetID(i int) *ClueCreate {
	cc.mutation.SetID(i)
	return cc
}

// SetCategory sets the "category" edge to the Category entity.
func (cc *ClueCreate) SetCategory(c *Category) *ClueCreate {
	return cc.SetCategoryID(c.ID)
}

// SetGame sets the "game" edge to the Game entity.
func (cc *ClueCreate) SetGame(g *Game) *ClueCreate {
	return cc.SetGameID(g.ID)
}

// Mutation returns the ClueMutation object of the builder.
func (cc *ClueCreate) Mutation() *ClueMutation {
	return cc.mutation
}

// Save creates the Clue in the database.
func (cc *ClueCreate) Save(ctx context.Context) (*Clue, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClueCreate) SaveX(ctx context.Context) *Clue {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClueCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClueCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClueCreate) check() error {
	if _, ok := cc.mutation.Question(); !ok {
		return &ValidationError{Name: "question", err: errors.New(`ent: missing required field "Clue.question"`)}
	}
	if _, ok := cc.mutation.Answer(); !ok {
		return &ValidationError{Name: "answer", err: errors.New(`ent: missing required field "Clue.answer"`)}
	}
	return nil
}

func (cc *ClueCreate) sqlSave(ctx context.Context) (*Clue, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ClueCreate) createSpec() (*Clue, *sqlgraph.CreateSpec) {
	var (
		_node = &Clue{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(clue.Table, sqlgraph.NewFieldSpec(clue.FieldID, field.TypeInt))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Question(); ok {
		_spec.SetField(clue.FieldQuestion, field.TypeString, value)
		_node.Question = value
	}
	if value, ok := cc.mutation.Answer(); ok {
		_spec.SetField(clue.FieldAnswer, field.TypeString, value)
		_node.Answer = value
	}
	if nodes := cc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clue.CategoryTable,
			Columns: []string{clue.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clue.GameTable,
			Columns: []string{clue.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GameID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ClueCreateBulk is the builder for creating many Clue entities in bulk.
type ClueCreateBulk struct {
	config
	builders []*ClueCreate
}

// Save creates the Clue entities in the database.
func (ccb *ClueCreateBulk) Save(ctx context.Context) ([]*Clue, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Clue, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClueMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ClueCreateBulk) SaveX(ctx context.Context) []*Clue {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ClueCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ClueCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
