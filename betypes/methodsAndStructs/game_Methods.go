package methodsAndStructs

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"team4_qgame/betypes"
	"time"
)

func (g *Game) MakeMoves(bot *tgbotapi.BotAPI) {
	go func() {
		log.Println("Game started, ID", g.GameID)
		var theClanLost bool
		for !theClanLost {
			for _, clan := range g.Battlefield.Clans {
				go func() {
					clans := make([]Clan, 0)
					for _, c := range g.Battlefield.Clans {
						if c.Number != clan.Number {
							clans = append(clans, c)
						}
					}
					clan.StartVotingWhereToGo(bot)
				}()
				go func() {
					//TODO: Here the logic of improvements while another clan thinks what cell to open
					fmt.Println("TODO: Here the logic of improvements while another clan thinks what cell to open")
				}()
				<-time.After(time.Second * betypes.TimeToMove)
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
