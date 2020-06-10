package betypes

type User struct {
	Id            int64  `json:"id"`
	FirstName     string `json:"first_name"`
	Username      string `json:"username"`
	Rank          int64  `json:"rank"`
	PlayingStatus bool   `json:"playing_status"`
	QueueStatus   bool   `json:"queue_status"`
}

//Point - We can save the location of the object with X and Y (For example, the location of the clan)
type Point struct {
	X, Y int
}

//Mine - It will give us resources depending on the level
type Mine struct {
	Location Point //Each mine has its location on the map (X and Y)
	Level    int   //The speed of resource extraction depends on it
	ToBelong int   //The number of the clan that owns the mine
}

//Clan - The main characteristics of the clan are stored here
type Clan struct {
	Location Point //Each clan has its location on the map (X and Y)
	ClanMine Mine  //The clan has one mine at the beginning
	Health   int   //Removed during an attack on a clan
	Users    []User
	Number   int //Clan number
}

//Field - The location of each clan and mine is preserved in the field
type Field struct {
	Mines []Mine
	Clans []Clan
}

//Game - Saves the ID of the game and the players waiting to start
type GameInLine struct {
	PlayersInLine map[int64]User //Players in line
	GameID        int64
}

type Game struct {
	Battlefield Field
	GameID      int64
}
