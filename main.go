package main

import (
	"bytes"
	"bufio"
	"encoding/json"
	"fmt"
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
	Follow []string `json:"follow"`
}

func main() {
	// set stream parameters
	streamParams := StreamParams {
		Follow: []string{"51091012,349094942"},
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
	consumer.Debug(true)
	//Roll your own AccessToken struct
	accessTok := &oauth.AccessToken{Token: accessToken,
		Secret: accessTokenSecret}
	TwitterEndpoint := os.Args[1]
	client,err := consumer.MakeHttpClient(accessTok)
	if err != nil {
		log.Fatal(err)
	}
	params, err := json.Marshal(streamParams)
	if err != nil {
		log.Fatal(err)
	}
	TwitterEndpoint += "?follow=51091012"
	req, err := http.NewRequest("POST", TwitterEndpoint, bytes.NewBuffer(params))
	if err != nil {
		panic(err)
	}
	//req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "go-twitter v0.1")
	//#test := url.Parse(TwitterEndpoint)
	log.Print(req)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err, response)
	}
	defer response.Body.Close()
	fmt.Println("Response:", response.StatusCode, response.Status)
	//go func() {
	reader := bufio.NewReader(response.Body)
	for {
		line, _ := reader.ReadBytes('\r')
		line = bytes.TrimSpace(line)
		fmt.Println(string(line))
	}
}
