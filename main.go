package main

import (

	"fmt"
	"os"
	"gopkg.in/redis.v3"
)

func main() {
	//set redis contexxt
	IPredis := os.Getenv("DB_PORT_6379_TCP_ADDR")

	if IPredis == "" {
		fmt.Println("var DB_PORT_6379_TCP_ADDR is null")
	} else {
		fmt.Println("var DB_PORT_6379_TCP_ADDR :=%s", IPredis)

		client := redis.NewClient(&redis.Options{
			Addr:     IPredis + ":6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		_, err := client.Ping().Result()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Connection to redis Ok")
		}
		err = client.Close()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Connection to redis closed")
		}
	}
}
