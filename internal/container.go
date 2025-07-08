package internal

import (
	"twitter-challenge-exercise/internal/adapter/handler/http"
	"twitter-challenge-exercise/internal/adapter/repository/database"
	"twitter-challenge-exercise/internal/config"
	"twitter-challenge-exercise/internal/core/service"
	"twitter-challenge-exercise/pkg"
	"twitter-challenge-exercise/pkg/mysql"
)

type Container struct {
	router http.Router
}

func StartContainer() (*Container, error) {
	// Here we could choose different configuration files, based on environment
	cfg := config.Configuration{}

	db, err := mysql.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	pkg.InitializeJWT(cfg.GetJwtTokenSecret(), cfg.GetJwtExpirationTime())

	userRepository := database.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	loginService := service.NewLoginService(userService)
	loginHandler := http.NewLoginHandler(loginService)

	tweetRepository := database.NewTweetRepository(db)
	tweetService := service.NewTweetService(tweetRepository)
	tweetHandler := http.NewTweetHandler(tweetService)

	followerRepository := database.NewFollowerRepository(db)
	followerService := service.NewFollowerService(followerRepository, userService)
	followerHandler := http.NewFollowerHandler(followerService)

	timelineRepository := database.NewTimelineRepository(db)
	timelineService := service.NewTimelineService(timelineRepository)
	timelineHandler := http.NewTimelineHandler(timelineService)

	router := http.NewRouter(*loginHandler, *userHandler, *tweetHandler, *followerHandler, *timelineHandler)

	return &Container{
		router: *router,
	}, nil
}

func (c *Container) ServeRouter() error {
	return c.router.Run()
}
