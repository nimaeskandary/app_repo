package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
)

// sqlDatabase wraps a validated database connection pool.
type sqlDatabase struct {
	db        *sql.DB
	dialect   db_types.Dialect
	stopOnce  sync.Once
	stopError error
}

// newSQLDatabase verifies a connection before making it available to callers.
func newSQLDatabase(db *sql.DB, dialect db_types.Dialect) (db_types.SQLDatabase, error) {
	if err := db.Ping(); err != nil {
		return nil, errors.Join(
			fmt.Errorf("ping %s database: %w", dialect, err),
			db.Close(),
		)
	}

	return &sqlDatabase{
		db:      db,
		dialect: dialect,
	}, nil
}

func (d *sqlDatabase) DB() *sql.DB {
	return d.db
}

func (d *sqlDatabase) Dialect() db_types.Dialect {
	return d.dialect
}

func (d *sqlDatabase) Stop(_ context.Context) error {
	d.stopOnce.Do(func() {
		d.stopError = d.db.Close()
	})

	return d.stopError
}
