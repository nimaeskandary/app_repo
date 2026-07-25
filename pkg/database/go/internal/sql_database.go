package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
)

// sqlDatabase wraps a database connection pool managed by the Fx lifecycle.
type sqlDatabase struct {
	open    func() (*sql.DB, error)
	db      *sql.DB
	dialect db_core.Dialect
}

// newSQLDatabase creates a database whose connection is opened during Start.
func newSQLDatabase(open func() (*sql.DB, error), dialect db_core.Dialect) db_core.SQLDatabase {
	return &sqlDatabase{
		open:    open,
		dialect: dialect,
	}
}

func (d *sqlDatabase) DB() *sql.DB {
	return d.db
}

func (d *sqlDatabase) Dialect() db_core.Dialect {
	return d.dialect
}

func (d *sqlDatabase) Start(ctx context.Context) error {
	if d.db != nil {
		return nil
	}

	db, err := d.open()
	if err != nil {
		return err
	}
	if err := db.PingContext(ctx); err != nil {
		return errors.Join(
			fmt.Errorf("ping %s database: %w", d.dialect, err),
			db.Close(),
		)
	}

	d.db = db
	return nil
}

func (d *sqlDatabase) Stop(context.Context) error {
	if d.db == nil {
		return nil
	}

	db := d.db
	d.db = nil
	return db.Close()
}
