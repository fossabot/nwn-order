package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/jasonlvhit/gocron"
)

func oneminutetask() {

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + os.Getenv("NWN_ORDER_REDIS_PORT"),
	})
	defer client.Close()

	pubsub := client.Subscribe("heartbeat")
	defer pubsub.Close()

	if err := client.Publish("heartbeat", "1minute").Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pubsub sent: 1 Minute heartbeat")
}

func fiveminutetask() {

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + os.Getenv("NWN_ORDER_REDIS_PORT"),
	})
	defer client.Close()

	pubsub := client.Subscribe("heartbeat")
	defer pubsub.Close()

	if err := client.Publish("heartbeat", "5minute").Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pubsub sent: 5 Minute heartbeat")
}

func thirtyminutetask() {

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + os.Getenv("NWN_ORDER_REDIS_PORT"),
	})
	defer client.Close()

	pubsub := client.Subscribe("heartbeat")
	defer pubsub.Close()

	if err := client.Publish("heartbeat", "30minute").Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pubsub sent: 30 Minute heartbeat")
}

func onehourtask() {

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + os.Getenv("NWN_ORDER_REDIS_PORT"),
	})
	defer client.Close()

	pubsub := client.Subscribe("heartbeat")
	defer pubsub.Close()

	if err := client.Publish("heartbeat", "1hour").Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pubsub sent: 1 Hour heartbeat")
}

func sixhourtask() {

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + os.Getenv("NWN_ORDER_REDIS_PORT"),
	})
	defer client.Close()

	pubsub := client.Subscribe("heartbeat")
	defer pubsub.Close()

	if err := client.Publish("heartbeat", "6hour").Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pubsub sent: 6 Hour heartbeat")
}

func twelvehourtask() {

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + os.Getenv("NWN_ORDER_REDIS_PORT"),
	})
	defer client.Close()

	pubsub := client.Subscribe("heartbeat")
	defer pubsub.Close()

	if err := client.Publish("heartbeat", "12hour").Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pubsub sent: 12 Hour heartbeat")
}

func twentyfourhourtask() {

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + os.Getenv("NWN_ORDER_REDIS_PORT"),
	})
	defer client.Close()

	pubsub := client.Subscribe("heartbeat")
	defer pubsub.Close()

	if err := client.Publish("heartbeat", "24hour").Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pubsub sent: 24 Hour heartbeat")
}

func main() {
	fmt.Println("Order has started")

	fmt.Println("NWN_ORDER_REDIS_PORT:", os.Getenv("NWN_ORDER_REDIS_PORT"))

	gocron.Every(1).Minute().Do(oneminutetask)
	gocron.Every(5).Minutes().Do(fiveminutetask)
	gocron.Every(30).Minutes().Do(thirtyminutetask)

	gocron.Every(1).Hour().Do(onehourtask)
	gocron.Every(6).Hours().Do(sixhourtask)
	gocron.Every(12).Hours().Do(twelvehourtask)
	gocron.Every(24).Hours().Do(twentyfourhourtask)

	<-gocron.Start()
}
