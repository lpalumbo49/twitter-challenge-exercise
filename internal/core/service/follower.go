package service

import (
	"context"
	"errors"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg"
)

type followerService struct {
	repository  port.FollowerRepository
	userService port.UserService
}

func NewFollowerService(repository port.FollowerRepository, userService port.UserService) port.FollowerService {
	return &followerService{
		repository:  repository,
		userService: userService,
	}
}

func (f *followerService) CreateFollower(ctx context.Context, follower domain.Follower) (domain.Follower, error) {
	if follower.UserID == follower.FollowedByUserID {
		return follower, pkg.NewBusinessError("cannot follow yourself")
	}

	// Check if followed user_id exists (double check with table foreign key)
	_, err := f.userService.GetUserByID(ctx, follower.FollowedByUserID)
	if err != nil {
		if pkg.IsEntityNotFoundError(err) {
			return follower, pkg.NewBusinessError("followed_by_user_id does not exist")
		}

		return follower, err
	}

	existingFollower, err := f.repository.GetFollowerByIDs(ctx, follower.UserID, follower.FollowedByUserID)
	if err != nil && !pkg.IsEntityNotFoundError(err) {
		return follower, err
	}

	if existingFollower.UserID != 0 && existingFollower.FollowedByUserID != 0 {
		// This follower already exists
		return existingFollower, nil
	}

	follower, err = f.repository.CreateFollower(ctx, follower)
	if err != nil {
		return follower, errors.Join(pkg.NewServerError("error creating follower"), err)
	}

	return f.repository.GetFollowerByIDs(ctx, follower.UserID, follower.FollowedByUserID)
}
