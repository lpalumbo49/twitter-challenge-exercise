package port

import (
	"context"
	"twitter-challenge-exercise/internal/core/domain"
)

type TweetService interface {
	CreateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error)
}

type TweetRepository interface {
	CreateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error)
}
