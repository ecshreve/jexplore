// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/ecshreve/jexplore/ent/game"
	"github.com/ecshreve/jexplore/ent/season"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ga *GameQuery) CollectFields(ctx context.Context, satisfies ...string) (*GameQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ga, nil
	}
	if err := ga.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ga, nil
}

func (ga *GameQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(game.Columns))
		selectedFields = []string{game.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "season":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SeasonClient{config: ga.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			ga.withSeason = query
		case "show":
			if _, ok := fieldSeen[game.FieldShow]; !ok {
				selectedFields = append(selectedFields, game.FieldShow)
				fieldSeen[game.FieldShow] = struct{}{}
			}
		case "airdate":
			if _, ok := fieldSeen[game.FieldAirDate]; !ok {
				selectedFields = append(selectedFields, game.FieldAirDate)
				fieldSeen[game.FieldAirDate] = struct{}{}
			}
		case "tapedate":
			if _, ok := fieldSeen[game.FieldTapeDate]; !ok {
				selectedFields = append(selectedFields, game.FieldTapeDate)
				fieldSeen[game.FieldTapeDate] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ga.Select(selectedFields...)
	}
	return nil
}

type gamePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GamePaginateOption
}

func newGamePaginateArgs(rv map[string]any) *gamePaginateArgs {
	args := &gamePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &GameOrder{Field: &GameOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithGameOrder(order))
			}
		case *GameOrder:
			if v != nil {
				args.opts = append(args.opts, WithGameOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*GameWhereInput); ok {
		args.opts = append(args.opts, WithGameFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *SeasonQuery) CollectFields(ctx context.Context, satisfies ...string) (*SeasonQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return s, nil
	}
	if err := s.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SeasonQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(season.Columns))
		selectedFields = []string{season.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "games":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&GameClient{config: s.config}).Query()
			)
			args := newGamePaginateArgs(fieldArgs(ctx, new(GameWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newGamePager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					s.loadTotal = append(s.loadTotal, func(ctx context.Context, nodes []*Season) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID string `sql:"season_games"`
							Count  int    `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(season.GamesColumn), ids...))
						})
						if err := query.GroupBy(season.GamesColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[string]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					s.loadTotal = append(s.loadTotal, func(_ context.Context, nodes []*Season) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Games)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "Game")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(season.GamesColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			s.WithNamedGames(alias, func(wq *GameQuery) {
				*wq = *query
			})
		case "number":
			if _, ok := fieldSeen[season.FieldNumber]; !ok {
				selectedFields = append(selectedFields, season.FieldNumber)
				fieldSeen[season.FieldNumber] = struct{}{}
			}
		case "startdate":
			if _, ok := fieldSeen[season.FieldStartDate]; !ok {
				selectedFields = append(selectedFields, season.FieldStartDate)
				fieldSeen[season.FieldStartDate] = struct{}{}
			}
		case "enddate":
			if _, ok := fieldSeen[season.FieldEndDate]; !ok {
				selectedFields = append(selectedFields, season.FieldEndDate)
				fieldSeen[season.FieldEndDate] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		s.Select(selectedFields...)
	}
	return nil
}

type seasonPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SeasonPaginateOption
}

func newSeasonPaginateArgs(rv map[string]any) *seasonPaginateArgs {
	args := &seasonPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &SeasonOrder{Field: &SeasonOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithSeasonOrder(order))
			}
		case *SeasonOrder:
			if v != nil {
				args.opts = append(args.opts, WithSeasonOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*SeasonWhereInput); ok {
		args.opts = append(args.opts, WithSeasonFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
