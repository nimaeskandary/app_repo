package db_types

import (
	"context"
	"database/sql"
)

type Dialect string

const (
	DialectSQLite   Dialect = "sqlite"
	DialectPostgres Dialect = "postgres"
)

type SQLDatabase interface {
	DB() *sql.DB
	Dialect() Dialect
	Stop(ctx context.Context) error
}
