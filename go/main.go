package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/glendc/go-external-ip"
	"github.com/go-redis/redis"
	"github.com/google/go-github/github"
	"github.com/jasonlvhit/gocron"
	. "github.com/logrusorgru/aurora"
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

func sendPubsub() {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + cfg.RedisPort,
	})
	defer client.Close()

	if err := client.Publish("event", "IncomingWebhook").Err(); err != nil {
		panic(err)
	}

	t := time.Now()
	fmt.Println("I [" + t.Format("15:04:05") + "] [NWN_Order] Github Webhook Event: channel=webhook")
}

func heartbeatTicker(ticker string) {

	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + os.Getenv("NWN_ORDER_REDIS_PORT"),
	})
	defer client.Close()

	if err := client.Publish("heartbeat", ticker).Err(); err != nil {
		panic(err)
	}

	if cfg.HbVerbose == true {
		t := time.Now()
		fmt.Println("I [" + t.Format("15:04:05") + "] [NWN_Order] Pubsub Event: channel=heartbeat message=" + ticker)
	}
}

func handleGithubWebhook(w http.ResponseWriter, r *http.Request) {
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
		go sendPubsub()
		fmt.Printf("%s made a commit to repo %s\n",
			*e.Sender.Login, *e.Repo.FullName)

	default:
		log.Printf("unknown event type %s\n", github.WebHookType(r))
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
	http.HandleFunc("/webhook", handleGithubWebhook)
	fmt.Println("webserver started with external IP of " + ip.String() + ":" + cfg.OrderPort + ". webhooks need to be sent to /webhook")

	http.HandleFunc("/", webpage)
	log.Fatal(http.ListenAndServe(":"+cfg.OrderPort, nil))
}

func main() {

	// just so we show up last on the boot timer.
	time.Sleep(3 * time.Second)

	fmt.Println(`  ________________ ___________  `)
	fmt.Println(` /  _  | ___ \  _  \  ___| ___\`)
	fmt.Println(`| | | | |_/ / | | | |__ | |_/ /`)
	fmt.Println(`| | | |    /| | | |  __||    / `)
	fmt.Println(`\ \_/ / |\ \| |/ /| |___| |\ \ `)
	fmt.Println(` \___/\_| \_|___/ \____/\_| \_|`)

	// load the env config
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	// start webhook reciever
	go webserver()

	// start the heartbeat timers
	if cfg.HbOneMinute == true {
		fmt.Println("heartbeat enabled:  1 minute")
		gocron.Every(1).Minute().Do(heartbeatTicker, "1m")
	} else {
		fmt.Println(Brown("heartbeat disabled: 1 minute"))
	}
	if cfg.HbFiveMinute == true {
		fmt.Println("heartbeat enabled:  5 minutes")
		gocron.Every(5).Minutes().Do(heartbeatTicker, "5m")
	} else {
		fmt.Println(Brown("heartbeat disabled: 5 minutes"))
	}
	if cfg.HbThirtyMinute == true {
		fmt.Println("heartbeat enabled:  30 minutes")
		gocron.Every(30).Minutes().Do(heartbeatTicker, "30m")
	} else {
		fmt.Println(Brown("heartbeat disabled: 30 minutes"))
	}
	if cfg.HbOneHour == true {
		fmt.Println("heartbeat enabled:  1 hour")
		gocron.Every(1).Hour().Do(heartbeatTicker, "1m")
	} else {
		fmt.Println(Brown("heartbeat disabled: 1 hour"))
	}
	if cfg.HbSixHour == true {
		fmt.Println("heartbeat enabled:  6 hours")
		gocron.Every(6).Hours().Do(heartbeatTicker, "1m")
	} else {
		fmt.Println(Brown("heartbeat disabled: 6 hours"))
	}
	if cfg.HbTwelveHour == true {
		fmt.Println("heartbeat enabled:  12 hours")
		gocron.Every(12).Hours().Do(heartbeatTicker, "1m")
	} else {
		fmt.Println(Brown("heartbeat disabled: 12 hours"))
	}
	if cfg.HbTwentyfourHour == true {
		fmt.Println("heartbeat enabled:  24 hours")
		gocron.Every(24).Hours().Do(heartbeatTicker, "1m")
	} else {
		fmt.Println(Brown("heartbeat disabled: 24 hours"))
	}
	if cfg.HbVerbose == true {
		fmt.Println("Verbose mode: on")
	} else {
		fmt.Println(Brown("Verbose mode: off"))
	}

	<-gocron.Start()
}
