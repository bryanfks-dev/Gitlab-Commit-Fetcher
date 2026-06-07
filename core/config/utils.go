package config

import (
	"fmt"
	"main/core/consts"
	"main/core/utils"
	"os"

	"github.com/goccy/go-yaml"
)

func initializeYAMLFile() error {
	defaultCfg := Config{
		Gitlab: GitlabConfig{
			BaseURL:    consts.DefaultGitlabBaseURL,
			APIVersion: consts.DefaultGitlabAPIVersion,
		},
	}
	return writeConfigToYAMLFile(&defaultCfg)
}

func writeConfigToYAMLFile(cfg *Config) error {
	exeDirpath, err := utils.GetRunningExecutableDirpath()
	if err != nil {
		panic("Error getting running executable dirpath: " + err.Error())
	}
	filepath := fmt.Sprintf("%s/%s", exeDirpath, consts.ConfigYAMLFilename)

	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath, bytes, 0644); err != nil {
		return err
	}
	return nil
}
