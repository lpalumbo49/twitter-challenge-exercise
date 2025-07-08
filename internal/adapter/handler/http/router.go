package http

import (
	"github.com/gin-gonic/gin"
	"twitter-challenge-exercise/internal/adapter/handler/http/middleware"
)

type Router struct {
	*gin.Engine
}

func NewRouter(loginHandler LoginHandler, userHandler UserHandler, tweetHandler TweetHandler, followerHandler FollowerHandler,
	timelineHandler TimelineHandler) *Router {
	router := gin.Default()

	// Tweets
	// ● Los usuarios deben poder publicar mensajes cortos (tweets) que no excedan un
	// límite de caracteres (por ejemplo, 280 caracteres).

	// Follow:
	// ● Los usuarios deben poder seguir a otros usuarios.

	// Timeline:
	// ● Deben poder ver una línea de tiempo que muestre los tweets de los usuarios a los
	// que siguen.

	// Todos los usuarios son válidos, no es necesario crear un módulo de signin ni
	// manejar sesiones. Se puede enviar el identificador de un usuario por header,
	// param, body o por donde crea más conveniente.

	routerGroup := router.Group("/api/v1")

	routerGroup.POST("/user", userHandler.CreateUser)
	routerGroup.POST("/login", loginHandler.UserLogin)

	secureRouter := routerGroup.Group("/").Use(middleware.Auth())

	secureRouter.GET("/user/:id", userHandler.GetUserByID)
	secureRouter.PUT("/user/:id", userHandler.UpdateUser)
	secureRouter.GET("/users", userHandler.GetUsers)

	secureRouter.POST("/tweet", tweetHandler.CreateTweet)
	secureRouter.GET("/tweet/:id", tweetHandler.GetTweetByID)
	secureRouter.PUT("/tweet/:id", tweetHandler.UpdateTweet)

	secureRouter.GET("/timeline", timelineHandler.GetTimelineByUserID)

	secureRouter.POST("/follower", followerHandler.CreateFollower)

	return &Router{router}
}
