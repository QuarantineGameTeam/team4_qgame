package game

import (
	"math/rand"
	"team4_qgame/betypes/methodsAndStructs"
	"time"
)

func CreateAField(width, height, numberOfMines int) methodsAndStructs.Field {
	clan := make([]methodsAndStructs.Clan, 0)
	mine := make([]methodsAndStructs.Mine, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numberOfMines; i++ {
		mine = append(mine, methodsAndStructs.Mine{
			Location: methodsAndStructs.Point{X: rand.Intn(width) + 1, Y: rand.Intn(height) + 1},
			Level:    1,
			ToBelong: 0,
		})
	}
	for i := 0; i < 3; /*Number of clans*/ i++ {
		clan = append(clan, methodsAndStructs.Clan{
			Location: methodsAndStructs.Point{X: rand.Intn(height) + 1, Y: rand.Intn(width) + 1},
			ClanMine: mine[i],
			Health:   0,
			Users:    nil,
			Number:   0,
		})
	}
	field := methodsAndStructs.Field{Mines: mine, Clans: clan}
	return field
}
