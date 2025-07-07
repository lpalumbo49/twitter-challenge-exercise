package http

import "github.com/gin-gonic/gin"

type Router struct {
	*gin.Engine
}

func NewRouter(userHandler UserHandler, tweetHandler TweetHandler) (*Router, error) {
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
	routerGroup.POST("/login", nil)

	// TODO LP: estas operaciones deben usar jwt
	// routerGroup.GET("/user/:id", userHandler.GetUserByID)
	// routerGroup.PUT("/user/:id", userHandler.UpdateUser)

	routerGroup.POST("/tweet", tweetHandler.CreateTweet)
	// routerGroup.GET("/tweet/:id", tweetHandler.GetTweetByID)
	// routerGroup.PUT("/tweet/:id", tweetHandler.UpdateTweet)

	// routerGroup.GET("/timeline", timelineHandler.GetTimelineByUserID)

	// routerGroup.POST("/follower", followerHandler.CreateFollower)

	return &Router{router}, nil
}
