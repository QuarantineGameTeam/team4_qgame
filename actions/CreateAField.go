package actions

import (
	"math/rand"
	"team4_qgame/betypes"
	"time"
)

func CreateAField(width, height int) betypes.Field {
	var (
		clan []betypes.Clan
		mine []betypes.Mine
	)
	clan = make([]betypes.Clan, 3)
	mine = make([]betypes.Mine, 5)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < cap(mine); i++ {
		mine[i] = betypes.Mine{Location: betypes.Point{X: rand.Intn(width) + 1, Y: rand.Intn(height) + 1}, Level: 1, ToBelong: 0}
	}
	for i := 0; i < cap(clan); i++ {
		clan[i] = betypes.Clan{Location: betypes.Point{X: rand.Intn(height) + 1, Y: rand.Intn(width) + 1}, ClanMine: mine[i], Health: 0, Users: nil, Number: 0}
	}
	field := betypes.Field{Mines: &mine, Clans: &clan}
	return field
}
