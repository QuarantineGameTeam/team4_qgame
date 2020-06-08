package mechanics

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"fmt"
	"time"
	"team4_qgame/betypes"
)

var (
	up 		= 0
	down 	= 0
	left 	= 0
	right 	= 0
	stay	= 0
	
	selectMove  = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(string(11013), "1"),
			tgbotapi.NewInlineKeyboardButtonData(string(11014), "2"),
			tgbotapi.NewInlineKeyboardButtonData(string(10145), "3"),
			tgbotapi.NewInlineKeyboardButtonData(string(11015), "4"),
		),
	)
)

func StepOfClan( clan betypes.Clan, bot *tgbotapi.BotAPI ) (string, int ){
	find := false
	findClan := 0

	go func(){
		StepOfUsers( clan betypes.Clan, bot *tgbotapi.BotAPI )
		directionOfStep := searchMax()
		generateNewPoint(directionOfStep, &clan.Location)

		find, findClan := isOtherClanInPoint( clan betypes.Clan, clan.Location )

	if find { 
		go func () {
			for _, user := range *clan.Users {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Your clan select direction  ' %s '. But in this point is located clan number %d", directionOfStep, findClan)))
			}
			<-time.After(time.Second * 60)
		}()
		return "OtherClan", findClan
	} 

	find, findClan := isMineInPoint( clan betypes.Clan, clan.Location )

	if find {
		go func () {
			for _, user := range *clan.Users {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Your clan select direction  ' %s '. But in this point is mine that beloved to clan number %d", directionOfStep, findClan)))
			}
			<-time.After(time.Second * 60)
		}()
		return "Mine", findClan
	}

	go func () {
		for _, user := range *clan.Users {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Your clan select direction  ' %s '. This is empty point", directionOfStep)))
		}
		<-time.After(time.Second * 60)
	}()
	return "Empty", 0

		<-time.After(time.Second * 60)
	}()
}



// Users selected direction
func StepOfUsers( clan betypes.Clan, bot *tgbotapi.BotAPI ) {
	for _, user := range *clan.Users {
		printMessage(user.Id, "Choose your next move!",  bot *tgbotapi.BotAPI, true)
		if  tgbotapi.Update.CallbackQuery != nil {
			clanLocation := clan.Location
			
			switch tgbotapi.Update.CallbackQuery.Data {
			case 1: 
				if clanLocation.X - 1 < 1 { 
					texts = "Oh no. Here is the ocean.Rather, choose another direction "
					printMessage(user.Id, texts,  bot *tgbotapi.BotAPI, true) 
				} else { left++ }
			case 2:
				if clanLocation.Y - 1 < 1 { 
					texts = "Oh no. Your path is blocked by mountains. Rather choose another direction"
					printMessage(user.Id, texts,  bot *tgbotapi.BotAPI, true) 
					
				} else { up++ }
			case 3:
				if clanLocation.X + 1 > betypes.Width { 
					texts = "Oh no. There is an active volcano. Rather choose another direction"
					printMessage(user.Id, texts,  bot *tgbotapi.BotAPI, true) 
				} else { right++ }

			case 4:
				if clanLocation.Y + 1 > betypes.Hight { 
					texts = "Oh no. This is a restricted area. You need permission.. Rather choose another direction"
					printMessage(user.Id, texts,  bot *tgbotapi.BotAPI, true) 
				} else { right++ }
			}
		}
}

//print Message for Users
func printMessage(id int64, msgs string bot *tgbotapi.BotAPI, keyboard bool) {
	msg := tgbotapi.NewMessage(id, "")
	msg.Text = msgs
	if keyboard {
		msg.ReplyMarkup = selectMove
	}
	bot.Send(msg)
}

// finds the direction of movement by a majority vote
func searchMax() string {

	maxWidth := 0
	maxDirWidth := ""
	maxHight := 0
	maxDirHight := ""

	if left - right < 0 {
		maxWidth = right
		maxDirWidth = "right"
	} else {
		maxWidth = left
		maxDirWidth = "left"
	}

	if up - down < 0 {
		maxHight = down
		maxDirHight = "down"
	} else {
		maxDirHight = up
		maxDirHight = "up"
	}

	if maxWidth - maxHight < 0 {
		return maxDirHight
	} else { return maxDirWidth }
}

//create newPoints for clan
func generateNewPoint(dir string, p *betypes.Point ) {
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

// check if in selected point is mine
func isMineInPoint( p betypes.Point ) (bool, int ) {
	mines := betypes.Field;
	for _, mine := range *mines.Mines{
		if p.X == mine.Location.X && p.Y == mine.Location.Y { return true, mine.ToBelong }
	}
	return false, 0
}

// check if in selected point is other clan
func isOtherClanInPoint(clan betypes.Clan, p betypes.Point) ( bool, int ) {
	clans := betypes.Field
	for _, otherClan := range *clans.Clan {
		if otherClan.Number != clan.Number {
			if otherClan.Location.X == p.X && otherClan.Location.Y == p.Y { return true, otherClan.Number }
		}
	}
	return false, 0
}
