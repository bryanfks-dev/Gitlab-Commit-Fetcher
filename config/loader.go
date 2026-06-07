package config

import (
	"fmt"
	"main/consts"
	"main/utils"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Gitlab GitlabConfig `yaml:"gitlab"`
}

func LoadFromYAMLFile() *Config {
	cfg := Config{}

	exeDirpath, err := utils.GetRunningExecutableDirpath()
	if err != nil {
		panic("Error getting running executable dirpath: " + err.Error())
	}

	configYAMLFilepath := fmt.Sprintf(
		"%s/%s",
		exeDirpath,
		consts.ConfigYAMLFilename,
	)
	if _, err := os.Stat(configYAMLFilepath); os.IsNotExist(err) {
		if err := initializeYAMLFile(configYAMLFilepath); err != nil {
			panic("Error initializing YAML file: " + err.Error())
		}
	}

	cfgFile, err := os.Open(configYAMLFilepath)
	if err != nil {
		msg := fmt.Sprintf("Error openning `%s`: %v", configYAMLFilepath, err)
		panic(msg)
	}
	defer cfgFile.Close()

	cfgDecoder := yaml.NewDecoder(cfgFile)
	if err := cfgDecoder.Decode(&cfg); err != nil {
		msg := fmt.Sprintf(
			"Error decoding YAML file `%s`: %v",
			configYAMLFilepath,
			err,
		)
		panic(msg)
	}

	return &cfg
}

func (c *Config) ReloadConfig() {
	cfg := LoadFromYAMLFile()
	c = cfg
}
