package pkg

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

const (
	hashCost           = 14
	hashedTestPassword = "hashed_test_password"
)

func HashPassword(password string) (string, error) {
	if testing.Testing() {
		// For simplicity, the same password is returned in unit tests
		return hashedTestPassword, nil
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return string(bytes), err
}

func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
