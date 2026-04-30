package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const DefaultURL = "https://acm.altinity.cloud/api/"

type Config struct {
	URL    string `yaml:"url"`
	Token  string `yaml:"token"`
	Output string `yaml:"output"`
}

func DefaultPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".acmctl.yaml")
}

func Load(path string) (*Config, error) {
	cfg := &Config{
		URL:    DefaultURL,
		Output: "table",
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	if cfg.URL == "" {
		cfg.URL = DefaultURL
	}
	if cfg.Output == "" {
		cfg.Output = "table"
	}
	return cfg, nil
}

func Save(path string, cfg *Config) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
