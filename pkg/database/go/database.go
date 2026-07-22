package database

import (
	"github.com/nimaeskandary/app_repo/pkg/database/go/internal"
	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	"go.uber.org/fx"
)

func NewSQLiteDatabaseModule() fx.Option {
	return di.NewFxModule[db_types.SQLDatabase]("sqlite_database", internal.NewSQLiteDatabase)
}
