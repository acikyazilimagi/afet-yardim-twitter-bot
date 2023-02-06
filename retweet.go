package main

import (
	"github.com/dghubble/go-twitter/twitter"
)

func Retweet(client twitter.Client, tweetID int64) (*twitter.Tweet, error) {
	tweet, _, err := client.Statuses.Retweet(tweetID, nil)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}
