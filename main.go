package main

import (
	"bytes"
	"bufio"
	"encoding/json"
	//"io/ioutil"
	"log"
	"os"
	"net/http"
	//"net/url"

	"github.com/gointo/oauth"
)
// Usage is the command line helper
//func Usage() {
//	fmt.Println("Usage:")
//	fmt.Print("go run examples/twitter/twitter.go")
//	fmt.Print("  --consumerkey <consumerkey>")
//	fmt.Println("  --consumersecret <consumersecret>")
//	fmt.Println("")
//	fmt.Println("In order to get your consumerkey and consumersecret, you must register an 'app' at twitter.com:")
//	fmt.Println("https://dev.twitter.com/apps/new")
//}

// StreamParams let give params to filter a public stream
type StreamParams struct {
	Follow string `json:"follow"`
}

func main() {
	// set stream parameters
	streamParams := StreamParams {
		Follow: "?follow=349094942,633673441,4221690875,3096291947,3096291947",
	}
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	if len(consumerKey) == 0 ||
			len(consumerSecret) == 0 ||
			len(accessToken) == 0 ||
			len(accessTokenSecret) == 0 {
		log.Fatalf("Config wrong!!!!")
	}
	consumer := oauth.NewConsumer(consumerKey,
		consumerSecret,
		oauth.ServiceProvider{})
	//NOTE: remove this line or turn off Debug if you don't
	//want to see what the headers look like
	// log.Println("Header: ", consumer.Debug(true))
	//Roll your own AccessToken struct
	accessTok := &oauth.AccessToken{Token: accessToken,
		Secret: accessTokenSecret}
	TwitterEndpoint := string(os.Args[1])
	client,err := consumer.MakeHttpClient(accessTok)
	if err != nil {
		log.Fatal(err)
	}
	params, err := json.Marshal(streamParams)
	if err != nil {
		log.Fatal(err)
	}
	TwitterEndpoint += streamParams.Follow
	req, err := http.NewRequest("POST", TwitterEndpoint, bytes.NewBuffer(params))
	if err != nil {
		panic(err)
	}
	//req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "go-twitter v0.1")
	//#test := url.Parse(TwitterEndpoint)
	log.Printf("Request: %v", req)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err, response)
	}
	defer response.Body.Close()
	log.Println("Response:", response.StatusCode, response.Status)
	//go func() {
	reader := bufio.NewReader(response.Body)
	var body message
	for {
		line, _ := reader.ReadBytes('\r')
		line = bytes.TrimSpace(line)
		//log.Printf("JSON answer lenght: %d", len(string(line)))
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
