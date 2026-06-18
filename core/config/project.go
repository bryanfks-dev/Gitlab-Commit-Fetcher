package config

type ProjectConfig struct {
	Name   string `yaml:"name"`
	ID     int    `yaml:"id"`
	WebURL string `yaml:"webUrl"`
}
