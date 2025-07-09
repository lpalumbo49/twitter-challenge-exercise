package database

import (
	"context"
	"github.com/stretchr/testify/mock"
	"twitter-challenge-exercise/internal/core/domain"
)

type UserMockRepository struct {
	mock.Mock
}

func NewUserMockRepository() *UserMockRepository {
	return &UserMockRepository{}
}

func (u *UserMockRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	args := u.Called(ctx, user)

	responseUser, _ := args.Get(0).(domain.User)
	err, _ := args.Get(1).(error)

	return responseUser, err
}

func (u *UserMockRepository) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	args := u.Called(ctx, user)

	responseUser, _ := args.Get(0).(domain.User)
	err, _ := args.Get(1).(error)

	return responseUser, err
}

func (u *UserMockRepository) GetUserByID(ctx context.Context, userID uint64) (domain.User, error) {
	args := u.Called(ctx, userID)

	responseUser, _ := args.Get(0).(domain.User)
	err, _ := args.Get(1).(error)

	return responseUser, err
}

func (u *UserMockRepository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	args := u.Called(ctx, username)

	responseUser, _ := args.Get(0).(domain.User)
	err, _ := args.Get(1).(error)

	return responseUser, err
}

func (u *UserMockRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	args := u.Called(ctx, email)

	responseUser, _ := args.Get(0).(domain.User)
	err, _ := args.Get(1).(error)

	return responseUser, err
}

func (u *UserMockRepository) SearchUsers() ([]domain.User, error) {
	args := u.Called()

	responseUsers, _ := args.Get(0).([]domain.User)
	err, _ := args.Get(1).(error)

	return responseUsers, err
}
