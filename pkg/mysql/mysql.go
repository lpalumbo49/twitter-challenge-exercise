package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// This is a simplified wrapper for handling MySQL connections. Normally, this should have enhanced support for more operations like transactions
// It's also a good practice to centralize these type of common functionalities in a cross-company toolkit repository
type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	// TODO LP: initialize!
	return &DB{}, nil
}
