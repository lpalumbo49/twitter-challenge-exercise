package database

import (
	"context"
	"github.com/stretchr/testify/mock"
	"twitter-challenge-exercise/internal/core/domain"
)

type TweetMockRepository struct {
	mock.Mock
}

func NewTweetMockRepository() *TweetMockRepository {
	return &TweetMockRepository{}
}

func (t *TweetMockRepository) CreateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error) {
	args := t.Called(ctx, tweet)

	responseTweet, _ := args.Get(0).(domain.Tweet)
	err, _ := args.Get(1).(error)

	return responseTweet, err
}

func (t *TweetMockRepository) UpdateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error) {
	args := t.Called(ctx, tweet)

	responseTweet, _ := args.Get(0).(domain.Tweet)
	err, _ := args.Get(1).(error)

	return responseTweet, err
}

func (t *TweetMockRepository) GetTweetByID(ctx context.Context, tweetID uint64) (domain.Tweet, error) {
	args := t.Called(ctx, tweetID)

	responseTweet, _ := args.Get(0).(domain.Tweet)
	err, _ := args.Get(1).(error)

	return responseTweet, err
}
