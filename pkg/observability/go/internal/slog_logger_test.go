package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"testing"

	obs_types "github.com/nimaeskandary/app_repo/pkg/observability/go/types"
	"github.com/stretchr/testify/require"
)

func TestNewSlogLogger(t *testing.T) {
	t.Parallel()

	logger, err := NewSlogLogger(obs_types.SlogLoggerConfig{Level: "DEBUG"})
	require.NoError(t, err)
	require.NoError(t, logger.Stop(t.Context()))

	_, err = NewSlogLogger(obs_types.SlogLoggerConfig{Level: "TRACE"})
	require.EqualError(t, err, "unknown log level TRACE")
}

func TestSlogger(t *testing.T) {
	t.Parallel()

	t.Run("Debug", func(t *testing.T) {
		t.Parallel()

		logger, buffer := newTestSlogger()
		logger.Debug(t.Context(), "debug message", "request_id", "debug")

		require.Equal(t, map[string]any{
			"level":      "DEBUG",
			"msg":        "debug message",
			"request_id": "debug",
		}, loggedRecord(t, buffer))
	})

	t.Run("Info", func(t *testing.T) {
		t.Parallel()

		logger, buffer := newTestSlogger()
		logger.Info(t.Context(), "info message", "request_id", "info")

		require.Equal(t, map[string]any{
			"level":      "INFO",
			"msg":        "info message",
			"request_id": "info",
		}, loggedRecord(t, buffer))
	})

	t.Run("Warn", func(t *testing.T) {
		t.Parallel()

		logger, buffer := newTestSlogger()
		logger.Warn(t.Context(), "warning message", "request_id", "warn")

		require.Equal(t, map[string]any{
			"level":      "WARN",
			"msg":        "warning message",
			"request_id": "warn",
		}, loggedRecord(t, buffer))
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		logger, buffer := newTestSlogger()
		logger.Error(t.Context(), "error message", "error", errors.New("failed"))

		require.Equal(t, map[string]any{
			"level": "ERROR",
			"msg":   "error message",
			"error": "failed",
		}, loggedRecord(t, buffer))
	})

	t.Run("CtxWithLogAttributes", func(t *testing.T) {
		t.Parallel()

		logger, buffer := newTestSlogger()
		ctx := logger.CtxWithLogAttributes(t.Context(), "request_id", "context")
		logger.Info(ctx, "context message", "operation", "test")

		require.Equal(t, map[string]any{
			"level":      "INFO",
			"msg":        "context message",
			"request_id": "context",
			"operation":  "test",
		}, loggedRecord(t, buffer))
	})

	t.Run("Stop", func(t *testing.T) {
		t.Parallel()

		logger, _ := newTestSlogger()
		require.NoError(t, logger.Stop(t.Context()))
	})
}

func newTestSlogger() (*slogger, *bytes.Buffer) {
	buffer := &bytes.Buffer{}
	return &slogger{
		logger: slog.New(slog.NewJSONHandler(buffer, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}, buffer
}

func loggedRecord(t *testing.T, buffer *bytes.Buffer) map[string]any {
	t.Helper()

	var record map[string]any
	require.NoError(t, json.Unmarshal(buffer.Bytes(), &record))
	delete(record, "time")
	return record
}
