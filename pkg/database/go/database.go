package database

import (
	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
	"github.com/nimaeskandary/app_repo/pkg/database/go/internal"
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	"go.uber.org/fx"
)

// NewSQLiteDatabaseModule provides a SQLite database managed by the Fx lifecycle.
func NewSQLiteDatabaseModule[
	Database db_core.SQLDatabase,
	Config db_core.SQLiteConfigProvider,
]() fx.Option {
	constructor := func(config Config) (db_core.SQLDatabase, error) {
		return internal.NewSQLiteDatabase(config.SQLiteConfig())
	}
	return di.NewFxModule[Database]("sqlite_database", constructor)
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
