package internal

import (
	app_database "github.com/nimaeskandary/app_repo/cmd/gordle_app/internal/database/app"
	app_migrations "github.com/nimaeskandary/app_repo/cmd/gordle_app/internal/database/app/migrations"
	config "github.com/nimaeskandary/app_repo/pkg/config/go"
	config_core "github.com/nimaeskandary/app_repo/pkg/config/go/core"
	database "github.com/nimaeskandary/app_repo/pkg/database/go"
	observability "github.com/nimaeskandary/app_repo/pkg/observability/go"
	obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"
	"go.uber.org/fx"
)

type Config struct {
	AppDatabase app_database.AppConfig    `json:"AppDatabase" validate:"required"`
	Logger      obs_core.SlogLoggerConfig `json:"Logger" validate:"required"`
}

func ModuleList(configBytes []byte) []fx.Option {
	return []fx.Option{
		config.NewJsonConfigLoaderModule[Config](configBytes, nil),
		fx.Provide(
			func(loader config_core.ConfigLoader[Config]) app_database.AppConfig {
				return loader.GetConfig().AppDatabase
			},
			// takes the AppConfig and does transformations before putting it back on dep graph
			app_database.NewAppSQLiteConfig,
			func(loader config_core.ConfigLoader[Config]) obs_core.SlogLoggerConfig {
				return loader.GetConfig().Logger
			},
		),
		database.NewSQLiteDatabaseModule[
			app_database.AppDatabase,
			app_database.AppSQLiteConfig,
		](),
		database.NewMigratorModule[
			app_database.AppDatabase,
			app_database.AppDatabaseMigrator,
		](app_migrations.MigrationSource()),
		database.NewMigrateAllOnStartModule[
			app_database.AppDatabaseMigrator,
			app_database.AppDatabaseMigrateAllOnStart,
		](),
		observability.NewSlogLoggerModule(),
	}
}
