package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg"
	"twitter-challenge-exercise/pkg/mysql"
)

const (
	insertFollowerQuery      = "INSERT INTO follower(user_id, followed_by_user_id, created_at) VALUES (?, ?, ?)"
	selectFollowerByIDsQuery = "SELECT user_id, followed_by_user_id, created_at from follower WHERE user_id = ? AND followed_by_user_id = ?"
)

type followerRepository struct {
	db *mysql.DB
}

func NewFollowerRepository(db *mysql.DB) port.FollowerRepository {
	return &followerRepository{
		db: db,
	}
}

func (f *followerRepository) CreateFollower(ctx context.Context, follower domain.Follower) (domain.Follower, error) {
	now := time.Now()

	_, err := f.db.Exec(insertFollowerQuery, follower.UserID, follower.FollowedByUserID, now)
	if err != nil {
		return follower, err
	}

	follower.CreatedAt = now

	return follower, nil
}

func (f *followerRepository) GetFollowerByIDs(ctx context.Context, userID, followedByUserID uint64) (domain.Follower, error) {
	var follower domain.Follower

	row := f.db.QueryRow(selectFollowerByIDsQuery, userID, followedByUserID)

	err := row.Scan(&follower.UserID, &follower.FollowedByUserID, &follower.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return follower, pkg.NewEntityNotFoundError(fmt.Sprintf("follower with user_id %d and followed_by_user_id %d not found", userID, followedByUserID))
		}

		return follower, err
	}

	return follower, nil
}
