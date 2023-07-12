// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GamesColumns holds the columns for the "games" table.
	GamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "show", Type: field.TypeInt},
		{Name: "air_date", Type: field.TypeTime},
		{Name: "tape_date", Type: field.TypeTime},
		{Name: "season_games", Type: field.TypeString, Nullable: true},
	}
	// GamesTable holds the schema information for the "games" table.
	GamesTable = &schema.Table{
		Name:       "games",
		Columns:    GamesColumns,
		PrimaryKey: []*schema.Column{GamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "games_seasons_games",
				Columns:    []*schema.Column{GamesColumns[4]},
				RefColumns: []*schema.Column{SeasonsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SeasonsColumns holds the columns for the "seasons" table.
	SeasonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "number", Type: field.TypeInt},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime},
	}
	// SeasonsTable holds the schema information for the "seasons" table.
	SeasonsTable = &schema.Table{
		Name:       "seasons",
		Columns:    SeasonsColumns,
		PrimaryKey: []*schema.Column{SeasonsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GamesTable,
		SeasonsTable,
	}
)

func init() {
	GamesTable.ForeignKeys[0].RefTable = SeasonsTable
}
