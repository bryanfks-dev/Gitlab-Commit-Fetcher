package utils

import (
	"os"
	"path/filepath"
)

func GetRunningExecutableDirpath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	exePath := filepath.Dir(exe)
	return exePath, nil
}
