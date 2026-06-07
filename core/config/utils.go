package config

import (
	"main/core/consts"
	"os"

	"github.com/goccy/go-yaml"
)

func initializeYAMLFile(filepath string) error {
	defaultCfg := Config{
		Gitlab: GitlabConfig{
			BaseURL:    consts.DefaultGitlabBaseURL,
			APIVersion: consts.DefaultGitlabAPIVersion,
		},
	}
	bytes, err := yaml.Marshal(&defaultCfg)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath, bytes, 0644); err != nil {
		return err
	}
	return nil
}
