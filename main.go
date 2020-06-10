package main

import (
	"log"
	"net/http"
	"team4_qgame/actions"
	"team4_qgame/betypes"
	"team4_qgame/game/mechanics"
	"team4_qgame/loger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(betypes.BOT_TOKEN)
)

func checkOnCommands(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil {
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case betypes.StartCommand:
				actions.StartCommand(update, bot)
			case betypes.HelpCommand:
				msg.Text = betypes.HelpText
			case betypes.StartANewGameCommand:
				actions.StartRecruitingForTheGame(update, bot)
			case betypes.ReregisterCommand:
				actions.Reregister(update, bot)
			default:
				msg.Text = betypes.UnclearCommandText
			}
			bot.Send(msg)
		} else if actions.FindPlayerInGame(int64(update.Message.From.ID)) && update.Message.Chat.Type == "private" /*Only in private chats*/ {
			gameCRT := actions.GetGame(actions.GetThePlayersGameID(int64(update.Message.From.ID)))
			for _, clan := range gameCRT.Battlefield.Clans {
				for _, user := range clan.Users {
					if user.Id == int64(update.Message.From.ID) {
						mechanics.SendMessageFormChat(bot, clan, update)
					}
				}
			}
		}
	}
	if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case actions.JoinToGame:
			actions.AddAPlayerToTheQueueForTheGame(update, bot)
		case mechanics.UpArrow, mechanics.DownArrow, mechanics.LeftArrow, mechanics.RightArrow:

		}
	}
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
