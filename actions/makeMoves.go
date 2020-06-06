package actions

import (
	"team4_qgame/betypes"
	"time"
)

//MakeMoves - Takes an array of clans that will make movements.
//Returns the victor clans
func MakeMoves(clans *[]betypes.Clan) string {
	var theClanLost bool
	for !theClanLost {
		for i := 0; i < len(*clans); i++ { //The clan that goes is "i"
			go func() {
				for j := 0; j < len((*clans)[i].Users); j++ {
					if j == i {
						//TODO: Make logic for sending votes
					}
				}
			}()
			go func() {
				for j := 0; j < len((*clans)[i].Users); j++ {
					if j != i {
						//TODO: Here the logic of improvements while another clan thinks what cell to open
					}
				}
			}()
			time.Sleep(betypes.TimeToMove * time.Second) //Time to move
			//TODO: You can also display how much time is left
		}
		//TODO: Here you should calculate the clan booty (or something like that)
	}
	return "NOT_IMPLEMENTED" //TODO: It should show which team won, who got what prey and the rating of players who played
}
