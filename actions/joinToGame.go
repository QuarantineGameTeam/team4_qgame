package actions

import (
	"strconv"
	"team4_qgame/betypes"
	"team4_qgame/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

)

// checks if the user is already in the database
func findUser(id int64) bool {
	if  db.GetUser(strconv.FormatInt(id, 10)).Id == id { return false }
	return true
}

//join new user to database
func JoinGame(update tgbotapi.Update) bool {
	if findUser(int64(update.Message.From.ID)) {
		userDate := betypes.User{
			Id: int64(update.Message.From.ID),
			FirstName: update.Message.From.FirstName,
			Username: update.Message.From.UserName,
			Rank: 0 }
		db.SaveUser(userDate)
		return true
	}
	return false
}