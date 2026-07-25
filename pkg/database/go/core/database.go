package db_core

import (
	"database/sql"

	di "github.com/nimaeskandary/app_repo/pkg/di/go"
)

// Dialect identifies the SQL dialect used by a database.
type Dialect string

const (
	// DialectSQLite identifies SQLite databases.
	DialectSQLite Dialect = "sqlite"
	// DialectPostgres identifies PostgreSQL databases.
	DialectPostgres Dialect = "postgres"
)

// SQLDatabase exposes a validated SQL connection and its lifecycle.
type SQLDatabase interface {
	di.Lifecycle
	// DB returns the underlying database connection pool.
	DB() *sql.DB
	// Dialect returns the database SQL dialect.
	Dialect() Dialect
}
