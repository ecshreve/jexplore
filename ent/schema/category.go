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

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
			entkit.TitleField(),
			entkit.FilterOperator(gen.Contains),
		),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("clues", Clue.Type).Annotations(entgql.RelayConnection()),
	}
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),

		entkit.Actions(entkit.ListAction, entkit.ShowAction),
		entkit.Icon("GroupOutlined"),
	}
}
