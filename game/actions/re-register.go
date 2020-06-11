package actions

import (
	"team4_qgame/betypes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Reregister(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	addUser(update)
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, betypes.ReReregisterSuccessfulText))
}
