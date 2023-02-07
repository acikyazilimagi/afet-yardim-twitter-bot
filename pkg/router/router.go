package router

import (
	"afet-yardim-twitter-bot/pkg/handler"
	"github.com/gin-gonic/gin"
)

// CORS middleware
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func NewRouter(handlers handler.Handlers) *gin.Engine {
	r := gin.New()

	r.Use(CORS())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	retweet := r.Group("/retweet")
	{
		retweet.GET("/", handlers.RetweetHandler)
	}

	return r
}
