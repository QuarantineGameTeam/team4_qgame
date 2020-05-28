package db

import (
	"encoding/json"
	"strconv"
	"team4_qgame/betypes"
	"team4_qgame/loger"

	"github.com/go-redis/redis/v8"
)

var (
	storage = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //No password set
		DB:       0,  //Use default DB
	})

	ctx = storage.Context()
)

//SaveUser - Writes the user to the database
func SaveUser(user betypes.User) {
	j, _ := json.Marshal(user)
	err := storage.Set(ctx, strconv.Itoa(int(user.Id)), string(j), 0).Err()
	loger.ForError(err, "Error writing to database error")
}

//GetUser - Returns the user from the database by ID
func GetUser(id string) betypes.User {
	u, err := storage.Get(ctx, id).Result()
	loger.ForError(err, "Error reading from database")
	var user betypes.User
	json.Unmarshal([]byte(u), &user)
	return user
}
