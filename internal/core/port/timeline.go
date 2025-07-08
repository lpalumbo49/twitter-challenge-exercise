package port

import (
	"context"
	"twitter-challenge-exercise/internal/core/domain"
)

type TimelineService interface {
	GetTimelineByUserID(ctx context.Context, userID uint64) ([]domain.TimelineTweet, error)
}

type TimelineRepository interface {
	GetTimelineByUserID(ctx context.Context, userID uint64) ([]domain.TimelineTweet, error)
}
