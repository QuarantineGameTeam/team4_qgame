package db

import (
	"encoding/json"
	"log"
	"strconv"
	"team4_qgame/betypes/methodsAndStructs"
	"team4_qgame/loger"

	"github.com/go-redis/redis/v8"
)

const USER = "USER_"

var (
	storage = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //No password set
		DB:       0,  //Use default DB
	})

	ctx = storage.Context()
)

//SaveUser - Writes the user to the database
func SaveUser(user methodsAndStructs.User) error {
	log.Println("Save user to the DB, ID", user.Id)
	j, err := json.Marshal(user)
	if err != nil {
		log.Println("Could not marshal user", err)
		return err
	}
	err = storage.Set(ctx, USER+strconv.FormatInt(user.Id, 10), string(j), 0).Err()
	if err != nil {
		log.Println("Could not save user", err)
		return err
	}
	log.Println("User successfully saved to DB, ID", user.Id)
	return nil
}

//GetUser - Returns the user from the database by ID
func GetUser(id int64) (*methodsAndStructs.User, error) {
	u := &methodsAndStructs.User{}
	log.Println("Get user from DB, user ID", id)
	r, err := storage.Get(ctx, USER+strconv.FormatInt(id, 10)).Result()
	if err == redis.Nil {
		log.Println("User not found", err)
		return nil, err
	} else if err != nil {
		loger.ForError(err, "Could not read user from DB")
	} else {
		err := json.Unmarshal([]byte(r), u)
		if err != nil {
			log.Println("Could not unmarshal user", err)
			return nil, err
		}
		log.Println("User successfully received from DB, ID", u.Id)
	}
	return u, nil
}

func GetAllKeys() []string {
	keys := make([]string, 0)
	iter := storage.Scan(ctx, 0, "", 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	return keys
}
