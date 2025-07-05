package service

import (
	"context"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"
)

type tweetService struct {
	repository port.TweetRepository
}

func NewTweetService(repository port.TweetRepository) port.TweetService {
	return &tweetService{
		repository: repository,
	}
}

func (t *tweetService) CreateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error) {
	return t.repository.CreateTweet(ctx, domain.Tweet{Text: "test"})
}
