package logger

import "log/slog"

func Debug(msg string, args ...any) {
	slog.Debug("DEBUG: "+msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error("ERROR: "+msg, args...)
}
