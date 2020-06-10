package actions

import (
	"fmt"
	"strconv"
	"team4_qgame/betypes"
	"team4_qgame/db"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const JoinToGame = "join_to_game"

var joinButton = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Join to game! ", JoinToGame),
	),
)

func sendTimeBeforeRegistrationEnds(toEnd int, message *tgbotapi.MessageConfig, bot *tgbotapi.BotAPI) {
	message.Text = fmt.Sprintf("Ends in %d seconds", toEnd)
	bot.Send(message)
	time.Sleep(1 * time.Second)
}

func StartRecruitingForTheGame(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	go func() {
		if update.Message.Chat.Type != "private" {
			if !FindGameInLine(update.Message.Chat.ID) {
				var passedFromTheBeginningOfRegistration int //It has passed since the beginning of registration
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\nEnds in %d seconds", betypes.StartANewGameText, betypes.TimeToRegister))
				go func() {
					msg.ReplyMarkup = joinButton
					bot.Send(msg)
					msg.ReplyMarkup = nil
					gamesInLine[update.Message.Chat.ID] = betypes.GameInLine{
						PlayersInLine: map[int64]betypes.User{},
						GameID:        update.Message.Chat.ID,
					}
					go func() {
						for !(passedFromTheBeginningOfRegistration > betypes.TimeToRegister) {
							switch betypes.TimeToRegister - passedFromTheBeginningOfRegistration {
							case 30, 10, 5:
								sendTimeBeforeRegistrationEnds(betypes.TimeToRegister-passedFromTheBeginningOfRegistration, &msg, bot)
								break
							case 0:
								bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint("Registration is complete!\nThe game begins!!!")))
								passedFromTheBeginningOfRegistration++
								if len(gamesInLine[update.Message.Chat.ID].PlayersInLine) >= 3 {
									gameInLine := gamesInLine[update.Message.Chat.ID]
									game := gameInLine.CreateAGame(createAField(betypes.FieldWidth, betypes.FieldHeight, betypes.NumberOfMines), update, bot, &games)
									game.MakeMoves()
								}
								delete(gamesInLine, update.Message.Chat.ID)
								break
							}
						}
					}()
					for ; passedFromTheBeginningOfRegistration < betypes.TimeToRegister; passedFromTheBeginningOfRegistration++ {
						<-time.After(time.Second * 1)
					}
				}()
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", betypes.RecruitmentForTheGameHasAlreadyBegun)))
			}
		} else {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", betypes.EnrollmentInTheGame)))
		}
	}()
}

func AddAPlayerToTheQueueForTheGame(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.CallbackQuery.Message.Chat.Type == "group" {
		if len(gamesInLine) != 0 && update.CallbackQuery != nil && update.CallbackQuery.Data == JoinToGame {
			if FindUser(int64(update.CallbackQuery.From.ID)) {
				user := db.GetUser(strconv.FormatInt(int64(update.CallbackQuery.From.ID), 10))
				if !user.QueueStatus {
					gamesInLine[update.CallbackQuery.Message.Chat.ID].PlayersInLine[int64(update.CallbackQuery.From.ID)] = user
					user.SetQueueStatus(true)
					db.SaveUser(user)
					bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("@%s joined!!", update.CallbackQuery.From.UserName)))
				} else {
					bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("@%s %s", update.CallbackQuery.From.UserName, betypes.AlreadyInLine)))
				}
			} else {
				bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,
					fmt.Sprintf("@%s %s \nFor registration @%s", update.CallbackQuery.From.UserName, betypes.UserIsNotRegisteredText, bot.Self.UserName)))
			}
		}
	} else {
		bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("%s", betypes.EnrollmentInTheGame)))
	}
}
