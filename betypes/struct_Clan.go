package betypes

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Clan - The main characteristics of the clan are stored here
type Clan struct {
	Location  Point  //Each clan has its location on the map (X and Y)
	Mines     []Mine //The clan has one mine at the beginning
	Health    int    //Removed during an attack on a clan
	Users     []User
	Number    int    //Clan number
	Name      string //Clan name
	WhereToGo WhereToGoVoting
}

func (c *Clan) AddUser(users User) {
	c.Users = make([]User, 0)
	c.Users = append(c.Users, users)
}

func (c *Clan) StartVotingWhereToGo(bot *tgbotapi.BotAPI) {
	for _, user := range c.Users {
		msg := tgbotapi.NewMessage(user.Id, "⚡Voting has started!! Please select where you want to go!!⚡")
		msg.ReplyMarkup = SelectMoveButton
		bot.Send(msg)
	}
}

func (c *Clan) StopVoting(bot *tgbotapi.BotAPI) {
	chose := c.WhereToGo.SearchMax()
	for _, user := range c.Users {
		bot.Send(tgbotapi.NewMessage(user.Id, fmt.Sprintf("⚡Your clan has chosen to go %s⚡", chose)))
	}
	c.WhereToGo.GenerateNewPoint(chose, &c.Location)
	c.WhereToGo.ClearVoting()
}

//check if in selected point is other clan
func (c *Clan) IsOtherClanInPoint(clans []Clan) (bool, *int, *string) {
	for _, otherClan := range clans {
		if otherClan.Number != otherClan.Number {
			if otherClan.Location.X == c.Location.X && otherClan.Location.Y == c.Location.Y {
				return true, &otherClan.Number, &otherClan.Name
			}
		}
	}
	return false, nil, nil
}
