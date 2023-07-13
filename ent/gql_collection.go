// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/ecshreve/jexplore/ent/category"
	"github.com/ecshreve/jexplore/ent/clue"
	"github.com/ecshreve/jexplore/ent/game"
	"github.com/ecshreve/jexplore/ent/season"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CategoryQuery) CollectFields(ctx context.Context, satisfies ...string) (*CategoryQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CategoryQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(category.Columns))
		selectedFields = []string{category.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "clues":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ClueClient{config: c.config}).Query()
			)
			args := newCluePaginateArgs(fieldArgs(ctx, new(ClueWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newCluePager(args.opts, args.last != nil)
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
					c.loadTotal = append(c.loadTotal, func(ctx context.Context, nodes []*Category) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"category_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(category.CluesColumn), ids...))
						})
						if err := query.GroupBy(category.CluesColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
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
					c.loadTotal = append(c.loadTotal, func(_ context.Context, nodes []*Category) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Clues)
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
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "Clue")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(category.CluesColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			c.WithNamedClues(alias, func(wq *ClueQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[category.FieldName]; !ok {
				selectedFields = append(selectedFields, category.FieldName)
				fieldSeen[category.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type categoryPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CategoryPaginateOption
}

func newCategoryPaginateArgs(rv map[string]any) *categoryPaginateArgs {
	args := &categoryPaginateArgs{}
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
				order      = &CategoryOrder{Field: &CategoryOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCategoryOrder(order))
			}
		case *CategoryOrder:
			if v != nil {
				args.opts = append(args.opts, WithCategoryOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*CategoryWhereInput); ok {
		args.opts = append(args.opts, WithCategoryFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *ClueQuery) CollectFields(ctx context.Context, satisfies ...string) (*ClueQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *ClueQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(clue.Columns))
		selectedFields = []string{clue.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "category":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CategoryClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			c.withCategory = query
			if _, ok := fieldSeen[clue.FieldCategoryID]; !ok {
				selectedFields = append(selectedFields, clue.FieldCategoryID)
				fieldSeen[clue.FieldCategoryID] = struct{}{}
			}
		case "game":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&GameClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			c.withGame = query
			if _, ok := fieldSeen[clue.FieldGameID]; !ok {
				selectedFields = append(selectedFields, clue.FieldGameID)
				fieldSeen[clue.FieldGameID] = struct{}{}
			}
		case "question":
			if _, ok := fieldSeen[clue.FieldQuestion]; !ok {
				selectedFields = append(selectedFields, clue.FieldQuestion)
				fieldSeen[clue.FieldQuestion] = struct{}{}
			}
		case "answer":
			if _, ok := fieldSeen[clue.FieldAnswer]; !ok {
				selectedFields = append(selectedFields, clue.FieldAnswer)
				fieldSeen[clue.FieldAnswer] = struct{}{}
			}
		case "categoryID":
			if _, ok := fieldSeen[clue.FieldCategoryID]; !ok {
				selectedFields = append(selectedFields, clue.FieldCategoryID)
				fieldSeen[clue.FieldCategoryID] = struct{}{}
			}
		case "gameID":
			if _, ok := fieldSeen[clue.FieldGameID]; !ok {
				selectedFields = append(selectedFields, clue.FieldGameID)
				fieldSeen[clue.FieldGameID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type cluePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CluePaginateOption
}

func newCluePaginateArgs(rv map[string]any) *cluePaginateArgs {
	args := &cluePaginateArgs{}
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
				order      = &ClueOrder{Field: &ClueOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithClueOrder(order))
			}
		case *ClueOrder:
			if v != nil {
				args.opts = append(args.opts, WithClueOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*ClueWhereInput); ok {
		args.opts = append(args.opts, WithClueFilter(v.Filter))
	}
	return args
}

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
			if _, ok := fieldSeen[game.FieldSeasonID]; !ok {
				selectedFields = append(selectedFields, game.FieldSeasonID)
				fieldSeen[game.FieldSeasonID] = struct{}{}
			}
		case "clues":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ClueClient{config: ga.config}).Query()
			)
			args := newCluePaginateArgs(fieldArgs(ctx, new(ClueWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newCluePager(args.opts, args.last != nil)
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
					ga.loadTotal = append(ga.loadTotal, func(ctx context.Context, nodes []*Game) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"game_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(game.CluesColumn), ids...))
						})
						if err := query.GroupBy(game.CluesColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					ga.loadTotal = append(ga.loadTotal, func(_ context.Context, nodes []*Game) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Clues)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
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
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "Clue")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(game.CluesColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			ga.WithNamedClues(alias, func(wq *ClueQuery) {
				*wq = *query
			})
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
		case "seasonID":
			if _, ok := fieldSeen[game.FieldSeasonID]; !ok {
				selectedFields = append(selectedFields, game.FieldSeasonID)
				fieldSeen[game.FieldSeasonID] = struct{}{}
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
							NodeID int `sql:"season_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(season.GamesColumn), ids...))
						})
						if err := query.GroupBy(season.GamesColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
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