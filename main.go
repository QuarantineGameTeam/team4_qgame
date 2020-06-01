package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"team4_qgame/betypes"
	"team4_qgame/db"
	"team4_qgame/loger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(betypes.BOT_TOKEN)
)

func getUser(msg *tgbotapi.Message) betypes.User {
	user := betypes.User{Id: msg.Chat.ID, FirstName: msg.Chat.FirstName, Username: msg.Chat.UserName, Rank: 0}
	fmt.Println(user)
	return user
}

func checkOnCommands(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var user betypes.User
	if update.Message.Text != "/start" {
		user = db.GetUser(strconv.Itoa(int(update.Message.Chat.ID)))
	}
	if update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case betypes.StartCommand:
			user = getUser(update.Message)
			db.SaveUser(user)
		case betypes.HelpCommand:
			msg.Text = betypes.HelpText

		default:
			msg.Text = betypes.UnclearCommandText
		}
		bot.Send(msg)

	}
	db.SaveUser(user)
}

func main() {

	go func() {
		log.Fatal(http.ListenAndServe(":"+betypes.BOT_PORT, nil))
	}()
	loger.ForError(BotErr, "BOT_TOKEN error")

	getUpdates(NewBot)
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
