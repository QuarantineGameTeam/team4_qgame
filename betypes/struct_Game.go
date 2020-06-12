package betypes

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Game struct {
	Battlefield Field
	GameID      int64
}

func (g *Game) MakeMoves(bot *tgbotapi.BotAPI) {
	go startMakeMoves(g, bot)
}

func startMakeMoves(g *Game, bot *tgbotapi.BotAPI) {
	log.Println("Game started, ID", g.GameID)
	g.SendTeamsColor(bot)

	var moveNumber int
	var theClanLost bool

	for !theClanLost {
		for _, clan := range g.Battlefield.Clans {
			log.Println("Game make move, ID", g.GameID)
			sendMoveInfo(bot, g.GameID, moveNumber, clan.Name)

			go clan.StartVotingWhereToGo(bot)
			<-time.After(time.Second * TimeToMove)
			clan.StopVoting(bot)

			for _, user := range clan.Users {
				bot.Send(tgbotapi.NewMessage(user.Id, "🔺Voting is over!🔺"))
			}

			//isThereAnother, clanNumber, clanName := clan.IsOtherClanInPoint(g.Battlefield.Clans)
			//if isThereAnother {
			//	bot.Send(tgbotapi.NewMessage(g.GameID, fmt.Sprintf("🔺%s found %s🔺", clan.Name, clanName)))
			//}
			//offerAnAllianceOrWar(clan, g.Battlefield.Clans, isThereAnother, *clanNumber, *clanName, bot)
		}
	}
}

func offerAnAllianceOrWar(sender Clan, clans []Clan, isThereAnother bool, clanNumber int, clanName string, bot *tgbotapi.BotAPI) {
	if isThereAnother {
		for _, clan := range clans {
			if clan.Number == clanNumber {

			}
		}
	}
}

func sendMoveInfo(bot *tgbotapi.BotAPI, gameId int64, moveNumber int, clanName string) {
	bot.Send(tgbotapi.NewMessage(gameId, fmt.Sprintf("🌟Move number %d\nWalking now %s🌟", moveNumber, clanName)))
}

func (g *Game) SendTeamsColor(bot *tgbotapi.BotAPI) {
	for _, clan := range g.Battlefield.Clans {
		for _, user := range clan.Users {
			bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintf("🔹You plaing in %v team🔹", clan.Name)))
		}
	}
}
