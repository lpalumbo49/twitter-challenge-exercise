package http

import "github.com/gin-gonic/gin"

type Router struct {
	*gin.Engine
}

func NewRouter(tweetHandler TweetHandler) (*Router, error) {
	router := gin.Default()

	// TODO LP: esto debe usar jwt
	routerGroup := router.Group("/api/v1")

	routerGroup.POST("/tweet", tweetHandler.CreateTweet)

	return &Router{router}, nil
}
