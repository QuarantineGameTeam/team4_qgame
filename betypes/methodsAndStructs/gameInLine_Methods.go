package methodsAndStructs

import (
	"math/rand"
	"time"
)

func (g *GameInLine) CreateAGame(field Field, allGames *[]Game) Game { //TODO: Rewrite
	firstClan := make([]User, 0)
	secondClan := make([]User, 0)
	thirdClan := make([]User, 0)
	rand.Seed(time.Now().UnixNano())
	for _, user := range g.PlayersInLine {
		firstClan = append(firstClan, user)
	}
	/*var i int
	for _, user := range g.PlayersInLine {
		switch i {
		case 0:
			bot.Send(tgbotapi.NewMessage(user.Id, "You are in the blue team!!"))
			firstClan = append(firstClan, user)
		case 1:
			bot.Send(tgbotapi.NewMessage(user.Id, "You are in the red team!!"))
			secondClan = append(secondClan, user)
		case 2:
			bot.Send(tgbotapi.NewMessage(user.Id, "You are in the green team!!"))
			thirdClan = append(thirdClan, user)
			i = 0
		}
		i++
	}*/
	for i := 0; i < len(field.Clans); i++ {
		switch i {
		case 0:
			field.Clans[i].Users = firstClan
		case 1:
			field.Clans[i].Users = secondClan
		case 2:
			field.Clans[i].Users = thirdClan
		}
	}
	newGame := Game{
		Battlefield: field,
		GameID:      g.GameID,
	}
	*allGames = append(*allGames, newGame)
	return newGame
}
