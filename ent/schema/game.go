package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/entc/gen"
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
		field.Int("id").Unique().Annotations(
			entgql.OrderField("ID"),
			entkit.FilterOperator(gen.Contains),
		),
		field.Int("show").Annotations(
			entgql.OrderField("SHOW"),
			entkit.FilterOperator(gen.Contains),
			entkit.TitleField(),
		),
		field.Time("air_date").Annotations(
			entgql.OrderField("AIR_DATE"),
			entkit.FilterOperator(gen.Contains),
		),
		field.Time("tape_date").Annotations(
			entgql.OrderField("TAPE_DATE"),
			entkit.FilterOperator(gen.Contains),
		),
		field.Int("season_id").Optional().
			Annotations(
				entgql.OrderField("SEASON_ID"),
				entkit.FilterOperator(gen.Contains),
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

		entkit.Actions(entkit.ListAction, entkit.ShowAction),
		entkit.Icon("PlusSquareOutlined"),
	}
}
