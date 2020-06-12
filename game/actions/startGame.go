package actions

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"team4_qgame/betypes"
	"team4_qgame/db"
	"team4_qgame/game"
	"time"
)

func StartRecruitingForTheGame(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if isAChatGroup(update, bot) && !іsTheGameAlreadyInLine(update.Message.Chat.ID, bot) && !іsTheGameAlreadyStarted(update.Message.Chat.ID, bot) {
		go startRecruitingForTheGame(update.Message.Chat.ID, bot)
	}
}

func AddAPlayerToTheQueueForTheGame(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if isAChatGroup(update, bot) && update.CallbackQuery.Data == betypes.JoinToGameButtonData && isTheUserRegistered(update, bot) && !hasAlreadyJoined(update, bot) {
		u, _ := db.GetUser(int64(update.CallbackQuery.From.ID))
		u.SetQueueStatus(true)

		game.GamesInLine[update.CallbackQuery.Message.Chat.ID].PlayersInLine[int64(update.CallbackQuery.From.ID)] = *u
		db.SaveUser(*u)

		bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("🔥@%s join!!🔥", update.CallbackQuery.From.UserName)))
	}
}

func startRecruitingForTheGame(gameID int64, bot *tgbotapi.BotAPI) {
	var passedFromTheBeginningOfRegistration int
	addGameToLine(gameID)
	msg := tgbotapi.NewMessage(gameID,
		fmt.Sprintf("%s\nEnds in %d seconds",
			betypes.StartANewGameText, betypes.TimeToRegister),
	)

	msg.ReplyMarkup = betypes.JoinButton
	bot.Send(msg)
	msg.ReplyMarkup = nil

	go func() {
		for !(passedFromTheBeginningOfRegistration > betypes.TimeToRegister) {
			switch betypes.TimeToRegister - passedFromTheBeginningOfRegistration {
			case 30, 10, 5:
				sendTimeBeforeRegistrationEnds(gameID, betypes.TimeToRegister-passedFromTheBeginningOfRegistration, bot)
				time.Sleep(1 * time.Second)
				break
			case 0:
				passedFromTheBeginningOfRegistration++
				startGame(gameID, bot)
			}
		}
	}()
	go func() {
		for ; passedFromTheBeginningOfRegistration < betypes.TimeToRegister; passedFromTheBeginningOfRegistration++ {
			<-time.After(time.Second * 1)
		}
	}()
}

func startGame(gameId int64, bot *tgbotapi.BotAPI) {
	if whetherThereAreEnoughPlayersInLine(gameId, bot) {
		gameInLine := game.GamesInLine[gameId]
		createAGame := gameInLine.CreateAGame(game.CreateAField(betypes.FieldWidth, betypes.FieldHeight, betypes.NumberOfMines), &game.Games)
		createAGame.MakeMoves(bot)
		for _, user := range game.GamesInLine[gameId].PlayersInLine {
			user.SetQueueStatus(false)
			user.SetPlayingStatus(true)
			db.SaveUser(user)
		}
		bot.Send(tgbotapi.NewMessage(gameId, fmt.Sprint("⭐The game began. Open chat with the team's players!⭐")))
	}
}

func sendTimeBeforeRegistrationEnds(gameId int64, toEnd int, bot *tgbotapi.BotAPI) {
	bot.Send(tgbotapi.NewMessage(gameId, fmt.Sprintf("💥Ends in %d seconds💥", toEnd)))
}

func isTheUserRegistered(update tgbotapi.Update, bot *tgbotapi.BotAPI) bool {
	if update.Message != nil {
		_, err := db.GetUser(int64(update.Message.From.ID))
		if err == nil {
			return true
		}
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s @%s", betypes.UserIsNotRegisteredText, update.Message.From.UserName)))
	} else if update.CallbackQuery != nil {
		_, err := db.GetUser(int64(update.CallbackQuery.From.ID))
		if err == nil {
			return true
		}
		bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("%s @%s", betypes.UserIsNotRegisteredText, update.CallbackQuery.From.UserName)))
	}
	return false
}

func hasAlreadyJoined(update tgbotapi.Update, bot *tgbotapi.BotAPI) bool {
	u, _ := db.GetUser(int64(update.CallbackQuery.From.ID))
	if u.QueueStatus == true {
		bot.Send(tgbotapi.NewMessage(int64(update.CallbackQuery.Message.Chat.ID), fmt.Sprintf("🔥@%s %s🔥", update.CallbackQuery.From.UserName, betypes.AlreadyInLine)))
		return true
	}
	return false
}

func isAChatGroup(update tgbotapi.Update, bot *tgbotapi.BotAPI) bool {
	if update.Message != nil {
		if update.Message.Chat.Type == "group" {
			return true
		}
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("💥%s💥", betypes.EnrollmentInTheGame)))
	} else if update.CallbackQuery != nil {
		if update.CallbackQuery.Message.Chat.Type == "group" {
			return true
		}
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("💥%s💥", betypes.EnrollmentInTheGame)))
	}
	return false
}

func іsTheGameAlreadyInLine(gameID int64, bot *tgbotapi.BotAPI) bool {
	if game.FindGameInLine(gameID) {
		bot.Send(tgbotapi.NewMessage(gameID, fmt.Sprintf("💥%s💥", betypes.GameHasAlreadyBegun)))
		return true
	}
	return false
}

func іsTheGameAlreadyStarted(gameID int64, bot *tgbotapi.BotAPI) bool {
	if game.FindGame(gameID) {
		bot.Send(tgbotapi.NewMessage(gameID, fmt.Sprintf("💥%s💥", betypes.GameHasAlreadyBegun)))
		return true
	}
	return false
}

func addGameToLine(gameID int64) {
	game.GamesInLine[gameID] = betypes.GameInLine{
		PlayersInLine: map[int64]betypes.User{},
		GameID:        gameID,
	}
}

func whetherThereAreEnoughPlayersInLine(gameID int64, bot *tgbotapi.BotAPI) bool {
	if len(game.GamesInLine[gameID].PlayersInLine) >= betypes.MinimumNumberOfPlayers {
		return true
	}
	bot.Send(tgbotapi.NewMessage(gameID, fmt.Sprint(betypes.NotEnoughPlayersToStartTheGame)))
	return false
}
