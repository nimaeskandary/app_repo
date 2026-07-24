package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
)

// sqlDatabase wraps a database connection pool managed by the Fx lifecycle.
type sqlDatabase struct {
	open      func() (*sql.DB, error)
	db        *sql.DB
	dialect   db_types.Dialect
	startOnce sync.Once
	startErr  error
	stopOnce  sync.Once
	stopError error
}

// newSQLDatabase creates a database whose connection is opened during Start.
func newSQLDatabase(open func() (*sql.DB, error), dialect db_types.Dialect) db_types.SQLDatabase {
	return &sqlDatabase{
		open:    open,
		dialect: dialect,
	}
}

func (d *sqlDatabase) DB() *sql.DB {
	return d.db
}

func (d *sqlDatabase) Dialect() db_types.Dialect {
	return d.dialect
}

func (d *sqlDatabase) Start(ctx context.Context) error {
	d.startOnce.Do(func() {
		db, err := d.open()
		if err != nil {
			d.startErr = err
			return
		}
		if err := db.PingContext(ctx); err != nil {
			d.startErr = errors.Join(
				fmt.Errorf("ping %s database: %w", d.dialect, err),
				db.Close(),
			)
			return
		}
		d.db = db
	})

	return d.startErr
}

func (d *sqlDatabase) Stop(context.Context) error {
	d.stopOnce.Do(func() {
		if d.db != nil {
			d.stopError = d.db.Close()
		}
	})

	return d.stopError
}
