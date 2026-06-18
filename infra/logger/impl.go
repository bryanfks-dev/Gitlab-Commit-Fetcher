package logger

import (
	"fmt"
)

func (l *logger) Infof(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.slog.Info(msg)
}

func (l *logger) Warnf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.slog.Warn(msg)
}

func (l *logger) Errorf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.slog.Error(msg)
}

func (l *logger) Print(msg string) {
	fmt.Print(msg)
}

func (l *logger) Printf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println(msg)
}
