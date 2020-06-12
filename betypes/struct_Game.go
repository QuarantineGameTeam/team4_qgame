package betypes

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

type Game struct {
	Battlefield Field
	GameID      int64
}

func (g *Game) MakeMoves(bot *tgbotapi.BotAPI) {
	go func() {
		log.Println("Game started, ID", g.GameID)
		g.SendTeamsColor(bot)

		var moveNumber int
		var theClanLost bool
		for !theClanLost {
			for _, clan := range g.Battlefield.Clans {
				log.Println("Game make move, ID", g.GameID)
				go func() {
					bot.Send(tgbotapi.NewMessage(g.GameID, fmt.Sprintf("Move number %d\nWalking now %s", moveNumber, clan.Name)))
					moveNumber++
					clan.StartVotingWhereToGo(bot)
				}()
				go func() {
					//TODO: Here the logic of improvements while another clan thinks what cell to open
				}()
				<-time.After(time.Second * TimeToMove)
				clan.StopVoting(bot)
			}
			//TODO: Here you should calculate the clan booty (or something like that)
		}
		//TODO: It should show which team won, who got what prey and the rating of players who played
	}()
}

func (g *Game) SendTeamsColor(bot *tgbotapi.BotAPI) {
	for _, clan := range g.Battlefield.Clans {
		for _, user := range clan.Users {
			bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintf("🔹You plaing in %v team🔹", clan.Name)))
		}
	}
}
