package observability

import (
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	"github.com/nimaeskandary/app_repo/pkg/observability/go/internal"
	obs_types "github.com/nimaeskandary/app_repo/pkg/observability/go/types"

	"go.uber.org/fx"
)

func NewSlogLoggerModule() fx.Option {
	return di.NewFxModule[obs_types.Logger]("slog_logger", internal.NewSlogLogger)
}

func NewNoopLoggerModule() fx.Option {
	return di.NewFxModule[obs_types.Logger]("noop_logger", internal.NewNoopLogger)
}
