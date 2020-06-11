package methodsAndStructs

import (
	"team4_qgame/betypes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Clan) AddUser(users User) {
	c.Users = make([]User, 0)
	c.Users = append(c.Users, users)
}

func (c *Clan) StartVotingWhereToGo(bot *tgbotapi.BotAPI) {
	go func() {
		for _, user := range c.Users {
			msg := tgbotapi.NewMessage(user.Id, "Voting has started!! Please select where you want to go!!")
			msg.ReplyMarkup = betypes.SelectMoveButton
			bot.Send(msg)
		}
	}()
}
