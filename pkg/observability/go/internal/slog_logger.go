package internal

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"
)

type attributesContextKey struct{}

type slogger struct {
	logger *slog.Logger
}

func NewSlogLogger(cfg obs_core.SlogLoggerConfig) (obs_core.Logger, error) {
	var level slog.Level
	switch cfg.Level {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		return nil, fmt.Errorf("unknown log level %v", cfg.Level)
	}

	return &slogger{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})),
	}, nil
}

func (l *slogger) Debug(ctx context.Context, msg string, attributes ...any) {
	l.logger.DebugContext(ctx, msg, append(attributesFromContext(ctx), attributes...)...)
}

func (l *slogger) Info(ctx context.Context, msg string, attributes ...any) {
	l.logger.InfoContext(ctx, msg, append(attributesFromContext(ctx), attributes...)...)
}

func (l *slogger) Warn(ctx context.Context, msg string, attributes ...any) {
	l.logger.WarnContext(ctx, msg, append(attributesFromContext(ctx), attributes...)...)
}

func (l *slogger) Error(ctx context.Context, msg string, attributes ...any) {
	l.logger.ErrorContext(ctx, msg, append(attributesFromContext(ctx), attributes...)...)
}

func (l *slogger) CtxWithLogAttributes(ctx context.Context, attributes ...any) context.Context {
	existingAttrs, _ := ctx.Value(attributesContextKey{}).([]any)
	newAttrs := make([]any, 0, len(existingAttrs)+len(attributes))
	newAttrs = append(newAttrs, existingAttrs...)
	newAttrs = append(newAttrs, attributes...)
	return context.WithValue(ctx, attributesContextKey{}, newAttrs)
}

func (l *slogger) Start(context.Context) error {
	return nil
}

func (l *slogger) Stop(context.Context) error {
	return nil
}

func attributesFromContext(ctx context.Context) []any {
	attrs, _ := ctx.Value(attributesContextKey{}).([]any)
	return attrs
}
