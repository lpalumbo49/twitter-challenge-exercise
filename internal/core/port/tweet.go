package port

import (
	"context"
	"twitter-challenge-exercise/internal/core/domain"
)

type TweetService interface {
	CreateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error)
	UpdateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error)
	GetTweetByID(ctx context.Context, tweetID uint64) (domain.Tweet, error)
}

type TweetRepository interface {
	CreateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error)
	UpdateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error)
	GetTweetByID(ctx context.Context, tweetID uint64) (domain.Tweet, error)
}
