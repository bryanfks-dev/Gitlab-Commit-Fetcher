package logger

import (
	"log/slog"
	"os"
)

type logger struct {
	slog *slog.Logger
}

type Logger interface {
	Infof(msg string, args ...any)
	Warnf(msg string, args ...any)
	Errorf(msg string, args ...any)

	Print(msg string)
	Printf(msg string, args ...any)
}

func NewLogger() Logger {
	slogLogger := slog.Default()
	logFilepath := newLoggerFilepath()
	logFile, err := os.OpenFile(
		logFilepath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		slogLogger.Error("Failed to open log file: " + err.Error())
	}
	slogLogger = slog.New(slog.NewJSONHandler(logFile, nil))

	return &logger{
		slog: slogLogger,
	}
}
