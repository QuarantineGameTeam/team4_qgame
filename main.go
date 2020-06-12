package main

import (
	"log"
	"net/http"
	"team4_qgame/betypes"
	"team4_qgame/game"
	"team4_qgame/game/actions"
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
		} else if game.FindPlayerInGame(int64(update.Message.From.ID)) && update.Message.Chat.Type == "private" /*Only in private chats*/ {
			for _, clan := range game.GetGame(game.GetGameIDByPlayerID(int64(update.Message.From.ID))).Battlefield.Clans { //Game Chat
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
		case betypes.JoinToGameButtonData:
			actions.AddAPlayerToTheQueueForTheGame(update, bot)
			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)) //In order for the Telegram to understand what we answered
		case betypes.UpArrow, betypes.DownArrow, betypes.LeftArrow, betypes.RightArrow:
			for _, clan := range game.GetGame(game.GetGameIDByPlayerID(int64(update.CallbackQuery.From.ID))).Battlefield.Clans {
				for _, user := range clan.Users {
					if user.Id == int64(update.CallbackQuery.From.ID) {
						clan.WhereToGo.AddAVoice(update.CallbackQuery.Data, &clan.Location, update, bot)
						bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)) //In order for the Telegram to understand what we answered
						actions.ReremoveKeyboard(bot, update)
					}
				}
			}
		}
	}
}

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":"+betypes.BOT_PORT, nil))
	}()
	log.Printf("Authorized on account %s", NewBot.Self.UserName)

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
