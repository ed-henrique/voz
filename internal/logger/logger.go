package logger

import "log/slog"

func Error(msg string, args ...any) {
	slog.Error("ERROR: "+msg, args...)
}
