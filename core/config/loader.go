package config

import (
	"fmt"
	"main/core/consts"
	"main/core/utils"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Gitlab             GitlabConfig    `yaml:"gitlab"`
	RegisteredProjects []ProjectConfig `yaml:"registeredProjects"`
}

func LoadFromYAMLFile() *Config {
	cfg := Config{}

	exeDirpath, err := utils.GetRunningExecutableDirpath()
	if err != nil {
		panic("Error getting running executable dirpath: " + err.Error())
	}

	filepath := fmt.Sprintf(
		"%s/%s",
		exeDirpath,
		consts.ConfigYAMLFilename,
	)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		if err := initializeYAMLFile(); err != nil {
			panic("Error initializing YAML file: " + err.Error())
		}
	}

	cfgFile, err := os.Open(filepath)
	if err != nil {
		msg := fmt.Sprintf("Error openning `%s`: %v", filepath, err)
		panic(msg)
	}
	defer cfgFile.Close()

	cfgDecoder := yaml.NewDecoder(cfgFile)
	if err := cfgDecoder.Decode(&cfg); err != nil {
		msg := fmt.Sprintf(
			"Error decoding YAML file `%s`: %v",
			filepath,
			err,
		)
		panic(msg)
	}

	return &cfg
}
