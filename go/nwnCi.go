package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func sendNwnxeeUpdate() {
	config, _ := LoadConfiguration("config.json")

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + config.Redis.Port,
	})
	defer client.Close()

	pubsub := client.Subscribe(config.Redis.PubsubNwnxee)
	defer pubsub.Close()

	if err := client.Publish(config.Redis.PubsubNwnxee, "true").Err(); err != nil {
		panic(err)
	}

	fmt.Println("Module update detected, sending alert to module")
}

func sendModuleUpdate() {
	config, _ := LoadConfiguration("config.json")

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + config.Redis.Port,
	})
	defer client.Close()

	pubsub := client.Subscribe(config.Redis.PubsubModule)
	defer pubsub.Close()

	if err := client.Publish(config.Redis.PubsubModule, "true").Err(); err != nil {
		panic(err)
	}

	fmt.Println("Module update detected, sending alert to module")
}
