package handler

import (
	apiError "afet-yardim-twitter-bot/pkg/error"
	"afet-yardim-twitter-bot/pkg/service"
	"afet-yardim-twitter-bot/pkg/structure"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	RetweetHandler gin.HandlerFunc
}

func RetweetHandler(s service.BotService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := structure.RetweetRequest{}
		err := ctx.ShouldBindQuery(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(400, apiError.NewBadRequestError(err).ToJson())
			return
		}

		resp, err := s.Retweet(ctx, req)

		if err != nil {
			sErr := err.(apiError.ApiError)
			ctx.AbortWithStatusJSON(sErr.StatusCode, sErr.ToJson())
			return
		}

		ctx.JSON(200, resp)
	}
}

func New(s service.BotService) Handlers {
	return Handlers{
		RetweetHandler: RetweetHandler(s),
	}
}
