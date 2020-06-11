package methodsAndStructs

import (
	"team4_qgame/betypes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (v *WhereToGoVoting) AddAVoice(forWhat string, clanLocation *Point, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), "<empty>")
		switch forWhat {
		case betypes.LeftArrow:
			if clanLocation.X-1 < 1 {
				msg.Text = "Oh no. Here is the ocean.Rather, choose another direction!!"
				msg.ReplyMarkup = betypes.SelectMoveButton
			} else {
				v.LeftNumberOfVotes++
			}
		case betypes.UpArrow:
			if clanLocation.Y-1 < 1 {
				msg.Text = "Oh no. Your path is blocked by mountains. Rather choose another direction!!"
				msg.ReplyMarkup = betypes.SelectMoveButton
			} else {
				v.UpNumberOfVotes++
			}
		case betypes.RightArrow:
			if clanLocation.X+1 > betypes.FieldWidth {
				msg.Text = "Oh no. There is an active volcano. Rather choose another direction!!"
				msg.ReplyMarkup = betypes.SelectMoveButton
			} else {
				v.RightNumberOfVotes++
			}
		case betypes.DownArrow:
			if clanLocation.Y+1 > betypes.FieldHeight {
				msg.Text = "Oh no. This is a restricted area. You need permission.. Rather choose another direction!!"
				msg.ReplyMarkup = betypes.SelectMoveButton
			} else {
				v.DownNumberOfVotes++
			}
		}
		bot.Send(msg)
	}
}

func (voting *WhereToGoVoting) ClearVoting() {
	voting.DownNumberOfVotes = 0
	voting.LeftNumberOfVotes = 0
	voting.RightNumberOfVotes = 0
	voting.UpNumberOfVotes = 0
}
