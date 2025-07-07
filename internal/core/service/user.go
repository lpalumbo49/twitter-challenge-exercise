package service

import (
	"context"
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
		return domain.User{}, err
	}

	if existingUser.ID != 0 {
		return domain.User{}, pkg.NewBusinessError(fmt.Sprintf("there is an existing user with email '%s'", user.Email))
	}

	existingUser, err = u.repository.GetUserByUsername(ctx, user.Username)
	if err != nil && !pkg.IsNotFoundError(err) {
		return domain.User{}, err
	}

	if existingUser.ID != 0 {
		return domain.User{}, pkg.NewBusinessError(fmt.Sprintf("there is an existing user with username '%s'", user.Username))
	}

	// TODO: hash password
	user.Password = "hash!!!"

	user, err = u.repository.CreateUser(ctx, user)
	if err != nil {
		// TODO LP: what could go wrong? only 500s?
		return domain.User{}, err
	}

	// Once created, this entity could have changed their values due to concurrent behaviour
	return u.repository.GetUserByID(ctx, user.ID)
}
