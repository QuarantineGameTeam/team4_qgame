package mechanics

import (
	"fmt"
	"team4_qgame/betypes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Watching another clan function

func Watch(sender *betypes.Clan, receiver *betypes.Clan, bot *tgbotapi.BotAPI) {
	go func() {
		for _, user := range receiver.Users {
			msg := tgbotapi.NewMessage(user.Id, fmt.Sprintf("Coordinates of clan #%d: (%d %d)", sender.Number, sender.Location.X, sender.Location.Y))
			bot.Send(msg)
		}
	}()
}
