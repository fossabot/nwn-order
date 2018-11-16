package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env"
	"github.com/glendc/go-external-ip"
	"github.com/go-redis/redis"
	"github.com/google/go-github/github"
	"github.com/jasonlvhit/gocron"
)

type config struct {
	RedisPort        string `env:"NWN_ORDER_REDIS_PORT" envDefault:"6379"`
	OrderModuleName  string `env:"NWN_ORDER_MODULE_NAME" envDefault:"my_module"`
	OrderPort        string `env:"NWN_ORDER_PORT" envDefault:"5750"`
	HbVerbose        bool   `env:"NWN_ORDER_HB_VERBOSE" envDefault:"true"`
	HbOneMinute      bool   `env:"NWN_ORDER_HB_ONE_MINUTE" envDefault:"true"`
	HbFiveMinute     bool   `env:"NWN_ORDER_HB_FIVE_MINUTE" envDefault:"true"`
	HbThirtyMinute   bool   `env:"NWN_ORDER_HB_THIRTY_MINUTE" envDefault:"true"`
	HbOneHour        bool   `env:"NWN_ORDER_HB_ONE_HOUR" envDefault:"true"`
	HbSixHour        bool   `env:"NWN_ORDER_HB_SIX_HOUR" envDefault:"true"`
	HbTwelveHour     bool   `env:"NWN_ORDER_HB_TWELVE_HOUR" envDefault:"true"`
	HbTwentyfourHour bool   `env:"NWN_ORDER_HB_TWENTYFOUR_HOUR" envDefault:"true"`
}

var (
	//RedisClient is a var
	RedisClient *redis.Client
)

func startPubsub() {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + cfg.RedisPort,
	})
	defer client.Close()

	pubSub := client.Subscribe(
		"heartbeat",
		"input",
		"debug",
	)
	for {
		msg, _ := pubSub.ReceiveMessage()
		switch msg.Channel {
		case "heartbeat":

		case "input":
			go uuidGeneration()
		case "debug":
		}
	}
}

func uuidGeneration() {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + cfg.RedisPort,
	})
	defer client.Close()

	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x%x%x%x%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	if err := client.Set("system:uuid", uuid, 0).Err(); err != nil {
		panic(err)
	}

	pub := client.Publish("output", "uuid")
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}
}

func sendPubsub(LogMessage string, PubsubChannel string, PubsubMessage string) {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + cfg.RedisPort,
	})
	defer client.Close()

	if err := client.Publish(PubsubChannel, PubsubMessage).Err(); err != nil {
		panic(err)
	}

	if cfg.HbVerbose == true {
		fmt.Println(LogMessage)
	}
}

func heartbeatWebhook(ticker string) {
	t := time.Now()
	msg := ("I [" + t.Format("15:04:05") + "] [NWN_Order] Pubsub Event: channel=heartbeat message=" + ticker)
	sendPubsub(msg, "heartbeat", ticker)
}

func githubWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error validating request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("could not parse webhook: err=%s\n", err)
		return
	}

	switch e := event.(type) {

	case *github.PushEvent:
		t := time.Now()
		msg := ("I [" + t.Format("15:04:05") + "] [NWN_Order] Webhook Event: channel=innwserver message=repoupdate | " + *e.Sender.Login + " made a commit to module repo")
		go sendPubsub(msg, "innwserver", "repoupdate")

	default:
		log.Printf("Only push events supported, unknown webhook event type %s\n", github.WebHookType(r))
		return
	}
}

func webpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func webserver() {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	consensus := externalip.DefaultConsensus(nil, nil)
	ip, _ := consensus.ExternalIP()
	http.HandleFunc("/webhook", githubWebhook)
	fmt.Println("webserver started with external IP of " + ip.String() + ":" + cfg.OrderPort + ". webhooks need to be sent to /webhook")

	http.HandleFunc("/", webpage)
	log.Fatal(http.ListenAndServe(":"+cfg.OrderPort, nil))
}

func main() {
	fmt.Println(`Order has started`)

	// start pubsub
	go startPubsub()
	fmt.Println(`Pubsub started`)

	// start webhook reciever
	go webserver()
	fmt.Println(`Webserver started`)

	// initial heartbeat
	go uuidGeneration()

	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	// start the heartbeat timers
	if cfg.HbOneMinute == true {
		fmt.Println("heartbeat enabled:  1 minute")
		gocron.Every(1).Second().Do(heartbeatWebhook, "1m")
	} else {
		fmt.Println("heartbeat disabled: 1 minute")
	}
	if cfg.HbFiveMinute == true {
		fmt.Println("heartbeat enabled:  5 minutes")
		gocron.Every(5).Minutes().Do(heartbeatWebhook, "5m")
	} else {
		fmt.Println("heartbeat disabled: 5 minutes")
	}
	if cfg.HbThirtyMinute == true {
		fmt.Println("heartbeat enabled:  30 minutes")
		gocron.Every(30).Minutes().Do(heartbeatWebhook, "30m")
	} else {
		fmt.Println("heartbeat disabled: 30 minutes")
	}
	if cfg.HbOneHour == true {
		fmt.Println("heartbeat enabled:  1 hour")
		gocron.Every(1).Hour().Do(heartbeatWebhook, "1m")
	} else {
		fmt.Println("heartbeat disabled: 1 hour")
	}
	if cfg.HbSixHour == true {
		fmt.Println("heartbeat enabled:  6 hours")
		gocron.Every(6).Hours().Do(heartbeatWebhook, "1m")
	} else {
		fmt.Println("heartbeat disabled: 6 hours")
	}
	if cfg.HbTwelveHour == true {
		fmt.Println("heartbeat enabled:  12 hours")
		gocron.Every(12).Hours().Do(heartbeatWebhook, "1m")
	} else {
		fmt.Println("heartbeat disabled: 12 hours")
	}
	if cfg.HbTwentyfourHour == true {
		fmt.Println("heartbeat enabled:  24 hours")
		gocron.Every(24).Hours().Do(heartbeatWebhook, "1m")
	} else {
		fmt.Println("heartbeat disabled: 24 hours")
	}
	if cfg.HbVerbose == true {
		fmt.Println("Verbose mode: on")
	} else {
		fmt.Println("Verbose mode: off")
	}

	<-gocron.Start()
}
