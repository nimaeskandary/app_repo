package db_core

import (
	"context"
	"database/sql"
	"io/fs"

	di "github.com/nimaeskandary/app_repo/pkg/di/go"
)

// MigrationFunc applies or reverts a migration in a transaction.
type MigrationFunc func(ctx context.Context, tx *sql.Tx) error

// Migration defines both directions of one versioned migration.
type Migration struct {
	// Version uniquely identifies the migration.
	Version int64
	// Name describes the migration.
	Name string
	// Up applies the migration.
	Up MigrationFunc
	// Down reverts the migration.
	Down MigrationFunc
}

// MigrationSource groups SQL files and code migrations for one database.
type MigrationSource struct {
	// SQLFiles contains paired .up.sql and .down.sql migration files.
	SQLFiles fs.FS
	// CodeMigrations contains migrations implemented as Go functions.
	CodeMigrations []Migration
}

// Migrator runs migrations in either direction.
//
//mockery:generate: true
type Migrator interface {
	di.Lifecycle
	// Up runs all pending migrations when version is nil, or the exact supplied version.
	Up(ctx context.Context, version *int64) error
	// Down reverts the exact supplied version.
	Down(ctx context.Context, version int64) error
}

// MigrateAllOnStart runs migration work performed during startup.
type MigrateAllOnStart interface {
	di.Lifecycle
}
