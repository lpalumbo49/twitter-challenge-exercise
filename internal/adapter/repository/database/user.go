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
	insertUserQuery           = "INSERT INTO user(name, surname, email, password, username, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)"
	updateUserQuery           = "UPDATE user SET name = ?, surname = ?, username = ?, updated_at = ? WHERE id = ?"
	selectUserByIDQuery       = "SELECT id, name, surname, email, username, created_at, updated_at FROM user WHERE id = ?" // No password querying!
	selectUserByEmailQuery    = "SELECT id, name, surname, email, username, created_at, updated_at FROM user WHERE email = ?"
	selectUserByUsernameQuery = "SELECT id, name, surname, email, username, created_at, updated_at FROM user WHERE username = ?"
)

type userRepository struct {
	db *mysql.DB
}

func NewUserRepository(db *mysql.DB) port.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	now := time.Now()

	result, err := u.db.Exec(insertUserQuery, user.Name, user.Surname, user.Email, user.Password, user.Username, now, now)
	if err != nil {
		return user, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.ID = uint64(lastInsertedID)
	user.CreatedAt = now
	user.UpdatedAt = now
	user.Password = ""

	return user, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	now := time.Now()

	_, err := u.db.Exec(updateUserQuery, user.Name, user.Surname, user.Username, now, user.ID)
	if err != nil {
		return user, err
	}

	user.UpdatedAt = now

	return user, nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	row := u.db.QueryRow(selectUserByEmailQuery, email)

	err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, pkg.NewNotFoundError(fmt.Sprintf("user with email %s not found", email))
		}

		return user, err
	}

	return user, nil
}

func (u *userRepository) GetUserByID(ctx context.Context, userID uint64) (domain.User, error) {
	var user domain.User

	row := u.db.QueryRow(selectUserByIDQuery, userID)

	err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, pkg.NewNotFoundError(fmt.Sprintf("user_id %d not found", userID))
		}

		return user, err
	}

	return user, nil
}

func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User

	row := u.db.QueryRow(selectUserByUsernameQuery, username)

	err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, pkg.NewNotFoundError(fmt.Sprintf("user with username %s not found", username))
		}

		return user, err
	}

	return user, nil
}

func (u *userRepository) SearchUsers() ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}
