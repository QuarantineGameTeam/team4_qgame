package game

import (
	"team4_qgame/betypes"
	"team4_qgame/db"
)

var (
	GamesInLine = make(map[int64]betypes.GameInLine)
	Games       = make([]betypes.Game, 0)
)

func FindUser(id int64) bool {
	u, _ := db.GetUser(id)
	if u.Id == id {
		return true
	}
	return false
}

func FindGameInLine(id int64) bool {
	for _, game := range GamesInLine {
		if game.GameID == id {
			return true
		}
	}
	return false
}

func FindGame(gameID int64) bool {
	for _, game := range Games {
		if game.GameID == gameID {
			return true
		}
	}
	return false
}

func FindPlayerInGame(playerID int64) bool {
	for _, game := range Games {
		for _, clan := range game.Battlefield.Clans {
			for _, user := range clan.Users {
				if user.Id == playerID {
					return true
				}
			}
		}
	}
	return false
}

func GetGameIDByPlayerID(playerID int64) int64 {
	for _, game := range Games {
		for _, clan := range game.Battlefield.Clans {
			for _, user := range clan.Users {
				if user.Id == playerID {
					return game.GameID
				}
			}
		}
	}
	return 0
}

func GetGame(gameID int64) betypes.Game {
	for _, game := range Games {
		if game.GameID == gameID {
			return game
		}
	}
	return betypes.Game{}
}
