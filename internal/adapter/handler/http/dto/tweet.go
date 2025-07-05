package dto

import (
	"time"
	"twitter-challenge-exercise/internal/core/domain"
)

type TweetResponse struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
