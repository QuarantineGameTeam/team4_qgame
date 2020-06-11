package methodsAndStructs

func (u *User) SetRank(newRank int64) {
	u.Rank = newRank
}

func (u *User) SetPlayingStatus(status bool) {
	u.PlayingStatus = status
}

func (u *User) SetQueueStatus(status bool) {
	u.QueueStatus = status
}
