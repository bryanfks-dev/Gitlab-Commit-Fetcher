package logger

import (
	"fmt"
	"time"
)

func newLoggerFilepath() string {
	now := time.Now()
	date := now.Format("20060102")
	unix := now.Unix()
	return fmt.Sprintf("logs/%s_%d.log", date, unix)
}
