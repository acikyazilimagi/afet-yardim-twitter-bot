package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

type Tweet_Data struct {
	tweet_id int64
}

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

var creds Credentials
var client *twitter.Client

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error
func getClient() (*twitter.Client, error) {

	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getTweetId(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data Tweet_Data
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	log.Println(data.tweet_id)

	retweet(data.tweet_id)

	fmt.Printf("got /tweet/:id request\n")
	io.WriteString(w, "This is a tweet!\n")
}

func retweet(tweet_id int64) {
	client.Statuses.Retweet(tweet_id, nil)
}

func main() {

	loadData()

	// make post only request
	http.HandleFunc("/tweet/", getTweetId)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func loadData() {
	// env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	creds.AccessToken = os.Getenv("ACCESS_TOKEN")
	creds.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	creds.ConsumerKey = os.Getenv("CONSUMER_KEY")
	creds.ConsumerSecret = os.Getenv("CONSUMER_SECRET")

	// in-memory cache the client
	// todo :: check for errors
	client, _ = getClient()
}
