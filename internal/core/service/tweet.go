package service

import (
	"context"
	"errors"
	"fmt"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg"
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
	tweet, err := t.repository.CreateTweet(ctx, tweet)
	if err != nil {
		return tweet, errors.Join(pkg.NewServerError("error creating tweet"), err)
	}

	return t.repository.GetTweetByID(ctx, tweet.ID)
}

func (t *tweetService) UpdateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error) {
	savedTweet, err := t.repository.GetTweetByID(ctx, tweet.ID)
	if err != nil {
		return tweet, err
	}

	if savedTweet.UserID != tweet.UserID {
		return tweet, pkg.NewForbiddenError(fmt.Sprintf("user %d does not own this tweet", tweet.UserID))
	}

	savedTweet.Text = tweet.Text

	savedTweet, err = t.repository.UpdateTweet(ctx, tweet)
	if err != nil {
		return tweet, errors.Join(pkg.NewServerError("error updating tweet"), err)
	}

	return t.repository.GetTweetByID(ctx, tweet.ID)
}

func (t *tweetService) GetTweetByID(ctx context.Context, tweetID uint64) (domain.Tweet, error) {
	return t.repository.GetTweetByID(ctx, tweetID)
}
