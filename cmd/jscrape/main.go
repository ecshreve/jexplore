package main

import (
	"context"

	"github.com/ecshreve/jexplore/ent"
	"github.com/ecshreve/jexplore/jscrape"
)

func main() {
	ctx := context.Background()
	client, err := jscrape.InitClient(ctx)
	if err != nil {
		panic(err)
	}

	for i := 11; i <= 25; i++ {
		games, err := jscrape.ScrapeSeasonGames(int64(i))
		if err != nil {
			panic(err)
		}

		toCreate := []*ent.GameCreate{}
		for _, g := range games {
			toCreate = append(toCreate, client.Game.Create().SetID(g.ID).SetSeasonID(g.SeasonID).SetShow(g.Show).SetAirDate(g.AirDate).SetTapeDate(g.TapeDate))
		}

		_, err = client.Game.CreateBulk(toCreate...).Save(ctx)
		if err != nil {
			panic(err)
		}
	}
}

// _, catMap := jscrape.ScrapeGameClues(8040)

// 	toCreateMap := map[string]*ent.CategoryCreate{}
// 	for _, cat := range catMap {
// 		c, err := client.Category.Query().Where(category.Name(cat)).Only(ctx)
// 		if err == nil && c != nil {
// 			continue
// 		}

// 		toCreateMap[cat] = client.Category.Create().SetName(cat)
// 	}

// 	toCreateList := []*ent.CategoryCreate{}
// 	for _, c := range toCreateMap {
// 		toCreateList = append(toCreateList, c)
// 	}

// 	freshClient, err := jscrape.InitClient(ctx)
// 	if err != nil {
// 		panic(err)
// 	}

// 	ret, err := freshClient.Category.CreateBulk(toCreateList...).Save(ctx)
// 	if err != nil {
// 		panic(err)
// 	}

// 	pretty.Println(ret)
// }
