package actions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"team4_qgame/betypes"
	"team4_qgame/db"
)

func StartCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, betypes.StartText)
	db.SaveUser(betypes.User{
		Id:        update.Message.Chat.ID,
		FirstName: update.Message.Chat.FirstName,
		Username:  update.Message.Chat.UserName,
		Rank:      0,
	})
	bot.Send(msg)
}
