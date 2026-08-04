package internal

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/ncruces/go-sqlite3"
	sqlite_driver "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/vfs/adiantum"
	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
)

const (
	sqliteMaxOpenConnections = 1
	sqliteEncryptionKeySize  = 32
)

// NewSQLiteReader creates a query-only SQLite database that opens during Start.
func NewSQLiteReader(config db_core.SQLiteConfig) (db_core.SQLDatabase, error) {
	return newSQLiteDatabase(config, configureSQLiteReader)
}

// NewSQLiteWriter creates a WAL-backed SQLite database that opens during Start.
func NewSQLiteWriter(config db_core.SQLiteConfig) (db_core.SQLDatabase, error) {
	return newSQLiteDatabase(config, configureSQLiteWriter)
}

func newSQLiteDatabase(
	config db_core.SQLiteConfig,
	configure func(*sqlite3.Conn) error,
) (db_core.SQLDatabase, error) {
	if config.Source == "" {
		return nil, errors.New("SQLite database source is required")
	}

	source := config.Source
	configureConnection := configure
	if config.IsEncrypted {
		if len(config.EncryptionKey) != sqliteEncryptionKeySize {
			return nil, fmt.Errorf("SQLite encryption key must be %d bytes", sqliteEncryptionKeySize)
		}

		var err error
		source, err = encryptedSQLiteSource(source)
		if err != nil {
			return nil, err
		}
		encryptionKey := append([]byte(nil), config.EncryptionKey...)
		configureConnection = func(conn *sqlite3.Conn) error {
			if err := conn.Exec("PRAGMA hexkey = '" + hex.EncodeToString(encryptionKey) + "'"); err != nil {
				return fmt.Errorf("set SQLite encryption key: %w", err)
			}
			return configure(conn)
		}
	}

	open := func() (*sql.DB, error) {
		db, err := sqlite_driver.Open(source, configureConnection)
		if err != nil {
			return nil, fmt.Errorf("open SQLite database: %w", err)
		}
		db.SetMaxOpenConns(sqliteMaxOpenConnections)
		return db, nil
	}

	return newSQLDatabase(open, db_core.DialectSQLite), nil
}

// encryptedSQLiteSource configures a SQLite source to use the Adiantum VFS.
// see https://github.com/ncruces/go-sqlite3/blob/main/vfs/adiantum/README.md
func encryptedSQLiteSource(source string) (string, error) {
	var parsed *url.URL
	var err error
	if strings.HasPrefix(source, "file:") {
		parsed, err = url.Parse(source)
		if err != nil {
			return "", fmt.Errorf("parse SQLite database source: %w", err)
		}
	} else {
		parsed = &url.URL{Scheme: "file", Path: filepath.ToSlash(source)}
	}

	parameters, err := url.ParseQuery(parsed.RawQuery)
	if err != nil {
		return "", fmt.Errorf("parse SQLite database source parameters: %w", err)
	}
	if configuredVFS := parameters.Get("vfs"); configuredVFS != "" && configuredVFS != "adiantum" {
		return "", fmt.Errorf("SQLite database source VFS is %q, encryption requires %q", configuredVFS, "adiantum")
	}
	parameters.Set("vfs", "adiantum")
	parsed.RawQuery = parameters.Encode()
	return parsed.String(), nil
}

func configureSQLiteReader(conn *sqlite3.Conn) error {
	if err := conn.Exec("PRAGMA query_only = ON"); err != nil {
		return fmt.Errorf("make SQLite connection query-only: %w", err)
	}
	return nil
}

func configureSQLiteWriter(conn *sqlite3.Conn) error {
	statement, _, err := conn.Prepare("PRAGMA journal_mode = WAL")
	if err != nil {
		return fmt.Errorf("enable SQLite WAL: %w", err)
	}
	defer func() {
		_ = statement.Close()
	}()

	if !statement.Step() {
		if err := statement.Err(); err != nil {
			return fmt.Errorf("enable SQLite WAL: %w", err)
		}
		return errors.New("enable SQLite WAL: journal mode was not returned")
	}
	if journalMode := statement.ColumnText(0); journalMode != "wal" {
		return fmt.Errorf("enable SQLite WAL: journal mode is %q", journalMode)
	}

	return configureSQLiteForeignKeys(conn)
}

func configureSQLiteForeignKeys(conn *sqlite3.Conn) error {
	if err := conn.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return fmt.Errorf("enable SQLite foreign keys: %w", err)
	}
	return nil
}
