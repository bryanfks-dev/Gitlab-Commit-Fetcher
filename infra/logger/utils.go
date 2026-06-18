package logger

import (
	"fmt"
	"main/core/utils"
	"time"
)

func newLoggerFilepath() string {
	date := time.Now().Format("2006-01-02")
	randomStr := utils.GenerateRandomString(16)
	return fmt.Sprintf("logs/%s-%s.log", date, randomStr)
}
