package port

import (
	"context"
	"twitter-challenge-exercise/internal/core/domain"
)

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByID(ctx context.Context, userID uint64) (domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}
