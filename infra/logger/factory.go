package logger

import (
	"log/slog"
)

type logger struct {
	slog    *slog.Logger
	showLog bool
}

type Logger interface {
	Infof(msg string, args ...any)
	Warnf(msg string, args ...any)
	Errorf(msg string, args ...any)

	Print(msg string)
	Printf(msg string, args ...any)
	Guidef(msg string, args ...any)
	Notef(msg string, args ...any)
}

func NewLogger(distributionLog bool) Logger {
	slogLogger := slog.Default()

	return &logger{
		slog:    slogLogger,
		showLog: !distributionLog,
	}
}
