package dto

import (
	"time"
	"twitter-challenge-exercise/internal/core/domain"
)

type CreateTweetRequest struct {
	UserID uint64 `json:"user_id" validate:"required"`
	Text   string `json:"text" validate:"required,min=1,max=280"`
}

type TweetResponse struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapCreateTweetRequestToTweet(request CreateTweetRequest) domain.Tweet {
	return domain.Tweet{
		Text:   request.Text,
		UserID: request.UserID,
	}
}

func MapTweetToTweetResponse(tweet domain.Tweet) TweetResponse {
	return TweetResponse{
		ID:        tweet.ID,
		UserID:    tweet.UserID,
		Text:      tweet.Text,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
	}
}

// --------------------------------------------------------------------------------
type UpdateTweetRequest struct {
	ID     uint64 `json:"id" validate:"required"`
	UserID uint64 `json:"user_id" validate:"required"`
	Text   string `json:"text" validate:"required,min=1,max=280"`
}

func MapUpdateTweetRequestToTweet(request UpdateTweetRequest) domain.Tweet {
	return domain.Tweet{
		ID:     request.ID,
		Text:   request.Text,
		UserID: request.UserID,
	}
}
