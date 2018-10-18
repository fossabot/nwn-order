package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	config, _ := LoadConfiguration("config.json")
	payload, err := github.ValidatePayload(r, []byte(config.External.GithubWebhookSecret))
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
		go sendModuleUpdate()
		fmt.Printf("%s made a commit to repo %s\n",
			*e.Sender.Login, *e.Repo.FullName)

	default:
		log.Printf("unknown event type %s\n", github.WebHookType(r))
		return
	}
}

func main() {
	//config, _ := LoadConfiguration("config.json")

	//consensus := externalip.DefaultConsensus(nil, nil)
	//ip, _ := consensus.ExternalIP()

	//fmt.Println("webserver started | external IP of " + ip.String() + ":" + config.Webserver.Port)

	//http.HandleFunc("/webhook", handleWebhook)
	//log.Fatal(http.ListenAndServe(":8080", nil))

	//go sendNwnxeeUpdate()

	fmt.Println(nwnxeeImageCheck)
}
