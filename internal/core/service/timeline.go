package service

import (
	"context"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"
)

type timelineService struct {
	repository port.TimelineRepository
}

func NewTimelineService(repository port.TimelineRepository) port.TimelineService {
	return &timelineService{
		repository: repository,
	}
}

func (t timelineService) GetTimelineByUserID(ctx context.Context, userID uint64) ([]domain.TimelineTweet, error) {
	return t.repository.GetTimelineByUserID(ctx, userID)
}
