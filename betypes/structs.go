package betypes

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"math/rand"
	"time"
)

/*type Update struct {
	Id  int64   `json:"update_id"`
	Msg Message `json:"message"`
}

type Message struct {
	Id   int64  `json:"message_id"`
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
	User User   `json:"from"` //Note: this is an optional field so it may be empty
}

type Chat struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}*/
type User struct {
	Id            int64  `json:"id"`
	FirstName     string `json:"first_name"`
	Username      string `json:"username"` //Note: another optional field
	Rank          int64
	PlayingStatus bool
	QueueStatus   bool
}

type UserMethods interface {
	SetRank(newRank int64)
	SetPlayingStatus(status bool)
}

func (u *User) SetRank(newRank int64) {
	u.Rank = newRank
}

func (u *User) SetPlayingStatus(status bool) {
	u.PlayingStatus = status
}

func (u *User) SetQueueStatus(status bool) {
	u.QueueStatus = status
}

//Point - We can save the location of the object with X and Y (For example, the location of the clan)
type Point struct {
	X, Y int
}

//Mine - It will give us resources depending on the level
type Mine struct {
	Location Point //Each mine has its location on the map (X and Y)
	Level    int   //The speed of resource extraction depends on it
	ToBelong int
}

//Clan - The main characteristics of the clan are stored here
type Clan struct {
	Location Point //Each clan has its location on the map (X and Y)
	ClanMine Mine  //The clan has one mine at the beginning
	Health   int   //Removed during an attack on a clan
	Users    []User
	Number   int //Clan number
}

type ClanMethods interface {
	AddUser(user User)
}

func (clan *Clan) AddUser(users User) {
	clan.Users = make([]User, 0)
	clan.Users = append(clan.Users, users)
}

//Field - The location of each clan and mine is preserved in the field
type Field struct {
	Mines []Mine
	Clans []Clan
}

type FieldMethods interface {
	AddMine(mine Mine)
	AddClans(clan Clan)
}

func (f *Field) AddMine(mine Mine) {
	f.Mines = append(f.Mines, mine)
}

func (f *Field) AddClan(clan Clan) {
	f.Clans = append(f.Clans, clan)
}

//Game - Saves the ID of the game and the players waiting to start
type GameInLine struct {
	PlayersInLine map[int64]User //Players in line
	GameID        int64
}

type GameInLineMethods interface {
	CreateAGame(field Field, update tgbotapi.Update, bot *tgbotapi.BotAPI)
}

func (game *GameInLine) CreateAGame(field Field, update tgbotapi.Update, bot *tgbotapi.BotAPI, allGames *[]Game) Game {
	firstClan := make([]User, 0)
	secondClan := make([]User, 0)
	thirdClan := make([]User, 0)
	rand.Seed(time.Now().UnixNano())
	for _, user := range game.PlayersInLine {
		switch rand.Intn(3) {
		case 0:
			firstClan = append(firstClan, user)
		case 1:
			secondClan = append(secondClan, user)
		case 2:
			thirdClan = append(thirdClan, user)
		}
	}
	for i := 0; i < len(field.Clans); i++ {
		switch i {
		case 0:
			field.Clans[i].Users = firstClan
		case 1:
			field.Clans[i].Users = secondClan
		case 2:
			field.Clans[i].Users = thirdClan
		}
	}

	newGame := Game{
		Battlefield: field,
		GameID:      game.GameID,
	}
	*allGames = append(*allGames, newGame)
	return newGame
}

type Game struct {
	Battlefield Field
	GameID      int64
}
