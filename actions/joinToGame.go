package actions

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"team4_qgame/betypes"
	"team4_qgame/db"
	"time"
)

type UserRequest struct {

}

var (
	users = []betypes.User{}
	joinButton =  tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Join to game! ", "1"),
		),
	)

	var UserQueue = make(chan)
)

func findUser(id int64) bool {
	if  db.GetUser(strconv.FormatInt(id, 10)).Id == id { return false }
	return true
}

func sendTimeBeforeRegistrationEnds(toEnd int, message *tgbotapi.MessageConfig, bot *tgbotapi.BotAPI){
	message.Text = fmt.Sprintf("Ends in %d seconds", toEnd)
	bot.Send(message)
	time.Sleep(1 * time.Second)
}

func StartRecruitingForTheGame (update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var passedFromTheBeginningOfRegistration int //It has passed since the beginning of registration
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("The set for the game has started!!\nEnds in %d seconds", betypes.TimeToRegister))
	
	go func() {
		msg.ReplyMarkup = joinButton
		bot.Send(msg)
		msg.ReplyMarkup = nil
		

		for !(passedFromTheBeginningOfRegistration > betypes.TimeToRegister) {
			switch betypes.TimeToRegister - passedFromTheBeginningOfRegistration {
			case 3, 2, 1:
				sendTimeBeforeRegistrationEnds(betypes.TimeToRegister - passedFromTheBeginningOfRegistration, &msg, bot)
				break
			case 0:
				msg.Text = fmt.Sprint("Registration is complete!\nThe game begins!!!")
				bot.Send(msg)
				passedFromTheBeginningOfRegistration++
				break
			}
		}
	}()

	for ; passedFromTheBeginningOfRegistration < betypes.TimeToRegister; passedFromTheBeginningOfRegistration++ {
		time.Sleep(1 * time.Second)
	}
}

func AddAPlayerToTheQueueForTheGame(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if findUser(int64(update.Message.From.ID)) {
		userDate := betypes.User{
			Id: int64(update.Message.From.ID),
			FirstName: update.Message.From.FirstName,
			Username: update.Message.From.UserName,
			Rank: 0 }
		db.SaveUser(userDate)

		users = append(users, userDate)
		return users
	} 
	return users
}
