package observability

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	obs_types "github.com/nimaeskandary/app_repo/pkg/observability/go/types"
	"github.com/stretchr/testify/require"
)

func TestAsSlog(t *testing.T) {
	t.Parallel()

	logger := &recordingLogger{}
	slogLogger := AsSlog(logger)
	errValue := errors.New("failed")

	slogLogger.DebugContext(t.Context(), "debug message", "request_id", "debug")
	slogLogger.InfoContext(t.Context(), "info message")
	slogLogger.WarnContext(t.Context(), "warning message")
	slogLogger.ErrorContext(t.Context(), "error message", "error", errValue)
	slogLogger.WithGroup("window").With("id", 1).InfoContext(t.Context(), "grouped message", "state", "open")

	require.Equal(t, []logEntry{
		{level: "debug", message: "debug message", attributes: []any{slog.String("request_id", "debug")}},
		{level: "info", message: "info message", attributes: []any{}},
		{level: "warn", message: "warning message", attributes: []any{}},
		{level: "error", message: "error message", attributes: []any{slog.Any("error", errValue)}},
		{level: "info", message: "grouped message", attributes: []any{slog.Group("window", slog.Int("id", 1)), slog.Group("window", slog.String("state", "open"))}},
	}, logger.entries)
}

type logEntry struct {
	level      string
	message    string
	attributes []any
}

type recordingLogger struct {
	entries []logEntry
}

var _ obs_types.Logger = (*recordingLogger)(nil)

func (l *recordingLogger) Debug(_ context.Context, msg string, attributes ...any) {
	l.entries = append(l.entries, logEntry{level: "debug", message: msg, attributes: attributes})
}

func (l *recordingLogger) Info(_ context.Context, msg string, attributes ...any) {
	l.entries = append(l.entries, logEntry{level: "info", message: msg, attributes: attributes})
}

func (l *recordingLogger) Warn(_ context.Context, msg string, attributes ...any) {
	l.entries = append(l.entries, logEntry{level: "warn", message: msg, attributes: attributes})
}

func (l *recordingLogger) Error(_ context.Context, msg string, attributes ...any) {
	l.entries = append(l.entries, logEntry{level: "error", message: msg, attributes: attributes})
}

func (l *recordingLogger) CtxWithLogAttributes(ctx context.Context, _ ...any) context.Context {
	return ctx
}

func (l *recordingLogger) Start(context.Context) error {
	return nil
}

func (l *recordingLogger) Stop(context.Context) error {
	return nil
}
