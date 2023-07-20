// package main is a WIP tool for scraping j-archive.com.
package main

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ecshreve/jexplore/ent"
	"github.com/ecshreve/jexplore/ent/category"
	"github.com/ecshreve/jexplore/ent/game"
	"github.com/gocolly/colly/v2"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting jscrape...")
	ctx := context.Background()
	client, err := InitClient(ctx)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 1; i++ {
		log.Info("exiting")
		continue

		games, err := ScrapeSeasonGames(int64(i))
		if err != nil {
			panic(err)
		}

		toCreate := []*ent.GameCreate{}
		for _, g := range games {
			toCreate = append(toCreate, client.Game.Create().SetID(g.ID).SetSeasonID(g.SeasonID).SetShow(g.Show).SetAirDate(g.AirDate).SetTapeDate(g.TapeDate))
		}

		_, err = client.Game.CreateBulk(toCreate...).Save(ctx)
		if err != nil {
			log.Infof("error creating games: %v", err)
		}
	}

	for seasonID := 1; seasonID <= 1; seasonID++ {
		log.Info("exiting")
		continue

		gameIDs := client.Game.Query().Where(game.SeasonIDEQ(int(seasonID))).IDsX(ctx)
		for ind, gameID := range gameIDs {
			clueMap, catMap := ScrapeGameClues(int64(gameID))

			for clueID, cat := range catMap {
				fromDB, err := client.Category.Query().Where(category.Name(cat)).Only(ctx)
				if err != nil || fromDB == nil {
					panic(err)
				}

				clueMap[clueID].CategoryID = int(fromDB.ID)
			}

			toCreateList := []*ent.ClueCreate{}
			for _, clue := range clueMap {
				toCreateList = append(toCreateList, client.Clue.Create().SetID(clue.ID).SetGameID(clue.GameID).SetCategoryID(clue.CategoryID).SetQuestion(clue.Question).SetAnswer(clue.Answer))
			}

			created, err := client.Clue.CreateBulk(toCreateList...).Save(ctx)
			if err != nil {
				log.Errorf("error creating clues: %v", err)
			}
			log.Infof("created %d clues", len(created))

			if ind%10 == 0 {
				log.Infof("finished %d/%d games -- season: %d", ind, len(gameIDs), seasonID)
			}
		}
		log.Infof("finished season %d", seasonID)
	}
}

// TODO: this is silly, should fix it
type Round int // 0 = Jeopardy, 1 = Double Jeopardy, 2 = Final Jeopardy

const (
	Jeopardy Round = iota + 1
	DoubleJeopardy
	FinalJeopardy
)

var RoundMap = map[string]int{
	"J":  int(Jeopardy),
	"DJ": int(DoubleJeopardy),
	"FJ": int(FinalJeopardy),
	"TB": int(FinalJeopardy),
}

