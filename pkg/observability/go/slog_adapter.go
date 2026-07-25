package observability

import (
	"context"
	"log/slog"

	obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"
)

// AsSlog adapts Logger for APIs, such as Wails, that require a slog.Logger.
func AsSlog(logger obs_core.Logger) *slog.Logger {
	return slog.New(&slogHandler{logger: logger})
}

type slogHandler struct {
	logger obs_core.Logger
	attrs  []slog.Attr
	groups []string
}

func (h *slogHandler) Enabled(context.Context, slog.Level) bool {
	return true
}

func (h *slogHandler) Handle(ctx context.Context, record slog.Record) error {
	attributes := make([]any, 0, len(h.attrs)+record.NumAttrs())
	for _, attr := range h.attrs {
		attributes = append(attributes, attr)
	}

	record.Attrs(func(attr slog.Attr) bool {
		// Apply groups configured on the derived logger to attributes for this record.
		attributes = append(attributes, h.withGroups(attr))
		return true
	})

	switch {
	case record.Level <= slog.LevelDebug:
		h.logger.Debug(ctx, record.Message, attributes...)
	case record.Level < slog.LevelWarn:
		h.logger.Info(ctx, record.Message, attributes...)
	case record.Level < slog.LevelError:
		h.logger.Warn(ctx, record.Message, attributes...)
	default:
		h.logger.Error(ctx, record.Message, attributes...)
	}
	return nil
}

func (h *slogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Retain these attributes for every record emitted by the derived logger.
	newAttrs := make([]slog.Attr, 0, len(h.attrs)+len(attrs))
	newAttrs = append(newAttrs, h.attrs...)
	for _, attr := range attrs {
		newAttrs = append(newAttrs, h.withGroups(attr))
	}
	return &slogHandler{logger: h.logger, attrs: newAttrs, groups: h.groups}
}

func (h *slogHandler) WithGroup(name string) slog.Handler {
	// Retain the group path so Handle can wrap attributes from the derived logger.
	if name == "" {
		return h
	}
	groups := make([]string, 0, len(h.groups)+1)
	groups = append(groups, h.groups...)
	groups = append(groups, name)
	return &slogHandler{logger: h.logger, attrs: h.attrs, groups: groups}
}

func (h *slogHandler) withGroups(attr slog.Attr) slog.Attr {
	for index := len(h.groups) - 1; index >= 0; index-- {
		attr = slog.Group(h.groups[index], attr)
	}
	return attr
}
