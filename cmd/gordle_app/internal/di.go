package internal

import (
	app_database "github.com/nimaeskandary/app_repo/cmd/gordle_app/internal/database/app"
	app_migrations "github.com/nimaeskandary/app_repo/cmd/gordle_app/internal/database/app/migrations"
	config "github.com/nimaeskandary/app_repo/pkg/config/go"
	config_types "github.com/nimaeskandary/app_repo/pkg/config/go/types"
	database "github.com/nimaeskandary/app_repo/pkg/database/go"
	greet "github.com/nimaeskandary/app_repo/pkg/greet/go"
	observability "github.com/nimaeskandary/app_repo/pkg/observability/go"
	obs_types "github.com/nimaeskandary/app_repo/pkg/observability/go/types"
	"go.uber.org/fx"
)

type Config struct {
	AppDatabase app_database.AppConfig     `json:"AppDatabase" validate:"required"`
	Logger      obs_types.SlogLoggerConfig `json:"Logger" validate:"required"`
}

func ModuleList(configBytes []byte) []fx.Option {
	return []fx.Option{
		config.NewJsonConfigLoaderModule[Config](configBytes, nil),
		fx.Provide(
			func(loader config_types.ConfigLoader[Config]) app_database.AppConfig {
				return loader.GetConfig().AppDatabase
			},
			// takes the AppConfig and does transformations before putting it back on dep graph
			app_database.NewAppSQLiteConfig,
			func(loader config_types.ConfigLoader[Config]) obs_types.SlogLoggerConfig {
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
		database.NewMigrateAllModule[
			app_database.AppDatabaseMigrator,
			app_database.AppDatabaseMigrateAll,
		](),
		observability.NewSlogLoggerModule(),
		greet.NewGreetModule(),
	}
}
