package internal

import (
	"twitter-challenge-exercise/internal/adapter/handler/http"
	"twitter-challenge-exercise/internal/adapter/repository/database"
	"twitter-challenge-exercise/internal/config"
	"twitter-challenge-exercise/internal/core/service"
	"twitter-challenge-exercise/pkg/mysql"
)

type Container struct {
	// TODO: acá van los puntos de entrada de todo (handlers, de adapter) revisar igual el acoplamiento?
	router http.Router
}

func StartContainer() (*Container, error) {
	// TODO LP: config.MySQL (env! mejor)
	db, err := mysql.NewDB(config.Configuration{})
	if err != nil {
		return nil, err
	}

	userRepository := database.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	tweetRepository := database.NewTweetRepository(db)
	tweetService := service.NewTweetService(tweetRepository)
	tweetHandler := http.NewTweetHandler(tweetService)

	followerRepository := database.NewFollowerRepository(db)
	followerService := service.NewFollowerService(followerRepository, userService)
	followerHandler := http.NewFollowerHandler(followerService)

	router, err := http.NewRouter(*userHandler, *tweetHandler, *followerHandler)
	if err != nil {
		return nil, err
	}

	return &Container{
		router: *router,
	}, nil
}

func (c *Container) ServeRouter() error {
	return c.router.Run()
}
