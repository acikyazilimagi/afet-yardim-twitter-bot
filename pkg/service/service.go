package service

import (
	apiError "afet-yardim-twitter-bot/pkg/error"
	"afet-yardim-twitter-bot/pkg/structure"
	"context"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/sirupsen/logrus"
)

type BotService interface {
	Retweet(ctx context.Context, req structure.RetweetRequest) (structure.RetweetResponse, error)
}

type botService struct {
	logger        *logrus.Logger
	twitterClient *twitter.Client
}

func (s botService) Retweet(_ context.Context, req structure.RetweetRequest) (structure.RetweetResponse, error) {
	mLogger := s.logger.WithFields(logrus.Fields{
		"method":  "Retweet",
		"tweetId": req.TweetID,
	})

	_, _, err := s.twitterClient.Statuses.Retweet(int64(req.TweetID), nil)
	if err != nil {
		mLogger.Errorf("error from twitter client: %v", err.Error())
		return structure.RetweetResponse{}, apiError.NewBadRequestError(err)
	}

	mLogger.Info("retweet service executed successfully")
	return structure.RetweetResponse{}, nil
}

func NewBotService(l *logrus.Logger, twitterClient *twitter.Client) BotService {
	return &botService{
		logger:        l,
		twitterClient: twitterClient,
	}
}
