package mechanics

import (
	"fmt"
	"math/rand"
	"strconv"
	"team4_qgame/betypes"
	"team4_qgame/db"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)
func OfferAnAlliance(sender *betypes.Clan, receiver *betypes.Clan, bot *tgbotapi.BotAPI) {
	var PollingKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Accept", "1"),
			tgbotapi.NewInlineKeyboardButtonData("Reject", "2"),
		),
	)
	pollCounterFor := 0
	pollCounterAgainst := 0
  //voting in clan of receiver
	go func() {
		for _, user := range *receiver.Users {
			bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintf("You've received an offer for alliance!")))
			msg := tgbotapi.NewMessage(user.Id, "")
			msg.ReplyMarkup = PollingKeyboard
			if tgbotapi.Update.CallbackQuery != nil {
				switch tgbotapi.Update.CallbackQuery.Data {
				case "Accept":
					pollCounterFor++
					msg.Text = fmt.Sprint("You voted for alliance,\n")
					break
				case "Reject":
					pollCounterAgainst++
					msg.Text = fmt.Sprint("You voted against alliance.\n")
					break
				}
				bot.Send(msg)
			}
		}
		<-time.After(time.Second * 60)
	}()
	if pollCounterFor >= pollCounterAgainst {
		for _, user := range *sender.Users {
			bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintf("Your offer for alliance has been accepted.")))
		}
    //ubite resources of both clans
		sender.Number += receiver.Number
		receiver.Number += sender.Number
		for _, mineR := range *receiver.Mines {
			*sender.Mines = append(*sender.Mines, mineR)
		}
		for _, mineS := range *sender.Mines {
			*receiver.Mines = append(*receiver.Mines, mineS)
		}
		for _, user := range *receiver.Users {
			*sender.Users = append(*sender.Users, user)
		}
		for _, user := range *sender.Users {
			*receiver.Users = append(*receiver.Users, user)
		}
	} else {
		for _, user := range *sender.Users {
			bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintf("Your offer for alliance has been rejected.")))
		}
	}
}

