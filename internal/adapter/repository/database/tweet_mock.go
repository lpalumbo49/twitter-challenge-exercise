package database

import (
	"context"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"

	"github.com/stretchr/testify/mock"
)

type tweetMockRepository struct {
	mock.Mock
}

func NewTweetMockRepository() port.TweetRepository {
	return &tweetMockRepository{}
}

func (t *tweetMockRepository) CreateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error) {
	args := t.Called(ctx, tweet)

	responseTweet, _ := args.Get(0).(domain.Tweet)
	err, _ := args.Get(1).(error)

	return responseTweet, err
}
