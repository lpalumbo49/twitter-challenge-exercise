package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"twitter-challenge-exercise/internal/config"
	"twitter-challenge-exercise/pkg"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driver = "mysql"
)

// DB This is a simplified wrapper for handling MySQL connections. Normally, this should have enhanced support for more operations like transactions
// It's also a good practice to centralize these type of common functionalities in a cross-company toolkit repository
type DB struct {
	*sql.DB
}

func NewDB(config config.Configuration) (*DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", config.GetDatabaseUser(),
		config.GetDatabasePassword(), config.GetDatabaseHost(), config.GetDatabasePort(), config.GetDatabaseName())

	db, err := sql.Open(driver, connectionString)
	if err != nil {
		return nil, errors.Join(pkg.NewServerError("error connecting to database"), err)
	}

	db.SetMaxOpenConns(config.GetDatabaseMaxOpenConns())
	db.SetMaxIdleConns(config.GetDatabaseMaxIdleConns())

	db.SetConnMaxLifetime(config.GetDatabaseConnMaxLifetimeSecs())

	return &DB{db}, nil
}
