package mechanics

import (
	"fmt"
	"team4_qgame/betypes"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendMessageFormChat(bot *tgbotapi.BotAPI, clan betypes.Clan, update tgbotapi.Update, users ...betypes.User) {
	go func() {
		if update.Message.Text != " " {
			text := ""
			//getMessage := update.Message.Text
			//getUserName := update.Message.From.UserName
			//getFirstName := update.Message.From.FirstName

			if len(users) == 0 {
				for _, user := range clan.Users {
					msg := tgbotapi.NewMessage(user.Id, "")
					text = textMassage(update, clan, false)
					msg.Text = text
					bot.Send(msg)
				}
			} else {
				for _, user := range users {
					msg := tgbotapi.NewMessage(user.Id, "")
					text = textMassage(update, clan, true)
					msg.Text = text
					bot.Send(msg)
				}
			}
		}
		<-time.After(time.Second * 60)
	}()
}

func textMassage(update tgbotapi.Update, clan betypes.Clan, writeClan bool) string {
	text := ""
	getMessage := update.Message.Text
	getUserName := update.Message.From.UserName
	getFirstName := update.Message.From.FirstName

	switch writeClan {
	case false:
		if getUserName == "" {
			text = fmt.Sprintf("*%s*\n  %s", getFirstName, getMessage)
		} else {
			text = fmt.Sprintf("*%s*\n  %s", getUserName, getMessage)
		}
	case true:
		getClanNumber := clan.Number

		if getUserName == "" {
			text = fmt.Sprintf("*%s* ( _%v_  clan ) \n  %s", getFirstName, getClanNumber, getMessage)
		} else {
			text = fmt.Sprintf("*%s* ( _%v_  clan ) \n  %s", getUserName, getClanNumber, getMessage)
		}
	}

	return text
}
