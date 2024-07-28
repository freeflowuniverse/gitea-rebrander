package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config is a struct that holds the configuration for the application.
type Config struct {
	GiteaRepoPath string            `yaml:"gitea-repo-path"`
	Files         map[string]string `yaml:"files"`
}

// ReadConfig is a function that reads the configuration from the config.yaml file.
func ReadConfig() (*Config, error) {
	var c Config

	out, err := os.Open("config.yaml")

	if err != nil {
		return nil, err
	}

	defer out.Close()

	err = yaml.NewDecoder(out).Decode(&c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
