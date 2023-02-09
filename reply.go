package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dghubble/oauth1"
)

const (
	consumerKey    = ""
	consumerSecret = ""
	accessToken    = ""
	accessSecret   = ""
)

func reply() {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Read tweet id from json file
	tweetIDBytes, err := ioutil.ReadFile("tweet_id.json")
	if err != nil {
		log.Fatalf("Error reading tweet id: %v", err)
	}
	var tweetID struct {
		ID int64 `json:"id"`
	}
	if err := json.Unmarshal(tweetIDBytes, &tweetID); err != nil {
		log.Fatalf("Error unmarshaling tweet id: %v", err)
	}

	// Send reply
	message := "your reply"
	status := fmt.Sprintf("@username %s", message)
	url := fmt.Sprintf("https://api.twitter.com/1.1/statuses/update.json?status=%s&in_reply_to_status_id=%d",
		status, tweetID.ID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending reply: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("Request failed: %s", body)
	}
	fmt.Println("Reply sent successfully!")
}
