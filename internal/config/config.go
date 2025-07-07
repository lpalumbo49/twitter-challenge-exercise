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

	databaseMaxIdleConns        = 25
	databaseMaxOpenConns        = 50
	databaseConnMaxLifetimeSecs = 300
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

func (*Configuration) GetDatabaseConnMaxLifetimeSecs() time.Duration {
	return databaseConnMaxLifetimeSecs * time.Second
}
