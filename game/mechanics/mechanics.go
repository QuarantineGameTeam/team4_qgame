package mechanics

import (
	"fmt"
	"math/rand"
	"team4_qgame/betypes"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func OfferAnAlliance(sender *betypes.Clan, receiver *betypes.Clan, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var PollingKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Accept", "accept_alliance"),
			tgbotapi.NewInlineKeyboardButtonData("Reject", "reject_alliance"),
		),
	)
	pollCounterFor := 0
	pollCounterAgainst := 0
	go func() {
		for _, user := range receiver.Users {
			bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintf("You've received an offer for alliance!")))
			msg := tgbotapi.NewMessage(user.Id, "")
			msg.ReplyMarkup = PollingKeyboard
			if update.CallbackQuery != nil {
				switch update.CallbackQuery.Data {
				case "accept_alliance":
					pollCounterFor++
					msg.Text = fmt.Sprint("You voted for alliance,\n")
					break
				case "reject_alliance":
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
		go func() {
			for _, user := range sender.Users {
				bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintf("Your offer for alliance has been accepted.")))
			}
			sender.Number += receiver.Number
			receiver.Number += sender.Number
		}()
		go func() {
			for _, mineR := range receiver.Mines {
				sender.Mines = append(sender.Mines, mineR)
			}
		}()
		for _, mineS := range sender.Mines {
			receiver.Mines = append(receiver.Mines, mineS)
		}
		go func() {
			for _, user := range receiver.Users {
				sender.Users = append(sender.Users, user)
			}
		}()
		go func() {
			for _, user := range sender.Users {
				receiver.Users = append(receiver.Users, user)
			}
		}()
	} else {
		go func() {
			for _, user := range sender.Users {
				bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintf("Your offer for alliance has been rejected.")))
			}
		}()
	}
}
func Attack(clan *betypes.Clan, bot *tgbotapi.BotAPI, damage int) bool {
	if clan.Health >= damage {
		clan.Health -= damage
		return true
	}
	return false
}

//Declaring a war function
func DeclareWar(sender *betypes.Clan, receiver *betypes.Clan, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var AttackButton = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Attack"),
		),
	)
	counterR := 0
	counterS := 0
	go func() {
		for _, user := range receiver.Users {
			msg := tgbotapi.NewMessage(user.Id, fmt.Sprintf("Clan #%d has just declared you a war!", sender.Number))
			bot.Send(msg)
			msg.ReplyMarkup = AttackButton
			if update.CallbackQuery != nil {
				switch update.CallbackQuery.Data {
				case "Attack":
					counterR++
					if counterR < len(receiver.Users) {
						msg.Text = "Wait for other members to accept war..."
						bot.Send(msg)
						break
					}
				}
			}
		}
	}()
	if counterR < len(receiver.Users) {
		go func() {
			for _, user := range receiver.Users {
				bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintln("Unfortunately, your clan's memebers decided not to attack.")))
			}
		}()
	} else {
		go func() {
			for _, user := range receiver.Users {
				bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintf("The war begins!💥")))
				msg := tgbotapi.NewMessage(user.Id, "")
				if update.CallbackQuery.Data == "Attack" {
					msg.Text = "The attack begins!"
					bot.Send(msg)
					damage := rand.Intn(counterR % 10)
					if Attack(sender, bot, damage) == true {
						msg.Text = fmt.Sprintf("Your clan has just attacked with damage of %d", damage)
						bot.Send(msg)
					} else {
						msg.Text = "You can't attack no more."
					}
				}
			}
		}()
	}
	go func() {
		for _, user := range sender.Users {
			msg := tgbotapi.NewMessage(user.Id, "")
			msg.ReplyMarkup = AttackButton
			if update.CallbackQuery != nil {
				switch update.CallbackQuery.Data {
				case "Attack":
					counterS++
					if counterS < len(sender.Users) {
						msg.Text = "Wait for other members to start attack..."
						bot.Send(msg)
					}
					break
				}
			}
		}
	}()
	if counterS < len(sender.Users) {
		go func() {
			for _, user := range sender.Users {
				bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintln("Unfortunately, your clan's memebers decided not to attack.")))
			}
		}()
	} else {
		go func() {
			for _, user := range sender.Users {
				bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintf("The war begins!💥")))
				msg := tgbotapi.NewMessage(user.Id, "")
				if update.CallbackQuery.Data == "Attack" {
					msg.Text = "The attack begins!"
					bot.Send(msg)
					damage := rand.Intn(counterS % 10)
					if Attack(receiver, bot, damage) == true {
						msg.Text = fmt.Sprintf("Your clan has just attacked with damage of %d", damage)
						bot.Send(msg)
					} else {
						msg.Text = "You can't attack no more."
						bot.Send(msg)
					}
				}
			}
		}()
	}
}
