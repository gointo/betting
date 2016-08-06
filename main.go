package main

import (
	"bytes"
	"bufio"
	"encoding/json"
	"log"
	"os"
	"net/http"

	"github.com/gointo/oauth"
)

// StreamParams let give params to filter a public stream
type StreamParams struct {
	Follow string `json:"follow"`
}

// GetURL return the url with the followers in query
func GetURL() string {
	streamParams := StreamParams {
		Follow: "?follow=349094942,633673441,4221690875,3096291947,3096291947",
	}
	TwitterEndpoint := string(os.Args[1])
	TwitterEndpoint += streamParams.Follow
	return TwitterEndpoint
}

// GetEnvCreds return credentials required from environement variables
func GetEnvCreds() (string, string, string, string){
	CK, CS, AT, ATS := os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"),
		os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET")
	if len(CK) == 0 || len(CS) == 0 || len(AT) == 0 || len(ATS) == 0 {
		log.Fatalf("Problem with environment variables credentials:\nConsumer Key Length: %d\nConsumer Secret: %d\nAccess Token Length: %d\nAccess Token Secret Length: %d",
					CK, CS, AT, ATS)
	}
	return CK, CS, AT, ATS
}

// GetClient return a client setted with the good credentials and token
func GetClient() *http.Client {
	consumerKey, consumerSecret, accessToken, accessTokenSecret := GetEnvCreds()
	consumer := oauth.NewConsumer(consumerKey,
		consumerSecret,
		oauth.ServiceProvider{})
	// log.Println("Header: ", consumer.Debug(true))
	accessTok := &oauth.AccessToken{Token: accessToken,
		Secret: accessTokenSecret}
	client, err := consumer.MakeHttpClient(accessTok)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// GetRequest return an HTTP request setted with the good method and url
func GetRequest() *http.Request {
	// GetURL stream parameters
	url := GetURL()
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "go-twitter v0.1")
	return req
}

// TreatResponse run on the stream to get json responses
func TreatResponse(reader *bufio.Reader, body *message) {
	for {
		line, _ := reader.ReadBytes('\r')
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		err := json.Unmarshal(line, &body)
		if err != nil {
			log.Fatal(err)
		}
		if body.User.ID == 349094942 || body.User.ID == 633673441 ||
				body.User.ID == 4221690875 || body.User.ID == 3096291947 ||
				body.User.ID == 3096291947 {
			log.Printf(
				"\nName: \033[1;31m%s\033[0m\n\033[1;32m%s\033[0m\n\n\n",
				body.User.ScreenName,
				body.Text)
		}
	}
}

// GetStream launch all the request logic
func GetStream() {
	client := GetClient()
	request := GetRequest()
	log.Printf("Request: %v", request)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err, response)
	}
	defer response.Body.Close()
	log.Println("Response:", response.StatusCode, response.Status)
	reader := bufio.NewReader(response.Body)
	var body message
	TreatResponse(reader, &body)
}

func main() {
	GetStream()
}
