package dto

import (
	"time"
	"twitter-challenge-exercise/internal/core/domain"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,max=32"`
	Surname  string `json:"surname" validate:"required,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Username string `json:"username" validate:"required,min=3,max=32"`
}

type UserResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapCreateUserRequestToUser(request CreateUserRequest) domain.User {
	return domain.User{
		Name:     request.Name,
		Surname:  request.Surname,
		Email:    request.Email,
		Password: request.Password,
		Username: request.Username,
	}
}

func MapUserToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// --------------------------------------------------------------------------------
type UpdateUserRequest struct {
	ID       uint64 `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required,max=32"`
	Surname  string `json:"surname" validate:"required,max=32"`
	Username string `json:"username" validate:"required,min=3,max=32"`
}

func MapUpdateUserRequestToUser(request UpdateUserRequest) domain.User {
	return domain.User{
		ID:       request.ID,
		Name:     request.Name,
		Surname:  request.Surname,
		Username: request.Username,
	}
}
