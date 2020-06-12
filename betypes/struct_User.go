package betypes

type User struct {
	Id            int64  `json:"id"`
	FirstName     string `json:"first_name"`
	Username      string `json:"username"`
	Rank          int64  `json:"rank"`
	PlayingStatus bool   `json:"playing_status"`
	QueueStatus   bool   `json:"queue_status"`
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
