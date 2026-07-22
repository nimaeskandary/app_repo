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

const sqliteMaxOpenConnections = 1

func NewSQLiteDatabase(cfg db_types.SQLiteConfig) (db_types.SQLDatabase, error) {
	if cfg.Source == "" {
		return nil, errors.New("SQLite database source is required")
	}

	db, err := sql.Open("sqlite", withSQLiteDefaults(cfg.Source))
	if err != nil {
		return nil, fmt.Errorf("open SQLite database: %w", err)
	}
	db.SetMaxOpenConns(sqliteMaxOpenConnections)

	if err := db.Ping(); err != nil {
		return nil, errors.Join(
			fmt.Errorf("ping SQLite database: %w", err),
			db.Close(),
		)
	}

	return newSQLDatabase(db, db_types.DialectSQLite), nil
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
