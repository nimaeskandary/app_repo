package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strings"

	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"

	_ "modernc.org/sqlite"
)

const sqliteMaxOpenConnections = 1

// NewSQLiteDatabase creates a SQLite database that opens during Start.
func NewSQLiteDatabase(config db_core.SQLiteConfig) (db_core.SQLDatabase, error) {
	if config.Source == "" {
		return nil, errors.New("SQLite database source is required")
	}

	open := func() (*sql.DB, error) {
		db, err := sql.Open("sqlite", withSQLiteDefaults(config.Source))
		if err != nil {
			return nil, fmt.Errorf("open SQLite database: %w", err)
		}
		db.SetMaxOpenConns(sqliteMaxOpenConnections)
		return db, nil
	}

	return newSQLDatabase(open, db_core.DialectSQLite), nil
}

func withSQLiteDefaults(source string) string {
	separator := "?"
	if strings.Contains(source, "?") {
		separator = "&"
	}

	query := url.Values{}
	query.Add("_pragma", "foreign_keys(1)")
	query.Add("_pragma", "busy_timeout(5000)")

	return source + separator + query.Encode()
}
