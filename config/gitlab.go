package config

import "fmt"

type GitlabConfig struct {
	BaseURL    string `yaml:"baseUrl"`
	APIVersion string `yaml:"apiVersion"`
}

func (c *GitlabConfig) GetGitlabAPIBaseURL() string {
	return fmt.Sprintf("%s/api/%s", c.BaseURL, c.APIVersion)
}
