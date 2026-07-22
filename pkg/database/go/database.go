package database

import (
	"github.com/nimaeskandary/app_repo/pkg/database/go/internal"
	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	"go.uber.org/fx"
)

// NewSQLiteDatabaseModule provides a SQLite database managed by the Fx lifecycle.
func NewSQLiteDatabaseModule[
	Database db_types.SQLDatabase,
	Config db_types.SQLiteConfigProvider,
]() fx.Option {
	constructor := func(cfg Config) (db_types.SQLDatabase, error) {
		return internal.NewSQLiteDatabase(cfg.SQLiteConfig())
	}
	return di.NewFxModule[Database]("sqlite_database", constructor)
}

// NewMigratorModule provides a tool-neutral migrator backed by Goose.
func NewMigratorModule[
	Database db_types.SQLDatabase,
	Migrator db_types.Migrator,
](source db_types.MigrationSource) fx.Option {
	constructor := func(database Database) (db_types.Migrator, error) {
		return internal.NewGooseMigrator(database, source)
	}
	return di.NewFxModule[Migrator]("migrator", constructor)
}

// NewMigrateAllModule runs all pending migrations during dependency construction.
func NewMigrateAllModule[
	Migrator db_types.Migrator,
	MigrateAll db_types.MigrateAll,
]() fx.Option {
	constructor := func(migrator Migrator) (db_types.MigrateAll, error) {
		return internal.NewMigrateAll(migrator)
	}
	return di.NewFxModule[MigrateAll]("migrate_all", constructor)
}
