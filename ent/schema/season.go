package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/entkit/entkit/v2"
)

// Season holds the schema definition for the Season entity.
type Season struct {
	ent.Schema
}

// Fields of the Season.
func (Season) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Annotations(
			entgql.OrderField("ID"),
		),
		field.Int("number").Annotations(
			entgql.OrderField("NUMBER"),
			entkit.TitleField(),
		),
		field.Time("startDate").Annotations(
			entgql.OrderField("START_DATE"),
		),
		field.Time("endDate").Annotations(
			entgql.OrderField("END_DATE"),
		),
	}
}

// Edges of the Season.
func (Season) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("games", Game.Type).Annotations(entgql.RelayConnection()),
	}
}

func (Season) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),

		entkit.IndexRoute(),
		entkit.Actions(entkit.ListAction, entkit.ShowAction),
		entkit.Icon("CalendarOutlined"),
	}
}
