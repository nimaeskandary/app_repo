package internal

import (
	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
	obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"
)

type AppConfig struct {
	DataDir string
}

type Config struct {
	AppConfig   AppConfig                 `json:"-"`
	AppDatabase db_core.SQLiteConfig      `json:"AppDatabase" validate:"required"`
	Logger      obs_core.SlogLoggerConfig `json:"Logger" validate:"required"`
}
