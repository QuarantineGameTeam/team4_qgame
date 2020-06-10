package actions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"team4_qgame/betypes"
)

func Reregister(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	addUser(update, bot)
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, betypes.ReReregisterSuccessfulText))
}
