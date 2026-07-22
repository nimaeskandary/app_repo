package obs_types

import "context"

type Logger interface {
	Debug(ctx context.Context, msg string, attributes ...any)
	Info(ctx context.Context, msg string, attributes ...any)
	Warn(ctx context.Context, msg string, attributes ...any)
	Error(ctx context.Context, msg string, attributes ...any)
	CtxWithLogAttributes(ctx context.Context, attributes ...any) context.Context
	Stop(ctx context.Context) error
}
