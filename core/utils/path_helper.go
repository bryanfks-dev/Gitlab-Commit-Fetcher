package utils

import (
	"fmt"
	"main/core/consts"
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

func NewCommitMessagesFilepath(suffixName string) (string, error) {
	dirPath, err := GetRunningExecutableDirpath()
	if err != nil {
		return "", err
	}
	filepath := fmt.Sprintf(
		"%s/%s/%s-%s.txt",
		dirPath,
		consts.CommitMessagesFileDirname,
		consts.CommitMessagesPrefixFilename,
		suffixName,
	)
	return filepath, nil
}
