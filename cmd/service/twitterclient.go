package service

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func initTwitterClient() (*twitter.Client, error) {
	oauthConfig := oauth1.NewConfig(cfg.Twitter.ConsumerKey, cfg.Twitter.ConsumerSecret)
	oauthToken := oauth1.NewToken(cfg.Twitter.AccessToken, cfg.Twitter.AccessTokenSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, oauthToken)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(false),
	}
	var _, _, err = client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, fmt.Errorf("account verification failed: %v", err)
	}

	return client, nil
}
