package pkg

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKeyBytes []byte
var jwtExpirationTime time.Duration

type CustomClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

func InitializeJWT(key string, expirationTime time.Duration) {
	jwtKeyBytes = []byte(key)
	jwtExpirationTime = expirationTime
}

func GenerateJWTToken(userID uint64) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpirationTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKeyBytes)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWTToken(token string) (*CustomClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signature method: %v", token.Header["alg"])
		}
		return jwtKeyBytes, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(*CustomClaims); ok && jwtToken.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
