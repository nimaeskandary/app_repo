package internal

import (
	app_database "github.com/nimaeskandary/app_repo/cmd/gordle_app/internal/database/app"
	app_migrations "github.com/nimaeskandary/app_repo/cmd/gordle_app/internal/database/app/migrations"
	database "github.com/nimaeskandary/app_repo/pkg/database/go"
	obs "github.com/nimaeskandary/app_repo/pkg/observability/go"
	"go.uber.org/fx"
)

func ModuleList(config Config) []fx.Option {
	return []fx.Option{
		database.NewSQLiteWriterModule[app_database.AppDBWriter](config.AppDatabase, config.AppConfig.DataDir),
		database.NewMigratorModule[
			app_database.AppDBWriter,
			app_database.AppDatabaseMigrator,
		](app_migrations.MigrationSource()),
		database.NewMigrateAllOnStartModule[
			app_database.AppDatabaseMigrator,
			app_database.AppDatabaseMigrateAllOnStart,
		](),
		database.NewSQLiteReaderModule[app_database.AppDBReader](config.AppDatabase, config.AppConfig.DataDir),
		obs.NewSlogLoggerModule(config.Logger),
	}
}
