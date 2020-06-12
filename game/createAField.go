package game

import (
	"math/rand"
	"team4_qgame/betypes"
	"time"
)

func CreateAField(width, height, numberOfMines int) betypes.Field {
	clan := make([]betypes.Clan, 0)
	mine := make([]betypes.Mine, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numberOfMines; i++ {
		mine = append(mine, betypes.Mine{
			Location: betypes.Point{X: rand.Intn(width) + 1, Y: rand.Intn(height) + 1},
			Level:    1,
			ToBelong: 0,
		})
	}
	for i := 0; i < 3; /*Number of clans*/ i++ {
		clan = append(clan, betypes.Clan{
			Location:  betypes.Point{X: rand.Intn(height) + 1, Y: rand.Intn(width) + 1},
			ClanMine:  mine[i],
			Health:    0,
			Users:     nil,
			Number:    0,
			Name:      betypes.NamesOfClans[i+1],
			WhereToGo: betypes.WhereToGoVoting{},
		})
	}
	field := betypes.Field{Mines: mine, Clans: clan}
	return field
}
