package obs

import (
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"
	"github.com/nimaeskandary/app_repo/pkg/observability/go/internal"

	"go.uber.org/fx"
)

func NewSlogLoggerModule() fx.Option {
	return di.NewFxModule[obs_core.Logger]("slog_logger", internal.NewSlogLogger)
}

func NewNoopLoggerModule() fx.Option {
	return di.NewFxModule[obs_core.Logger]("noop_logger", internal.NewNoopLogger)
}
