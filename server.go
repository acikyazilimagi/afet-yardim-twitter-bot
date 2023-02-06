package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

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

// var client *twitter.Client

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error

var client = &http.Client{
	// Timeout: 10 * time.Second,
}

func getTweetId(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data Tweet_Data
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal((err))
	}
	log.Println(data.tweet_id)

	retweet(data.tweet_id)

	fmt.Printf("got /tweet/:id request\n")
	io.WriteString(w, "This is a tweet!\n")
}

func retweet(tweet_id int64) {

	// config := oauth1.Config{
	// 	ConsumerKey:    "consumerKey",
	// 	ConsumerSecret: "consumerSecret",
	// 	CallbackURL:    "http://mysite.com/oauth/twitter/callback",
	// 	Endpoint:       twitter.AuthorizeEndpoint,
	// }

	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("token", "tokenSecret")

	// httpClient will automatically authorize http.Request's
	httpClient := config.Client(oauth1.NoContext, token)

	// example Twitter API request
	path := "https://api.twitter.com/1.1/statuses/home_timeline.json?count=2"
	resp, _ := httpClient.Get(path)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Raw Response Body:\n%v\n", string(body))

	url := fmt.Sprintf("https://api.twitter.com/1.1/statuses/retweet/%d.json", tweet_id)

	client := http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		//Handle Error
		log.Fatal(err)
	}

	// todo :: load from env
	strr := fmt.Sprintf("OAuth oauth_consumer_key=%s, oauth_nonce=AUTO_GENERATED_NONCE, oauth_signature=AUTO_GENERATED_SIGNATURE, oauth_signature_method=HMAC-SHA1, oauth_timestamp=AUTO_GENERATED_TIMESTAMP, oauth_token=%s, oauth_version=1.0", creds.ConsumerKey, token)

	req.Header = http.Header{
		"authorization": []string{strr},
		"content-type":  []string{"application/json"},
	}

	res, err := client.Do(req)

	if err != nil {
		//Handle Error
		panic(err)
	}

	// close body
	defer res.Body.Close()

	sr, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %s\n", string(sr))
}

func handleAuth(w http.ResponseWriter, r *http.Request) {

	// config := oauth1.NewConfig("consumerKey", "consumerSecret")
	// token := oauth1.NewToken("token", "tokenSecret")

	// // httpClient will automatically authorize http.Request's
	// httpClient := config.Client(oauth1.NoContext, token)

	// // example Twitter API request
	// path := "https://api.twitter.com/1.1/statuses/home_timeline.json?count=2"
	// resp, _ := httpClient.Get(path)
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Printf("Raw Response Body:\n%v\n", string(body))

	config := oauth1.NewConfig("3nVuSoBZnx6U4vzUxf5w", "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys")

	callbackURL := "http://localhost:8080/callback"
	log.Print(callbackURL)

	requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		// Handle error
	}

	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		// Handle error
	}

	// Redirect user to the authorizationURL to grant access
	http.Redirect(w, r, authorizationURL.String(), http.StatusFound)

	// The user will be redirected back to the callbackURL with the oauth_token and oauth_verifier
	// in the query parameters
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}

	oauthToken := ("78264296-XUtCd2Ovy7VeXLrEUPQf2IYdYYjlVaELyi10G6xIt")
	oauthVerifier := values.Get("oauth_verifier")

	log.Print("oauthToken:", oauthToken)
	log.Print("values:", values)

	// Use the oauthToken and oauthVerifier to get an access token
	accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, oauthVerifier)
	if err != nil {
		// Handle error
	}

	log.Print("access token:", (accessToken), "access secret:", (accessSecret))

}

func main() {

	// // load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	creds.ConsumerKey = os.Getenv("CONSUMER_KEY")
	creds.ConsumerSecret = os.Getenv("CONSUMER_SECRET")
	creds.AccessToken = os.Getenv("ACCESS_TOKEN")
	creds.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")

	// make post only request
	http.HandleFunc("/tweet/", getTweetId)
	http.HandleFunc("/auth", handleAuth)

	err = http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
