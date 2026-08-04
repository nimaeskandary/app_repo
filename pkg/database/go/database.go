package database

import (
	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
	"github.com/nimaeskandary/app_repo/pkg/database/go/internal"
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	"go.uber.org/fx"
)

// NewSQLiteReaderModule provides a query-only SQLite database managed by the Fx lifecycle.
func NewSQLiteReaderModule[
	Database db_core.SQLDatabase,
](config db_core.SQLiteConfig, dbDir string) fx.Option {
	constructor := func() (db_core.SQLDatabase, error) {
		return internal.NewSQLiteReader(config, dbDir)
	}
	return di.NewFxModule[Database]("sqlite_reader", constructor)
}

// NewSQLiteWriterModule provides a WAL-backed SQLite database managed by the Fx lifecycle.
func NewSQLiteWriterModule[
	Database db_core.SQLDatabase,
](config db_core.SQLiteConfig, dbDir string) fx.Option {
	constructor := func() (db_core.SQLDatabase, error) {
		return internal.NewSQLiteWriter(config, dbDir)
	}
	return di.NewFxModule[Database]("sqlite_writer", constructor)
}

// NewMigratorModule provides a tool-neutral migrator backed by Goose.
func NewMigratorModule[
	Database db_core.SQLDatabase,
	Migrator db_core.Migrator,
](source db_core.MigrationSource) fx.Option {
	constructor := func(database Database) (db_core.Migrator, error) {
		return internal.NewGooseMigrator(database, source)
	}
	return di.NewFxModule[Migrator]("migrator", constructor)
}

// NewMigrateAllOnStartModule runs all pending migrations during Fx startup.
func NewMigrateAllOnStartModule[
	Migrator db_core.Migrator,
	MigrateAllOnStart db_core.MigrateAllOnStart,
]() fx.Option {
	constructor := func(migrator Migrator) (db_core.MigrateAllOnStart, error) {
		return internal.NewMigrateAllOnStart(migrator)
	}
	return di.NewFxModule[MigrateAllOnStart]("migrate_all", constructor)
}
