package internal

import (
	"context"

	obs_types "github.com/nimaeskandary/app_repo/pkg/observability/go/types"
)

type noopLogger struct{}

func NewNoopLogger() obs_types.Logger {
	return &noopLogger{}
}

func (l *noopLogger) Debug(context.Context, string, ...any) {}

func (l *noopLogger) Info(context.Context, string, ...any) {}

func (l *noopLogger) Warn(context.Context, string, ...any) {}

func (l *noopLogger) Error(context.Context, string, ...any) {}

func (l *noopLogger) CtxWithLogAttributes(ctx context.Context, _ ...any) context.Context {
	return ctx
}

func (l *noopLogger) Start(context.Context) error {
	return nil
}

func (l *noopLogger) Stop(context.Context) error {
	return nil
}
