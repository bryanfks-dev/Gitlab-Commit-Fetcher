package config

type ProjectConfig struct {
	Name   string `yaml:"name"`
	ID     int    `yaml:"id"`
	WebURL string `yaml:"webUrl"`
}

func NewProjectConfig(
	id int,
	name string,
	webURL string,
) *ProjectConfig {
	return &ProjectConfig{
		ID:     id,
		Name:   name,
		WebURL: webURL,
	}
}
