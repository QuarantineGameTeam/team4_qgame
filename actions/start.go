package actions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"team4_qgame/betypes"
	"team4_qgame/db"
)

func StartCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if !findUser(int64(update.Message.From.ID)) {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, betypes.StartText+"\n"+betypes.RegistrationIsSuccessfulText))
		db.SaveUser(betypes.User{
			Id:        update.Message.Chat.ID,
			FirstName: update.Message.Chat.FirstName,
			Username:  update.Message.Chat.UserName,
			Rank:      0,
		})
	} else {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, betypes.AlreadyRegisteredText))
	}
}
