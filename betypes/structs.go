package betypes

type Update struct {
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
}

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"` //Note: another optional field
	Rank      int64
}

//Point - We can save the location of the object with X and Y (For example, the location of the clan)
type Point struct {
	X, Y int
}

//Mine - It will give us resources depending on the level
type Mine struct {
	Point     //Each mine has its location on the map (X and Y)
	Level int //The speed of resource extraction depends on it
}

//Clan - The main characteristics of the clan are stored here
type Clan struct {
	Point      //Each clan has its location on the map (X and Y)
	Mine       //The clan has one mine at the beginning
	Health int //Removed during an attack on a clan
	Users  *[]User
	Number int //Clan number
}

//Field - The location of each clan and mine is preserved in the field
type Field struct {
	Mines *[]Mine
	Clans *[]Clan
}

//Game - Saves the ID of the game and the players waiting to start
type Game struct {
	PlayersInLine map[int64]User //Players in line
	GameID        int64
	GameStarted   bool
}
