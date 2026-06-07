package logger

import (
	"fmt"
)

func (l *logger) Infof(format string, args ...any) {
	if !l.showLog {
		return
	}
	msg := fmt.Sprintf(format+"\n", args...)
	l.slog.Info(msg)
}

func (l *logger) Warnf(format string, args ...any) {
	if !l.showLog {
		return
	}
	msg := fmt.Sprintf(format+"\n", args...)
	l.slog.Warn(msg)
}

func (l *logger) Errorf(format string, args ...any) {
	if !l.showLog {
		return
	}
	msg := fmt.Sprintf(format+"\n", args...)
	l.slog.Error(msg)
}

func (l *logger) Print(msg string) {
	fmt.Print(msg)
}

func (l *logger) Printf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println(msg)
}

func (l *logger) Guidef(format string, args ...any) {
	msg := fmt.Sprintf("[GUIDE] "+format, args...)
	fmt.Println(msg)
}

func (l *logger) Notef(format string, args ...any) {
	msg := fmt.Sprintf("[NOTE] "+format, args...)
	fmt.Println(msg)
}
