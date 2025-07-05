package internal

import (
	"twitter-challenge-exercise/internal/adapter/handler/http"
	"twitter-challenge-exercise/internal/adapter/repository/database"
	"twitter-challenge-exercise/internal/core/service"
	"twitter-challenge-exercise/pkg/mysql"
)

type Container struct {
	// TODO: acá van los puntos de entrada de todo (handlers, de adapter) revisar igual el acoplamiento?
	router http.Router
}

func StartContainer() (*Container, error) {
	// TODO LP: config.MySQL (env!)
	db, err := mysql.NewDB()
	if err != nil {
		// TODO: error handling!
		return nil, err
	}

	tweetRepository := database.NewTweetRepository(db)
	tweetService := service.NewTweetService(tweetRepository)
	tweetHandler := http.NewTweetHandler(tweetService)

	router, err := http.NewRouter(*tweetHandler)
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
