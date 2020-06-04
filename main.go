package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"team4_qgame/actions"
	"team4_qgame/betypes"
	"team4_qgame/loger"
)

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(betypes.BOT_TOKEN)
)

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":"+betypes.BOT_PORT, nil))
	}()
	loger.ForError(BotErr, "BOT_TOKEN error")

	getUpdates(NewBot)
}

func checkOnCommands(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil {
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case betypes.StartCommand:
				msg.Text = betypes.StartText
			case betypes.HelpCommand:
				msg.Text = betypes.HelpText
			case "startgame":
				actions.StartRecruitingForTheGame(update, bot)
			default:
				msg.Text = betypes.UnclearCommandText
			}
			bot.Send(msg)
		}
	}
	if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case "join_to_game":
			actions.AddAPlayerToTheQueueForTheGame(update, bot)
		}
	}
}

func setWebhook(bot *tgbotapi.BotAPI) {
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(betypes.WEB_HOOK))
	loger.ForError(err, "setting WEB_HOOK", betypes.WEB_HOOK, "error")
}

func getUpdates(bot *tgbotapi.BotAPI) {
	setWebhook(bot)
	updates := bot.ListenForWebhook("/")
	for update := range updates {
		checkOnCommands(update, bot)
	}
}
