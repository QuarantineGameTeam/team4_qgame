package actions

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"team4_qgame/betypes"
	"team4_qgame/db"
	"time"
)

var (
	gamesInLine = make(map[int64]betypes.Game)
	joinButton  = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Join to game! ", "join_to_game"),
		),
	)
)

func findUser(id int64) bool {
	if db.GetUser(strconv.FormatInt(id, 10)).Id == id {
		return true
	}
	return false
}

func findGame(id int64) bool {
	for _, game := range gamesInLine {
		if game.GameID == id {
			return true
		}
	}
	return false
}

func sendTimeBeforeRegistrationEnds(toEnd int, message *tgbotapi.MessageConfig, bot *tgbotapi.BotAPI) {
	message.Text = fmt.Sprintf("Ends in %d seconds", toEnd)
	bot.Send(message)
	time.Sleep(1 * time.Second)
}

func StartRecruitingForTheGame(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var passedFromTheBeginningOfRegistration int //It has passed since the beginning of registration
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("The set for the game has started!!\nEnds in %d seconds", betypes.TimeToRegister))
	go func() {
		msg.ReplyMarkup = joinButton
		bot.Send(msg)
		msg.ReplyMarkup = nil
		gamesInLine[update.Message.Chat.ID] = betypes.Game{
			PlayersInLine: nil,
			GameID:        update.Message.Chat.ID,
			GameStarted:   false,
		}
		go func() {
			for !(passedFromTheBeginningOfRegistration > betypes.TimeToRegister) {
				switch betypes.TimeToRegister - passedFromTheBeginningOfRegistration {
				case 30, 10, 5:
					sendTimeBeforeRegistrationEnds(betypes.TimeToRegister-passedFromTheBeginningOfRegistration, &msg, bot)
					break
				case 0:
					msg.Text = fmt.Sprint("Registration is complete!\nThe game begins!!!")
					bot.Send(msg)
					passedFromTheBeginningOfRegistration++
					//TODO: Logic that divides into teams
					delete(gamesInLine, update.Message.Chat.ID)
					break
				}
			}
		}()
		for ; passedFromTheBeginningOfRegistration < betypes.TimeToRegister; passedFromTheBeginningOfRegistration++ {
			time.Sleep(1 * time.Second)
		}
	}()
}

func AddAPlayerToTheQueueForTheGame(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.CallbackQuery != nil && update.CallbackQuery.Data == "join_to_game" {
		if findUser(int64(update.CallbackQuery.From.ID)) {
			user := db.GetUser(strconv.FormatInt(int64(update.CallbackQuery.From.ID), 10))
			gamesInLine[update.CallbackQuery.Message.Chat.ID].PlayersInLine[int64(update.CallbackQuery.From.ID)] = user
			bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("%s joined", update.CallbackQuery.From.FirstName)))
		}
	}
}
