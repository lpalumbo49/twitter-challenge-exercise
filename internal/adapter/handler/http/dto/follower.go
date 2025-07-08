package dto

import (
	"time"
	"twitter-challenge-exercise/internal/core/domain"
)

type CreateFollowerRequest struct {
	UserID           uint64 `json:"user_id" validate:"required"`
	FollowedByUserID uint64 `json:"followed_by_user_id" validate:"required"`
}

type FollowerResponse struct {
	UserID           uint64    `json:"user_id"`
	FollowedByUserID uint64    `json:"followed_by_user_id"`
	CreatedAt        time.Time `json:"created_at"`
}

func MapCreateFollowerRequestToFollower(request CreateFollowerRequest) domain.Follower {
	return domain.Follower{
		UserID:           request.UserID,
		FollowedByUserID: request.FollowedByUserID,
	}
}

func MapFollowerToFollowerResponse(follower domain.Follower) FollowerResponse {
	return FollowerResponse{
		UserID:           follower.UserID,
		FollowedByUserID: follower.FollowedByUserID,
		CreatedAt:        follower.CreatedAt,
	}
}
