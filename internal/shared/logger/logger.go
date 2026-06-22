package logger

import (
	"log/slog"
	"os"
	"strings"
)

func New(level string, format string) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: parseLevel(level),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := a.Value.Time()
				return slog.String(a.Key, t.Format("2026-04-05 15:04:05"))
			}
			return a
		}}

	return slog.New(getHandler(format, opts))
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func getHandler(format string, opts *slog.HandlerOptions) slog.Handler {
	switch format {
	case "json":
		return slog.NewJSONHandler(os.Stdout, opts)
	default:
		return slog.NewTextHandler(os.Stdout, opts)
	}
}
