package db_types

import (
	"context"
	"database/sql"
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
	// DB returns the underlying database connection pool.
	DB() *sql.DB
	// Dialect returns the database SQL dialect.
	Dialect() Dialect
	// Start opens and validates the database connection pool.
	Start(ctx context.Context) error
	// Stop closes the database connection pool.
	Stop(ctx context.Context) error
}
