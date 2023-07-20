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

// Season holds the schema definition for the Season entity.
type Season struct {
	ent.Schema
}

// Fields of the Season.
func (Season) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Annotations(
			entgql.OrderField("ID"),
			entkit.FilterOperator(gen.Contains),
		),
		field.Int("number").Annotations(
			entgql.OrderField("NUMBER"),
			entkit.TitleField(),
			entkit.FilterOperator(gen.Contains),
		),
		field.Time("start_date").Annotations(
			entgql.OrderField("START_DATE"),
			entkit.FilterOperator(gen.Contains),
		),
		field.Time("end_date").Annotations(
			entgql.OrderField("END_DATE"),
			entkit.FilterOperator(gen.Contains),
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
