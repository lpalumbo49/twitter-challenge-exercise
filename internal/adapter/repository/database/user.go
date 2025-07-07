package database

import (
	"context"
	"time"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg/mysql"
)

type userRepository struct {
	db *mysql.DB
}

func NewUserRepository(db *mysql.DB) port.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (t *userRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	// TODO LP: save in db, of course
	user.CreatedAt = time.Now()

	return user, nil
}

func (t *userRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	panic("unimplemented")
}

func (t *userRepository) GetUserByID(ctx context.Context, userID uint64) (domain.User, error) {
	panic("unimplemented")
}

func (t *userRepository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	panic("unimplemented")
}
