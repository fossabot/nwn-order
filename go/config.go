package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configfile struct
type Configfile struct {
	External struct {
		GithubWebhookSecret string `json:"github_webhook_secret"`
	} `json:"external"`
	Webserver struct {
		Port string `json:"server_port"`
	} `json:"webserver"`
	Redis struct {
		Port         string `json:"redis_port"`
		PubsubNwnxee string `json:"nwnxee_pubsub_channel"`
		PubsubModule string `json:"module_pubsub_channel"`
	} `json:"redis"`
	Docker struct {
		NwserverName string `json:"nwserver_container_name"`
		Port         string `json:"nwserver_port"`
	} `json:"docker"`
}

// LoadConfiguration func
func LoadConfiguration(file string) (Configfile, error) {
	var config Configfile
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}
