package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"lazzytchk/council/internal/model"
	"lazzytchk/council/internal/session"
	"time"
)

func main() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password123",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	token := session.Token{
		User: model.User{
			Id:       40,
			Username: "dude",
			Email:    "ehfhf@jfkf.kfj",
			Password: "my_password",
		},
		ExpireTime: time.Hour * 4,
	}

	err = client.Set("dude", token, time.Hour*4).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get("dude").Result()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

}
