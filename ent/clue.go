// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/jexplore/ent/category"
	"github.com/ecshreve/jexplore/ent/clue"
	"github.com/ecshreve/jexplore/ent/game"
)

// Clue is the model entity for the Clue schema.
type Clue struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Question holds the value of the "question" field.
	Question string `json:"question,omitempty"`
	// Answer holds the value of the "answer" field.
	Answer string `json:"answer,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID int `json:"category_id,omitempty"`
	// GameID holds the value of the "game_id" field.
	GameID int `json:"game_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClueQuery when eager-loading is set.
	Edges        ClueEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ClueEdges holds the relations/edges for other nodes in the graph.
type ClueEdges struct {
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// Game holds the value of the game edge.
	Game *Game `json:"game,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClueEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[0] {
		if e.Category == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// GameOrErr returns the Game value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClueEdges) GameOrErr() (*Game, error) {
	if e.loadedTypes[1] {
		if e.Game == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: game.Label}
		}
		return e.Game, nil
	}
	return nil, &NotLoadedError{edge: "game"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Clue) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case clue.FieldID, clue.FieldCategoryID, clue.FieldGameID:
			values[i] = new(sql.NullInt64)
		case clue.FieldQuestion, clue.FieldAnswer:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Clue fields.
func (c *Clue) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case clue.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case clue.FieldQuestion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field question", values[i])
			} else if value.Valid {
				c.Question = value.String
			}
		case clue.FieldAnswer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field answer", values[i])
			} else if value.Valid {
				c.Answer = value.String
			}
		case clue.FieldCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value.Valid {
				c.CategoryID = int(value.Int64)
			}
		case clue.FieldGameID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field game_id", values[i])
			} else if value.Valid {
				c.GameID = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Clue.
// This includes values selected through modifiers, order, etc.
func (c *Clue) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryCategory queries the "category" edge of the Clue entity.
func (c *Clue) QueryCategory() *CategoryQuery {
	return NewClueClient(c.config).QueryCategory(c)
}

// QueryGame queries the "game" edge of the Clue entity.
func (c *Clue) QueryGame() *GameQuery {
	return NewClueClient(c.config).QueryGame(c)
}

// Update returns a builder for updating this Clue.
// Note that you need to call Clue.Unwrap() before calling this method if this Clue
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Clue) Update() *ClueUpdateOne {
	return NewClueClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Clue entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Clue) Unwrap() *Clue {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Clue is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Clue) String() string {
	var builder strings.Builder
	builder.WriteString("Clue(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("question=")
	builder.WriteString(c.Question)
	builder.WriteString(", ")
	builder.WriteString("answer=")
	builder.WriteString(c.Answer)
	builder.WriteString(", ")
	builder.WriteString("category_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CategoryID))
	builder.WriteString(", ")
	builder.WriteString("game_id=")
	builder.WriteString(fmt.Sprintf("%v", c.GameID))
	builder.WriteByte(')')
	return builder.String()
}

// Clues is a parsable slice of Clue.
type Clues []*Clue