// scrapeGameClues scrapes a game from j-archive.com.
func ScrapeGameClues(gameID int64) (map[int64]*ent.Clue, map[int64]string) {
	clueMap := map[int64]*ent.Clue{}
	clueStrings := map[int64]string{}
	cats := map[Round][]string{}

	c := colly.NewCollector(
		colly.CacheDir("./cache"),
	)

	// collect and parse the clues
	c.OnHTML("td.clue", func(e *colly.HTMLElement) {
		cid := e.ChildAttr("td.clue_text", "id")
		if cid == "" {
			return
		}

		clueText := e.ChildText(fmt.Sprintf("td#%s", cid))
		clueAnswer := e.ChildText(fmt.Sprintf("td#%s_r em.correct_response", cid))
		clueId := ParseClueID(cid, gameID, RoundMap)

		clueMap[clueId] = &ent.Clue{ID: int(clueId), GameID: int(gameID), Question: clueText, Answer: clueAnswer}
		clueStrings[clueId] = cid
	})

	// collect and parse the categories for single jepp
	c.OnHTML("div[id=jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[Jeopardy] = append(cats[Jeopardy], cc...)
	})

	// collect and parse the categories for double jepp
	c.OnHTML("div[id=double_jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[DoubleJeopardy] = append(cats[DoubleJeopardy], cc...)
	})

	// collect and parse the categories for final jepp
	c.OnHTML("div[id=final_jeopardy_round]", func(e *colly.HTMLElement) {
		cc := []string{}
		e.ForEach("td.category_name", func(_ int, el *colly.HTMLElement) {
			cc = append(cc, el.Text)
		})
		cats[FinalJeopardy] = append(cats[FinalJeopardy], cc...)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.Visit(fmt.Sprintf("https://www.j-archive.com/showgame.php?game_id=%d", gameID))

	catMap := map[int64]string{}

	for clueId, clueStr := range clueStrings {
		rd, col := ParseRoundAndColumn(clueStr)
		catName := cats[Round(rd)][col-1]
		catMap[clueId] = catName
	}

	return clueMap, catMap
}

func InitClient(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open(
		"sqlite3",
		"file:file.db?mode=rwc&cache=shared&_fk=1",
	)
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(
		ctx,
		//migrate.WithGlobalUniqueID(true),
	); err != nil {
		return nil, err
	}
	return client, nil
}

// ParseRoundAndColumn parses a clue string and returns the round and column.
func ParseRoundAndColumn(clueString string) (int64, int64) {
	tokens := strings.Split(clueString, "_")

	var rd int64
	if tokens[1] == "J" {
		rd = 1
	} else if tokens[1] == "DJ" {
		rd = 2
	} else if tokens[1] == "FJ" {
		return 3, 1
	} else if tokens[1] == "TB" {
		return 3, 2
	}

	if len(tokens) != 4 {
		return -1, -1
	}

	col, _ := strconv.ParseInt(tokens[2], 10, 64)
	return rd, col
}

// ParseClueID converts a clue string to a clue ID.
// Clue strings have the format "clue_DJ_1_2", "clue_FJ" and the parsed int64
// is of the form <game_id>0<round>0<clue_index>.
func ParseClueID(clueString string, gameId int64, rm map[string]int) int64 {
	baseVal := gameId * 100000
	tokens := strings.Split(clueString, "_")
	if len(tokens) == 2 {
		if tokens[1] == "FJ" {
			return baseVal + 3061
		}
		return baseVal + 3062
	}

	// TODO: this is hacky, and I forget how it works
	round := rm[tokens[1]]                           // val = 804000000 round = DJ = 2
	baseVal = baseVal + (int64(round) * 1000)        // val = 804002000
	column, _ := strconv.ParseInt(tokens[2], 10, 64) // column = 1, row = 2
	row, _ := strconv.ParseInt(tokens[3], 10, 64)    // val = 804002032
	return baseVal + ((int64(round) - 1) * 30) + (((column - 1) * 5) + row)
}

func ScrapeSeasonGames(seasonID int64) ([]*ent.Game, error) {
	var gameRE = regexp.MustCompile(`game_id=([0-9]+)`)
	var metaRE = regexp.MustCompile(`#([0-9]+),.*aired.*([0-9]{4}-[0-9]{2}-[0-9]{2})`)
	var tapedRE = regexp.MustCompile(`Taped.*([0-9]{4}-[0-9]{2}-[0-9]{2})`)

	gameIDs := []int64{}
	showNums := map[int64]int64{}
	airedDates := map[int64]time.Time{}
	tapedDates := map[int64]time.Time{}

	c := colly.NewCollector(
		colly.CacheDir("./cache"),
	)

	c.OnHTML("div#content tr", func(e *colly.HTMLElement) {
		var gid int64
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			gameID := gameRE.FindStringSubmatch(el.Attr("href"))
			if gameID == nil {
				return
			}
			gid, _ = strconv.ParseInt(gameID[1], 10, 64)
			gameIDs = append(gameIDs, gid)

			taped := tapedRE.FindStringSubmatch(el.Attr("title"))
			if taped == nil {
				return
			}
			tapedDates[gid], _ = time.Parse("2006-01-02", taped[1])
		})

		e.ForEach("td", func(_ int, el *colly.HTMLElement) {
			meta := metaRE.FindStringSubmatch(el.Text)
			if meta == nil {
				return
			}
			showNums[gid], _ = strconv.ParseInt(meta[1], 10, 64)
			airedDates[gid], _ = time.Parse("2006-01-02", meta[2])
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.Visit(fmt.Sprintf("https://www.j-archive.com/showseason.php?season=%d", seasonID))

	games := []*ent.Game{}
	for _, gid := range gameIDs {
		games = append(games, &ent.Game{
			ID:       int(gid),
			SeasonID: int(seasonID),
			Show:     int(showNums[gid]),
			AirDate:  airedDates[gid],
			TapeDate: tapedDates[gid],
		})
	}

	sort.Slice(games, func(i, j int) bool {
		return games[i].Show < games[j].Show
	})

	return games, nil
}
