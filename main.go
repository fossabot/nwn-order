package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/caarlos0/env"
	"github.com/glendc/go-external-ip"
	"github.com/go-redis/redis"
	"github.com/google/go-github/github"
	"github.com/jasonlvhit/gocron"
)

type config struct {
	RedisPort        string `env:"NWN_ORDER_REDIS_PORT" envDefault:"6379"`
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
		"github",
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
	msg := ("O [" + t.Format("15:04:05") + "] [NWN_Order] Pubsub Event: channel=heartbeat message=" + ticker)
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
		msg := ("O [" + t.Format("15:04:05") + "] [NWN_Order] Webhook Event: channel=innwserver message=repoupdate | " + *e.Sender.Login + " made a commit to module repo")
		go sendPubsub(msg, "github", "commit")

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
	t := time.Now()
	fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: webserver started with external IP of " + ip.String() + ":" + cfg.OrderPort + ". webhooks need to be sent to /webhook")

	http.HandleFunc("/", webpage)
	log.Fatal(http.ListenAndServe(":"+cfg.OrderPort, nil))
}

func main() {
	t := time.Now()
	log.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: Order has Started")

	conn, err := net.Dial("udp", "redis:6379")
	for retry := 1; err != nil; retry++ {
		trds := time.Now()
		s := strconv.Itoa(retry)
		log.Println("O [" + trds.Format("15:04:05") + "] [NWN_Order] Boot Event: Redis not connected | Retry attempt: " + s + " | 5 second sleep")
		if retry > 4 {
			log.Println("O [" + trds.Format("15:04:05") + "] [NWN_Order] Boot Event: Redis not connected | Exiting")
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
	}
	conn.Close()
	t = time.Now()
	log.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: Redis connected")

	conn, err = net.Dial("udp", "nwserver:5121")
	for retry := 1; err != nil; retry++ {
		trds := time.Now()
		s := strconv.Itoa(retry)
		log.Println("O [" + trds.Format("15:04:05") + "] [NWN_Order] Boot Event: nwserver not started | Retry attempt: " + s + " | 5 second sleep")
		if retry > 4 {
			log.Println("O [" + trds.Format("15:04:05") + "] [NWN_Order] Boot Event: nwserver not started | Exiting")
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
	}
	conn.Close()
	t = time.Now()
	log.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: nwserver connected")

	// start pubsub
	go startPubsub()
	fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: Pubsub started")

	// start webhook reciever
	go webserver()
	fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: Webserver started")

	// initial heartbeat
	// initial uuid
	go uuidGeneration()

	cfg := config{}
	err = env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	// start the heartbeat timers
	if cfg.HbOneMinute == true {
		fmt.Println("0 [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat enabled:  1 minute")
		gocron.Every(1).Minute().Do(heartbeatWebhook, "1")
	} else {
		fmt.Println("0 [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat disabled: 1 minute")
	}
	if cfg.HbFiveMinute == true {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat enabled:  5 minutes")
		gocron.Every(5).Minutes().Do(heartbeatWebhook, "5")
	} else {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat disabled: 5 minutes")
	}
	if cfg.HbThirtyMinute == true {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat enabled:  30 minutes")
		gocron.Every(30).Minutes().Do(heartbeatWebhook, "30")
	} else {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat disabled: 30 minutes")
	}
	if cfg.HbOneHour == true {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat enabled:  1 hour")
		gocron.Every(1).Hour().Do(heartbeatWebhook, "60")
	} else {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat disabled: 1 hour")
	}
	if cfg.HbSixHour == true {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat enabled:  6 hours")
		gocron.Every(6).Hours().Do(heartbeatWebhook, "360")
	} else {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat disabled: 6 hours")
	}
	if cfg.HbTwelveHour == true {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat enabled:  12 hours")
		gocron.Every(12).Hours().Do(heartbeatWebhook, "720")
	} else {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat disabled: 12 hours")
	}
	if cfg.HbTwentyfourHour == true {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: heartbeat enabled:  24 hours")
		gocron.Every(24).Hours().Do(heartbeatWebhook, "1440")
	} else {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event:heartbeat disabled: 24 hours")
	}
	if cfg.HbVerbose == true {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: Verbose mode: on")
	} else {
		fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: Verbose mode: off")
	}

	<-gocron.Start()
}
