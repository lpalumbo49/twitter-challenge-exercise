package domain

import "time"

type User struct {
	ID        uint64
	Name      string
	Surname   string
	Email     string
	Password  string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
