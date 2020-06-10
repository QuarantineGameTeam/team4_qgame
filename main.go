package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"team4_qgame/betypes"
	"team4_qgame/loger"
)

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(betypes.BOT_TOKEN)
)

func setWebhook(bot *tgbotapi.BotAPI) {
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(betypes.WEB_HOOK))
	loger.ForrError(err, "setting WEB_HOOK", betypes.WEB_HOOK, "error")
}

func getUpdates(bot *tgbotapi.BotAPI) {
	setWebhook(bot)
	updates := bot.ListenForWebhook("/")
	for update := range updates {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)); err != nil {
			loger.LogFile.Fatal(err)
		}
	}
}

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":"+betypes.BOT_PORT, nil))
	}()
	loger.ForrError(BotErr, "BOT_TOKEN error")

	getUpdates(NewBot)
}
