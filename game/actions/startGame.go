package actions

import (
	"fmt"
	"team4_qgame/betypes"
	"team4_qgame/betypes/methodsAndStructs"
	"team4_qgame/db"
	"team4_qgame/game"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func StartRecruitingForTheGame(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	go func() {
		if update.Message.Chat.Type != "private" {
			if !game.FindGameInLine(update.Message.Chat.ID) && !game.FindGame(update.Message.Chat.ID) {
				var passedFromTheBeginningOfRegistration int //It has passed since the beginning of registration
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\nEnds in %d seconds", betypes.StartANewGameText, betypes.TimeToRegister))
				go func() {
					msg.ReplyMarkup = betypes.JoinButton
					bot.Send(msg)
					msg.ReplyMarkup = nil
					game.GamesInLine[update.Message.Chat.ID] = methodsAndStructs.GameInLine{
						PlayersInLine: map[int64]methodsAndStructs.User{},
						GameID:        update.Message.Chat.ID,
					}
					go func() {
						for !(passedFromTheBeginningOfRegistration > betypes.TimeToRegister) {
							switch betypes.TimeToRegister - passedFromTheBeginningOfRegistration {
							case 30, 10, 5:
								sendTimeBeforeRegistrationEnds(betypes.TimeToRegister-passedFromTheBeginningOfRegistration, &msg, bot)
								break
							case 0:
								passedFromTheBeginningOfRegistration++
								startGame(bot, update)
							}
						}
					}()
					for ; passedFromTheBeginningOfRegistration < betypes.TimeToRegister; passedFromTheBeginningOfRegistration++ {
						<-time.After(time.Second * 1)
					}
				}()
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", betypes.GameHasAlreadyBegun)))
			}
		} else {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", betypes.EnrollmentInTheGame)))
		}
	}()
}

func AddAPlayerToTheQueueForTheGame(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.CallbackQuery.Message.Chat.Type == "group" {
		if len(game.GamesInLine) != 0 && update.CallbackQuery != nil && update.CallbackQuery.Data == betypes.JoinToGameButtonData {
			u, err := db.GetUser(int64(update.CallbackQuery.From.ID))
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,
					fmt.Sprintf("@%s %s \nFor registration @%s",
						update.CallbackQuery.From.UserName, betypes.UserIsNotRegisteredText, bot.Self.UserName)))
			} else {
				if !u.QueueStatus {
					game.GamesInLine[update.CallbackQuery.Message.Chat.ID].PlayersInLine[int64(update.CallbackQuery.From.ID)] = *u
					u.SetQueueStatus(true)
					db.SaveUser(*u)
					bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("@%s joined!!", update.CallbackQuery.From.UserName)))
				} else {
					bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("@%s %s", update.CallbackQuery.From.UserName, betypes.AlreadyInLine)))
				}
			}
		}
	} else {
		bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("%s", betypes.EnrollmentInTheGame)))
	}
}

func sendTimeBeforeRegistrationEnds(toEnd int, message *tgbotapi.MessageConfig, bot *tgbotapi.BotAPI) {
	message.Text = fmt.Sprintf("Ends in %d seconds", toEnd)
	bot.Send(message)
	time.Sleep(1 * time.Second)
}

func startGame(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint("Registration is complete!\nThe game begins!!!")))
	if len(game.GamesInLine[update.Message.Chat.ID].PlayersInLine) >= 1 /*If there are enough players the game will start*/ {
		gameInLine := game.GamesInLine[update.Message.Chat.ID]
		game := gameInLine.CreateAGame(game.CreateAField(betypes.FieldWidth, betypes.FieldHeight, betypes.NumberOfMines), &game.Games)
		game.MakeMoves(bot) //Starteds
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint("The game began. Open chat with the team's players!")))
	} else {
		for _, user := range game.GamesInLine[update.Message.Chat.ID].PlayersInLine {
			user.SetPlayingStatus(false)
			user.SetQueueStatus(false)
			db.SaveUser(user)
		}
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint("Not enough players!!")))
	}
	delete(game.GamesInLine, update.Message.Chat.ID) //Remove game from the queue
}
