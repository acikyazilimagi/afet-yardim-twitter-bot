package endpoint

import (
	"afet-yardim-twitter-bot/pkg/service"
	"afet-yardim-twitter-bot/pkg/structure"
	"context"
)

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

type Endpoints struct {
	RetweetEndpoint Endpoint
}

func New(s service.BotService) Endpoints {
	return Endpoints{
		RetweetEndpoint: RetweetEndpoint(s),
	}
}

func RetweetEndpoint(s service.BotService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(structure.RetweetRequest)
		return s.Retweet(ctx, req)
	}
}
