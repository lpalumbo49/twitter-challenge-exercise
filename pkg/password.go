package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	HashCost = 14
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), HashCost)
	return string(bytes), err
}

func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
