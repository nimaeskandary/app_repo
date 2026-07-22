package internal

import (
	"context"
	"database/sql"
	"sync"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
)

type sqlDatabase struct {
	db        *sql.DB
	dialect   db_types.Dialect
	stopOnce  sync.Once
	stopError error
}

func newSQLDatabase(db *sql.DB, dialect db_types.Dialect) db_types.SQLDatabase {
	return &sqlDatabase{
		db:      db,
		dialect: dialect,
	}
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
