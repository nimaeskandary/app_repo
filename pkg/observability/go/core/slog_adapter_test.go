package obs_core_test

import (
	"errors"
	"log/slog"
	"testing"

	obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"
	obs_core_mocks "github.com/nimaeskandary/app_repo/pkg/observability/go/core/mocks"
	"github.com/stretchr/testify/mock"
)

func TestAsSlog(t *testing.T) {
	t.Parallel()

	logger := obs_core_mocks.NewMockLogger(t)
	slogLogger := obs_core.AsSlog(logger)
	errValue := errors.New("failed")

	debugCall := logger.EXPECT().Debug(mock.Anything, "debug message", []any{slog.String("request_id", "debug")}).Once()
	infoCall := logger.EXPECT().Info(mock.Anything, "info message").Once()
	warnCall := logger.EXPECT().Warn(mock.Anything, "warning message").Once()
	errorCall := logger.EXPECT().Error(mock.Anything, "error message", []any{slog.Any("error", errValue)}).Once()
	groupedCall := logger.EXPECT().Info(
		mock.Anything,
		"grouped message",
		[]any{slog.Group("window", slog.Int("id", 1)), slog.Group("window", slog.String("state", "open"))},
	).Once()
	mock.InOrder(debugCall, infoCall, warnCall, errorCall, groupedCall)

	slogLogger.DebugContext(t.Context(), "debug message", "request_id", "debug")
	slogLogger.InfoContext(t.Context(), "info message")
	slogLogger.WarnContext(t.Context(), "warning message")
	slogLogger.ErrorContext(t.Context(), "error message", "error", errValue)
	slogLogger.WithGroup("window").With("id", 1).InfoContext(t.Context(), "grouped message", "state", "open")
}
