package betypes

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type WhereToGoVoting struct {
	UpNumberOfVotes    int
	DownNumberOfVotes  int
	LeftNumberOfVotes  int
	RightNumberOfVotes int
	//StayNumberOfVotes  int
}

func (v *WhereToGoVoting) AddAVoice(forWhat string, clanLocation *Point, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), "<empty>")
		switch forWhat {
		case LeftArrow:
			if clanLocation.X-1 < 1 {
				msg.Text = "Oh no. Here is the ocean.Rather, choose another direction!!"
				msg.ReplyMarkup = SelectMoveButton
			} else {
				v.LeftNumberOfVotes++
			}
		case UpArrow:
			if clanLocation.Y-1 < 1 {
				msg.Text = "Oh no. Your path is blocked by mountains. Rather choose another direction!!"
				msg.ReplyMarkup = SelectMoveButton
			} else {
				v.UpNumberOfVotes++
			}
		case RightArrow:
			if clanLocation.X+1 > FieldWidth {
				msg.Text = "Oh no. There is an active volcano. Rather choose another direction!!"
				msg.ReplyMarkup = SelectMoveButton
			} else {
				v.RightNumberOfVotes++
			}
		case DownArrow:
			if clanLocation.Y+1 > FieldHeight {
				msg.Text = "Oh no. This is a restricted area. You need permission.. Rather choose another direction!!"
				msg.ReplyMarkup = SelectMoveButton
			} else {
				v.DownNumberOfVotes++
			}
		}
		bot.Send(msg)
	}
}

func (v *WhereToGoVoting) SearchMax() string {
	maxWidth := 0
	maxDirWidth := ""
	maxHight := 0
	maxDirHight := ""

	if v.LeftNumberOfVotes-v.RightNumberOfVotes < 0 {
		maxWidth = v.RightNumberOfVotes
		maxDirWidth = "right"
	} else {
		maxWidth = v.LeftNumberOfVotes
		maxDirWidth = "left"
	}

	if v.UpNumberOfVotes-v.DownNumberOfVotes < 0 {
		maxHight = v.DownNumberOfVotes
		maxDirHight = "down"
	} else {
		maxHight = v.UpNumberOfVotes
		maxDirHight = "up"
	}

	if maxWidth-maxHight < 0 {
		return maxDirHight
	} else {
		return maxDirWidth
	}
}

func (v *WhereToGoVoting) GenerateNewPoint(dir string, p *Point) {
	switch dir {
	case "left":
		p.X = p.X - 1
	case "right":
		p.X = p.X + 1
	case "up":
		p.Y = p.Y - 1
	case "down":
		p.Y = p.Y + 1
	}
}

func (v *WhereToGoVoting) ClearVoting() {
	v.DownNumberOfVotes = 0
	v.LeftNumberOfVotes = 0
	v.RightNumberOfVotes = 0
	v.UpNumberOfVotes = 0
}
