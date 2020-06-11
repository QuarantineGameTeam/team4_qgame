package actions

import (
	"team4_qgame/betypes"
	"team4_qgame/betypes/methodsAndStructs"
	"team4_qgame/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func addUser(update tgbotapi.Update) {
	db.SaveUser(methodsAndStructs.User{
		Id:        int64(update.Message.From.ID),
		FirstName: update.Message.From.FirstName,
		Username:  update.Message.From.UserName,
		Rank:      0,
	})
}

func StartCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, betypes.StartText+"\n"+betypes.RegistrationIsSuccessfulText))
	addUser(update)
}
