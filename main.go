package main

import (
	"bytes"
	"bufio"
	"encoding/json"
	//"io"
	"log"
	"os"
	"net/http"
	"net/url"
	"strings"
	"regexp"

	"github.com/gointo/oauth"
)

// StreamParams let give params to filter a public stream
type StreamParams struct {
	With string `json:"follow"`
}

// GetURL return the url with the followers in query
func GetURL() string {
	streamParams := StreamParams {
		//Follow: "?follow=349094942,4197365524",
		With: "?with=followings",
	}
	TwitterEndpoint := os.Args[1]
	TwitterEndpoint += streamParams.With
	return TwitterEndpoint
}

// GetEnvCreds return credentials required from environement variables
func GetEnvCreds() (string, string, string, string){
	CK, CS, AT, ATS := os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"),
		os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET")
	if len(CK) == 0 || len(CS) == 0 || len(AT) == 0 || len(ATS) == 0 {
		log.Fatalf("Problem with environment variables credentials:\nConsumer Key Length: %s\nConsumer Secret: %s\nAccess Token Length: %s\nAccess Token Secret Length: %s",
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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "go-twitter v0.1")
	return req
}

func sendTelegram(text string) {
	v := url.Values{"chat_id": {os.Getenv("TELEGRAM_CHAT_ID")}, "text": {text}}
	req := "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_TOKEN") + "/sendMessage"
	log.Printf("telegram post request: %v", req)
	resp, err := http.PostForm(req, v)
	if err != nil {
	  log.Fatal(err)
	}
	debug := bufio.NewReader(resp.Body)
	res, _, _ := debug.ReadLine()
	var f interface{}
	_ = json.Unmarshal(res, &f)
	log.Printf("telegram post response: %v\nchat_id: %d", f, 42)
}

// TODO func getBet(body string) (string, string)
// TODO getGame(team string, body string) string
// TODO getStake(body string) string

func isBet(count *int, msg string) bool {
	//if strings.Contains(strings.ToLower(msg), "stake") {
	if res, _ := regexp.MatchString("stake.*[0-9](?:.[0-9](?:[0-9])?)?u", strings.ToLower(msg)); res == true {
		*count += 1
		return true
	}
	return false
}

// TreatResponse run on the stream to get json responses
func TreatResponse(count *int, useless *int, reader *bufio.Reader, body *message) {
	//var data map[string]interface{}
	sendTelegram("bot starting")
	var prevID string
	for {
		line, _ := reader.ReadBytes('\n')
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		err := json.Unmarshal(line, &body)
		//err := json.Unmarshal(line, &data)
		//log.Printf("line: %v\ndata: %v", line, data)
		if err != nil {
			log.Printf("err: %v", err)
		//	log.Fatal(err)
		}
		if (body.User.ID == 349094942 || body.User.ID == 4197365524) && prevID != body.IDStr {
			// var dn io.Closer
			log.Printf("msg_id: %s name: \033[1;31m%s\033[0m \033[1;32m%s\033[0m",
				body.IDStr,
				body.User.ScreenName,
				body.Text)
			if isBet(count, body.Text) {
				sendTelegram("BET " + string(*count) + "\n\n" + body.Text)
			} else {
				sendTelegram("USELESS " + string(*useless) + "\n\n" + body.Text)
				*useless += 1
			}
			prevID = body.IDStr
		//log.Printf("\n%v\n", data)
		}
	}
}

// GetStream launch all the request logic
func GetStream() {
	count := 0
	useless := 1
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
	TreatResponse(&count, &useless, reader, &body)
}

func main() {
	GetStream()
}
