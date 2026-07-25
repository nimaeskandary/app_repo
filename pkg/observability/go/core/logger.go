package obs_core

import (
	"context"

	di "github.com/nimaeskandary/app_repo/pkg/di/go"
)

//mockery:generate: true
type Logger interface {
	di.Lifecycle
	Debug(ctx context.Context, msg string, attributes ...any)
	Info(ctx context.Context, msg string, attributes ...any)
	Warn(ctx context.Context, msg string, attributes ...any)
	Error(ctx context.Context, msg string, attributes ...any)
	CtxWithLogAttributes(ctx context.Context, attributes ...any) context.Context
}
