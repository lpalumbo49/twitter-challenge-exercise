package port

import (
	"context"
	"twitter-challenge-exercise/internal/core/domain"
)

type FollowerService interface {
	CreateFollower(ctx context.Context, follower domain.Follower) (domain.Follower, error)
}

type FollowerRepository interface {
	CreateFollower(ctx context.Context, follower domain.Follower) (domain.Follower, error)
	GetFollowerByIDs(ctx context.Context, userID, followedByUserID uint64) (domain.Follower, error)
}
