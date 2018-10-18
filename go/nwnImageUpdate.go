package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	userAgent = "nwn-go 0.0.1"
)

func get(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Accept", "application/vnd.api+json")
	request.Header.Set("Content-Type", "application/vnd.api+json")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Expected status %d; Got %d\nResponse: %#v", 200, response.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}

type nwnxeeImageData struct {
	LatestSHA   string    `json:"LatestSHA"`
	WebhookURL  string    `json:"WebhookURL"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
	Description string    `json:"Description"`
	LastUpdated time.Time `json:"LastUpdated"`
	PullCount   int       `json:"PullCount"`
	StarCount   int       `json:"StarCount"`
	Versions    []struct {
		SHA  string `json:"SHA"`
		Tags []struct {
			Tag string `json:"tag"`
		} `json:"Tags"`
		ImageName string `json:"ImageName"`
		Author    string `json:"Author"`
		Labels    struct {
			Maintainer string `json:"maintainer"`
		} `json:"Labels"`
		LayerCount   int       `json:"LayerCount"`
		DownloadSize int       `json:"DownloadSize"`
		Created      time.Time `json:"Created"`
	} `json:"Versions"`
	ID           string `json:"Id"`
	ImageName    string `json:"ImageName"`
	ImageURL     string `json:"ImageURL"`
	LayerCount   int    `json:"LayerCount"`
	DownloadSize int    `json:"DownloadSize"`
	Labels       struct {
		Maintainer string `json:"maintainer"`
	} `json:"Labels"`
	LatestVersion string `json:"LatestVersion"`
}

func nwnxeeImageCheck() (*nwnxeeImageData, error) {

	byt, er := get("https://api.microbadger.com/v1/images/nwnxee/unified")
	if er != nil {
		return nil, er
	}

	imageData := new(nwnxeeImageData)

	return imageData, er
}
