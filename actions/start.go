package actions

import (
	"fmt"
	"team4_qgame/betypes"
	"team4_qgame/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func addUser(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, betypes.StartText+"\n"+betypes.RegistrationIsSuccessfulText))
	db.SaveUser(betypes.User{
		Id:        int64(update.Message.From.ID),
		FirstName: update.Message.From.FirstName,
		Username:  update.Message.From.UserName,
		Rank:      0,
	})
}

func StartCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if len(db.GetAllKeys()) == 0 {
		addUser(update, bot)
		return
	}
	if !findUser(int64(update.Message.From.ID)) {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, betypes.StartText+"\n"+betypes.RegistrationIsSuccessfulText))
		addUser(update, bot)
		return
	}
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("@%s \n%s", update.Message.From.UserName, betypes.AlreadyRegisteredText)))
}
