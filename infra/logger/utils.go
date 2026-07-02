package logger

import (
	"fmt"
	"main/core/utils"
	"time"
)

func newLoggerFilepath() string {
	exeDirpath, err := utils.GetRunningExecutableDirpath()
	if err != nil {
		panic(err)
	}

	now := time.Now()
	date := now.Format("20060102")
	unix := now.Unix()
	return fmt.Sprintf("%s/logs/%s_%d.log", exeDirpath, date, unix)
}
