package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strings"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"

	_ "modernc.org/sqlite"
)

// sqliteMaxOpenConnections avoids concurrent SQLite writers in this application.
const sqliteMaxOpenConnections = 1

// NewSQLiteDatabase opens a validated SQLite database connection.
func NewSQLiteDatabase(cfg db_types.SQLiteConfig) (db_types.SQLDatabase, error) {
	if cfg.Source == "" {
		return nil, errors.New("SQLite database source is required")
	}

	db, err := sql.Open("sqlite", withSQLiteDefaults(cfg.Source))
	if err != nil {
		return nil, fmt.Errorf("open SQLite database: %w", err)
	}
	db.SetMaxOpenConns(sqliteMaxOpenConnections)

	return newSQLDatabase(db, db_types.DialectSQLite)
}

// withSQLiteDefaults adds connection-level pragmas to every pooled connection.
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
