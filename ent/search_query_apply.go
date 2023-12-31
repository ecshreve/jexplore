// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/ecshreve/jexplore/ent/category"
	"github.com/ecshreve/jexplore/ent/clue"
	"github.com/ecshreve/jexplore/ent/game"
	"github.com/ecshreve/jexplore/ent/predicate"
	"github.com/ecshreve/jexplore/ent/season"
)

func (cwi *CategoryWhereInput) ApplySearchQuery(q *string) *CategoryWhereInput {
	if cwi == nil {
		cwi = &CategoryWhereInput{}
	}

	if q == nil {
		return cwi
	}

	var orPredicates []predicate.Category

	orPredicates = append(orPredicates, category.NameContains(*q))

	cwi.AddPredicates(category.Or(orPredicates...))
	return cwi
}

func (cwi *ClueWhereInput) ApplySearchQuery(q *string) *ClueWhereInput {
	if cwi == nil {
		cwi = &ClueWhereInput{}
	}

	if q == nil {
		return cwi
	}

	var orPredicates []predicate.Clue

	orPredicates = append(orPredicates, clue.QuestionContains(*q))

	orPredicates = append(orPredicates, clue.AnswerContains(*q))

	cwi.AddPredicates(clue.Or(orPredicates...))
	return cwi
}

func (gwi *GameWhereInput) ApplySearchQuery(q *string) *GameWhereInput {
	if gwi == nil {
		gwi = &GameWhereInput{}
	}

	if q == nil {
		return gwi
	}

	var orPredicates []predicate.Game

	gwi.AddPredicates(game.Or(orPredicates...))
	return gwi
}

func (swi *SeasonWhereInput) ApplySearchQuery(q *string) *SeasonWhereInput {
	if swi == nil {
		swi = &SeasonWhereInput{}
	}

	if q == nil {
		return swi
	}

	var orPredicates []predicate.Season

	swi.AddPredicates(season.Or(orPredicates...))
	return swi
}
