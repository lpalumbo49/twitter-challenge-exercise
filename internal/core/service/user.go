package service

import (
	"context"
	"errors"
	"fmt"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg"
)

type userService struct {
	repository port.UserRepository
}

func NewUserService(repository port.UserRepository) port.UserService {
	return &userService{
		repository: repository,
	}
}

func (u *userService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	existingUser, err := u.repository.GetUserByEmail(ctx, user.Email)
	if err != nil && !pkg.IsNotFoundError(err) {
		return user, err
	}

	if existingUser.ID != 0 {
		return user, pkg.NewBusinessError(fmt.Sprintf("there is an existing user with email '%s'", user.Email))
	}

	existingUser, err = u.repository.GetUserByUsername(ctx, user.Username)
	if err != nil && !pkg.IsNotFoundError(err) {
		return user, err
	}

	if existingUser.ID != 0 {
		return user, pkg.NewBusinessError(fmt.Sprintf("there is an existing user with username '%s'", user.Username))
	}

	hashedPassword, err := pkg.HashPassword(user.Password)
	if err != nil {
		return user, errors.Join(pkg.NewServerError("error encrypting user password"), err)
	}

	user.Password = hashedPassword

	user, err = u.repository.CreateUser(ctx, user)
	if err != nil {
		return user, errors.Join(pkg.NewServerError("error creating user"), err)
	}

	// Once created, this entity could have changed their values due to concurrent behaviour
	return u.repository.GetUserByID(ctx, user.ID)
}

func (u *userService) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	existingUser, err := u.repository.GetUserByUsername(ctx, user.Username)
	if err != nil && !pkg.IsNotFoundError(err) {
		return user, err
	}

	// Already existing user, different user_id than the one that we are trying to update
	if existingUser.ID != 0 && user.ID != existingUser.ID {
		return user, pkg.NewBusinessError(fmt.Sprintf("there is an existing user with username '%s'", user.Username))
	}

	savedUser, err := u.repository.GetUserByID(ctx, user.ID)
	if err != nil {
		return user, err
	}

	savedUser.Username = user.Username
	savedUser.Name = user.Name
	savedUser.Surname = user.Surname

	savedUser, err = u.repository.UpdateUser(ctx, savedUser)
	if err != nil {
		return user, errors.Join(pkg.NewServerError("error updating user"), err)
	}

	// Once created, this entity could have changed their values due to concurrent behaviour
	return u.repository.GetUserByID(ctx, savedUser.ID)
}

func (u *userService) GetUserByID(ctx context.Context, userID uint64) (domain.User, error) {
	return u.repository.GetUserByID(ctx, userID)
}

func (u *userService) SearchUsers(ctx context.Context) ([]domain.User, error) {
	// There should be some type of criteria for user filtering here
	// For this exercise, it is simplified by searching all users
	return u.repository.SearchUsers()
}
