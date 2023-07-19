package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/entkit/entkit/v2"
)

// Clue holds the schema definition for the Clue entity.
type Clue struct {
	ent.Schema
}

// Fields of the Clue.
func (Clue) Fields() []ent.Field {
	return []ent.Field{
		field.Text("question").Annotations(
			entgql.OrderField("QUESTION"),
		),
		field.Text("answer").Annotations(
			entgql.OrderField("ANSWER"),
		),
		field.Int("category_id").Optional().
			Annotations(
				entgql.OrderField("CATEGORY_ID"),
			),
		field.Int("game_id").Optional().
			Annotations(
				entgql.OrderField("GAME_ID"),
			),
	}
}

// Edges of the Game.
func (Clue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).Ref("clues").Field("category_id").Unique(),
		edge.From("game", Game.Type).Ref("clues").Field("game_id").Unique(),
	}
}

func (Clue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entkit.Actions(append(entkit.DefaultActions, entkit.EdgesDiagramAction)...),
		entkit.Icon("QuestionCircleOutlined"),
	}
}
