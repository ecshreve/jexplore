package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/entkit/entkit/v2"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.Int("show").Annotations(
			entgql.OrderField("SHOW"),
		),
		field.Time("airDate").Annotations(
			entgql.OrderField("AIR_DATE"),
		),
		field.Time("tapeDate").Annotations(
			entgql.OrderField("TAPE_DATE"),
		),
		field.Int("season_id").Optional().
			Annotations(
				entgql.OrderField("SEASON_ID"),
			),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("season", Season.Type).Ref("games").Field("season_id").Unique(),
		edge.To("clues", Clue.Type).Annotations(entgql.RelayConnection()),
	}
}

func (Game) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entkit.Actions(append(entkit.DefaultActions, entkit.EdgesDiagramAction)...),
		entkit.Icon("PlusSquareOutlined"),
	}
}
