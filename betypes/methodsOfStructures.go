package betypes

import (
	"fmt"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (u *User) SetRank(newRank int64) {
	u.Rank = newRank
}

func (u *User) SetPlayingStatus(status bool) {
	u.PlayingStatus = status
}

func (u *User) SetQueueStatus(status bool) {
	u.QueueStatus = status
}

func (c *Clan) AddUser(users User) {
	c.Users = make([]User, 0)
	c.Users = append(c.Users, users)
}

func (f *Field) AddMine(mine Mine) {
	f.Mines = append(f.Mines, mine)
}

func (f *Field) AddClan(clan Clan) {
	f.Clans = append(f.Clans, clan)
}

func (g *GameInLine) CreateAGame(field Field, update tgbotapi.Update, bot *tgbotapi.BotAPI, allGames *[]Game) Game { //TODO: Rewrite
	firstClan := make([]User, 0)
	secondClan := make([]User, 0)
	thirdClan := make([]User, 0)
	rand.Seed(time.Now().UnixNano())
	var i int
	for _, user := range g.PlayersInLine {
		switch i {
		case 0:
			firstClan = append(firstClan, user)
		case 1:
			secondClan = append(secondClan, user)
		case 2:
			thirdClan = append(thirdClan, user)
			i = 0
		}
		i++
	}

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

func (g *Game) MakeMoves() {
	go func() {
		var theClanLost bool
		for !theClanLost {
			for i := 0; i < len(g.Battlefield.Clans); i++ {
				go func() {
					//TODO: Make logic for sending votes
					fmt.Println("TODO: Make logic for sending votes")
				}()
				go func() {
					//TODO: Here the logic of improvements while another clan thinks what cell to open
					fmt.Println("TODO: Here the logic of improvements while another clan thinks what cell to open")
				}()
				<-time.After(time.Second * TimeToMove)
				//TODO: You can also display how much time is left
				fmt.Println("TODO: You can also display how much time is left")
			}
			//TODO: Here you should calculate the clan booty (or something like that)
			fmt.Println("TODO: Here you should calculate the clan booty (or something like that)")
		}
		//TODO: It should show which team won, who got what prey and the rating of players who played
		fmt.Println("TODO: It should show which team won, who got what prey and the rating of players who played")
	}()
}
