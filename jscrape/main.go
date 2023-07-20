// package jscrape is a WIP tool for scraping j-archive.com.
package jscrape

import (
	"context"

	"github.com/ecshreve/jexplore/ent"
	"github.com/ecshreve/jexplore/ent/category"
	"github.com/ecshreve/jexplore/ent/game"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	client, err := InitClient(ctx)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 0; i++ {
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

	for seasonID := 1; seasonID <= 0; seasonID++ {
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
