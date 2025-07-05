package database

import (
	"context"
	"time"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg/mysql"
)

type tweetRepository struct {
	db *mysql.DB
}

func NewTweetRepository(db *mysql.DB) port.TweetRepository {
	return &tweetRepository{
		db: db,
	}
}

func (t *tweetRepository) CreateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error) {
	// TODO LP: save in db, of course
	tweet.CreatedAt = time.Now()

	return tweet, nil
}
