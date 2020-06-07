package db

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"strconv"
	"team4_qgame/betypes"
	"team4_qgame/loger"
	"time"
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

//SaveField - save new field's data to the database
func SaveField(field betypes.Field) {
	j, _ := json.Marshal(field)
	rand.Seed(time.Now().UnixNano())
	fieldID := rand.Intn(1000000)
	err := storage.Set(ctx, strconv.Itoa(fieldID), string(j), 0).Err()
	loger.ForError(err, "Error writing to database error")
}

//GetField - Returns the field from the database by ID
func GetField(id string) betypes.Field {
	f, err := storage.Get(ctx, id).Result()
	loger.ForError(err, "Error reading from database")
	var field betypes.Field
	json.Unmarshal([]byte(f), &field)
	return field
}

func GetAllKeys() []string {
	keys := make([]string, 0)
	iter := storage.Scan(ctx, 0, "", 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	return keys
}
