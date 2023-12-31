// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"time"

	"github.com/ecshreve/jexplore/ent/category"
	"github.com/ecshreve/jexplore/ent/clue"
	"github.com/ecshreve/jexplore/ent/game"
	"github.com/ecshreve/jexplore/ent/predicate"
	"github.com/ecshreve/jexplore/ent/season"
)

// CategoryWhereInput represents a where input for filtering Category queries.
type CategoryWhereInput struct {
	Predicates []predicate.Category  `json:"-"`
	Not        *CategoryWhereInput   `json:"not,omitempty"`
	Or         []*CategoryWhereInput `json:"or,omitempty"`
	And        []*CategoryWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "clues" edge predicates.
	HasClues     *bool             `json:"hasClues,omitempty"`
	HasCluesWith []*ClueWhereInput `json:"hasCluesWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *CategoryWhereInput) AddPredicates(predicates ...predicate.Category) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the CategoryWhereInput filter on the CategoryQuery builder.
func (i *CategoryWhereInput) Filter(q *CategoryQuery) (*CategoryQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyCategoryWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyCategoryWhereInput is returned in case the CategoryWhereInput is empty.
var ErrEmptyCategoryWhereInput = errors.New("ent: empty predicate CategoryWhereInput")

// P returns a predicate for filtering categories.
// An error is returned if the input is empty or invalid.
func (i *CategoryWhereInput) P() (predicate.Category, error) {
	var predicates []predicate.Category
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, category.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Category, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, category.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Category, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, category.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, category.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, category.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, category.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, category.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, category.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, category.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, category.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, category.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, category.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, category.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, category.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, category.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, category.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, category.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, category.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, category.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, category.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, category.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, category.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, category.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, category.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasClues != nil {
		p := category.HasClues()
		if !*i.HasClues {
			p = category.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasCluesWith) > 0 {
		with := make([]predicate.Clue, 0, len(i.HasCluesWith))
		for _, w := range i.HasCluesWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasCluesWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, category.HasCluesWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyCategoryWhereInput
	case 1:
		return predicates[0], nil
	default:
		return category.And(predicates...), nil
	}
}

// ClueWhereInput represents a where input for filtering Clue queries.
type ClueWhereInput struct {
	Predicates []predicate.Clue  `json:"-"`
	Not        *ClueWhereInput   `json:"not,omitempty"`
	Or         []*ClueWhereInput `json:"or,omitempty"`
	And        []*ClueWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "question" field predicates.
	Question             *string  `json:"question,omitempty"`
	QuestionNEQ          *string  `json:"questionNEQ,omitempty"`
	QuestionIn           []string `json:"questionIn,omitempty"`
	QuestionNotIn        []string `json:"questionNotIn,omitempty"`
	QuestionGT           *string  `json:"questionGT,omitempty"`
	QuestionGTE          *string  `json:"questionGTE,omitempty"`
	QuestionLT           *string  `json:"questionLT,omitempty"`
	QuestionLTE          *string  `json:"questionLTE,omitempty"`
	QuestionContains     *string  `json:"questionContains,omitempty"`
	QuestionHasPrefix    *string  `json:"questionHasPrefix,omitempty"`
	QuestionHasSuffix    *string  `json:"questionHasSuffix,omitempty"`
	QuestionEqualFold    *string  `json:"questionEqualFold,omitempty"`
	QuestionContainsFold *string  `json:"questionContainsFold,omitempty"`

	// "answer" field predicates.
	Answer             *string  `json:"answer,omitempty"`
	AnswerNEQ          *string  `json:"answerNEQ,omitempty"`
	AnswerIn           []string `json:"answerIn,omitempty"`
	AnswerNotIn        []string `json:"answerNotIn,omitempty"`
	AnswerGT           *string  `json:"answerGT,omitempty"`
	AnswerGTE          *string  `json:"answerGTE,omitempty"`
	AnswerLT           *string  `json:"answerLT,omitempty"`
	AnswerLTE          *string  `json:"answerLTE,omitempty"`
	AnswerContains     *string  `json:"answerContains,omitempty"`
	AnswerHasPrefix    *string  `json:"answerHasPrefix,omitempty"`
	AnswerHasSuffix    *string  `json:"answerHasSuffix,omitempty"`
	AnswerEqualFold    *string  `json:"answerEqualFold,omitempty"`
	AnswerContainsFold *string  `json:"answerContainsFold,omitempty"`

	// "category_id" field predicates.
	CategoryID       *int  `json:"categoryID,omitempty"`
	CategoryIDNEQ    *int  `json:"categoryIDNEQ,omitempty"`
	CategoryIDIn     []int `json:"categoryIDIn,omitempty"`
	CategoryIDNotIn  []int `json:"categoryIDNotIn,omitempty"`
	CategoryIDIsNil  bool  `json:"categoryIDIsNil,omitempty"`
	CategoryIDNotNil bool  `json:"categoryIDNotNil,omitempty"`

	// "game_id" field predicates.
	GameID       *int  `json:"gameID,omitempty"`
	GameIDNEQ    *int  `json:"gameIDNEQ,omitempty"`
	GameIDIn     []int `json:"gameIDIn,omitempty"`
	GameIDNotIn  []int `json:"gameIDNotIn,omitempty"`
	GameIDIsNil  bool  `json:"gameIDIsNil,omitempty"`
	GameIDNotNil bool  `json:"gameIDNotNil,omitempty"`

	// "category" edge predicates.
	HasCategory     *bool                 `json:"hasCategory,omitempty"`
	HasCategoryWith []*CategoryWhereInput `json:"hasCategoryWith,omitempty"`

	// "game" edge predicates.
	HasGame     *bool             `json:"hasGame,omitempty"`
	HasGameWith []*GameWhereInput `json:"hasGameWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *ClueWhereInput) AddPredicates(predicates ...predicate.Clue) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the ClueWhereInput filter on the ClueQuery builder.
func (i *ClueWhereInput) Filter(q *ClueQuery) (*ClueQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyClueWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyClueWhereInput is returned in case the ClueWhereInput is empty.
var ErrEmptyClueWhereInput = errors.New("ent: empty predicate ClueWhereInput")

// P returns a predicate for filtering clues.
// An error is returned if the input is empty or invalid.
func (i *ClueWhereInput) P() (predicate.Clue, error) {
	var predicates []predicate.Clue
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, clue.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Clue, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, clue.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Clue, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, clue.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, clue.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, clue.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, clue.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, clue.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, clue.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, clue.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, clue.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, clue.IDLTE(*i.IDLTE))
	}
	if i.Question != nil {
		predicates = append(predicates, clue.QuestionEQ(*i.Question))
	}
	if i.QuestionNEQ != nil {
		predicates = append(predicates, clue.QuestionNEQ(*i.QuestionNEQ))
	}
	if len(i.QuestionIn) > 0 {
		predicates = append(predicates, clue.QuestionIn(i.QuestionIn...))
	}
	if len(i.QuestionNotIn) > 0 {
		predicates = append(predicates, clue.QuestionNotIn(i.QuestionNotIn...))
	}
	if i.QuestionGT != nil {
		predicates = append(predicates, clue.QuestionGT(*i.QuestionGT))
	}
	if i.QuestionGTE != nil {
		predicates = append(predicates, clue.QuestionGTE(*i.QuestionGTE))
	}
	if i.QuestionLT != nil {
		predicates = append(predicates, clue.QuestionLT(*i.QuestionLT))
	}
	if i.QuestionLTE != nil {
		predicates = append(predicates, clue.QuestionLTE(*i.QuestionLTE))
	}
	if i.QuestionContains != nil {
		predicates = append(predicates, clue.QuestionContains(*i.QuestionContains))
	}
	if i.QuestionHasPrefix != nil {
		predicates = append(predicates, clue.QuestionHasPrefix(*i.QuestionHasPrefix))
	}
	if i.QuestionHasSuffix != nil {
		predicates = append(predicates, clue.QuestionHasSuffix(*i.QuestionHasSuffix))
	}
	if i.QuestionEqualFold != nil {
		predicates = append(predicates, clue.QuestionEqualFold(*i.QuestionEqualFold))
	}
	if i.QuestionContainsFold != nil {
		predicates = append(predicates, clue.QuestionContainsFold(*i.QuestionContainsFold))
	}
	if i.Answer != nil {
		predicates = append(predicates, clue.AnswerEQ(*i.Answer))
	}
	if i.AnswerNEQ != nil {
		predicates = append(predicates, clue.AnswerNEQ(*i.AnswerNEQ))
	}
	if len(i.AnswerIn) > 0 {
		predicates = append(predicates, clue.AnswerIn(i.AnswerIn...))
	}
	if len(i.AnswerNotIn) > 0 {
		predicates = append(predicates, clue.AnswerNotIn(i.AnswerNotIn...))
	}
	if i.AnswerGT != nil {
		predicates = append(predicates, clue.AnswerGT(*i.AnswerGT))
	}
	if i.AnswerGTE != nil {
		predicates = append(predicates, clue.AnswerGTE(*i.AnswerGTE))
	}
	if i.AnswerLT != nil {
		predicates = append(predicates, clue.AnswerLT(*i.AnswerLT))
	}
	if i.AnswerLTE != nil {
		predicates = append(predicates, clue.AnswerLTE(*i.AnswerLTE))
	}
	if i.AnswerContains != nil {
		predicates = append(predicates, clue.AnswerContains(*i.AnswerContains))
	}
	if i.AnswerHasPrefix != nil {
		predicates = append(predicates, clue.AnswerHasPrefix(*i.AnswerHasPrefix))
	}
	if i.AnswerHasSuffix != nil {
		predicates = append(predicates, clue.AnswerHasSuffix(*i.AnswerHasSuffix))
	}
	if i.AnswerEqualFold != nil {
		predicates = append(predicates, clue.AnswerEqualFold(*i.AnswerEqualFold))
	}
	if i.AnswerContainsFold != nil {
		predicates = append(predicates, clue.AnswerContainsFold(*i.AnswerContainsFold))
	}
	if i.CategoryID != nil {
		predicates = append(predicates, clue.CategoryIDEQ(*i.CategoryID))
	}
	if i.CategoryIDNEQ != nil {
		predicates = append(predicates, clue.CategoryIDNEQ(*i.CategoryIDNEQ))
	}
	if len(i.CategoryIDIn) > 0 {
		predicates = append(predicates, clue.CategoryIDIn(i.CategoryIDIn...))
	}
	if len(i.CategoryIDNotIn) > 0 {
		predicates = append(predicates, clue.CategoryIDNotIn(i.CategoryIDNotIn...))
	}
	if i.CategoryIDIsNil {
		predicates = append(predicates, clue.CategoryIDIsNil())
	}
	if i.CategoryIDNotNil {
		predicates = append(predicates, clue.CategoryIDNotNil())
	}
	if i.GameID != nil {
		predicates = append(predicates, clue.GameIDEQ(*i.GameID))
	}
	if i.GameIDNEQ != nil {
		predicates = append(predicates, clue.GameIDNEQ(*i.GameIDNEQ))
	}
	if len(i.GameIDIn) > 0 {
		predicates = append(predicates, clue.GameIDIn(i.GameIDIn...))
	}
	if len(i.GameIDNotIn) > 0 {
		predicates = append(predicates, clue.GameIDNotIn(i.GameIDNotIn...))
	}
	if i.GameIDIsNil {
		predicates = append(predicates, clue.GameIDIsNil())
	}
	if i.GameIDNotNil {
		predicates = append(predicates, clue.GameIDNotNil())
	}

	if i.HasCategory != nil {
		p := clue.HasCategory()
		if !*i.HasCategory {
			p = clue.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasCategoryWith) > 0 {
		with := make([]predicate.Category, 0, len(i.HasCategoryWith))
		for _, w := range i.HasCategoryWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasCategoryWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, clue.HasCategoryWith(with...))
	}
	if i.HasGame != nil {
		p := clue.HasGame()
		if !*i.HasGame {
			p = clue.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasGameWith) > 0 {
		with := make([]predicate.Game, 0, len(i.HasGameWith))
		for _, w := range i.HasGameWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasGameWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, clue.HasGameWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyClueWhereInput
	case 1:
		return predicates[0], nil
	default:
		return clue.And(predicates...), nil
	}
}

// GameWhereInput represents a where input for filtering Game queries.
type GameWhereInput struct {
	Predicates []predicate.Game  `json:"-"`
	Not        *GameWhereInput   `json:"not,omitempty"`
	Or         []*GameWhereInput `json:"or,omitempty"`
	And        []*GameWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "show" field predicates.
	Show      *int  `json:"show,omitempty"`
	ShowNEQ   *int  `json:"showNEQ,omitempty"`
	ShowIn    []int `json:"showIn,omitempty"`
	ShowNotIn []int `json:"showNotIn,omitempty"`
	ShowGT    *int  `json:"showGT,omitempty"`
	ShowGTE   *int  `json:"showGTE,omitempty"`
	ShowLT    *int  `json:"showLT,omitempty"`
	ShowLTE   *int  `json:"showLTE,omitempty"`

	// "air_date" field predicates.
	AirDate      *time.Time  `json:"airDate,omitempty"`
	AirDateNEQ   *time.Time  `json:"airDateNEQ,omitempty"`
	AirDateIn    []time.Time `json:"airDateIn,omitempty"`
	AirDateNotIn []time.Time `json:"airDateNotIn,omitempty"`
	AirDateGT    *time.Time  `json:"airDateGT,omitempty"`
	AirDateGTE   *time.Time  `json:"airDateGTE,omitempty"`
	AirDateLT    *time.Time  `json:"airDateLT,omitempty"`
	AirDateLTE   *time.Time  `json:"airDateLTE,omitempty"`

	// "tape_date" field predicates.
	TapeDate      *time.Time  `json:"tapeDate,omitempty"`
	TapeDateNEQ   *time.Time  `json:"tapeDateNEQ,omitempty"`
	TapeDateIn    []time.Time `json:"tapeDateIn,omitempty"`
	TapeDateNotIn []time.Time `json:"tapeDateNotIn,omitempty"`
	TapeDateGT    *time.Time  `json:"tapeDateGT,omitempty"`
	TapeDateGTE   *time.Time  `json:"tapeDateGTE,omitempty"`
	TapeDateLT    *time.Time  `json:"tapeDateLT,omitempty"`
	TapeDateLTE   *time.Time  `json:"tapeDateLTE,omitempty"`

	// "season_id" field predicates.
	SeasonID       *int  `json:"seasonID,omitempty"`
	SeasonIDNEQ    *int  `json:"seasonIDNEQ,omitempty"`
	SeasonIDIn     []int `json:"seasonIDIn,omitempty"`
	SeasonIDNotIn  []int `json:"seasonIDNotIn,omitempty"`
	SeasonIDIsNil  bool  `json:"seasonIDIsNil,omitempty"`
	SeasonIDNotNil bool  `json:"seasonIDNotNil,omitempty"`

	// "season" edge predicates.
	HasSeason     *bool               `json:"hasSeason,omitempty"`
	HasSeasonWith []*SeasonWhereInput `json:"hasSeasonWith,omitempty"`

	// "clues" edge predicates.
	HasClues     *bool             `json:"hasClues,omitempty"`
	HasCluesWith []*ClueWhereInput `json:"hasCluesWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *GameWhereInput) AddPredicates(predicates ...predicate.Game) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the GameWhereInput filter on the GameQuery builder.
func (i *GameWhereInput) Filter(q *GameQuery) (*GameQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyGameWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyGameWhereInput is returned in case the GameWhereInput is empty.
var ErrEmptyGameWhereInput = errors.New("ent: empty predicate GameWhereInput")

// P returns a predicate for filtering games.
// An error is returned if the input is empty or invalid.
func (i *GameWhereInput) P() (predicate.Game, error) {
	var predicates []predicate.Game
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, game.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Game, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, game.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Game, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, game.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, game.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, game.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, game.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, game.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, game.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, game.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, game.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, game.IDLTE(*i.IDLTE))
	}
	if i.Show != nil {
		predicates = append(predicates, game.ShowEQ(*i.Show))
	}
	if i.ShowNEQ != nil {
		predicates = append(predicates, game.ShowNEQ(*i.ShowNEQ))
	}
	if len(i.ShowIn) > 0 {
		predicates = append(predicates, game.ShowIn(i.ShowIn...))
	}
	if len(i.ShowNotIn) > 0 {
		predicates = append(predicates, game.ShowNotIn(i.ShowNotIn...))
	}
	if i.ShowGT != nil {
		predicates = append(predicates, game.ShowGT(*i.ShowGT))
	}
	if i.ShowGTE != nil {
		predicates = append(predicates, game.ShowGTE(*i.ShowGTE))
	}
	if i.ShowLT != nil {
		predicates = append(predicates, game.ShowLT(*i.ShowLT))
	}
	if i.ShowLTE != nil {
		predicates = append(predicates, game.ShowLTE(*i.ShowLTE))
	}
	if i.AirDate != nil {
		predicates = append(predicates, game.AirDateEQ(*i.AirDate))
	}
	if i.AirDateNEQ != nil {
		predicates = append(predicates, game.AirDateNEQ(*i.AirDateNEQ))
	}
	if len(i.AirDateIn) > 0 {
		predicates = append(predicates, game.AirDateIn(i.AirDateIn...))
	}
	if len(i.AirDateNotIn) > 0 {
		predicates = append(predicates, game.AirDateNotIn(i.AirDateNotIn...))
	}
	if i.AirDateGT != nil {
		predicates = append(predicates, game.AirDateGT(*i.AirDateGT))
	}
	if i.AirDateGTE != nil {
		predicates = append(predicates, game.AirDateGTE(*i.AirDateGTE))
	}
	if i.AirDateLT != nil {
		predicates = append(predicates, game.AirDateLT(*i.AirDateLT))
	}
	if i.AirDateLTE != nil {
		predicates = append(predicates, game.AirDateLTE(*i.AirDateLTE))
	}
	if i.TapeDate != nil {
		predicates = append(predicates, game.TapeDateEQ(*i.TapeDate))
	}
	if i.TapeDateNEQ != nil {
		predicates = append(predicates, game.TapeDateNEQ(*i.TapeDateNEQ))
	}
	if len(i.TapeDateIn) > 0 {
		predicates = append(predicates, game.TapeDateIn(i.TapeDateIn...))
	}
	if len(i.TapeDateNotIn) > 0 {
		predicates = append(predicates, game.TapeDateNotIn(i.TapeDateNotIn...))
	}
	if i.TapeDateGT != nil {
		predicates = append(predicates, game.TapeDateGT(*i.TapeDateGT))
	}
	if i.TapeDateGTE != nil {
		predicates = append(predicates, game.TapeDateGTE(*i.TapeDateGTE))
	}
	if i.TapeDateLT != nil {
		predicates = append(predicates, game.TapeDateLT(*i.TapeDateLT))
	}
	if i.TapeDateLTE != nil {
		predicates = append(predicates, game.TapeDateLTE(*i.TapeDateLTE))
	}
	if i.SeasonID != nil {
		predicates = append(predicates, game.SeasonIDEQ(*i.SeasonID))
	}
	if i.SeasonIDNEQ != nil {
		predicates = append(predicates, game.SeasonIDNEQ(*i.SeasonIDNEQ))
	}
	if len(i.SeasonIDIn) > 0 {
		predicates = append(predicates, game.SeasonIDIn(i.SeasonIDIn...))
	}
	if len(i.SeasonIDNotIn) > 0 {
		predicates = append(predicates, game.SeasonIDNotIn(i.SeasonIDNotIn...))
	}
	if i.SeasonIDIsNil {
		predicates = append(predicates, game.SeasonIDIsNil())
	}
	if i.SeasonIDNotNil {
		predicates = append(predicates, game.SeasonIDNotNil())
	}

	if i.HasSeason != nil {
		p := game.HasSeason()
		if !*i.HasSeason {
			p = game.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasSeasonWith) > 0 {
		with := make([]predicate.Season, 0, len(i.HasSeasonWith))
		for _, w := range i.HasSeasonWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasSeasonWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, game.HasSeasonWith(with...))
	}
	if i.HasClues != nil {
		p := game.HasClues()
		if !*i.HasClues {
			p = game.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasCluesWith) > 0 {
		with := make([]predicate.Clue, 0, len(i.HasCluesWith))
		for _, w := range i.HasCluesWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasCluesWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, game.HasCluesWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyGameWhereInput
	case 1:
		return predicates[0], nil
	default:
		return game.And(predicates...), nil
	}
}

// SeasonWhereInput represents a where input for filtering Season queries.
type SeasonWhereInput struct {
	Predicates []predicate.Season  `json:"-"`
	Not        *SeasonWhereInput   `json:"not,omitempty"`
	Or         []*SeasonWhereInput `json:"or,omitempty"`
	And        []*SeasonWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "number" field predicates.
	Number      *int  `json:"number,omitempty"`
	NumberNEQ   *int  `json:"numberNEQ,omitempty"`
	NumberIn    []int `json:"numberIn,omitempty"`
	NumberNotIn []int `json:"numberNotIn,omitempty"`
	NumberGT    *int  `json:"numberGT,omitempty"`
	NumberGTE   *int  `json:"numberGTE,omitempty"`
	NumberLT    *int  `json:"numberLT,omitempty"`
	NumberLTE   *int  `json:"numberLTE,omitempty"`

	// "start_date" field predicates.
	StartDate      *time.Time  `json:"startDate,omitempty"`
	StartDateNEQ   *time.Time  `json:"startDateNEQ,omitempty"`
	StartDateIn    []time.Time `json:"startDateIn,omitempty"`
	StartDateNotIn []time.Time `json:"startDateNotIn,omitempty"`
	StartDateGT    *time.Time  `json:"startDateGT,omitempty"`
	StartDateGTE   *time.Time  `json:"startDateGTE,omitempty"`
	StartDateLT    *time.Time  `json:"startDateLT,omitempty"`
	StartDateLTE   *time.Time  `json:"startDateLTE,omitempty"`

	// "end_date" field predicates.
	EndDate      *time.Time  `json:"endDate,omitempty"`
	EndDateNEQ   *time.Time  `json:"endDateNEQ,omitempty"`
	EndDateIn    []time.Time `json:"endDateIn,omitempty"`
	EndDateNotIn []time.Time `json:"endDateNotIn,omitempty"`
	EndDateGT    *time.Time  `json:"endDateGT,omitempty"`
	EndDateGTE   *time.Time  `json:"endDateGTE,omitempty"`
	EndDateLT    *time.Time  `json:"endDateLT,omitempty"`
	EndDateLTE   *time.Time  `json:"endDateLTE,omitempty"`

	// "games" edge predicates.
	HasGames     *bool             `json:"hasGames,omitempty"`
	HasGamesWith []*GameWhereInput `json:"hasGamesWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *SeasonWhereInput) AddPredicates(predicates ...predicate.Season) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the SeasonWhereInput filter on the SeasonQuery builder.
func (i *SeasonWhereInput) Filter(q *SeasonQuery) (*SeasonQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptySeasonWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptySeasonWhereInput is returned in case the SeasonWhereInput is empty.
var ErrEmptySeasonWhereInput = errors.New("ent: empty predicate SeasonWhereInput")

// P returns a predicate for filtering seasons.
// An error is returned if the input is empty or invalid.
func (i *SeasonWhereInput) P() (predicate.Season, error) {
	var predicates []predicate.Season
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, season.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Season, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, season.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Season, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, season.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, season.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, season.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, season.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, season.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, season.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, season.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, season.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, season.IDLTE(*i.IDLTE))
	}
	if i.Number != nil {
		predicates = append(predicates, season.NumberEQ(*i.Number))
	}
	if i.NumberNEQ != nil {
		predicates = append(predicates, season.NumberNEQ(*i.NumberNEQ))
	}
	if len(i.NumberIn) > 0 {
		predicates = append(predicates, season.NumberIn(i.NumberIn...))
	}
	if len(i.NumberNotIn) > 0 {
		predicates = append(predicates, season.NumberNotIn(i.NumberNotIn...))
	}
	if i.NumberGT != nil {
		predicates = append(predicates, season.NumberGT(*i.NumberGT))
	}
	if i.NumberGTE != nil {
		predicates = append(predicates, season.NumberGTE(*i.NumberGTE))
	}
	if i.NumberLT != nil {
		predicates = append(predicates, season.NumberLT(*i.NumberLT))
	}
	if i.NumberLTE != nil {
		predicates = append(predicates, season.NumberLTE(*i.NumberLTE))
	}
	if i.StartDate != nil {
		predicates = append(predicates, season.StartDateEQ(*i.StartDate))
	}
	if i.StartDateNEQ != nil {
		predicates = append(predicates, season.StartDateNEQ(*i.StartDateNEQ))
	}
	if len(i.StartDateIn) > 0 {
		predicates = append(predicates, season.StartDateIn(i.StartDateIn...))
	}
	if len(i.StartDateNotIn) > 0 {
		predicates = append(predicates, season.StartDateNotIn(i.StartDateNotIn...))
	}
	if i.StartDateGT != nil {
		predicates = append(predicates, season.StartDateGT(*i.StartDateGT))
	}
	if i.StartDateGTE != nil {
		predicates = append(predicates, season.StartDateGTE(*i.StartDateGTE))
	}
	if i.StartDateLT != nil {
		predicates = append(predicates, season.StartDateLT(*i.StartDateLT))
	}
	if i.StartDateLTE != nil {
		predicates = append(predicates, season.StartDateLTE(*i.StartDateLTE))
	}
	if i.EndDate != nil {
		predicates = append(predicates, season.EndDateEQ(*i.EndDate))
	}
	if i.EndDateNEQ != nil {
		predicates = append(predicates, season.EndDateNEQ(*i.EndDateNEQ))
	}
	if len(i.EndDateIn) > 0 {
		predicates = append(predicates, season.EndDateIn(i.EndDateIn...))
	}
	if len(i.EndDateNotIn) > 0 {
		predicates = append(predicates, season.EndDateNotIn(i.EndDateNotIn...))
	}
	if i.EndDateGT != nil {
		predicates = append(predicates, season.EndDateGT(*i.EndDateGT))
	}
	if i.EndDateGTE != nil {
		predicates = append(predicates, season.EndDateGTE(*i.EndDateGTE))
	}
	if i.EndDateLT != nil {
		predicates = append(predicates, season.EndDateLT(*i.EndDateLT))
	}
	if i.EndDateLTE != nil {
		predicates = append(predicates, season.EndDateLTE(*i.EndDateLTE))
	}

	if i.HasGames != nil {
		p := season.HasGames()
		if !*i.HasGames {
			p = season.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasGamesWith) > 0 {
		with := make([]predicate.Game, 0, len(i.HasGamesWith))
		for _, w := range i.HasGamesWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasGamesWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, season.HasGamesWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptySeasonWhereInput
	case 1:
		return predicates[0], nil
	default:
		return season.And(predicates...), nil
	}
}
