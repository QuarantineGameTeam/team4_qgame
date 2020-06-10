package actions

import (
	"strconv"
	"team4_qgame/betypes"
	"team4_qgame/db"
)

var (
	gamesInLine = make(map[int64]betypes.GameInLine)
	games       = make([]betypes.Game, 0)
)

func FindUser(id int64) bool {
	if db.GetUser(strconv.FormatInt(id, 10)).Id == id {
		return true
	}
	return false
}

func FindGameInLine(id int64) bool {
	for _, game := range gamesInLine {
		if game.GameID == id {
			return true
		}
	}
	return false
}

func FindGame(gameID int64) bool {
	for _, game := range games {
		if game.GameID == gameID {
			return true
		}
	}
	return false
}

func FindPlayerInGame(playerID int64) bool {
	for _, game := range games {
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

func GetThePlayersGameID(playerID int64) int64 {
	for _, game := range games {
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
	for _, game := range games {
		if game.GameID == gameID {
			return game
		}
	}
	return betypes.Game{}
}

//func AtakClan(clanToAttack betypes.Clan, )
