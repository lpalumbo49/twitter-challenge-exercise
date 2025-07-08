package service

import (
	"context"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg"
)

type loginService struct {
	userService port.UserService
}

func NewLoginService(userService port.UserService) port.LoginService {
	return &loginService{
		userService: userService,
	}
}

func (l *loginService) UserLogin(ctx context.Context, email string, password string) (string, error) {
	// Check if user exists, check their password, and then generate JWT token
	existingUser, err := l.userService.GetUserByEmail(ctx, email)
	if err != nil {
		return "", pkg.NewBusinessError("invalid email or password")
	}

	if !pkg.VerifyPassword(password, existingUser.Password) {
		return "", pkg.NewBusinessError("invalid email or password")
	}

	token, err := pkg.GenerateJWTToken(existingUser.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
