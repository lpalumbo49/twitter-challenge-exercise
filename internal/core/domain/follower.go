package domain

import "time"

type Follower struct {
	UserID           uint64
	FollowedByUserID uint64
	CreatedAt        time.Time
}
