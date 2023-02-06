package main

import (
	"flag"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	flags          = flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey    = flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret = flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken    = flags.String("access-token", "", "Twitter Access Token")
	accessSecret   = flags.String("access-secret", "", "Twitter Access Secret")
)

var client *twitter.Client

func init() {
	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client = twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(false),
	}
	var _, _, err = client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		log.Fatal("Account verification failed: ", err)
	}
}

func main() {
	http.HandleFunc("/retweet", handler)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	tweetID := 0
	if ok {
		tweetID, _ = strconv.Atoi(keys[0])
	}

	if tweetID != 0 {
		_, err := Retweet(*client, int64(tweetID))
		if err != nil {
			io.WriteString(w, "Retweet failed: "+err.Error())
		}

		io.WriteString(w, "Retweeted: "+strconv.Itoa(tweetID))
	}
}
