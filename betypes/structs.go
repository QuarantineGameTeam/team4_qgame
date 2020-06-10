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
}
