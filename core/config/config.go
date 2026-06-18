package config

func (c *Config) ReloadConfig() {
	cfg := LoadFromYAMLFile()
	c = cfg
}

func (c *Config) RegisterProject(projectCfg ProjectConfig) error {
	if c.CheckProjectExists(projectCfg.ID) {
		return nil
	}
	var projects []ProjectConfig
	projects = append(projects, projectCfg)
	projects = append(projects, c.RegisteredProjects...)
	return writeConfigToYAMLFile(c)
}

func (c *Config) UnregisterProject(projectID int) error {
	for i, project := range c.RegisteredProjects {
		if project.ID != projectID {
			continue
		}
		c.RegisteredProjects = append(
			c.RegisteredProjects[:i],
			c.RegisteredProjects[i+1:]...,
		)
	}
	return writeConfigToYAMLFile(c)
}

func (c *Config) CheckProjectExists(projectID int) bool {
	for _, project := range c.RegisteredProjects {
		if project.ID == projectID {
			return true
		}
	}
	return false
}
