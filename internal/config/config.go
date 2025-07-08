package config

import (
	"os"
	"strconv"
	"time"
)

const (
	databaseHost     = "DATABASE_HOST"
	databasePort     = "DATABASE_PORT"
	databaseUser     = "DATABASE_USER"
	databasePassword = "DATABASE_PASSWORD"
	databaseName     = "DATABASE_NAME"

	databaseMaxIdleConns    = 25
	databaseMaxOpenConns    = 50
	databaseConnMaxLifetime = 300 * time.Second

	jwtTokenSecret    = "JWT_TOKEN_SECRET"
	jwtExpirationTime = time.Hour * 24
)

type Configuration struct{}

func (*Configuration) GetDatabaseHost() string {
	return os.Getenv(databaseHost)
}

func (*Configuration) GetDatabasePort() int {
	port, err := strconv.ParseInt(os.Getenv(databasePort), 10, 64)
	if err != nil {
		panic("invalid database port number")
	}

	return int(port)
}

func (*Configuration) GetDatabaseUser() string {
	return os.Getenv(databaseUser)
}

func (*Configuration) GetDatabasePassword() string {
	return os.Getenv(databasePassword)
}

func (*Configuration) GetDatabaseName() string {
	return os.Getenv(databaseName)
}

func (*Configuration) GetDatabaseMaxIdleConns() int {
	return databaseMaxIdleConns
}

func (*Configuration) GetDatabaseMaxOpenConns() int {
	return databaseMaxOpenConns
}

func (*Configuration) GetDatabaseConnMaxLifetime() time.Duration {
	return databaseConnMaxLifetime
}

func (*Configuration) GetJwtTokenSecret() string {
	return os.Getenv(jwtTokenSecret)
}

func (*Configuration) GetJwtExpirationTime() time.Duration {
	return jwtExpirationTime
}
